package utils

import (
	"bevelctl/tpls"
	"bytes"
	"html/template"
	"log"

	"github.com/Masterminds/sprig/v3"
)

func CreateNetworkConfig(mode string, platform string) string {
	if mode == "dev" {
		return CreateDevModeNetworkConfig(platform)
	}
	return "no mode detected"
}

func CreateDevModeNetworkConfig(platform string) string {
	if platform == "fabric" {
		return FabricNetworkConfig()
	}
	return CordaNetworkConfig()
}

func CreateProdModeNetworkConfig() {

}

func FabricNetworkConfig() string {
	var Count int = 2
	var FabricConfigFile bytes.Buffer
	t := template.New("Fabric Template").Funcs(sprig.FuncMap())
	t, err := t.Parse(tpls.Fabric)
	if err != nil {
		log.Fatal("Parse: ", err)
	}
	err = t.Execute(&FabricConfigFile, Count+1)
	if err != nil {
		log.Fatal("Parse: ", err)
	}
	return FabricConfigFile.String()
}

func CordaNetworkConfig() string {
	return "Corda"
}
