package main

import (
	"bevelctl/config"
	"bevelctl/docker"
	"bevelctl/k8ind"
	"bevelctl/support"
	"bevelctl/vault"
	"fmt"
	"os"

	"go.uber.org/zap"
)

func main() {

	fmt.Println("--------------------------------------------")
	fmt.Println("             WELCOME TO BevelCtl            ")
	fmt.Println("              Cli-fying Bevel               ")
	fmt.Println("--------------------------------------------")

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	for {
		environment := support.EnvironmentSelect(logger)
		if environment == support.SupportedEnvironments[len(support.SupportedEnvironments)-1] {
			os.Exit(0)
		}
		selectedOS := support.SelectOS(logger)
		if selectedOS == support.SupportedOS[len(support.SupportedOS)-1] {
			os.Exit(0)
		}
		platform := support.PlatformSelect(logger)
		if platform != support.SupportedPlatforms[len(support.SupportedPlatforms)-1] {
			networkyaml := config.CreateNetworkConfig(environment, platform, selectedOS, logger)
			fmt.Println(networkyaml)
			docker.InstallDocker(selectedOS, logger)
			k8ind.SetupKind(selectedOS, logger)
			k8ind.KindConfig(selectedOS, logger)
			vault.SetupVault(selectedOS, logger)
			break
		}
	}

}
