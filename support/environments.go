package support

import (
	"github.com/manifoldco/promptui"
	"go.uber.org/zap"
)

var SupportedEnvironments = []string{"Dev mode", "Production mode", "Option: Exit"}

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
