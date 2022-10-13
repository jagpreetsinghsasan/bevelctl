package vault

import (
	"bevelctl/tpls/vault"
	"bytes"
	"os"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"go.uber.org/zap"
)

func CreateVaultConfig(logger *zap.Logger) {
	var VaultConfigFile bytes.Buffer
	vaultTemplate := template.New("Vault Config File").Funcs(template.FuncMap(sprig.FuncMap()))
	vaultTemplate, err := vaultTemplate.Parse(vault.Vault)
	if err != nil {
		logger.Fatal("Error during parsing the vault config file", zap.Any("ERROR", err))
	}
	err = vaultTemplate.Execute(&VaultConfigFile, "")
	if err != nil {
		logger.Fatal("Error during executing the tpl file with vars", zap.Any("ERROR", err))
	}

	os.Mkdir("build", os.ModePerm)
	file, err := os.Create("build/vaultconfig.yaml")
	if err != nil {
		logger.Fatal("Failed while creating the build/vaultconfig.yaml file", zap.Any("ERROR", err))
	}
	defer file.Close()
	file.WriteString(VaultConfigFile.String())

}
