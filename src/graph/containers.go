package graph

import (
	"FetchBox/src/core"
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
		uninstall.Choco.Verbose = check
	})
	if data.Choco_Uninstall_Configs.Verbose {
		verboseUninstallChocoCheck.SetChecked(true)
	}
	forceUninstallChocoCheck := widget.NewCheck("Force", func(check bool) {
		uninstall.Choco.Force = check
	})
	if data.Choco_Uninstall_Configs.Force {
		forceUninstallChocoCheck.SetChecked(true)
	}
	chocoUninstallButton := widget.NewButtonWithIcon("Uninstall Packages", UninstallICON, func() {
		err := basicChocoCheck(app, editedTextChocoUninstall.Text)
		if err != nil {
			return
		}
		err = uninstall.UninstallChocoPkgs()
		if err != nil {
			ErrWin(app, err, nil)
		}
	})
	// Scoop
	scoopLabelUtab := widget.NewLabel("Scoop packages to uninstall")
	scoopLabelUtab.TextStyle.Bold = true
	editedTextScoopUninstall := widget.NewMultiLineEntry()
	editedTextScoopUninstall.SetText(data.Scoop_Uninstall)
	scoopUninstallButton := widget.NewButtonWithIcon("Uninstall packages", UninstallICON, func() {
		err := uninstall.UninstallScoopPkgs()
		if err != nil {
			ErrWin(app, err, nil)
		}
	})

	// Both Interface

	saveButton := widget.NewButtonWithIcon("Save configs", SaveICON, func() {
		// Set to save the multine edited text
		data.Scoop_Uninstall = editedTextScoopUninstall.Text
		data.Choco_Uninstall = editedTextChocoUninstall.Text
		// Save the choco checks
		data.Choco_Uninstall_Configs.Force = forceUninstallChocoCheck.Checked
		data.Choco_Uninstall_Configs.Verbose = verboseUninstallChocoCheck.Checked
		// Scoop no have checks :/
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
	chocoLabel.TextStyle.Bold = true // Set bold to label
	// Multi Line Entry
	editedTextChoco := widget.NewMultiLineEntry()
	editedTextChoco.SetText(fmt.Sprintf("%v", data.Choco_Install)) // Set yaml data
	// CheckBoxs
	chocoForceCheckBox1 := widget.NewCheck("Force", func(check bool) {
		install.Choco.Force = check // Install choco parameters <- check value
	})
	if data.Choco_Install_Configs.Force { // Set the variable in accordance with the yaml setting
		chocoForceCheckBox1.SetChecked(true) // Set yaml config
	}
	chocoVerboseCheckBox := widget.NewCheck("Verbose", func(check bool) {
		install.Choco.Verbose = check // <- choco verbose parameter <- check value
	})
	if data.Choco_Install_Configs.Verbose { // Set the variable in accordance with the yaml setting
		chocoVerboseCheckBox.SetChecked(true) // set yaml config
	}
	chocoUpgradeCheckBox := widget.NewCheck("Upgrade", func(check bool) {
		install.Choco.Upgrade = check // <- choco upgrade parameter <- check value

	})
	if data.Choco_Install_Configs.Upgrade { //Set the variable in accordance with the yaml setting
		chocoUpgradeCheckBox.SetChecked(true) // Set yaml config
	}

	// Set button for install with icon
	installChocoPackBtn := widget.NewButtonWithIcon("Install Choco packages", DownloadICON, func() {
		err := basicChocoCheck(app, editedTextChoco.Text)
		if err != nil {
			return
		}
		InstallWindow(app, "Choco", func() error { // Install window with install function
			err := install.ChocoPkgInstall()
			return err
		})
	})
	// Scoop Interface
	scoopLabel := widget.NewLabel("Scoop packages to install:")
	scoopLabel.TextStyle.Bold = true // Set label text to bold style
	// Declare multine entry
	editedTextScoop := widget.NewMultiLineEntry()
	editedTextScoop.SetText(data.Scoop_Install) // Set multine entry with yaml data

	// Checks
	scoopUpgradeCheckButton := widget.NewCheck("Upgrade", func(check bool) { // Scoop Upgrade check
		install.Scoop.Upgrade = check // Set upgrade scoop parameter
	})
	//
	if data.Scoop_Install_Configs.Upgrade { //Set the variable in accordance with the yaml setting
		scoopUpgradeCheckButton.SetChecked(true)
	}
	// Scoop install init button with icon
	installScoopPackBtn := widget.NewButtonWithIcon("Install Scoop Packages", DownloadICON, func() {
		if err := core.CheckOS(); err != nil { // Check the OS
			ErrWin(app, err, nil)
			return
		}
		// Check for nil string values
		if editedTextScoop.Text == "" || editedTextScoop.Text == "<nil>" {
			ErrWin(app, errors.New("scoop package list is empty"), nil)
			return
		}
		// Check if scoop is installed
		if !core.CheckScoop() {
			InstallPkgManagerWin(app, "Scoop", core.InstallScoop)
			return
		}
		// Init install window
		InstallWindow(app, "Scoop", func() error {
			err := install.ScoopPkgInstall()
			return err
		})
	})

	// Declare button for save the configs
	saveButton := widget.NewButtonWithIcon("Save configs", SaveICON, func() {
		data.Choco_Install = editedTextChoco.Text
		data.Scoop_Install = editedTextScoop.Text
		// Set the choco checks
		data.Choco_Install_Configs.Verbose = chocoVerboseCheckBox.Checked
		data.Choco_Install_Configs.Force = chocoForceCheckBox1.Checked
		data.Choco_Install_Configs.Upgrade = chocoUpgradeCheckBox.Checked
		// Set the scoop checks
		data.Scoop_Install_Configs.Upgrade = scoopUpgradeCheckButton.Checked
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
