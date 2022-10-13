package utils

import (
	"fmt"
	"runtime"
)

// Utility function which can be called by any function to get the caller function name
func GetFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
}
