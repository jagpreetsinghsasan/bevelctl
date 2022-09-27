package support

import (
	"github.com/manifoldco/promptui"
	"go.uber.org/zap"
)

var SupportedOS = []string{"Ubuntu or Debian", "None of the above (More OS to be supported in the future)"}

func SelectOS(logger *zap.Logger) string {
	osSelect := promptui.Select{
		Label: "Please select the machine environment",
		Items: SupportedOS,
	}
	_, osSelectResult, err := osSelect.Run()

	if err != nil {
		logger.Fatal("Prompt failed while selecting the OS", zap.Any("ERROR", err))
	}
	return osSelectResult
}
