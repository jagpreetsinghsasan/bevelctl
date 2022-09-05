package support

import (
	"github.com/manifoldco/promptui"
	"fmt"
	"os"
)

var SupportedPlatforms = []string{"Hyperledger Fabric", "R3 Corda", "Option: Go Back to the Main Menu"}

func PlatformSelect() string {
	platSelect := promptui.Select{
		Label: "Please select the required platform",
		Items: SupportedPlatforms,
	}

	_, platSelectResult, err := platSelect.Run()

	if err != nil {
		fmt.Printf("Prompt failed while selecting platform %v\n", err)
		os.Exit(1)
	}

	return platSelectResult
}
