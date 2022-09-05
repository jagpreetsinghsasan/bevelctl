package utils

import (
	"time"
	"github.com/da0x/golang/olog"
)

type BoxData struct {
	Binary string
	Status string
}

type BiggerBoxData struct {
	Argument string
	Value    string
}

func PrintBox(binaryName string, currentStatus string) {
	time.Sleep(1 * time.Second)
	var boxData = []BoxData{
		{Binary: binaryName, Status: currentStatus},
	}
	olog.PrintHStrong(boxData)
}

func PrintBiggerBox(biggerBoxData []BiggerBoxData) {
	time.Sleep(1 * time.Second)
	olog.PrintHStrong(biggerBoxData)
}
