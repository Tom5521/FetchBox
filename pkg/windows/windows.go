/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package windows

import (
	icon "FetchBox/pkg/icons"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	// Declare winsizes
	InstallSize = fyne.NewSize(400, 70)
	MainSize    = fyne.NewSize(344, 237)
	ErrSize     = fyne.NewSize(400, 80)
)

func ProcessWindow(app fyne.App, windowName string, f func() error) {
	window := app.NewWindow(windowName)
	window.Resize(InstallSize)
	window.SetFixedSize(true)
	window.SetIcon(icon.InstallICON)
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
			ErrWin(app, err.Error(), window)
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

func ErrWin(app fyne.App, err string, clWindow ...fyne.Window) {
	window := app.NewWindow("Error")
	window.Resize(ErrSize)
	//window.SetFixedSize(true)
	window.SetIcon(icon.ErrorICON)
	errlabel := widget.NewLabel(err)
	errlabel.TextStyle.Bold = true
	errlabel.Alignment = fyne.TextAlignCenter
	acceptButton := widget.NewButton("Accept", func() {
		window.Close()
		if len(clWindow) != 0 {
			if clWindow != nil {
				clWindow[0].Close()
			}
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
	window.SetIcon(icon.InstallICON)
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
			ErrWin(app, err.Error(), window)
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
	window.SetIcon(icon.RestartICON)

	label := widget.NewLabel(restartTXT)
	label.TextStyle.Bold = true
	label.Alignment = fyne.TextAlignCenter
	InfoWindow(app, "You need to restart the program!")
	restartButton := widget.NewButtonWithIcon("Restart", icon.RestartICON, func() {
		app.Quit()
	})

	window.SetContent(container.NewVBox(
		label,
		restartButton,
	))
	window.Show()
}

func InfoWindow(app fyne.App, InfoLabelTXT string, infoTitle ...string) {
	var info = "Info"
	if len(infoTitle) >= 1 {
		if infoTitle[0] != "" {
			info = infoTitle[0]
		}
	}
	window := app.NewWindow(info)
	window.SetIcon(icon.InfoICON)

	InfoLabel := widget.NewLabel(InfoLabelTXT)
	InfoLabel.Alignment = fyne.TextAlignCenter
	InfoLabel.TextStyle.Bold = true

	AcceptButton := widget.NewButton("Accept", window.Close)

	window.SetContent(container.NewVBox(InfoLabel, AcceptButton))

	window.Show()
}
