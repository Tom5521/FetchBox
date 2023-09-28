package dev

import (
	"Windows-package-autoinstaller/src/core"
	"Windows-package-autoinstaller/src/icon"
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func DevWindow(app fyne.App, RestartWindow func(fyne.App, string), ErrWin func(fyne.App, error, fyne.Window), InstallPkgManagerWin func(fyne.App, string, func() error)) { // This statement of function is horrendous and fucking hellish, maybe I'll fix it later.
	window := app.NewWindow("Dev")
	window.SetIcon(icon.DevICON)
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

	content := container.NewVBox(
		restartButtom,
		errButtom,
		installpkgmngBT_Choco,
		installpkgmngBT_Scoop,
		ShowIconsButton,
	)
	window.SetContent(content)
	window.Show()
}

func ShowIconsWin(app fyne.App) {
	w := app.NewWindow("IconsPreview")

	size := fyne.NewSize(64, 64)

	AppIcon := canvas.NewImageFromResource(icon.AppICON)
	AppIcon.Resize(size)

	icon1 := canvas.NewImageFromResource(icon.DevICON_Dark)
	icon2 := canvas.NewImageFromResource(icon.DownloadICON_Dark)
	icon3 := canvas.NewImageFromResource(icon.ErrorICON_Dark)
	icon4 := canvas.NewImageFromResource(icon.InstallICON_Dark)
	icon5 := canvas.NewImageFromResource(icon.SaveICON_Dark)
	icon6 := canvas.NewImageFromResource(icon.RestartICON_Dark)

	//icon1.Resize(size)
	//icon1.MinSize().Width = 64
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

	icon7.FillMode = canvas.ImageFillOriginal
	icon8.FillMode = canvas.ImageFillOriginal
	icon9.FillMode = canvas.ImageFillOriginal
	icon10.FillMode = canvas.ImageFillOriginal
	icon11.FillMode = canvas.ImageFillOriginal
	icon12.FillMode = canvas.ImageFillOriginal

	contentH1 := container.NewHBox(
		icon1,
		icon2,
		icon3,
		icon4,
		icon5,
		icon6,
	)

	contentH2 := container.NewHBox(
		icon7,
		icon8,
		icon9,
		icon10,
		icon11,
		icon12,
	)

	w.SetContent(container.NewVBox(contentH1, contentH2, AppIcon))
	w.Show()
}
