/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package graph

import (
	"FetchBox/cmd/core"
	"FetchBox/pkg/checks"
	"FetchBox/pkg/windows"
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
	if err := checks.CheckOS(); err != nil { // Check the OS for show error in linux
		windows.ErrWin(app, err.Error())
		return err
	}

	if editedTextChoco == "" || editedTextChoco == "<nil>" { // Check the text for nil values
		err := errors.New("choco package list is empty")
		windows.ErrWin(app, err.Error())
		return err
	}
	if !core.IsAdmin { // Check admin permissions
		err := checks.CheckSudo_External()
		err1 := errors.New("sudo not detected!\nRestart the program with administrator permissions")
		if err != nil {
			windows.ErrWin(
				app,
				err1.Error(),
			)
			return err1
		}
	}
	if !checks.CheckChoco() { // Check Choco package manager
		windows.InstallPkgManagerWin(app, "Choco", core.InstallChoco)
		return errors.New("choco is'nt installed")
	}
	return nil
}

func basicScoopCheck(app fyne.App, editedTextScoop string) error {
	if err := checks.CheckOS(); err != nil { // Check the OS
		windows.ErrWin(app, err.Error())
		return err
	}
	// Check for nil string values
	if editedTextScoop == "" || editedTextScoop == "<nil>" {
		err := errors.New("scoop package list is empty")
		windows.ErrWin(app, err.Error())
		return err
	}
	// Check if scoop is installed
	if !checks.CheckScoop() {
		windows.InstallPkgManagerWin(app, "Scoop", core.InstallScoop)
		return nil
	}
	return nil

}
