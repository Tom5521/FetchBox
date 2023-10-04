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
	// Declare sizes
	MainSize    = fyne.NewSize(344, 237)
	ErrSize     = fyne.NewSize(400, 80)
	InstallSize = fyne.NewSize(400, 70)

	// Declare Icons
	DevICON      fyne.Resource
	DownloadICON fyne.Resource
	ErrorICON    fyne.Resource
	InstallICON  fyne.Resource
	SaveICON     fyne.Resource
	RestartICON  fyne.Resource
	InfoICON     fyne.Resource

	// Declare structures
	install   = core.Install{}
	uninstall = core.Uninstall{}
	data      = core.GetYamldata()

	// Declare global var (
	Scoop_InstallTXT   string
	Choco_InstallTXT   string
	Scoop_UninstallTXT string
	Choco_UninstallTXT string
)

func Init() {
	app := app.New()
	window := app.NewWindow("Windows Package AutoInstaller")
	window.Resize(MainSize)
	window.SetMaster()
	window.SetIcon(icon.AppICON)

	// Load the Icons and set the theme

	func() {
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

	GeneralTabs := container.NewAppTabs(
		container.NewTabItem("Install", InstallTab(app)),
		container.NewTabItem("Uninstall", UninstallTab(app)),
	)

	window.SetContent(GeneralTabs)
	window.ShowAndRun()
}
func InstallTab(app fyne.App) *container.AppTabs {
	chocoLabel := widget.NewLabel("Choco packages to install:")
	chocoLabel.TextStyle.Bold = true
	editedTextChoco := widget.NewMultiLineEntry()
	editedTextChoco.SetText(fmt.Sprintf("%v", data.Choco_Install)) // Set yaml data
	chocoForceCheckBox1 := widget.NewCheck("Force", func(check bool) {
		if check {
			install.Choco.Force = true
		}
	})
	if data.Choco_Install_Configs.Force {
		chocoForceCheckBox1.SetChecked(true) // Set yaml config
	}
	chocoVerboseCheckBox := widget.NewCheck("Verbose", func(check bool) {
		if check {
			install.Choco.Verbose = true
		}
	})
	if data.Choco_Install_Configs.Verbose {
		chocoVerboseCheckBox.SetChecked(true) // set yaml config
	}
	chocoUpgradeCheckBox := widget.NewCheck("Upgrade", func(check bool) {
		if check {
			install.Choco.Upgrade = true
		}
	})
	if data.Choco_Install_Configs.Upgrade {
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
			err := install.ChocoPkgInstall()
			return err
		})
	})
	// Scoop Interface
	scoopLabel := widget.NewLabel("Scoop packages to install:")
	scoopLabel.TextStyle.Bold = true
	editedTextScoop := widget.NewMultiLineEntry()
	editedTextScoop.SetText(fmt.Sprintf("%v", data.Scoop_Install))
	scoopUpgradeCheckButton := widget.NewCheck("Upgrade", func(check bool) {
		if check {
			install.Scoop.Upgrade = true
		}
	})
	if data.Scoop_Install_Configs.Upgrade {
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
			err := install.ScoopPkgInstall()
			return err
		})
	})

	saveButton := widget.NewButtonWithIcon("Save configs", SaveICON, func() {
		editedTextChoco.Text = Choco_InstallTXT
		editedTextScoop.Text = Scoop_InstallTXT
		saveText()
	})

	label := widget.NewLabel("Select any option")
	label.TextStyle.Italic = true

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
	InstallTabs := container.NewAppTabs(
		container.NewTabItem("Choco", chocoTab),
		container.NewTabItem("Scoop", scoopTab),
	)
	return InstallTabs
}

func UninstallTab(app fyne.App) *container.AppTabs {
	chocoLabelUtab := widget.NewLabel("Choco packages to uninstall")
	chocoLabelUtab.TextStyle.Bold = true
	editedTextChocoUninstall := widget.NewMultiLineEntry()
	editedTextChocoUninstall.SetText(data.Choco_Uninstall)

	verboseUninstallChocoCheck := widget.NewCheck("Verbose", func(check bool) {
		if check {
			uninstall.Choco.Verbose = true
		}
	})
	if data.Choco_Uninstall_Configs.Verbose {
		verboseUninstallChocoCheck.SetChecked(true)
	}
	forceUninstallChocoCheck := widget.NewCheck("Force", func(check bool) {
		if check {
			uninstall.Choco.Force = true
		}
	})
	if data.Choco_Uninstall_Configs.Force {
		forceUninstallChocoCheck.SetChecked(true)
	}
	chocoUninstallButton := widget.NewButton("Uninstall Packages", func() {
		err := uninstall.UninstallChocoPkgs()
		if err != nil {
			ErrWin(app, err, nil)
		}
	})
	// Scoop
	scoopLabelUtab := widget.NewLabel("Scoop packages to uninstall")
	scoopLabelUtab.TextStyle.Bold = true
	editedTextScoopUninstall := widget.NewMultiLineEntry()
	editedTextScoopUninstall.SetText(data.Scoop_Uninstall)
	scoopUninstallButton := widget.NewButton("Uninstall packages", func() {
		err := uninstall.UninstallScoopPkgs()
		if err != nil {
			ErrWin(app, err, nil)
		}
	})

	// Both Interface

	saveButton := widget.NewButtonWithIcon("Save configs", SaveICON, func() {
		editedTextChocoUninstall.Text = Choco_UninstallTXT
		editedTextScoopUninstall.Text = Scoop_UninstallTXT
		saveText()
	})

	label := widget.NewLabel("Select any option")
	label.TextStyle.Italic = true
	// Declare the containers
	chocoUninstallTab := container.NewVBox(
		chocoLabelUtab,
		editedTextChocoUninstall,
		label,
		container.NewHBox(
			verboseUninstallChocoCheck,
			forceUninstallChocoCheck,
		),
		container.NewHBox(
			saveButton,
			chocoUninstallButton),
	)
	scoopUninstallTab := container.NewVBox(
		scoopLabelUtab,
		editedTextChocoUninstall,
		label,
		container.NewHBox(
			saveButton,
			scoopUninstallButton),
	)

	UninstallTab := container.NewAppTabs(
		container.NewTabItem("Choco", chocoUninstallTab),
		container.NewTabItem("Scoop", scoopUninstallTab),
	)
	return UninstallTab
}

// Internal functions
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
