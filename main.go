package main

import (
	"bevelctl/config"
	"bevelctl/docker"
	"bevelctl/k8ind"
	"bevelctl/support"
	"fmt"
	"os"
)

func main() {

	fmt.Println("--------------------------------------------")
	fmt.Println("             WELCOME TO BevelCtl            ")
	fmt.Println("              Cli-fying Bevel               ")
	fmt.Println("--------------------------------------------")

	for {
		environment := support.EnvironmentSelect()
		if environment == support.SupportedEnvironments[len(support.SupportedEnvironments)-1] {
			os.Exit(0)
		}
		platform := support.PlatformSelect()
		if platform != support.SupportedPlatforms[len(support.SupportedPlatforms)-1] {
			networkyaml := config.CreateNetworkConfig(environment, platform)
			fmt.Println(networkyaml)
			docker.InstallDocker()
			k8ind.SetupKind()
			k8ind.KindConfig()
			break
		}
	}

}
