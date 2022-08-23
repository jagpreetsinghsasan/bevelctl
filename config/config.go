package config

import (
	"bevelctl/support"
)

func CreateNetworkConfig(environment string, platform string) string {
	if environment == support.SupportedEnvironments[0] {
		return CreateDevModeNetworkConfig(platform)
	} else {
		return "Production ready nahi hai abhi"
	}
}
