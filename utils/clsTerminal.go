package utils

import (
	"time"

	"github.com/inancgumus/screen"
)

// Utility function to clear the stdout
func ClearScreen() {
	screen.Clear()
	screen.MoveTopLeft()
	time.Sleep(500 * time.Millisecond)
}
