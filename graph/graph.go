package graph

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/Windows-package-autoinstaller/src"
	"gopkg.in/yaml.v3"
)

func Init() {
	app := app.New()
	window := app.NewWindow("Windows Package AutoInstaller")
	window.Resize(fyne.NewSize(344, 327))

	yamlData := getYmlData()

	chocoLabel := widget.NewLabel("Choco packages to install:")
	editedTextChoco := widget.NewMultiLineEntry()
	editedTextChoco.SetText(fmt.Sprintf("%v", yamlData["choco"]))

	scoopLabel := widget.NewLabel("Scoop packages to install:")
	editedTextScoop := widget.NewMultiLineEntry()
	editedTextScoop.SetText(fmt.Sprintf("%v", yamlData["scoop"]))

	saveButton := widget.NewButton("Save package lists", func() {
		saveText(editedTextChoco.Text, editedTextScoop.Text)
	})

	label := widget.NewLabel("Select any option")

	installChocoPackBtn := widget.NewButton("Install Choco packages", src.ChocoInstall)
	installScoopPackBtn := widget.NewButton("Install Scoop Packages", src.ScoopInstall)

	content := container.NewVBox(
		chocoLabel,
		editedTextChoco,
		scoopLabel,
		editedTextScoop,
		saveButton,
		label,
		installChocoPackBtn,
		installScoopPackBtn,
	)

	window.SetContent(content)
	window.ShowAndRun()
}

func saveText(chocoText, scoopText string) {
	yamlData := map[string]interface{}{
		"choco": chocoText,
		"scoop": scoopText,
	}

	data, err := yaml.Marshal(yamlData)
	if err != nil {
		fmt.Println("Error marshalling YAML:", err)
		return
	}

	err = os.WriteFile("packages.yml", data, 0644)
	if err != nil {
		fmt.Println("Error writing YAML file:", err)
	}
}

func getYmlData() map[string]interface{} {
	yamlData := make(map[string]interface{})
	data, err := os.ReadFile("packages.yml")
	if err == nil {
		err := yaml.Unmarshal(data, &yamlData)
		if err != nil {
			fmt.Println("Error parsing YAML:", err)
		}
	}
	return yamlData
}
