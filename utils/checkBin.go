package utils

import (
	"fmt"
	"os/exec"
)

func CheckBinary(binaryChkCmd string) bool {
	fmt.Println("Checking if the binary is present or not")
	_, err := exec.LookPath(binaryChkCmd)
	if err != nil {
		fmt.Println("Binary not found")
		return true
	} else {
		fmt.Println("Binary found. Skipping installation of the same.")
		return false
	}
}
