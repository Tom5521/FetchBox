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
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Tom5521/MyGolangTools/commands"
	"github.com/gookit/color"
	"gopkg.in/yaml.v3"
)

var (
	Version string = "2.0"
	Red            = color.FgRed.Render
	//bgyellow        = color.BgYellow.Render
	Yellow         = color.FgYellow.Render
	linuxCH        = CheckOS()
	Root    string = func() string {
		binpath, _ := filepath.Abs(os.Args[0])
		return filepath.Dir(binpath)
	}()
)

func getYamldata() yamlfile {
	yamldata := yamlfile{}
	if !CheckDir("packages.yml") {
		fmt.Printf(Red("packages.yml not found...") + Yellow("Creating a new one...\n"))
		NewYamlFile()
		if CheckDir("packages.yml") {
			color.Green.Println("packages.yml file created!!!")
		}
		return getYamldata()
	}
	file, err := os.ReadFile("packages.yml")
	if err != nil {
		color.Red.Println("Error reading packages.yml")
		End()
	}
	err = yaml.Unmarshal(file, &yamldata)
	if err != nil {
		color.Red.Println("Error Unmarshalling the data")
		End()
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
	yamlstruct := yamlfile{}
	file, err := os.Create("packages.yml")
	if err != nil {
		color.Red.Println("Error creating packages.yml")
		End()
		return
	}
	defer file.Close()
	data, err := yaml.Marshal(yamlstruct)
	if err != nil {
		color.Red.Println("Error Marshalling packages file")
		End()
		return
	}
	_, err = file.WriteString(string(data))
	if err != nil {
		color.Red.Println("Error writing the data in the new yml file")
		End()
		return
	}
}

type yamlfile struct {
	Scoop string `yaml:"scoop"`
	Choco string `yaml:"choco"`
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

func End() {
	fmt.Println("Process Completed.")
}

var sh commands.Sh = commands.Sh{}

func ScoopBucketInstall(bucket string) {
	if _, check := sh.Out("git --version"); check != nil {
		color.Yellow.Println("Git is not installed... Installing git...")
		err := sh.Cmd("scoop install git")
		if err != nil {
			color.Red.Println("Error installing git...")
			End()
			return
		}
		color.Green.Println("Git Installed!")
	}
	color.Yellow.Printf("Adding %v bucket...", bucket)
	err := sh.Cmd(fmt.Sprintf("scoop bucket add %v", bucket))
	if err != nil {
		color.Red.Printf("Error adding %v bucket", bucket)
	}
	color.Green.Printf("%v bucket added!", bucket)
}
func ScoopPkgInstall() error {
	data := getYamldata()
	if linuxCH != nil {
		return linuxCH
	}
	if data.Scoop == "" {
		color.Red.Println("No package for scoop written in packages.yml")
		End()
		return errors.New("no package for scoop written in packages.yml")
	}
	if check := CheckScoop(); !check {
		err := InstallScoop()
		if err != nil {
			return err
		}
	}
	if IsAdmin {
		return errors.New("Scoop must be run without administrator permissions")
	}
	if strings.Contains(data.Scoop, "np") {
		ScoopBucketInstall("nonportable")
	}
	fmt.Printf(Yellow("Installing with scoop ")+"%v\n", data.Scoop)
	err := sh.Cmd("scoop install " + data.Scoop)
	if err != nil {
		color.Red.Println("Prossess Completed with errors.")
		return err
	} else {
		color.Green.Println("Prosess Completed without errors!!!")
		return nil
	}
}

func ChocoPkgInstall() error {
	var (
		checksudo bool
		sudotype  string
		data      = getYamldata()
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
			End()
			return errors.New("sudo or gsudo not detected.")
		}
	} else if checksudo {
		color.Yellow.Println("Running as administrator")
	}
	if check := CheckChoco(); !check {
		err := InstallChoco()
		if err != nil {
			return err
		}
	}
	fmt.Printf(Yellow("Installing with choco ")+"%v\n", data.Choco)
	if checksudo {
		color.Yellow.Println("Using " + sudotype)
	}
	command := fmt.Sprintf("%vchoco install -y %v", sudotype, data.Choco)
	err := sh.Cmd(command)
	if err != nil {
		color.Red.Println("Prossess Completed with errors.")
		return err
	} else {
		color.Green.Println("Prosess Completed without errors!!!")
		return nil
	}
}

func Clear() {
	tempcmd := commands.Sh{}
	tempcmd.Windows.PowerShell = true
	err := tempcmd.Cmd("clear")
	if err != nil {
		color.Red.Println("Error Cleaning the terminal")
	}
}

func CheckSudo() (bool, string) {
	var (
		err1, err2 bool
		sudotype   string
	)
	color.Yellow.Println("Checking gsudo...")
	_, err := sh.Out("gsudo echo ...")
	if err == nil {
		color.Green.Println("gsudo detected!!!")
		err2 = true
		sudotype = "gsudo "
	} else {
		color.Yellow.Println("gsudo not detected...")
	}
	color.Yellow.Println("Checking sudo...")
	_, err = sh.Out("sudo echo ...")
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
	tempshell.Windows.PowerShell = true
	err1 := tempshell.Cmd("Set-Exe*cutionPolicy RemoteSigned -Scope CurrentUser")
	err2 := tempshell.Cmd("irm get.scoop.sh | iex")
	ScoopBucketInstall("extras")
	if err1 != nil || err2 != nil {
		return errors.New("Error Installing scoop")
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
