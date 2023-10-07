/*
 * Copyright (c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package dev

import (
	"FetchBox/src/core"
	"FetchBox/src/icon"
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func DevWindow(app fyne.App, RestartWindow func(fyne.App, string), ErrWin func(fyne.App, error, fyne.Window), InstallPkgManagerWin func(fyne.App, string, func() error), InfoWin func(fyne.App, string)) {
	// This statement of function is horrendous and fucking hellish, maybe I'll fix it later.
	window := app.NewWindow("Dev")
	window.SetIcon(icon.DevICON)
	checksudobutton := widget.NewButton("Check Sudo", func() {
		err := core.CheckSudo_External()
		if err != nil {
			InfoWin(app, "NOT Sudo")
		} else {
			InfoWin(app, "YES Sudo")
		}
	})
	restartButtom := widget.NewButton("Restart", func() {
		RestartWindow(app, "Dev")
	})
	errButtom := widget.NewButton("Custom Error buttom", func() {
		ErrWin(app, errors.New("Development"), nil)
	})
	installpkgmngBT_Scoop := widget.NewButton("Install Scoop pkg manager", func() {
		InstallPkgManagerWin(app, "Development -Scoop-", core.InstallScoop)
	})
	installpkgmngBT_Choco := widget.NewButton("Install choco pkg manager", func() {
		InstallPkgManagerWin(app, "Development -Choco-", core.InstallChoco)
	})
	ShowIconsButton := widget.NewButton("Show all icons", func() {
		ShowIconsWin(app)
	})
	ShowInfoWinButton := widget.NewButton("Show Info window", func() {
		InfoWin(app, "This is a dev window!")
	})
	CheckScoopButton := widget.NewButton("Check Scoop", func() {
		check := core.CheckScoop()
		if check {
			InfoWin(app, "Scoop is installed")
		} else {
			InfoWin(app, "Scoop isn't installed")
		}
	})

	content := container.NewVBox(
		restartButtom,
		errButtom,
		installpkgmngBT_Choco,
		installpkgmngBT_Scoop,
		ShowIconsButton,
		ShowInfoWinButton,
		checksudobutton,
		CheckScoopButton,
	)
	window.SetContent(content)
	window.Show()
}

func ShowIconsWin(app fyne.App) {
	w := app.NewWindow("IconsPreview")

	AppIcon := canvas.NewImageFromResource(icon.AppICON)
	AppIcon.FillMode = canvas.ImageFillOriginal

	icon1 := canvas.NewImageFromResource(icon.DevICON_Dark)
	icon2 := canvas.NewImageFromResource(icon.DownloadICON_Dark)
	icon3 := canvas.NewImageFromResource(icon.ErrorICON_Dark)
	icon4 := canvas.NewImageFromResource(icon.InstallICON_Dark)
	icon5 := canvas.NewImageFromResource(icon.SaveICON_Dark)
	icon6 := canvas.NewImageFromResource(icon.RestartICON_Dark)
	icon14 := canvas.NewImageFromResource(icon.InfoICON_Dark)
	icon15 := canvas.NewImageFromResource(icon.UninstallICON_Dark)

	icon1.FillMode = canvas.ImageFillOriginal
	icon2.FillMode = canvas.ImageFillOriginal
	icon3.FillMode = canvas.ImageFillOriginal
	icon4.FillMode = canvas.ImageFillOriginal
	icon5.FillMode = canvas.ImageFillOriginal
	icon6.FillMode = canvas.ImageFillOriginal

	icon7 := canvas.NewImageFromResource(icon.DevICON_Light)
	icon8 := canvas.NewImageFromResource(icon.DownloadICON_Light)
	icon9 := canvas.NewImageFromResource(icon.ErrorICON_Light)
	icon10 := canvas.NewImageFromResource(icon.InstallICON_Light)
	icon11 := canvas.NewImageFromResource(icon.SaveICON_Light)
	icon12 := canvas.NewImageFromResource(icon.RestartICON_Light)
	icon13 := canvas.NewImageFromResource(icon.InfoICON_Light)
	icon16 := canvas.NewImageFromResource(icon.UninstallICON_Light)

	icon7.FillMode = canvas.ImageFillOriginal
	icon8.FillMode = canvas.ImageFillOriginal
	icon9.FillMode = canvas.ImageFillOriginal
	icon10.FillMode = canvas.ImageFillOriginal
	icon11.FillMode = canvas.ImageFillOriginal
	icon12.FillMode = canvas.ImageFillOriginal

	icon13.FillMode = canvas.ImageFillOriginal
	icon14.FillMode = canvas.ImageFillOriginal
	icon15.FillMode = canvas.ImageFillOriginal
	icon16.FillMode = canvas.ImageFillOriginal

	contentH1 := container.NewHBox(
		icon1,
		icon2,
		icon3,
		icon4,
		icon5,
		icon6,
		icon14,
		icon15,
	)

	contentH2 := container.NewHBox(
		icon7,
		icon8,
		icon9,
		icon10,
		icon11,
		icon12,
		icon13,
		icon16,
	)
	w.SetContent(container.NewVBox(contentH1, contentH2, AppIcon))
	w.Show()
}
