package vault

// Template containing values for patching the vault chart to provide extra functionality
const Vault = `
ui:
  enabled: true
  serviceType: "NodePort"
`
