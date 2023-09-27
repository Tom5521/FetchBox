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
	"strings"

	"github.com/Tom5521/MyGolangTools/commands"
	"github.com/Tom5521/MyGolangTools/file"
	"github.com/gookit/color"
	"gopkg.in/yaml.v3"
)

var (
	Version string = "v2.2"
	Red            = color.FgRed.Render
	//bgyellow        = color.BgYellow.Render
	Yellow  = color.FgYellow.Render
	linuxCH = CheckOS()
	sh      = func() commands.Sh {
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
	Scoop        string `yaml:"scoop"`
	Choco        string `yaml:"choco"`
	ChocoConfigs struct {
		Verbose bool `yaml:"verbose"`
		Force   bool `yaml:"force"`
		Upgrade bool `yaml:"upgrade"`
	} `yaml:"Choco_configs"`
	ScoopConfigs struct {
		Upgrade bool `yaml:"upgrade"`
	} `yaml:"Scoop_Configs"`
}

func GetYamldata() Yamlfile {
	yamldata := Yamlfile{}
	if !CheckDir("packages.yml") {
		fmt.Printf(Red("packages.yml not found...") + Yellow("Creating a new one...\n"))
		NewYamlFile()
		if CheckDir("packages.yml") {
			color.Green.Println("packages.yml file created!!!")
		}
		return GetYamldata()
	}
	file, err := os.ReadFile("packages.yml")
	if err != nil {
		color.Red.Println("Error reading packages.yml")
	}
	err = yaml.Unmarshal(file, &yamldata)
	if err != nil {
		color.Red.Println("Error Unmarshalling the data")
	}
	return yamldata
}

func CheckOS() error {
	if runtime.GOOS == "linux" {
		return errors.New("You're on linux")
	} else {
		return nil
	}
}

func NewYamlFile() {
	yamlData := Yamlfile{}
	data, err := yaml.Marshal(yamlData)
	if err != nil {
		return
	}
	err = file.ReWriteFile("packages.yml", string(data))
	if err != nil {
		color.Red.Println("Error writing the data in the new yml file")
		return
	}
}

var IsAdmin bool = func() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false
	} else {
		return true
	}
}()

func CheckDir(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func ScoopBucketInstall(bucket string) error {
	if _, check := sh.Out("git --version"); check != nil {
		color.Yellow.Println("Git is not installed... Installing git...")
		err := sh.Cmd("scoop install git")
		if err != nil {
			color.Red.Println("Error installing git...")
			return err
		}
		color.Green.Println("Git Installed!")
	}
	color.Yellow.Printf("Adding %v bucket...", bucket)
	err := sh.Cmd(fmt.Sprintf("scoop bucket add %v", bucket))
	if err != nil {
		color.Red.Printf("Error adding %v bucket", bucket)
		return err
	}
	color.Green.Printf("%v bucket added!", bucket)
	return nil
}
func ScoopPkgInstall(optArsg ...string) error {
	var err error
	data := GetYamldata()
	if linuxCH != nil {
		return linuxCH
	}
	if data.Scoop == "" {
		color.Red.Println("No package for scoop written in packages.yml")
		return errors.New("no package for scoop written in packages.yml")
	}
	if IsAdmin {
		return errors.New("Scoop must be run without administrator permissions")
	}
	if strings.Contains(data.Scoop, "np") {
		err := ScoopBucketInstall("nonportable")
		if err != nil {
			return err
		}
	}
	fmt.Printf(Yellow("Installing with scoop ")+"%v\n", data.Scoop)
	if optArsg[0] != "upgrade" {
		err = sh.Cmd(fmt.Sprintf("scoop install %v %v", strings.Join(optArsg, " "), data.Scoop))
	} else {
		err = sh.Cmd(fmt.Sprintf("scoop upgrade %v %v", strings.Join(optArsg[0:], " "), data.Scoop))
	}
	if err != nil {
		color.Red.Println("Prossess Completed with errors.")
		return err
	} else {
		color.Green.Println("Prosess Completed without errors!!!")
		return nil
	}
}

func ChocoPkgInstall(args ...string) error {
	var (
		checksudo         bool
		sudotype, command string
		data              = GetYamldata()
	)
	if linuxCH != nil {
		return linuxCH
	}
	if data.Choco == "" {
		color.Red.Println("No package for choco written in packages.yml")
		return errors.New("No package for choco written in packages.yml")
	}

	if !IsAdmin {
		color.Red.Print("Running without administrator permissions... ")
		color.Yellow.Println("Checking sudo or gsudo...")
		checksudo, sudotype = CheckSudo()
		if !checksudo {
			color.Red.Println("sudo or gsudo not detected.")
			return errors.New("sudo or gsudo not detected.")
		}
	} else if checksudo {
		color.Yellow.Println("Running as administrator")
	}
	fmt.Printf(Yellow("Installing with choco ")+"%v\n", data.Choco)
	if checksudo {
		color.Yellow.Println("Using " + sudotype)
	}
	if args[2] == "upgrade" {
		command = fmt.Sprintf("%vchoco upgrade -y %v %v", sudotype, fmt.Sprintf("%v %v", args[1], args[0]), data.Choco)
	} else {
		command = fmt.Sprintf("%vchoco install -y %v %v", sudotype, strings.Join(args, " "), data.Choco)
	}
	err := sh.Cmd(command)
	if err != nil {
		color.Red.Println("Prossess Completed with errors.")
		return err
	} else {
		color.Green.Println("Prosess Completed without errors!!!")
		return nil
	}
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
		return errors.New(fmt.Sprintf("Error installing scoop:\nCmd1:%v\nCmd2:%v", err1.Error(), err2.Error()))
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
	_, err := sh.Out("scoop --version")
	if err != nil {
		return false
	}
	return true
}

func CheckChoco() bool {
	_, err := sh.Out("choco --version")
	if err != nil {
		return false
	}
	return true
}

func CheckSudo_External() error {
	if check, _ := CheckSudo(); !check {
		return errors.New("Sudo not detected!")
	}
	return nil
}
