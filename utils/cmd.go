package utils

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
)

func ExecuteCmd(cmdWithArgs []string) string {
	cmd := exec.Command(cmdWithArgs[0], cmdWithArgs[1:]...)
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr := string(stdoutBuf.Bytes())
	return outStr
}
