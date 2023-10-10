package core

import (
	"os"

	"github.com/Tom5521/MyGolangTools/commands"
	"github.com/gookit/color"
)

var (
	Version string = "v2.2"
	Red            = color.FgRed.Render
	//bgyellow        = color.BgYellow.Render
	Yellow         = color.FgYellow.Render
	linuxCH        = CheckOS()
	ConfigFilename = "FetchBox-conf.yml"
	sudotype       string
	sh             = commands.Sh{}
	Root           = func() string {
		dir, _ := os.Executable()
		return dir
	}()
	IsAdmin bool = func() bool {
		if _, err := os.Open("\\\\.\\PHYSICALDRIVE0"); err != nil {
			return false
		}
		return true
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
