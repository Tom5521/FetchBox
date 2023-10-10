/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package core

import (
	"errors"
	"os"
	"runtime"
)

func CheckOS() error {
	if runtime.GOOS == "linux" {
		return errors.New("you're on linux")
	}
	return nil
}

func CheckDir(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}

func CheckSudo() (bool, string) {
	tmpsh := sh
	//tmpsh.CustomStd.Enable = true
	//tmpsh.Windows.RunWithPowerShell = true
	//tmpsh.SetWindowPSMode(commands.WinmodeHidden)

	var (
		check    bool
		sudotype string
	)

	if _, err := tmpsh.Out("gsudo -v"); err == nil {
		check = true
		sudotype = "gsudo "
		return check, sudotype
	}
	if _, err := tmpsh.Out("sudo /?"); err == nil {
		check = true
		sudotype = "sudo "
		return check, sudotype
	}
	if _, err := tmpsh.Out("sudo -h"); err == nil {
		check = true
		sudotype = "sudo "
		return check, sudotype
	}
	return false, ""
}

// Check if the pkg managers exists
func CheckScoop() bool {
	if _, err := sh.Out("scoop --version"); err != nil {
		return false
	}
	return true
}

func CheckChoco() bool {
	if _, err := sh.Out("choco --version"); err != nil {
		return false
	}
	return true
}

func CheckSudo_External() error {
	if check, _ := CheckSudo(); !check {
		return errors.New("sudo not detected")
	}
	return nil
}
