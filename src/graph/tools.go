package graph

import (
	"Windows-package-autoinstaller/src/core"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func saveText() {
	data, err := yaml.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling YAML:", err)
		return
	}

	err = os.WriteFile(core.ConfigFilename, data, 0644)
	if err != nil {
		fmt.Println("Error writing YAML file:", err)
	}
}
