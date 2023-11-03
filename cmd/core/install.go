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
		checksudo, sudotype = checks.CheckSudo()
		if !checksudo {
			return errors.New("sudo or gsudo not detected")
		}
	}
	command = fmt.Sprintf("%vchoco %v %v -y %v %v", sudotype, mode, force, verbose, data.Choco_Install)
	err := cmd.SetAndRun(command)
	if err != nil {
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
	var (
		mode string
	)
	if i.Scoop.Upgrade {
		mode = "upgrade"
	} else {
		mode = "install"
	}
	command := fmt.Sprintf("scoop %v %v", mode, data.Scoop_Install)
	err = cmd.SetAndRun(command)
	if err != nil {
		return err
	}
	return nil
}

func ScoopBucketInstall(bucket string) error {
	if check := cmd.SetAndRun("git --version"); check != nil {
		err := cmd.SetAndRun("scoop install git")
		if err != nil {
			return err
		}
	}
	err := cmd.SetAndRun(fmt.Sprintf("scoop bucket add %v", bucket))
	if err != nil {
		return err
	}
	return nil
}
