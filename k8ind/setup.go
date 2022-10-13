package k8ind

import (
	"bevelctl/support"
	"bevelctl/utils"

	"go.uber.org/zap"
)

// This function sets up kubectl binary required to perform kubectl commands
// TODO: Remove the dependency of this binary completely and use k8s go-client code
func setupKubectl(selectedOS string, logger *zap.Logger) {
	// utils.ClearScreen()
	utils.PrintBox("kubectl", "Installing...")
	if selectedOS == support.SupportedOS[0] && utils.CheckBinary("kubectl", logger) {
		utils.ExecuteCmd([]string{"bash", "-c", "sudo snap install kubectl --classic"}, logger)
		utils.PrintBox("kubectl", "Installation complete...")
	} else {
		utils.PrintBox("kubectl", "Skipped...")
	}
}

// This function sets up kind - kubernetes in docker
// TODO: Remove sudo from the involved commands
func SetupKind(selectedOS string, logger *zap.Logger) {
	setupKubectl(selectedOS, logger)
	// utils.ClearScreen()
	utils.PrintBox("Kind k8s", "Installing...")
	if selectedOS == support.SupportedOS[0] && utils.CheckBinary("kind",logger) {
		utils.ExecuteCmd([]string{"bash", "-c", "curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.14.0/kind-linux-amd64"}, logger)
		utils.ExecuteCmd([]string{"bash", "-c", "chmod +x ./kind"}, logger)
		utils.ExecuteCmd([]string{"bash", "-c", "sudo mv ./kind /usr/local/bin/kind"}, logger)
		utils.PrintBox("Kind k8s", "Installation complete...")
	} else {
		utils.PrintBox("Kind k8s", "Skipped...")
	}
}
