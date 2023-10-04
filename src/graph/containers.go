package graph

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

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
