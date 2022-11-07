package utils

import (
	"bytes"
	"os"
	"os/exec"

	"go.uber.org/zap"
)

// Utility function to execute the mentioned command
// and return the command output as a string
func ExecuteCmd(cmdWithArgs []string, logger *zap.Logger) string {
	cmd := exec.Command(cmdWithArgs[0], cmdWithArgs[1:]...)
	var stdoutBuf bytes.Buffer

	fileName := "log/cmd.txt"
	logFile, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	cmd.Stdout = logFile
	cmd.Stderr = logFile
	err := cmd.Run()
	if err != nil {
		logger.Fatal("Command run failed", zap.Any("ERROR", err))
	}
	outStr := string(stdoutBuf.Bytes())
	return outStr
}
