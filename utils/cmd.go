package utils

import (
	"bytes"
	"io"
	"os"
	"os/exec"

	"go.uber.org/zap"
)

// Utility function to execute the mentioned command 
// and return the command output as a string
func ExecuteCmd(cmdWithArgs []string, logger *zap.Logger) string {
	cmd := exec.Command(cmdWithArgs[0], cmdWithArgs[1:]...)
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Run()
	if err != nil {
		logger.Fatal("Command run failed", zap.Any("ERROR", err))
	}
	outStr := string(stdoutBuf.Bytes())
	return outStr
}
