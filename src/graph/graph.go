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

	data := core.GetYamldata()

	// Choco Interface
	chocoLabel := widget.NewLabel("Choco packages to install:")
	chocoLabel.TextStyle.Bold = true
	editedTextChoco := widget.NewMultiLineEntry()
	editedTextChoco.SetText(fmt.Sprintf("%v", data.Choco))
	installChocoPackBtn := widget.NewButton("Install Choco packages", func() {
		if err := core.CheckOS(); err != nil {
			ErrWin(app, err, nil)
			return
		}
		if editedTextChoco.Text == "" || editedTextChoco.Text == "<nil>" {
			ErrWin(app, errors.New("Choco package list is empty"), nil)
			return
		}
		if !core.IsAdmin {
			err := core.CheckSudo_External()
			if err != nil {
				ErrWin(app, errors.New("Sudo not detected!\nRestart the program with administrator permissions"), nil)
				return
			}
		}
		InstallWindow(app, "choco", core.ChocoPkgInstall)
	})
	// Scoop Interface
	scoopLabel := widget.NewLabel("Scoop packages to install:")
	scoopLabel.TextStyle.Bold = true
	editedTextScoop := widget.NewMultiLineEntry()
	editedTextScoop.SetText(fmt.Sprintf("%v", data.Scoop))

	installScoopPackBtn := widget.NewButton("Install Scoop Packages", func() {
		if err := core.CheckOS(); err != nil {
			ErrWin(app, err, nil)
			return
		}
		if editedTextScoop.Text == "" || editedTextScoop.Text == "<nil>" {
			ErrWin(app, errors.New("Scoop package list is empty"), nil)
			return
		}
		InstallWindow(app, "scoop", core.ScoopPkgInstall)
	})

	// Both Interface

	saveButton := widget.NewButton("Save package lists", func() {
		saveText(editedTextChoco.Text, editedTextScoop.Text)
	})

	label := widget.NewLabel("Select any option")
	label.TextStyle.Italic = true

	// Set tabs and windows content
	chocoTab := container.NewVBox(
		chocoLabel,
		editedTextChoco,
		label,
		saveButton,
		installChocoPackBtn,
	)
	scoopTab := container.NewVBox(
		scoopLabel,
		editedTextScoop,
		label,
		saveButton,
		installScoopPackBtn,
	)
	tabs := container.NewAppTabs(
		container.NewTabItem("Choco", chocoTab),
		container.NewTabItem("Scoop", scoopTab),
	)

	content := container.NewVBox(
		tabs,
	)

	window.SetContent(content)
	window.ShowAndRun()
}

// Internal functions
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

func InstallWindow(app fyne.App, pkgManager string, f func() error) {
	window := app.NewWindow(fmt.Sprintf("Installing %v packages", pkgManager))
	window.Resize(fyne.NewSize(400, 70))
	window.SetFixedSize(true)
	infinite := widget.NewProgressBarInfinite()
	acceptButton := widget.NewButton("Continue", func() {
		window.Close()
	})
	acceptButton.Disable()
	go func() {
		err := f()
		if err != nil {
			window.SetTitle("Completed with errors")
			infinite.Stop()
			ErrWin(app, err, window)
		} else {
			infinite.Stop()
			window.SetTitle("Completed.")
			acceptButton.Enable()
		}
	}()
	content := container.NewVBox(
		infinite,
		acceptButton,
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
