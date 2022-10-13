package utils

import (
	"bytes"
	"os/exec"
	"strings"

	"go.uber.org/zap"
)

// Function to check if the bevelcluster is already present or not
// If present, the user is requested to back up any important data in the bevelcluster and delete it using
// kind delete clusters bevelcluster
func CheckCluster(logger *zap.Logger) bool {
	logger.Info("Checking if the bevelcluster exists or not")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	bashCmd := exec.Command("bash", "-c", "kind get clusters")
	bashCmd.Stdout = &stdout
	bashCmd.Stderr = &stderr
	bashCmd.Run()
	clusterStatus := strings.Contains(stdout.String()+stderr.String(), "bevelcluster")
	if clusterStatus {
		logger.Info("bevelcluster already present. Please delete this cluster (and back up anything important in it if required) to proceed.")
		logger.Info("To delete the cluster, use this command:")
		logger.Info("kind delete clusters bevelcluster")
		logger.Info("Once deleted, you can re-run the cli")
		return true
	} else {
		return false
	}
}
