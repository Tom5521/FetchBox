package graph

import (
	"Windows-package-autoinstaller/src/core"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func saveText() {
	yamlData := core.Yamlfile{}

	// Set choco install values
	yamlData.Choco_Install = Choco_InstallTXT
	yamlData.Choco_Install_Configs.Force = install.Choco.Force
	yamlData.Choco_Install_Configs.Verbose = install.Choco.Verbose
	yamlData.Choco_Install_Configs.Upgrade = install.Choco.Upgrade
	// Set scoop install values
	yamlData.Scoop_Install_Configs.Upgrade = install.Scoop.Upgrade
	yamlData.Scoop_Install = Scoop_InstallTXT
	// Set choco uninstall values
	yamlData.Choco_Uninstall = Choco_UninstallTXT
	yamlData.Choco_Uninstall_Configs.Force = uninstall.Choco.Force
	yamlData.Choco_Uninstall_Configs.Verbose = uninstall.Choco.Verbose
	// Set Scoop uninstall values
	yamlData.Scoop_Uninstall = Scoop_UninstallTXT
	data, err := yaml.Marshal(yamlData)
	if err != nil {
		fmt.Println("Error marshalling YAML:", err)
		return
	}

	err = os.WriteFile(core.ConfigFilename, data, 0644)
	if err != nil {
		fmt.Println("Error writing YAML file:", err)
	}
}
