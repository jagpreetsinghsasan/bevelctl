package support

import (
	"github.com/manifoldco/promptui"
	"go.uber.org/zap"
)

// String array including all the supported platforms
var SupportedPlatforms = []string{"fabric", "corda", "Option: Go Back to the Main Menu"}

// Helper function to select one of the supported platforms
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
