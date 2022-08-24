package main

import (
	"bevelctl/config"
	"bevelctl/support"
	"fmt"
	"os"
	"github.com/manifoldco/promptui"
)

func main() {

	fmt.Println("--------------------------------------------")
	fmt.Println("             WELCOME TO BevelCtl            ")
	fmt.Println("              Cli-fying Bevel               ")
	fmt.Println("--------------------------------------------")

	for {
		environment := environmentSelect()
		if environment == support.SupportedEnvironments[len(support.SupportedEnvironments)-1] {
			os.Exit(0)
		}
		platform := platformSelect()
		if platform != support.SupportedPlatforms[len(support.SupportedPlatforms)-1] {
			networkyaml := config.CreateNetworkConfig(environment, platform)
			fmt.Println(networkyaml)
		}
	}

}

func environmentSelect() string {
	envSelect := promptui.Select{
		Label: "Please select the required environment",
		Items: support.SupportedEnvironments,
	}
	_, envSelectResult, err := envSelect.Run()

	if err != nil {
		fmt.Printf("Prompt failed while selecting environment %v\n", err)
		os.Exit(1)
	}

	return envSelectResult
}

func platformSelect() string {
	platSelect := promptui.Select{
		Label: "Please select the required platform",
		Items: support.SupportedPlatforms,
	}

	_, platSelectResult, err := platSelect.Run()

	if err != nil {
		fmt.Printf("Prompt failed while selecting platform %v\n", err)
		os.Exit(1)
	}

	return platSelectResult
}
