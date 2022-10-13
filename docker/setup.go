package docker

import (
	"bevelctl/support"
	"bevelctl/utils"

	"go.uber.org/zap"
)

// This function setups snap which is required for docker to work
// TODO: removal of sudo
func setupSnap(selectedOS string, logger *zap.Logger) {
	utils.PrintBox("Snap", "Installing...")
	if selectedOS == support.SupportedOS[0] && utils.CheckBinary("snap", logger) {
		logger.Info("Installing snap app store for Linux")
		utils.ExecuteCmd([]string{"bash", "-c", "sudo apt update && sudo apt install snapd"}, logger)
		utils.PrintBox("Snap", "Installation complete...")
	} else {
		utils.PrintBox("Snap", "Skipped...")
	}

}

// This function setups docker on the machine for the selected OS
// TODO: removal of sudo
func InstallDocker(selectedOS string, logger *zap.Logger) {
	// utils.ClearScreen()
	utils.PrintBox("Docker", "Installing...")
	if selectedOS == support.SupportedOS[0] && utils.CheckBinary("docker", logger) {
		setupSnap(selectedOS, logger)
		logger.Info("Installing docker using snap")
		utils.ExecuteCmd([]string{"bash", "-c", "sudo snap install docker"}, logger)
		utils.PrintBox("Docker", "Installation complete...")
	} else {
		utils.PrintBox("Docker", "Skipped...")
	}
}
