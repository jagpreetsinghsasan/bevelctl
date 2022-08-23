package config

import (
	"bevelctl/tpls"
	"bytes"
	"html/template"
	"log"
	"github.com/Masterminds/sprig/v3"
)

func FabricNetworkConfig() string {
	var Count int = 2
	var FabricConfigFile bytes.Buffer
	fabricTemplate := template.New("Fabric Template").Funcs(sprig.FuncMap())
	fabricTemplate, err := fabricTemplate.Parse(tpls.Fabric)
	if err != nil {
		log.Fatal("Parse: ", err)
	}
	err = fabricTemplate.Execute(&FabricConfigFile, Count+1)
	if err != nil {
		log.Fatal("Parse: ", err)
	}
	return FabricConfigFile.String()
}
