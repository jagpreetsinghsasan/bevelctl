package main

import (
	"bevelctl/bevel"
	"bevelctl/support"
	"fmt"
	"os"
	"strings"

	"github.com/devfacet/gocmd/v3"
	"github.com/manifoldco/promptui"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var InstallationStatus []string

type BevelctlInputs struct {
	environment string `json:environment`
	platform    string `json:platform`
}

func main() {

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewConsoleEncoder(config)
	os.RemoveAll("log")
	os.Mkdir("log", os.ModePerm)
	fileName := "log/general.txt"
	logFile, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	defer logger.Sync()

	var bevelctlInputs = BevelctlInputs{
		platform:    "not_selected",
		environment: "not_selected",
	}

	flags := struct {
		Help        bool   `short:"h" long:"help" description:"Bevelctl usage" global:"true"`
		Version     bool   `short:"v" long:"version" description:"Bevelctl version"`
		Environment string `short:"e" long:"environment" description:"Select the environment: [dev]"`
		Platform    string `short:"p" long:"platform" description:"Select the blockchain platform: [fabric, corda]"`
	}{}

	gocmd.HandleFlag("Environment", func(cmd *gocmd.Cmd, args []string) error {
		if strings.ToLower(flags.Environment) == "dev" {
			bevelctlInputs.environment = flags.Environment
		}
		return nil
	})

	gocmd.HandleFlag("Platform", func(cmd *gocmd.Cmd, args []string) error {
		if strings.ToLower(flags.Platform) == "fabric" || strings.ToLower(flags.Platform) == "corda" {
			bevelctlInputs.platform = flags.Platform
		}
		return nil
	})

	gocmd.New(gocmd.Options{
		Name:        "bevelctl",
		Description: "Cli for Hyperledger Bevel",
		Version:     "0.1",
		Flags:       &flags,
		ConfigType:  gocmd.ConfigTypeAuto,
	})

	if bevelctlInputs.environment == "not_selected" {
		bevelctlInputs.environment = support.SupportedEnvironments[0]
	}
	if bevelctlInputs.platform == "not_selected" {
		bevelctlInputs.platform = support.SupportedPlatforms[0]
	}

	fmt.Printf("-------------------------------\nYour selected choices are: %+v\n-------------------------------\n", bevelctlInputs)
	continueSelect := promptui.Select{
		Label: "Do you want to continue with the above choices?",
		Items: []string{"Yes", "No"},
	}
	_, continueSelectResult, err := continueSelect.Run()
	if err != nil || continueSelectResult == "No" {
		logger.Info("Exiting...")
		os.Exit(0)
	} else {
		bevel.ExecuteBevel(bevelctlInputs, logger)
	}
}
