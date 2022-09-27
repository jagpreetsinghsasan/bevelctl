package support

import (
	"github.com/manifoldco/promptui"
	"go.uber.org/zap"
)

var SupportedPlatforms = []string{"Hyperledger Fabric", "R3 Corda", "Option: Go Back to the Main Menu"}

func PlatformSelect(logger *zap.Logger) string {
	platSelect := promptui.Select{
		Label: "Please select the required platform",
		Items: SupportedPlatforms,
	}

	_, platSelectResult, err := platSelect.Run()

	if err != nil {
		logger.Fatal("Prompt failed while selecting the platform", zap.Any("ERROR", err))
	}

	return platSelectResult
}
