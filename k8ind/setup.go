package k8ind

import (
	"bevelctl/utils"

	"go.uber.org/zap"
)

// This function sets up kubectl binary required to perform kubectl commands
// TODO: Remove the dependency of this binary completely and use k8s go-client code
func setupKubectl(logger *zap.Logger) {
	// utils.ClearScreen()
	logger.Info("Setting up kubectl")
	if utils.CheckBinary("kubectl", logger) {
		utils.ExecuteCmd([]string{"bash", "-c", "sudo snap install kubectl --classic"}, logger)
		logger.Info("kubectl installed")
	} else {
		logger.Info("Installation of kubectl skipped")
	}
}

// This function sets up kind - kubernetes in docker
// TODO: Remove sudo from the involved commands
func SetupKind(logger *zap.Logger) {
	setupKubectl(logger)
	// utils.ClearScreen()
	logger.Info("Setting up kind")
	if utils.CheckBinary("kind", logger) {
		utils.ExecuteCmd([]string{"bash", "-c", "curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.14.0/kind-linux-amd64"}, logger)
		utils.ExecuteCmd([]string{"bash", "-c", "chmod +x ./kind"}, logger)
		utils.ExecuteCmd([]string{"bash", "-c", "sudo mv ./kind /usr/local/bin/kind"}, logger)
		logger.Info("Kind installed")
	} else {
		logger.Info("Installation of kind skipped")
	}
}
