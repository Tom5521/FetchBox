/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package core

import (
	"FetchBox/pkg/checks"
	"os"

	"github.com/Tom5521/CmdRunTools/command"
)

var (
	Version string = "v2.3.2"
	cmd            = command.Cmd{}
	//Red            = color.FgRed.Render
	//bgyellow        = color.BgYellow.Render
	//Yellow         = color.FgYellow.Render
	linuxCH        = checks.CheckOS()
	ConfigFilename = "FetchBox-conf.yml"
	sudotype       string
	Root           = func() string {
		dir, _ := os.Executable()
		return dir
	}()
	IsAdmin = checks.IsAdmin()
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
