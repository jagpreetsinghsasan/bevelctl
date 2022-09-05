package vault

import (
	"bevelctl/support"
	"bevelctl/utils"
	"fmt"
	"os"
	"path/filepath"
)

func SetupVault() {
	osSelectResult := support.SelectOS()
	if osSelectResult == support.SupportedOS[0] {
		utils.ExecuteCmd([]string{"bash", "-c", "sudo apt update && sudo apt install gpg"})
		// executeCmd(exec.Command("bash", "-c", "wget -O- https://apt.releases.hashicorp.com/gpg | gpg --dearmor | sudo tee /usr/share/keyrings/hashicorp-archive-keyring.gpg >/dev/null"))
		// executeCmd(exec.Command("bash", "-c", "echo \"deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main\" | sudo tee /etc/apt/sources.list.d/hashicorp.list"))
		// executeCmd(exec.Command("bash", "-c", "sudo apt update && sudo apt install vault"))
		// executeCmd(exec.Command("bash", "-c", "vault server -dev -dev-root-token-id=\"root\" > vault_output.txt 2>&1 &"))
		// unsealKey := executeCmd(exec.Command("bash", "-c", "cat vault_output.txt | grep Unseal | cut -d \":\" -f 2 | cut -d \" \" -f 2 > unsealKey"))
		// fmt.Println(unsealKey)
		unsealKey := ""
		for unsealKey == "" {
			vaultOutputAbsPath, _ := filepath.Abs("vault_output.txt")
			data, err := os.ReadFile(vaultOutputAbsPath)
			if err != nil {
				fmt.Printf("Cannot read the unseal key. Please setup the vault manually. %v", data)
			}
		}
	} else {
		fmt.Println("Unsupported OS")
	}
}
