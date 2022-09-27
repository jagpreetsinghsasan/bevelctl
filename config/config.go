package config

import (
	"bevelctl/support"

	"go.uber.org/zap"
)

func CreateNetworkConfig(environment string, platform string, selectedOS string, logger *zap.Logger) string {
	if environment == support.SupportedEnvironments[0] {
		return CreateDevModeNetworkConfig(platform, selectedOS, logger)
	} else {
		return "Production ready features soon!"
	}
}
