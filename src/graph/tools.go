package graph

import (
	"FetchBox/src/core"
	"errors"
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"gopkg.in/yaml.v3"
)

func saveText() {
	data, err := yaml.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling YAML:", err)
		return
	}

	err = os.WriteFile(core.ConfigFilename, data, 0644)
	if err != nil {
		fmt.Println("Error writing YAML file:", err)
	}
}

func basicChocoCheck(app fyne.App, editedTextChoco string) error {
	if err := core.CheckOS(); err != nil { // Check the OS for show error in linux
		ErrWin(app, err, nil)
		return err
	}

	if editedTextChoco == "" || editedTextChoco == "<nil>" { // Check the text for nil values
		err := errors.New("choco package list is empty")
		ErrWin(app, err, nil)
		return err
	}
	if !core.IsAdmin { // Check admin permissions
		err := core.CheckSudo_External()
		err1 := errors.New("sudo not detected!\nRestart the program with administrator permissions")
		if err != nil {
			ErrWin(
				app,
				err1,
				nil,
			)
			return err1
		}
	}
	if !core.CheckChoco() { // Check Choco package manager
		InstallPkgManagerWin(app, "Choco", core.InstallChoco)
		return errors.New("choco is'nt installed")
	}
	return nil
}

func basicScoopCheck(app fyne.App, editedTextScoop string) error {
	if err := core.CheckOS(); err != nil { // Check the OS
		ErrWin(app, err, nil)
		return err
	}
	// Check for nil string values
	if editedTextScoop == "" || editedTextScoop == "<nil>" {
		err := errors.New("scoop package list is empty")
		ErrWin(app, err, nil)
		return err
	}
	// Check if scoop is installed
	if !core.CheckScoop() {
		InstallPkgManagerWin(app, "Scoop", core.InstallScoop)
		return nil
	}
	return nil

}
