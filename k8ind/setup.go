package k8ind

import (
	"bevelctl/support"
	"bevelctl/utils"
)

func setupKubectl() {
	utils.ClearScreen()
	utils.PrintBox("kubectl", "Installing...")
	osSelectResult := support.SelectOS()
	if osSelectResult == support.SupportedOS[0] && utils.CheckBinary("kubectl") {
		utils.ExecuteCmd([]string{"bash", "-c", "sudo snap install kubectl --classic"})
		utils.PrintBox("kubectl", "Installation complete...")
	} else {
		utils.PrintBox("kubectl", "Skipped...")
	}
}

func SetupKind() {
	setupKubectl()
	utils.ClearScreen()
	utils.PrintBox("Kind k8s", "Installing...")
	osSelectResult := support.SelectOS()
	if osSelectResult == support.SupportedOS[0] && utils.CheckBinary("kind") {
		utils.ExecuteCmd([]string{"bash", "-c", "curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.14.0/kind-linux-amd64"})
		utils.ExecuteCmd([]string{"bash", "-c", "chmod +x ./kind"})
		utils.ExecuteCmd([]string{"bash", "-c", "sudo mv ./kind /usr/local/bin/kind"})
		utils.PrintBox("Kind k8s", "Installation complete...")
	} else {
		utils.PrintBox("Kind k8s", "Skipped...")
	}
}
