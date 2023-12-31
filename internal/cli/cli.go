package cli

import (
	"FetchBox/cmd/core"
	"fmt"
	"os"

	"github.com/gookit/color"
)

func Init() {
	var (
		install = core.Install{}
	)
	switch os.Args[1] {
	case "scoop":
		err := install.ScoopPkgInstall()
		if err != nil {
			fmt.Println(err.Error())
		}
	case "choco":
		err := install.ChocoPkgInstall()
		if err != nil {
			fmt.Println(err.Error())
		}
	case "newyamlfile":
		core.NewYamlFile() // Create a new yaml file
	case "version":
		fmt.Println(
			"Windows-package-autoinstaller v"+core.Version,
			"\nCreated by "+color.Yellow.Render("Angel(Tom5521)")+"\nUnder the "+color.Red.Render("MIT")+" License",
		)
	case "addCustomBuckets":
		if len(os.Args) >= 3 {
			err := core.ScoopBucketInstall(os.Args[2])
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
