package docker

import (
	"bevelctl/utils"

	"go.uber.org/zap"
)

// This function setups snap which is required for docker to work
// TODO: removal of sudo
func setupSnap(logger *zap.Logger) {
	logger.Info("Setting up snap")
	if utils.CheckBinary("snap", logger) {
		logger.Info("Installing snap app store for Linux")
		utils.ExecuteCmd([]string{"bash", "-c", "sudo apt update && sudo apt install snapd"}, logger)
		logger.Info("Snap installed")
	} else {
		logger.Info("Installation of snap skipped")
	}

}

// This function setups docker on the machine for the selected OS
// TODO: removal of sudo
func InstallDocker(logger *zap.Logger) {
	// utils.ClearScreen()
	logger.Info("Setting up docker")
	if utils.CheckBinary("docker", logger) {
		setupSnap(logger)
		logger.Info("Installing docker using snap")
		utils.ExecuteCmd([]string{"bash", "-c", "sudo snap install docker"}, logger)
		logger.Info("Docker installed")
	} else {
		logger.Info("Installation of docker skipped")
		
	}
}
