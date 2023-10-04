/*
 * Copyright (c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package graph

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

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
	Isdev     = true
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

func CallDevWindow(app fyne.App) {
	dev.DevWindow(app, RestartWindow, ErrWin, InstallPkgManagerWin, InfoWindow)
}
