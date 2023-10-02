/*
 * Copyright (c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package graph

import (
	"errors"
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gopkg.in/yaml.v3"

	"Windows-package-autoinstaller/src/core"
	"Windows-package-autoinstaller/src/dev"
	"Windows-package-autoinstaller/src/icon"
)

var (
	MainSize    = fyne.NewSize(344, 237)
	ErrSize     = fyne.NewSize(400, 80)
	InstallSize = fyne.NewSize(400, 70)
)
var (
	DevICON      fyne.Resource
	DownloadICON fyne.Resource
	ErrorICON    fyne.Resource
	InstallICON  fyne.Resource
	SaveICON     fyne.Resource
	RestartICON  fyne.Resource
	InfoICON     fyne.Resource
)

func Init() {
	var (
		arg1, arg2, arg3                                     string
		scoopArg1                                            string
		chocoVerbose, ChocoForce, ChocoUpgrade, ScoopUpgrade bool
	)
	app := app.New()
	window := app.NewWindow("Windows Package AutoInstaller")
	window.Resize(MainSize)
	window.SetMaster()
	window.SetIcon(icon.AppICON)

	// Load the Icons and set the theme

	func() {
		icon.LoadIcons(app, ErrWin)
		icon.SetThemeIcons(app, ErrWin)
		DevICON = icon.DevICON
		DownloadICON = icon.DownloadICON
		ErrorICON = icon.ErrorICON
		InstallICON = icon.InstallICON
		SaveICON = icon.SaveICON
		RestartICON = icon.RestartICON
		InfoICON = icon.InfoICON
	}()

	if len(os.Args) > 1 {
		if os.Args[1] == "dev" {
			go CallDevWindow(app)
		}
	}

	// Get the yaml data
	data := core.GetYamldata()

	// Choco Interface
	chocoLabel := widget.NewLabel("Choco packages to install:")
	chocoLabel.TextStyle.Bold = true
	editedTextChoco := widget.NewMultiLineEntry()
	editedTextChoco.SetText(fmt.Sprintf("%v", data.Choco)) // Set yaml data
	chocoForceCheckBox1 := widget.NewCheck("Force", func(check bool) {
		if check {
			arg1 = "-f"
			ChocoForce = true
		}
	})
	if data.ChocoConfigs.Force {
		chocoForceCheckBox1.SetChecked(true) // Set yaml config
	}
	chocoVerboseCheckBox := widget.NewCheck("Verbose", func(check bool) {
		if check {
			arg2 = "-v"
			chocoVerbose = true
		}
	})
	if data.ChocoConfigs.Verbose {
		chocoVerboseCheckBox.SetChecked(true) // set yaml config
	}
	chocoUpgradeCheckBox := widget.NewCheck("Upgrade", func(check bool) {
		if check {
			arg3 = "upgrade"
			ChocoUpgrade = true
		}
	})
	if data.ChocoConfigs.Upgrade {
		chocoUpgradeCheckBox.SetChecked(true) // Set yaml config
	}
	installChocoPackBtn := widget.NewButtonWithIcon("Install Choco packages", DownloadICON, func() {
		if err := core.CheckOS(); err != nil {
			ErrWin(app, err, nil)
			return
		}
		if editedTextChoco.Text == "" || editedTextChoco.Text == "<nil>" {
			ErrWin(app, errors.New("choco package list is empty"), nil)
			return
		}
		if !core.IsAdmin {
			err := core.CheckSudo_External()
			if err != nil {
				ErrWin(
					app,
					errors.New(
						"sudo not detected!\nRestart the program with administrator permissions",
					),
					nil,
				)
				return
			}
		}
		if !core.CheckChoco() {
			InstallPkgManagerWin(app, "Choco", core.InstallChoco)
			return
		}
		InstallWindow(app, "Choco", func() error {
			err := core.ChocoPkgInstall(arg1, arg2, arg3)
			return err
		})
	})
	// Scoop Interface
	scoopLabel := widget.NewLabel("Scoop packages to install:")
	scoopLabel.TextStyle.Bold = true
	editedTextScoop := widget.NewMultiLineEntry()
	editedTextScoop.SetText(fmt.Sprintf("%v", data.Scoop))
	scoopUpgradeCheckButton := widget.NewCheck("Upgrade", func(check bool) {
		if check {
			scoopArg1 = "upgrade"
			ScoopUpgrade = true
		}
	})
	if data.ScoopConfigs.Upgrade {
		scoopUpgradeCheckButton.SetChecked(true)
	}
	installScoopPackBtn := widget.NewButtonWithIcon("Install Scoop Packages", DownloadICON, func() {
		if err := core.CheckOS(); err != nil {
			ErrWin(app, err, nil)
			return
		}
		if editedTextScoop.Text == "" || editedTextScoop.Text == "<nil>" {
			ErrWin(app, errors.New("scoop package list is empty"), nil)
			return
		}
		if !core.CheckScoop() {
			InstallPkgManagerWin(app, "Scoop", core.InstallScoop)
			return
		}
		InstallWindow(app, "Scoop", func() error {
			err := core.ScoopPkgInstall(scoopArg1)
			return err
		})
	})

	// Both Interface

	saveButton := widget.NewButtonWithIcon("Save configs", SaveICON, func() {
		saveText(
			editedTextChoco.Text,
			editedTextScoop.Text,
			chocoVerbose,
			ChocoForce,
			ChocoUpgrade,
			ScoopUpgrade,
		)
	})

	label := widget.NewLabel("Select any option")
	label.TextStyle.Italic = true

	// Set tabs and windows content
	chocoTab := container.NewVBox(
		chocoLabel,
		editedTextChoco,
		label,
		container.NewHBox(
			chocoVerboseCheckBox,
			chocoForceCheckBox1,
		),
		container.NewHBox(
			chocoUpgradeCheckBox,
		),
		container.NewHBox(
			saveButton,
			installChocoPackBtn,
		),
	)
	scoopTab := container.NewVBox(
		scoopLabel,
		editedTextScoop,
		label,
		scoopUpgradeCheckButton,
		container.NewHBox(
			saveButton,
			installScoopPackBtn,
		),
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
func saveText(
	chocoText, scoopText string,
	chocoVerbose, chocoForce, chocoUpgrade, scoopUpgrade bool,
) {
	yamlData := core.Yamlfile{}

	// Set choco values
	yamlData.Choco = chocoText
	yamlData.Scoop = scoopText
	yamlData.ChocoConfigs.Force = chocoForce
	yamlData.ChocoConfigs.Verbose = chocoVerbose
	yamlData.ChocoConfigs.Upgrade = chocoUpgrade
	// Set scoop values
	yamlData.ScoopConfigs.Upgrade = scoopUpgrade
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

func InstallWindow(app fyne.App, windowName string, f func() error) {
	window := app.NewWindow(windowName)
	window.Resize(InstallSize)
	window.SetFixedSize(true)
	window.SetIcon(InstallICON)
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
	window.Resize(ErrSize)
	//window.SetFixedSize(true)
	window.SetIcon(ErrorICON)
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

func InstallPkgManagerWin(app fyne.App, pkgman_name string, f func() error) {
	var restart bool
	window := app.NewWindow("Installing " + pkgman_name)
	window.Resize(InstallSize)
	window.SetFixedSize(true)
	window.SetIcon(InstallICON)
	continueButton := widget.NewButton("Accept", func() {
		window.Close()
		if restart {
			app.Quit()
		}
	})
	continueButton.Disable()

	infinite := widget.NewProgressBarInfinite()

	go func() {
		err := f()
		if err != nil {
			ErrWin(app, err, window)
			infinite.Stop()
			continueButton.Enable()
		} else {
			continueButton.Enable()
			infinite.Stop()
			RestartWindow(app, "Restart the program to use this function!!!")
		}
	}()

	content := container.NewVBox(
		infinite,
		continueButton,
	)
	window.SetContent(content)
	window.Show()
}

func RestartWindow(app fyne.App, restartTXT string) {
	window := app.NewWindow("Restart")
	window.Resize(ErrSize)
	window.SetFixedSize(true)
	window.SetIcon(RestartICON)

	label := widget.NewLabel(restartTXT)
	label.TextStyle.Bold = true
	label.Alignment = fyne.TextAlignCenter
	InfoWindow(app, "You need to restart the program!")
	restartButton := widget.NewButtonWithIcon("Restart", RestartICON, func() {
		app.Quit()
	})

	window.SetContent(container.NewVBox(
		label,
		restartButton,
	))
	window.Show()
}

func InfoWindow(app fyne.App, text string) {
	window := app.NewWindow("Info")
	window.SetIcon(InfoICON)

	InfoLabel := widget.NewLabel(text)
	InfoLabel.Alignment = fyne.TextAlignCenter
	InfoLabel.TextStyle.Bold = true

	AcceptButton := widget.NewButton("Accept", window.Close)

	window.SetContent(container.NewVBox(InfoLabel, AcceptButton))

	window.Show()
}

func CallDevWindow(app fyne.App) {
	dev.DevWindow(app, RestartWindow, ErrWin, InstallPkgManagerWin, InfoWindow)
}
