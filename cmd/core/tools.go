/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package core

import (
	"FetchBox/pkg/checks"
	"errors"
	"fmt"
	"os"

	"github.com/Tom5521/CmdRunTools/command"
	"github.com/Tom5521/MyGolangTools/file"
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
	if !checks.CheckDir(ConfigFilename) {
		NewYamlFile()
		return GetYamldata()
	}

	file, _ := os.ReadFile(ConfigFilename)
	yaml.Unmarshal(file, &yamldata)
	return yamldata
}

//Install pkgmanagers functions

func InstallScoop() error {
	cmd := command.Cmd{}
	cmd.RunWithPS(true)
	err1 := cmd.SetAndRun("Set-ExecutionPolicy RemoteSigned -Scope CurrentUser")
	err2 := cmd.SetAndRun("irm get.scoop.sh | iex")

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

	command := "Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))"
	cmd.SetInput(command)
	cmd.RunWithPS(true)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
