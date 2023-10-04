package graph

import (
	"Windows-package-autoinstaller/src/core"
	"errors"
	"fmt"

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
		data.Choco_Uninstall_Configs.Verbose = check
		uninstall.Choco.Verbose = check
	})
	if data.Choco_Uninstall_Configs.Verbose {
		verboseUninstallChocoCheck.SetChecked(true)
	}
	forceUninstallChocoCheck := widget.NewCheck("Force", func(check bool) {
		data.Choco_Uninstall_Configs.Force = check
		uninstall.Choco.Force = data.Choco_Uninstall_Configs.Force
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
		data.Scoop_Uninstall = editedTextScoopUninstall.Text
		data.Choco_Uninstall = editedTextChocoUninstall.Text
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

func InstallTab(app fyne.App) *container.AppTabs {
	// First label
	chocoLabel := widget.NewLabel("Choco packages to install:")
	chocoLabel.TextStyle.Bold = true
	// Multi Line Entry
	editedTextChoco := widget.NewMultiLineEntry()
	editedTextChoco.SetText(fmt.Sprintf("%v", data.Choco_Install)) // Set yaml data
	// CheckBoxs
	chocoForceCheckBox1 := widget.NewCheck("Force", func(check bool) {
		data.Choco_Install_Configs.Force = check
		install.Choco.Force = check
	})
	if data.Choco_Install_Configs.Force {
		chocoForceCheckBox1.SetChecked(true) // Set yaml config
	}
	chocoVerboseCheckBox := widget.NewCheck("Verbose", func(check bool) {
		data.Choco_Install_Configs.Verbose = check
		install.Choco.Verbose = check
	})
	if data.Choco_Install_Configs.Verbose {
		chocoVerboseCheckBox.SetChecked(true) // set yaml config
	}
	chocoUpgradeCheckBox := widget.NewCheck("Upgrade", func(check bool) {
		if check {
			data.Choco_Install_Configs.Upgrade = check
			install.Choco.Upgrade = check
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
	editedTextScoop.SetText(data.Scoop_Install)
	scoopUpgradeCheckButton := widget.NewCheck("Upgrade", func(check bool) {
		data.Scoop_Install_Configs.Upgrade = check
		install.Scoop.Upgrade = check
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
		data.Choco_Install = editedTextChoco.Text
		data.Scoop_Install = editedTextScoop.Text
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
