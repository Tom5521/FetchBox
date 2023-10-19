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
	"strings"

	win "github.com/Tom5521/CmdRunTools/windows"
	"github.com/gookit/color"
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
		checksudo, sudotype = checks.CheckSudo()
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
	err := win.Cmd(command).Run()
	if err != nil {
		color.Red.Println("Prossess Completed with errors.")
		return err
	}
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
	err = win.Cmd(command).Run()
	if err != nil {
		return err
	}
	return nil
}

func ScoopBucketInstall(bucket string) error {
	cmd := win.Cmd("")
	cmd.SetInput("git --version")
	if check := cmd.Run(); check != nil {
		color.Yellow.Println("Git is not installed... Installing git...")
		cmd.SetInput("scoop install git")
		err := cmd.Run()
		if err != nil {
			color.Red.Println("Error installing git...")
			return err
		}
		color.Green.Println("Git Installed!")
	}
	color.Yellow.Printf("Adding %v bucket...", bucket)
	cmd.SetInput(fmt.Sprintf("scoop bucket add %v", bucket))
	err := cmd.Run()
	if err != nil {
		color.Red.Printf("Error adding %v bucket", bucket)
		return err
	}
	color.Green.Printf("%v bucket added!", bucket)
	return nil
}
