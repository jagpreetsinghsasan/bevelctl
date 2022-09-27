package vault

import (
	"bevelctl/support"
	"bevelctl/utils"
	"os"

	"go.uber.org/zap"
)

func setupHelm(selectedOS string, logger *zap.Logger) {
	utils.PrintBox("kubectl", "Installing...")
	if selectedOS == support.SupportedOS[0] {
		utils.ExecuteCmd([]string{"bash", "-c", "sudo snap install helm --classic"}, logger)
		utils.PrintBox("helm", "Installation complete...")
	} else {
		utils.PrintBox("helm", "Skipped...")
	}
}

func SetupVault(selectedOS string, logger *zap.Logger) {
	setupHelm(selectedOS, logger)
	if selectedOS == support.SupportedOS[0] {
		oldVaultFullLabel := "app.kubernetes.io/name=vault"
		newVaultLabelKey := "bevelabel"
		newVaultLabelValue := "bevelvault"
		utils.ExecuteCmd([]string{"bash", "-c", "helm repo add hashicorp https://helm.releases.hashicorp.com"}, logger)
		utils.ExecuteCmd([]string{"bash", "-c", "helm install vault hashicorp/vault --version 0.13.0"}, logger)
		kubeClient := utils.GetKubeClient(os.Getenv("HOME")+"/.kube/config", logger)
		utils.WaitForPodToRun(kubeClient, "default", oldVaultFullLabel, logger)
		utils.AddLabelToARunningPod(kubeClient, "default", "app.kubernetes.io/name=vault", newVaultLabelKey, newVaultLabelValue, logger)

	} else {
		logger.Fatal("Unsupported OS")
	}
}
