package k8ind

import (
	"bevelctl/tpls/kind"
	"bevelctl/utils"
	"bytes"
	"html/template"
	"os"

	"github.com/Masterminds/sprig/v3"
	"go.uber.org/zap"
)


// This struct is used to take user inputs where
// ControlPlaneCount refers to the k8s master node count
// WorkerNodeCount referes to the k8s worker node count
type KindClusterConfig struct {
	ControlPlaneCount int
	WorkerNodeCount   int
}

// This function creates the config file required by Kind to create cluster
// and outputs the config file under build directory
func KindConfig(logger *zap.Logger) {
	// utils.ClearScreen()
	logger.Info("Setting up the kind cluster")
	var KindConfigFile bytes.Buffer
	kindTemplate := template.New("Kind Config File").Funcs(sprig.FuncMap())
	kindTemplate, err := kindTemplate.Parse(kind.Kind)
	if err != nil {
		logger.Fatal("Error during parsing the kind config file", zap.Any("ERROR", err))
	}
	err = kindTemplate.Execute(&KindConfigFile, KindClusterConfig{ControlPlaneCount: 1, WorkerNodeCount: 3})
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
