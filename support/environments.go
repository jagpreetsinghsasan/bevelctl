package support

import (
	"github.com/manifoldco/promptui"
	"fmt"
	"os"
)

var SupportedEnvironments = []string{"Dev mode", "Production mode", "Option: Exit"}

func EnvironmentSelect() string {
	envSelect := promptui.Select{
		Label: "Please select the required environment",
		Items: SupportedEnvironments,
	}
	_, envSelectResult, err := envSelect.Run()

	if err != nil {
		fmt.Printf("Prompt failed while selecting environment %v\n", err)
		os.Exit(1)
	}

	return envSelectResult
}