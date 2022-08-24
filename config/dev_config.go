package config

import (
	"bevelctl/config/fabric"
	"bevelctl/config/corda"
)

import "bevelctl/support"

func CreateDevModeNetworkConfig(platform string) string {
	if platform == support.SupportedPlatforms[0] {
		return fabric.DevFabricNetworkConfig(platform)
	}else{
		return corda.ProdCordaNetworkConfig()
	}
}
