/*
 * Copyright (c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package main

import (
	"fmt"
	"os"

	"github.com/gookit/color"

	"Windows-package-autoinstaller/base"
)

var sh base.Sh

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "scoop":
			base.ScoopInstall()
		case "choco":
			base.ChocoInstall()
		}
		return
	}
	sh.Cmd("cls")
	var option string
	red := color.FgRed.Render
	yellow := color.FgYellow.Render
	fmt.Printf("Select an option\n1:Install Choco packages%v\n2:Install Scoop packages%v\n0:%v\n:", yellow("(requires administrator permissions)"), yellow("(requires not to use administrator)"), red("Exit"))
	fmt.Scanln(&option)
	switch option {
	case "":
		base.Clear()
		main()
	case "0":
		base.End()
		os.Exit(0)
	case "1":
		base.ChocoInstall()
	case "2":
		base.ScoopInstall()
	}
}
