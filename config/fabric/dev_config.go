package fabric

import (
	"bevelctl/tpls/fabric"
	"bytes"
	"html/template"
	"strconv"

	"github.com/Masterminds/sprig/v3"
	"github.com/manifoldco/promptui"
	"go.uber.org/zap"
)

// This struct includes the user input variables with default values as
// orderer count as 1 and organization count as 1
type DevFabric struct {
	OrdererCount int
	OrgCount     int
}

// This function takes the DevFabric struct input from the user and returns the constructed struct
// to the caller function
func getInputs(logger *zap.Logger) DevFabric {

	ordererCount := promptui.Prompt{
		Label:   "Enter the orderer count: ",
		Default: "1",
	}
	ordererCountResult, err := ordererCount.Run()
	if err != nil {
		logger.Fatal("Prompt failed entering the orderer count", zap.Any("ERROR", err))
	}

	orgCount := promptui.Prompt{
		Label:   "Enter the organization count: ",
		Default: "1",
	}
	orgCountResult, err := orgCount.Run()
	if err != nil {
		logger.Fatal("Prompt failed entering the organization count", zap.Any("ERROR", err))
	}

	ordererCountRes, _ := strconv.Atoi(ordererCountResult)
	orgCountRes, _ := strconv.Atoi(orgCountResult)

	return DevFabric{OrdererCount: ordererCountRes, OrgCount: orgCountRes}
}

// This function constructs the network.yaml file customized as per the DevFabric struct
// and returns the network.yaml as a string
// TODO: Fix this code to include auto usage of the network.yaml from Bevel repository in absence of custom input
func DevFabricNetworkConfig(platform string, selectedOS string, logger *zap.Logger) string {
	var inputVars = getInputs(logger)
	var FabricConfigFile bytes.Buffer
	fabricTemplate := template.New("Dev Fabric Template").Funcs(sprig.FuncMap())
	fabricTemplate, err := fabricTemplate.Parse(fabric.Fabric)
	if err != nil {
		logger.Fatal("Failed to parse the network.yaml file", zap.Any("ERROR", err))
	}
	err = fabricTemplate.Execute(&FabricConfigFile, inputVars)
	if err != nil {
		logger.Fatal("Failed to execute the operation", zap.Any("ERROR", err))
	}
	return FabricConfigFile.String()
}
