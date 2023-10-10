package core

import (
	"errors"
	"fmt"
	"os"

	"github.com/Tom5521/MyGolangTools/file"
	"github.com/gookit/color"
	"gopkg.in/yaml.v3"
)

func NewYamlFile() {
	yamlData := Yamlfile{}
	data, err := yaml.Marshal(yamlData)
	if err != nil {
		return
	}
	err = file.ReWriteFile(ConfigFilename, string(data))
	if err != nil {
		return
	}
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

//Install pkgmanagers functions

func InstallScoop() error {
	sh := sh
	sh.Windows.RunWithPowerShell = true
	err1 := sh.Cmd("Set-ExecutionPolicy RemoteSigned -Scope CurrentUser")
	err2 := sh.Cmd("irm get.scoop.sh | iex")
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
	if !IsAdmin {
		return errors.New("To install choco, run this as administrator")
	}
	sh := sh
	sh.Windows.RunWithPowerShell = true
	command := "Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))"
	err := sh.Cmd(command)
	if err != nil {
		return err
	}
	return nil
}
