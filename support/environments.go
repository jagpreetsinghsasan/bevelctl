package support

import (
	"github.com/manifoldco/promptui"
	"go.uber.org/zap"
)

// String array listing the supported environments
var SupportedEnvironments = []string{"Dev mode", "Production mode", "Option: Exit"}


// Helper function to select one of the supported environments
func EnvironmentSelect(logger *zap.Logger) string {
	envSelect := promptui.Select{
		Label: "Please select the required environment",
		Items: SupportedEnvironments,
	}
	_, envSelectResult, err := envSelect.Run()

	if err != nil {
		logger.Fatal("Prompt failed while selecting the environment", zap.Any("ERROR", err))
	}

	return envSelectResult
}
