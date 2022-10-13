package utils

import (
	"os/exec"

	"go.uber.org/zap"
)

// Function to check if the mentioned binary is present or not
func CheckBinary(binaryChkCmd string, logger *zap.Logger) bool {
	logger.Info("Checking if the binary is present or not")
	_, err := exec.LookPath(binaryChkCmd)
	if err != nil {
		logger.Info("Binary not found")
		return true
	} else {
		logger.Info("Binary found. Skipping its installation.")
		return false
	}
}
