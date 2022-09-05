package support

import (
	"github.com/manifoldco/promptui"
	"fmt"
	"os"
)

var SupportedOS = []string{"Ubuntu or Debian", "None of the above (More OS to be supported in the future)"}

func SelectOS() string{
	osSelect := promptui.Select{
		Label: "Please select the machine environment",
		Items: SupportedOS,
	}
	_, osSelectResult, err := osSelect.Run()

	if err != nil {
		fmt.Printf("Prompt failed while selecting OS %v\n", err)
		os.Exit(1)
	}
	return osSelectResult
}