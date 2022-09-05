package fabric

import (
	"bevelctl/tpls/fabric"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"strconv"

	"github.com/Masterminds/sprig/v3"
	"github.com/manifoldco/promptui"
)

type DevFabric struct {
	OrdererCount int
	OrgCount     int
}

func getInputs() DevFabric {

	ordererCount := promptui.Prompt{
		Label:   "Enter the orderer count: ",
		Default: "1",
	}
	ordererCountResult, err := ordererCount.Run()
	if err != nil {
		fmt.Printf("Prompt failed entering the orderer count %v\n", err)
		os.Exit(1)
	}

	orgCount := promptui.Prompt{
		Label:   "Enter the organization count: ",
		Default: "1",
	}
	orgCountResult, err := orgCount.Run()
	if err != nil {
		fmt.Printf("Prompt failed entering the organization count %v\n", err)
		os.Exit(1)
	}

	ordererCountRes, _ := strconv.Atoi(ordererCountResult)
	orgCountRes, _ := strconv.Atoi(orgCountResult)

	return DevFabric{OrdererCount: ordererCountRes, OrgCount: orgCountRes}
}

func DevFabricNetworkConfig(platform string) string {
	var inputVars = getInputs()
	var FabricConfigFile bytes.Buffer
	fabricTemplate := template.New("Dev Fabric Template").Funcs(sprig.FuncMap())
	fabricTemplate, err := fabricTemplate.Parse(fabric.Fabric)
	if err != nil {
		log.Fatal("Parse: ", err)
	}
	err = fabricTemplate.Execute(&FabricConfigFile, inputVars)
	if err != nil {
		log.Fatal("Execute: ", err)
	}
	return FabricConfigFile.String()
}
