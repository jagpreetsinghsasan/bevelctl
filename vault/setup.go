package vault

import (
	"bevelctl/support"
	"log"
	"os/exec"
	"fmt"
	"os"
	"github.com/manifoldco/promptui"
)


func selectOS() string{
	osSelect := promptui.Select{
		Label: "Please select the machine environment",
		Items: support.SupportedOS,
	}
	_, osSelectResult, err := osSelect.Run()

	if err != nil {
		fmt.Printf("Prompt failed while selecting OS %v\n", err)
		os.Exit(1)
	}
	return osSelectResult
}