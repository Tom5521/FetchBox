package graph

import (
	"errors"
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/Windows-package-autoinstaller/src/core"
	"gopkg.in/yaml.v3"
)

func Init() {
	app := app.New()
	window := app.NewWindow("Windows Package AutoInstaller")
	window.Resize(fyne.NewSize(344, 327))
	window.SetMaster()

	yamlData := getYmlData()

	chocoLabel := widget.NewLabel("Choco packages to install:")
	chocoLabel.TextStyle.Bold = true
	editedTextChoco := widget.NewMultiLineEntry()
	editedTextChoco.SetText(fmt.Sprintf("%v", yamlData["choco"]))

	scoopLabel := widget.NewLabel("Scoop packages to install:")
	scoopLabel.TextStyle.Bold = true
	editedTextScoop := widget.NewMultiLineEntry()
	editedTextScoop.SetText(fmt.Sprintf("%v", yamlData["scoop"]))

	saveButton := widget.NewButton("Save package lists", func() {
		saveText(editedTextChoco.Text, editedTextScoop.Text)
	})

	label := widget.NewLabel("Select any option")
	label.TextStyle.Italic = true

	installChocoPackBtn := widget.NewButton("Install Choco packages", func() {
		if editedTextChoco.Text == "" || editedTextChoco.Text == "<nil>" {
			ErrWin(app, errors.New("Choco package list is empty"), nil)
			return
		}
		err := core.CheckSudo_External()
		if err != nil {
			ErrWin(app, errors.New("Sudo not detected!\nRestart the program with administrator permissions"), nil)
			return
		}
		ChocoInstall(app)
	})
	installScoopPackBtn := widget.NewButton("Install Scoop Packages", func() {
		if editedTextScoop.Text == "" || editedTextScoop.Text == "<nil>" {
			ErrWin(app, errors.New("Scoop package list is empty"), nil)
			return
		}
		ScoopInstall(app)
	})

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

func ChocoInstall(app fyne.App) {
	window := app.NewWindow("Installing choco packages")
	window.Resize(fyne.NewSize(400, 70))
	window.SetFixedSize(true)
	infinite := widget.NewProgressBarInfinite()
	acpBT := widget.NewButton("Continue", func() {
		window.Close()
	})
	acpBT.Disable()
	go func() {
		err := core.ChocoPkgInstall()
		if err != nil {
			infinite.Stop()
			window.SetTitle("Completed with errors.")
			ErrWin(app, err, window)
		} else {
			infinite.Stop()
			window.SetTitle("Completed.")
			acpBT.Enable()
		}
	}()
	window.Close()
	content := container.NewVBox(
		infinite,
		acpBT,
	)
	window.SetContent(content)
	window.Show()

}

func ScoopInstall(app fyne.App) {
	window := app.NewWindow("Installing scoop packages")
	window.Resize(fyne.NewSize(400, 70))
	window.SetFixedSize(true)
	infinite := widget.NewProgressBarInfinite()
	acpBT := widget.NewButton("Continue", func() {
		window.Close()
	})
	acpBT.Disable()
	go func() {
		err := core.ScoopPkgInstall()
		if err != nil {
			window.SetTitle("Completed with errors")
			infinite.Stop()
			ErrWin(app, err, window)
		} else {
			infinite.Stop()
			window.SetTitle("Completed.")
			acpBT.Enable()
		}
	}()
	window.Close()
	content := container.NewVBox(
		infinite,
		acpBT,
	)
	window.SetContent(content)
	window.Show()

}

func ErrWin(app fyne.App, err error, clWindow fyne.Window) {
	window := app.NewWindow("Error")
	window.Resize(fyne.NewSize(400, 50))
	window.SetFixedSize(true)
	errlabel := widget.NewLabel(err.Error())
	errlabel.TextStyle.Bold = true
	errlabel.Alignment = fyne.TextAlignCenter
	acceptButton := widget.NewButton("Accept", func() {
		window.Close()
		if clWindow != nil {
			clWindow.Close()
		}
	})

	content := container.NewVBox(
		errlabel,
		acceptButton,
	)
	window.SetContent(content)
	window.SetMainMenu(window.MainMenu())
	window.Show()

}
