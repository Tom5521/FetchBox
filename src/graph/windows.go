package graph

import (
	"Windows-package-autoinstaller/src/core"
	"errors"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

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
