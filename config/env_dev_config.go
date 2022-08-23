package config

import "bevelctl/support"

func CreateDevModeNetworkConfig(platform string) string {
	if platform == support.SupportedPlatforms[0] {
		return FabricNetworkConfig()
	}
	return CordaNetworkConfig()
}
