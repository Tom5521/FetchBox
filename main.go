/*
 * Copyright (c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package main

import (
	"fmt"
	"os"

	src "github.com/Tom5521/Windows-package-autoinstaller/src/core"
	"github.com/Tom5521/Windows-package-autoinstaller/src/graph"
)

func main() {
	if len(os.Args) > 1 {
		TermMode()
		return
	}
	graph.Init()
}

func TermMode() {
	switch os.Args[1] {
	case "tui":
		TUIMode()
	case "scoop":
		src.ScoopInstall()
	case "choco":
		src.ChocoInstall()
	case "test":
		fmt.Println("true")
	case "newyamlfile":
		src.NewYamlFile()
	case "version":
		fmt.Println("Windows-package-autoinstaller v"+src.Version, "\nCreated by "+src.Yellow("Angel(Tom5521)")+"\nUnder the "+src.Red("MIT")+" License")
	case "addCustomBuckets":
		if len(os.Args) >= 3 {
			src.ScoopBucketInstall(os.Args[2])
		}
	}
}

func TUIMode() {
	src.Clear()
	var option string
	fmt.Printf(
		"Select an option\n1:Install Choco packages%v\n2:Install Scoop packages%v\n0:%v\n:",
		src.Yellow("(requires administrator permissions)"),
		src.Yellow("(requires not to use administrator)"),
		src.Red("Exit"),
	)
	fmt.Scanln(&option)
	switch option {
	case "":
		src.Clear()
		TUIMode()
	case "0":
		src.End()
		os.Exit(0)
	case "1":
		src.ChocoInstall()
	case "2":
		src.ScoopInstall()
	}
}
