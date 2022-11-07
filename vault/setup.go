package vault

import (
	"bevelctl/utils"
	"fmt"
	"os"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

// Function to setup helm binary on the machine
func setupHelm(logger *zap.Logger) {
	logger.Info("Setting up Helm")
	if utils.CheckBinary("helm", logger) {
		logger.Info("Installing helm using snap")
		utils.ExecuteCmd([]string{"bash", "-c", "sudo snap install helm --classic"}, logger)
		logger.Info("Helm installed")
	} else {
		logger.Info("Installation of helm skipped")
	}
}

// Function to get the formatted vault unseal key string
func getUnsealKey(vaultConfig string) string {
	unsealkeyString := strings.Split(vaultConfig, "\n")[0]
	unsealKeyValue := strings.Split(unsealkeyString, ": ")[1]
	unsealKeyValue = strings.ReplaceAll(unsealKeyValue, " ", "")
	return unsealKeyValue
}

// Function to get the formatted vault inital root token string
func getInitalRootToken(vaultConfig string) string {
	initialRootTokenString := strings.Split(vaultConfig, "\n")[2]
	initialRootTokenValue := strings.Split(initialRootTokenString, ":")[1]
	initialRootTokenValue = strings.ReplaceAll(initialRootTokenValue, " ", "")
	return initialRootTokenValue
}

// Function to store the vault initial root token and unseal key in a json file
// for user reference
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

// Function to setup vault, initialize and unseal it
func SetupVault(logger *zap.Logger) {
	setupHelm(logger)
	oldVaultFullLabel := "app.kubernetes.io/name=vault"
	newVaultLabelKey := "bevelabel"
	newVaultLabelValue := "bevelvault"
	clusterContext := "kind-bevelcluster"
	CreateVaultConfig(logger)
	utils.ExecuteCmd([]string{"bash", "-c", "helm repo add hashicorp https://helm.releases.hashicorp.com"}, logger)
	utils.ExecuteCmd([]string{"bash", "-c", "helm install vault hashicorp/vault --version 0.13.0 -f build/vaultconfig.yaml"}, logger)
	kubeClient := utils.GetKubeClient(os.Getenv("HOME")+"/.kube/config", clusterContext, logger)
	utils.WaitForPodToRun(kubeClient, "default", oldVaultFullLabel, logger)
	utils.AddLabelToARunningPod(kubeClient, "default", "app.kubernetes.io/name=vault", newVaultLabelKey, newVaultLabelValue, logger)
	restConfig := utils.GetK8sRestConfig(os.Getenv("HOME")+"/.kube/config", "kind-bevelcluster", logger)
	vaultConfigString := utils.KubectlExecCmd(restConfig, "vault-0", "vault", "default", "vault operator init -key-shares=1 -key-threshold=1 -format=table", logger)
	storeVaultCredsInFile(vaultConfigString, logger)
	vaultEnvVarsString := `export VAULT_ADDR=http://` + utils.GetK8sNodeIP(kubeClient, logger)[0] + `:` + strconv.FormatInt(int64(utils.GetK8sServicePort(kubeClient, "default", "vault-ui", logger)[0].NodePort), 10) + `; export VAULT_TOKEN=` + getInitalRootToken(vaultConfigString) + `; `
	unsealVaultCmdString := vaultEnvVarsString + `vault operator unseal ` + getUnsealKey(vaultConfigString)
	utils.ExecuteCmd([]string{"bash", "-c", unsealVaultCmdString}, logger)
	enableSecretsEngineCmdString := vaultEnvVarsString + `vault secrets enable -version=2 -path=secret kv`
	utils.ExecuteCmd([]string{"bash", "-c", enableSecretsEngineCmdString}, logger)
}
