package config

import (
	"bevelctl/config/corda"
	"bevelctl/config/fabric"
	"bevelctl/support"

	"go.uber.org/zap"
)

// This function generates the network.yaml file when the selected mode == dev
// TODO: Corda dev mode support
func CreateDevModeNetworkConfig(platform string, selectedOS string, logger *zap.Logger) string {
	if platform == support.SupportedPlatforms[0] {
		return fabric.DevFabricNetworkConfig(platform, selectedOS, logger)
	}else{
		return corda.ProdCordaNetworkConfig()
	}
}
