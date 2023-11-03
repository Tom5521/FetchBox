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
)

type Uninstall struct {
	Choco struct {
		Verbose bool
		Force   bool
	}
}

func (u Uninstall) UninstallScoopPkgs() error {
	var data = GetYamldata()
	command := fmt.Sprintf("scoop uninstall %v", data.Scoop_Uninstall)
	err := cmd.SetAndRun(command)
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
		force = " -f "
	}
	if u.Choco.Verbose {
		verbose = "-v "
	}
	if !IsAdmin {
		checksudo, sudotype = checks.CheckSudo()
		if !checksudo {
			return errors.New("sudo or gsudo not detected")
		}
	}

	command := fmt.Sprintf("%vchoco uninstall -y%v%v %v ", sudotype, force, verbose, data.Choco_Uninstall)
	err := cmd.SetAndRun(command)
	if err != nil {
		return err
	}
	return nil
}
