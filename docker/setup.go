package docker

import (
	"bevelctl/support"
	"bevelctl/utils"
	"fmt"
)

func setupSnap() {
	utils.PrintBox("Snap", "Installing...")
	osSelectResult := support.SelectOS()
	if osSelectResult == support.SupportedOS[0] && utils.CheckBinary("snap") {
		fmt.Println("Installing snap app store for Linux")
		utils.ExecuteCmd([]string{"bash", "-c", "sudo apt update && sudo apt install snapd"})
		utils.PrintBox("Snap", "Installation complete...")
	} else {
		utils.PrintBox("Snap", "Skipped...")
	}

}

func InstallDocker() {
	utils.ClearScreen()
	utils.PrintBox("Docker", "Installing...")
	osSelectResult := support.SelectOS()
	if osSelectResult == support.SupportedOS[0] && utils.CheckBinary("docker") {
		setupSnap()
		fmt.Println("Installing docker using snap")
		utils.ExecuteCmd([]string{"bash", "-c", "sudo snap install docker"})
		utils.PrintBox("Docker", "Installation complete...")
	} else {
		utils.PrintBox("Docker", "Skipped...")
	}
}
