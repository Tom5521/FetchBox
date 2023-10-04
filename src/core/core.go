/*
 * Copyright (c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package core

import (
	"errors"
	"fmt"
	"os"
	"runtime"

	"github.com/Tom5521/MyGolangTools/commands"
	"github.com/Tom5521/MyGolangTools/file"
	"github.com/gookit/color"
	"gopkg.in/yaml.v3"
)

var (
	Version string = "v2.2"
	Red            = color.FgRed.Render
	//bgyellow        = color.BgYellow.Render
	Yellow         = color.FgYellow.Render
	linuxCH        = CheckOS()
	ConfigFilename = "FetchBox-conf.yml"
	sudotype       string
	sh             = func() commands.Sh {
		internal_sh := commands.Sh{}
		internal_sh.RunWithShell = false
		return internal_sh
	}()
	Root = func() string {
		dir, _ := os.Executable()
		return dir
	}()
)

type Yamlfile struct {
	Scoop_Install         string `yaml:"Scoop-Install"`
	Choco_Install         string `yaml:"Choco-Install"`
	Scoop_Uninstall       string `yaml:"Scoop-Uninstall"`
	Choco_Uninstall       string `yaml:"Choco-Uninstall"`
	Choco_Install_Configs struct {
		Verbose bool `yaml:"verbose"`
		Force   bool `yaml:"force"`
		Upgrade bool `yaml:"upgrade"`
	} `yaml:"Choco Install Configs"`
	Scoop_Install_Configs struct {
		Upgrade bool `yaml:"upgrade"`
	} `yaml:"Scoop Install Configs"`
	Choco_Uninstall_Configs struct {
		Verbose bool `yaml:"verbose"`
		Force   bool `yaml:"force"`
	} `yaml:"Choco uninstall Configs"`
}

func GetYamldata() Yamlfile {
	yamldata := Yamlfile{}
	if !CheckDir(ConfigFilename) {
		fmt.Printf(Red(ConfigFilename+" not found...") + Yellow("Creating a new one...\n"))
		NewYamlFile()
		if CheckDir(ConfigFilename) {
			color.Green.Println(ConfigFilename + " file created!!!")
		}
		return GetYamldata()
	}

	file, err := os.ReadFile(ConfigFilename)
	if err != nil {
		color.Red.Println("Error reading " + ConfigFilename)
	}
	err = yaml.Unmarshal(file, &yamldata)
	if err != nil {
		color.Red.Println("Error Unmarshalling the data")
	}
	return yamldata
}

func CheckOS() error {
	if runtime.GOOS == "linux" {
		return errors.New("you're on linux")
	}
	return nil
}

func NewYamlFile() {
	yamlData := Yamlfile{}
	data, err := yaml.Marshal(yamlData)
	if err != nil {
		return
	}
	err = file.ReWriteFile(ConfigFilename, string(data))
	if err != nil {
		color.Red.Println("Error writing the data in the new yml file")
		return
	}
}

var IsAdmin bool = func() bool {
	if _, err := os.Open("\\\\.\\PHYSICALDRIVE0"); err != nil {
		return false
	}
	return true
}()

func CheckDir(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}

func CheckSudo() (bool, string) {
	var (
		err1, err2 bool
		sudotype   string
	)
	color.Yellow.Println("Checking gsudo...")
	_, err := sh.Out("gsudo -v")
	if err == nil {
		color.Green.Println("gsudo detected!!!")
		err2 = true
		sudotype = "gsudo "
	} else {
		color.Yellow.Println("gsudo not detected...")
	}
	color.Yellow.Println("Checking sudo...")
	_, err = sh.Out("sudo /?")
	if err == nil {
		color.Green.Println("sudo detected!!!")
		err1 = true
		sudotype = "sudo "
	} else {
		color.Yellow.Println("sudo not detected...")
	}
	return err1 || err2, sudotype
}

//Install pkgmanagers functions

func InstallScoop() error {
	tempshell := sh
	tempshell.RunWithShell = true
	tempshell.Windows.PowerShell = true
	err1 := tempshell.Cmd("Set-ExecutionPolicy RemoteSigned -Scope CurrentUser")
	err2 := tempshell.Cmd("irm get.scoop.sh | iex")
	if err1 != nil || err2 != nil {
		return fmt.Errorf(
			fmt.Sprintf("Error installing scoop:\nCmd1:%v\nCmd2:%v", err1.Error(), err2.Error()),
		)
	}
	err := ScoopBucketInstall("extras")
	if err != nil {
		return err
	}
	return nil
}

func InstallChoco() error {
	tempshell := sh
	tempshell.Windows.PowerShell = true
	err := tempshell.Cmd(
		"Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))",
	)
	if err != nil {
		return err
	}
	return nil
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
