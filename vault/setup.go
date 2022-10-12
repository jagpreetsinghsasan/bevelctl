package vault

import (
	"bevelctl/support"
	"bevelctl/utils"
	"fmt"
	"os"
	"strings"

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

func getUnsealKey(vaultConfig string) string {
	unsealkeyString := strings.Split(vaultConfig, "\n")[0]
	unsealKeyValue := strings.Split(unsealkeyString, ": ")[1]
	unsealKeyValue = strings.ReplaceAll(unsealKeyValue, " ", "")
	return unsealKeyValue
}

func getInitalRootToken(vaultConfig string) string {
	initialRootTokenString := strings.Split(vaultConfig, "\n")[2]
	initialRootTokenValue := strings.Split(initialRootTokenString, ":")[1]
	initialRootTokenValue = strings.ReplaceAll(initialRootTokenValue, " ", "")
	return initialRootTokenValue
}

func storeVaultCredsInFile(vaultConfig string, logger *zap.Logger) {
	rootToken := getInitalRootToken(vaultConfig)
	unsealKey := getUnsealKey(vaultConfig)

	os.Mkdir("build", os.ModePerm)
	file, err := os.Create("build/vaultconfig.json")
	if err != nil {
		logger.Fatal("Failed while creating the build/vaultconfig.json file", zap.Any("ERROR", err))
	}
	defer file.Close()
	fileData := fmt.Sprintf("{\n  \"unseal_key\" : \"%s\",\n  \"root_token\" : \"%s\"\n}", unsealKey, rootToken)
	file.WriteString(fileData)
}

func SetupVault(selectedOS string, logger *zap.Logger) {
	setupHelm(selectedOS, logger)
	if selectedOS == support.SupportedOS[0] {
		oldVaultFullLabel := "app.kubernetes.io/name=vault"
		newVaultLabelKey := "bevelabel"
		newVaultLabelValue := "bevelvault"
		clusterContext := "kind-bevelcluster"
		utils.ExecuteCmd([]string{"bash", "-c", "helm repo add hashicorp https://helm.releases.hashicorp.com"}, logger)
		utils.ExecuteCmd([]string{"bash", "-c", "helm install vault hashicorp/vault --version 0.13.0"}, logger)
		kubeClient := utils.GetKubeClient(os.Getenv("HOME")+"/.kube/config", clusterContext, logger)
		utils.WaitForPodToRun(kubeClient, "default", oldVaultFullLabel, logger)
		utils.AddLabelToARunningPod(kubeClient, "default", "app.kubernetes.io/name=vault", newVaultLabelKey, newVaultLabelValue, logger)
		restConfig := utils.GetK8sRestConfig(os.Getenv("HOME")+"/.kube/config", "kind-bevelcluster", logger)
		vaultConfigString := utils.KubectlExecCmd(restConfig, "vault-0", "vault", "default", "vault operator init -key-shares=1 -key-threshold=1 -format=table", logger)
		storeVaultCredsInFile(vaultConfigString, logger)

	} else {
		logger.Fatal("Unsupported OS")
	}
}
