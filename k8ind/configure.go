package k8ind

import (
	"bevelctl/tpls/kind"
	"bevelctl/utils"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"strconv"

	"github.com/Masterminds/sprig/v3"
	"github.com/manifoldco/promptui"
)

type KindClusterConfig struct {
	ControlPlaneCount int
	WorkerNodeCount   int
}

func getInputs() KindClusterConfig {
	controlPlaneCount := promptui.Prompt{
		Label:   "Enter the number of control plane nodes: ",
		Default: "1",
	}
	controlPlaneCountResult, err := controlPlaneCount.Run()
	if err != nil {
		fmt.Printf("Prompt failed entering the control plane node count %v\n", err)
		os.Exit(1)
	}

	workerNodeCount := promptui.Prompt{
		Label:   "Enter the number of worker nodes: ",
		Default: "3",
	}
	workerNodeCountCountResult, err := workerNodeCount.Run()
	if err != nil {
		fmt.Printf("Prompt failed entering the worker node count %v\n", err)
		os.Exit(1)
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

func KindConfig() {
	utils.ClearScreen()
	fmt.Println("Setting up the kind cluster")
	var inputVars = getInputs()
	var KindConfigFile bytes.Buffer
	kindTemplate := template.New("Kind Config File").Funcs(sprig.FuncMap())
	kindTemplate, err := kindTemplate.Parse(kind.Kind)
	if err != nil {
		log.Fatal("Parse: ", err)
	}
	err = kindTemplate.Execute(&KindConfigFile, inputVars)
	if err != nil {
		log.Fatal("Execute: ", err)
	}

	os.Mkdir("build", os.ModePerm)
	file, err := os.Create("build/kindconfig.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(KindConfigFile.String())
	if(utils.CheckCluster()){
		os.Exit(1)
	}else{
		utils.ExecuteCmd([]string{"bash", "-c", "kind create cluster --config build/kindconfig.yaml --name bevelcluster"})
	}
}
