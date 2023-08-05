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

	"github.com/Tom5521/Windows-package-autoinstaller/src"
)



func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "scoop":
			src.ScoopInstall()
		case "choco":
			src.ChocoInstall()
		case "test":
			fmt.Println("true")
		case "newyamlfile":
			src.NewYamlFile()
		case "version":
			fmt.Println("Windows-package-autoinstaller v",src.Version,"\nCreated by Angel(Tom5521)\nUnder the MIT License")
		}
		return
	}
	src.Clear()
	var option string
	red := color.FgRed.Render
	yellow := color.FgYellow.Render
	fmt.Printf(
		"Select an option\n1:Install Choco packages%v\n2:Install Scoop packages%v\n0:%v\n:",
		yellow("(requires administrator permissions)"),
		yellow("(requires not to use administrator)"),
		red("Exit"),
	)
	fmt.Scanln(&option)
	switch option {
	case "":
		src.Clear()
		main()
	case "0":
		src.End()
		os.Exit(0)
	case "1":
		src.ChocoInstall()
	case "2":
		src.ScoopInstall()
	}
}
