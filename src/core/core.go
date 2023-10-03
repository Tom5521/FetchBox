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
	checkVarBool        = true
	Version      string = "v2.2"
	Red                 = color.FgRed.Render
	//bgyellow        = color.BgYellow.Render
	Yellow         = color.FgYellow.Render
	linuxCH        = CheckOS()
	ConfigFilename = "wpa-config.yml"
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

type Install struct {
	Choco struct {
		Verbose bool
		Force   bool
		Upgrade bool
	}
	Scoop struct {
		Upgrade bool
	}
}

type Uninstall struct {
	Choco struct {
		Verbose bool
		Force   bool
	}
}

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
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
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
func (i Install) ScoopPkgInstall() error {
	var err error
	data := GetYamldata()
	if linuxCH != nil {
		return linuxCH
	}
	if data.Scoop_Install == "" {
		return errors.New("no package for scoop written in " + ConfigFilename)
	}
	if IsAdmin {
		return errors.New("scoop must be run without administrator permissions")
	}
	if strings.Contains(data.Scoop_Install, "np") {
		err := ScoopBucketInstall("nonportable")
		if err != nil {
			return err
		}
	}
	fmt.Printf(Yellow("Installing with scoop ")+"%v\n", data.Scoop_Install)
	var (
		mode string
	)
	if i.Scoop.Upgrade {
		mode = "upgrade"
	} else {
		mode = "install"
	}
	command := fmt.Sprintf("scoop %v %v", mode, data.Scoop_Install)
	err = sh.Cmd(command)
	if err != nil {
		return err
	}
	return nil
}

func (i Install) ChocoPkgInstall() error {
	var (
		checksudo                               bool
		sudotype, command, mode, force, verbose string
		data                                    = GetYamldata()
	)

	if i.Choco.Upgrade {
		mode = "upgrade"
	} else {
		mode = "install"
	}

	if i.Choco.Force {
		force = "-f"
	}

	if i.Choco.Verbose {
		verbose = "-v"
	}

	if linuxCH != nil {
		return linuxCH
	}
	if data.Choco_Install == "" {
		return errors.New("No package for choco written in " + ConfigFilename)
	}

	if !IsAdmin {
		color.Red.Print("Running without administrator permissions... ")
		color.Yellow.Println("Checking sudo or gsudo...")
		checksudo, sudotype = CheckSudo()
		if !checksudo {
			return errors.New("sudo or gsudo not detected")
		}
	} else if checksudo {
		color.Yellow.Println("Running as administrator")
	}
	fmt.Printf(Yellow("Installing with choco ")+"%v\n", data.Choco_Install)
	if checksudo {
		color.Yellow.Println("Using " + sudotype)
	}
	command = fmt.Sprintf("%vchoco %v %v -y %v %v", sudotype, mode, force, verbose, data.Choco_Install)
	err := sh.Cmd(command)
	if err != nil {
		color.Red.Println("Prossess Completed with errors.")
		return err
	}
	return nil
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
	_, err := sh.Out("scoop --version")
	if err != nil {
		return !checkVarBool
	}
	return checkVarBool
}

func CheckChoco() bool {
	_, err := sh.Out("choco --version")
	if err != nil {
		return !checkVarBool
	}
	return checkVarBool
}

func CheckSudo_External() error {
	if check, _ := CheckSudo(); !check {
		return errors.New("sudo not detected")
	}
	return nil
}

func (u Uninstall) UninstallScoopPkgs() error {
	var data = GetYamldata()
	command := fmt.Sprintf("scoop uninstall %v", data.Scoop_Uninstall)
	err := sh.Cmd(command)
	if err != nil {
		return err
	}
	return nil
}

func (u Uninstall) UninstallChocoPkgs() error {
	var (
		data      = GetYamldata()
		force     string
		verbose   string
		checksudo bool
	)
	if u.Choco.Force {
		force = "-f"
	}
	if u.Choco.Verbose {
		verbose = "-v"
	}
	if !IsAdmin {
		checksudo, sudotype = CheckSudo()
		if !checksudo {
			return errors.New("sudo or gsudo not detected")
		}
	}

	command := fmt.Sprintf("%vchoco uninstall -y %v %v %v ", sudotype, force, verbose, data.Choco_Uninstall)
	err := sh.Cmd(command)
	if err != nil {
		return err
	}
	return nil
}
