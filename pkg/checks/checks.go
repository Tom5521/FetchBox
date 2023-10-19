/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package checks

import (
	"errors"
	"os"
	"runtime"

	win "github.com/Tom5521/CmdRunTools/windows"
	"github.com/Tom5521/MyGolangTools/file"
)

func CheckOS() error {
	if runtime.GOOS == "linux" {
		return errors.New("you're on linux")
	}
	return nil
}

func CheckDir(dir string) bool {
	check, err := file.CheckFile(dir)
	if err != nil {
		return false
	}
	return check
}

func CheckSudo() (bool, string) {
	cmd := win.Cmd("")
	cmd.HideCmdWindow(true)
	cmd.SetInput("gsudo -v")
	if err := cmd.Run(); err == nil {
		return true, "gsudo "
	}
	cmd.SetInput("sudo /?")
	if err := cmd.Run(); err == nil {
		return true, "sudo "
	}
	cmd.SetInput("sudo -h")
	if err := cmd.Run(); err == nil {
		return true, "sudo "
	}
	return false, ""
}

// Check if the pkg managers exists
func CheckScoop() bool {
	cmd := win.Cmd("scoop --version")
	cmd.HideCmdWindow(true)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func CheckChoco() bool {
	cmd := win.Cmd("choco --version")
	cmd.HideCmdWindow(true)
	if err := cmd.Run(); err != nil {
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

func IsAdmin() bool {
	if _, err := os.Open("\\\\.\\PHYSICALDRIVE0"); err != nil {
		return false
	}
	return true
}
