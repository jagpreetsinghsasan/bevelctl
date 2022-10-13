package k8ind

import (
	"bevelctl/tpls/kind"
	"bevelctl/utils"
	"bytes"
	"html/template"
	"os"
	"strconv"

	"github.com/Masterminds/sprig/v3"
	"github.com/manifoldco/promptui"
	"go.uber.org/zap"
)

// This struct is used to take user inputs where
// ControlPlaneCount refers to the k8s master node count
// WorkerNodeCount referes to the k8s worker node count
type KindClusterConfig struct {
	ControlPlaneCount int
	WorkerNodeCount   int
}

// This function takes the user input to construct the KindClusterConfig struct
// and return this struct
func getInputs(logger *zap.Logger) KindClusterConfig {
	controlPlaneCount := promptui.Prompt{
		Label:   "Enter the number of control plane nodes: ",
		Default: "1",
	}
	controlPlaneCountResult, err := controlPlaneCount.Run()
	if err != nil {
		logger.Fatal("Prompt failed entering the control plane node count", zap.Any("ERROR", err))
	}

	workerNodeCount := promptui.Prompt{
		Label:   "Enter the number of worker nodes: ",
		Default: "3",
	}
	workerNodeCountCountResult, err := workerNodeCount.Run()
	if err != nil {
		logger.Fatal("Prompt failed entering the worker node count", zap.Any("ERROR", err))
	}

	controlPlaneCountRes, _ := strconv.Atoi(controlPlaneCountResult)
	workerNodeCountRes, _ := strconv.Atoi(workerNodeCountCountResult)

	printData := []utils.BiggerBoxData{
		{
			Argument: "Control Plane Count",
			Value:    controlPlaneCountResult,
		},
		{
			Argument: "Worker Node Count",
			Value:    workerNodeCountCountResult,
		},
	}
	utils.PrintBiggerBox(printData)

	return KindClusterConfig{ControlPlaneCount: controlPlaneCountRes, WorkerNodeCount: workerNodeCountRes}
}

// This function creates the config file required by Kind to create cluster
// and outputs the config file under build directory
func KindConfig(selectedOS string, logger *zap.Logger) {
	// utils.ClearScreen()
	logger.Info("Setting up the kind cluster")
	var inputVars = getInputs(logger)
	var KindConfigFile bytes.Buffer
	kindTemplate := template.New("Kind Config File").Funcs(sprig.FuncMap())
	kindTemplate, err := kindTemplate.Parse(kind.Kind)
	if err != nil {
		logger.Fatal("Error during parsing the kind config file", zap.Any("ERROR", err))
	}
	err = kindTemplate.Execute(&KindConfigFile, inputVars)
	if err != nil {
		logger.Fatal("Error during executing the tpl file with vars", zap.Any("ERROR", err))
	}

	os.Mkdir("build", os.ModePerm)
	file, err := os.Create("build/kindconfig.yaml")
	if err != nil {
		logger.Fatal("Failed while creating the build/kindconfig.yaml file", zap.Any("ERROR", err))
	}
	defer file.Close()
	file.WriteString(KindConfigFile.String())
	if utils.CheckCluster(logger) {
		logger.Fatal("Exiting...")
	} else {
		utils.ExecuteCmd([]string{"bash", "-c", "kind create cluster --config build/kindconfig.yaml --name bevelcluster"}, logger)
	}
}
