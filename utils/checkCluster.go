package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func CheckCluster() bool {
	fmt.Println("Checking if the bevelcluster exists or not")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	bashCmd := exec.Command("bash", "-c", "kind get clusters")
	bashCmd.Stdout = &stdout
	bashCmd.Stderr = &stderr
	bashCmd.Run()
	clusterStatus := strings.Contains(stdout.String()+stderr.String(), "bevelcluster")
	if clusterStatus {
		fmt.Println()
		fmt.Println("bevelcluster already present. Please delete this cluster to proceed using the command:")
		fmt.Println("kind delete clusters bevelcluster")
		fmt.Println("Once done, you can start over the deployment process")
		fmt.Println()
		return true
	} else {
		return false
	}
}
