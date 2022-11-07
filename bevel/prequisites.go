package bevel

import (
	"bevelctl/docker"
	"bevelctl/k8ind"
	"bevelctl/vault"
	"time"

	"github.com/briandowns/spinner"
	"go.uber.org/zap"
)

func ExecuteBevel(bevelctlInputs interface{}, logger *zap.Logger) {
	spinnerSet := []string{"◤◢", "◤ ◢", "◤   ◢", "◤ ◢", "◤◢"}
	s := spinner.New(spinnerSet, 100*time.Millisecond)
	s.Color("magenta")
	s.Start()
	// networkyaml := config.CreateNetworkConfig(bevelctlInputs.environment, bevelctlInputs.platform, logger)
	// fmt.Println(networkyaml)
	docker.InstallDocker(logger)
	k8ind.SetupKind(logger)
	k8ind.KindConfig(logger)
	vault.SetupVault(logger)
	s.Stop()
}

func CheckBinaryStatus(logger *zap.Logger) {
	
}
