package utils

import (
	"time"
	"github.com/da0x/golang/olog"
)

// Helper struct for printing the boxed data with predefined column name
type BoxData struct {
	Binary string
	Status string
}

// Helper struct to print the custom box data and custom column description
type BiggerBoxData struct {
	Argument string
	Value    string
}

// Utility function to print a key,value sort of data in a boxed design
// as Binary: < > , Status: < >
func PrintBox(binaryName string, currentStatus string) {
	time.Sleep(1 * time.Second)
	var boxData = []BoxData{
		{Binary: binaryName, Status: currentStatus},
	}
	olog.PrintHStrong(boxData)
}

// Utility function to print a boxed data with custom key,value pairs and headings
func PrintBiggerBox(biggerBoxData []BiggerBoxData) {
	time.Sleep(1 * time.Second)
	olog.PrintHStrong(biggerBoxData)
}
