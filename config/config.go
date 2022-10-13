package config

import (
	"bevelctl/support"

	"go.uber.org/zap"
)

// This function creates network.yaml file needed as the single input to the Bevel playbooks
// TODO: Production ready features
func CreateNetworkConfig(environment string, platform string, selectedOS string, logger *zap.Logger) string {
	if environment == support.SupportedEnvironments[0] {
		return CreateDevModeNetworkConfig(platform, selectedOS, logger)
	} else {
		return "Production ready features soon!"
	}
}
