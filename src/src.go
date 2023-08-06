/*
 * Copyright (c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package src

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Tom5521/MyGolangTools/commands"
	"github.com/gookit/color"
	"gopkg.in/yaml.v3"
)

var (
	Version    string = "1.2"
	red               = color.FgRed.Render
	bgyellow          = color.BgYellow.Render
	yellow            = color.FgYellow.Render
	ConfigData        = getYamldata()
	Root       string = func() string {
		binpath, _ := filepath.Abs(os.Args[0])
		return filepath.Dir(binpath)
	}()
)

func getYamldata() yamlfile {
	yamldata := yamlfile{}
	if !CheckDir("packages.yml") {
		fmt.Printf(red("packages.yml not found...") + yellow("Creating a new one...\n"))
		NewYamlFile()
		if CheckDir("packages.yml") {
			color.Green.Println("packages.yml file created!!!")
		}
		End()
		os.Exit(0)
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
	sh := commands.Sh{}
	_, err := sh.Out("net session")
	if err != nil {
		return false
	} else {
		return true
	}
}()

func CheckPackageManagers(tested string) {
	tempshell := commands.Sh{}
	tempshell.Windows.PowerShell = true
	install_choco := func() {
		err := tempshell.Cmd(
			"Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))",
		)
		if err != nil {
			color.Red.Println("Error installing choco")
			End()
			return
		}
	}
	install_scoop := func() {
		err1 := tempshell.Cmd("Set-ExecutionPolicy RemoteSigned -Scope CurrentUser")
		err2 := tempshell.Cmd("irm get.scoop.sh | iex")
		if err1 != nil || err2 != nil {
			color.Red.Println("Error installing scoop")
			End()
			return
		}
	}
	if strings.Contains(tested, "choco") {
		color.Yellow.Println("Checking choco...")
		_, err := sh.Out("choco --version")
		if err != nil {
			fmt.Printf("%v... %v...\n", red("Choco not detected"), yellow("Trying to install choco"))
			install_choco()
		} else {
			color.Green.Println("Choco is Installed!")
		}
	}
	if strings.Contains(tested, "scoop") {
		color.Yellow.Println("Checking Scoop...")
		_, err := sh.Out("scoop --version")
		if err != nil {

			fmt.Printf("%v... %v...\n", red("Scoop not detected"), yellow("Trying to install Scoop"))
			install_scoop()
		} else {
			color.Green.Println("Scoop is Installed!")
		}
	}
}

func CheckDir(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func End() {
	if len(os.Args) > 2 {
		if os.Args[2] == "noend" {
			return
		}
	}
	fmt.Println("Press " + bgyellow("enter") + " to exit...")
	fmt.Scanln()
}

var sh commands.Sh = commands.Sh{}

func ScoopInstall() {
	CheckPackageManagers("scoop")
	if IsAdmin {
		var option string
		color.Yellow.Println("You really want to run scoop with administrator permissions? y/N")
		fmt.Scanln(&option)
		option = strings.ToUpper(option)
		if option != "Y" {
			color.Red.Println("Scoop will not run with administrator permissions.")
			End()
			return
		}
	}
	if ConfigData.Scoop == "" {
		color.Red.Println("No package for scoop written in packages.yml")
		End()
		return
	}
	fmt.Printf(yellow("Installing with scoop ")+"%v\n", ConfigData.Scoop)
	err := sh.Cmd("scoop install " + ConfigData.Scoop)
	if err != nil {
		color.Red.Println("Prossess Completed with errors.")
	} else {
		color.Green.Println("Prosess Completed without errors!!!")
	}
	End()
}

func ChocoInstall() {
	var (
		checksudo bool
		sudotype  string
	)
	if !IsAdmin {
		color.Red.Print("Running without administrator permissions... ")
		color.Yellow.Println("Checking sudo or gsudo...")
		checksudo, sudotype = CheckSudo()
		if !checksudo {
			color.Red.Println("sudo or gsudo not detected.")
			End()
			return
		}
	} else if checksudo {
		color.Yellow.Println("Running as administrator")
	}
	if ConfigData.Choco == "" {
		color.Red.Println("No package for choco written in packages.yml")
		End()
		return
	}
	CheckPackageManagers("choco")
	fmt.Printf(yellow("Installing with choco ")+"%v\n", ConfigData.Choco)
	if checksudo {
		color.Yellow.Println("Using " + sudotype)
	}
	command := fmt.Sprintf("%vchoco install -y %v", sudotype, ConfigData.Choco)
	err := sh.Cmd(command)
	if err != nil {
		color.Red.Println("Prossess Completed with errors.")
	} else {
		color.Green.Println("Prosess Completed without errors!!!")
	}
	End()
}

func Clear() {
	err := sh.Cmd("cls")
	if err != nil {
		color.Red.Println("Error Cleaning the terminal")
	}
}

func CheckSudo() (bool, string) {
	var err1, err2 bool
	var sudotype string
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
