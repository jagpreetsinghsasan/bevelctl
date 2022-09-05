package utils

import (
	"time"

	"github.com/inancgumus/screen"
)

func ClearScreen() {
	screen.Clear()
	screen.MoveTopLeft()
	time.Sleep(500 * time.Millisecond)
}
