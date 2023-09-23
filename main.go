/*
 * Copyright (c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package main

import (
	"fmt"
	"os"

	"github.com/Tom5521/Windows-package-autoinstaller/src/core"
	"github.com/Tom5521/Windows-package-autoinstaller/src/graph"
)

func main() {
	if len(os.Args) > 1 {
		TermMode() // Check the cmd args
		return
	}
	graph.Init() // Initialize the graphical mode
}

func TermMode() {
	switch os.Args[1] {
	case "scoop":
		core.ScoopPkgInstall()
	case "choco":
		core.ChocoPkgInstall()
	case "newyamlfile":
		core.NewYamlFile() // Create a new yaml file
	case "version":
		fmt.Println(
			"Windows-package-autoinstaller v"+core.Version,
			"\nCreated by "+core.Yellow("Angel(Tom5521)")+"\nUnder the "+core.Red("MIT")+" License",
		)
	case "addCustomBuckets":
		if len(os.Args) >= 3 {
			core.ScoopBucketInstall(os.Args[2])
		}
	}
}
