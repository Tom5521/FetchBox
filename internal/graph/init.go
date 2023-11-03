/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package graph

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"FetchBox/cmd/core"
	"FetchBox/internal/dev"
	icon "FetchBox/pkg/icons"
	"FetchBox/pkg/windows"
)

var (
	// Declare structures
	install   = core.Install{}
	uninstall = core.Uninstall{}
	data      = core.GetYamldata()
)

func Init() {
	app := app.New()
	window := app.NewWindow("FetchBox")
	window.Resize(windows.MainSize)
	window.SetMaster()

	// Load the Icons and set the theme
	icon.SetThemeIcons(app)

	window.SetIcon(icon.AppICON)

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
	dev.DevWindow(app, windows.RestartWindow, windows.ErrWin, windows.InstallPkgManagerWin, windows.InfoWindow)
}
