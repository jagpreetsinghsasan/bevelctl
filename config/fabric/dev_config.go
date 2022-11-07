package fabric

import (
	"bevelctl/tpls/fabric"
	"bytes"
	"html/template"

	"github.com/Masterminds/sprig/v3"
	"go.uber.org/zap"
)

// This struct includes the user input variables with default values as
// orderer count as 1 and organization count as 1
type DevFabric struct {
	OrdererCount int
	OrgCount     int
}

// This function constructs the network.yaml file customized as per the DevFabric struct
// and returns the network.yaml as a string
// TODO: Fix this code to include auto usage of the network.yaml from Bevel repository in absence of custom input
func DevFabricNetworkConfig(platform string, logger *zap.Logger) string {
	var FabricConfigFile bytes.Buffer
	fabricTemplate := template.New("Dev Fabric Template").Funcs(sprig.FuncMap())
	fabricTemplate, err := fabricTemplate.Parse(fabric.Fabric)
	if err != nil {
		logger.Fatal("Failed to parse the network.yaml file", zap.Any("ERROR", err))
	}
	err = fabricTemplate.Execute(&FabricConfigFile, DevFabric{OrdererCount: 1, OrgCount: 3})
	if err != nil {
		logger.Fatal("Failed to execute the operation", zap.Any("ERROR", err))
	}
	return FabricConfigFile.String()
}
