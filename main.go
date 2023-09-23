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
		TermMode() // Check the cmd args
		return
	}
	graph.Init() // Initialize the graphical mode
}

func TermMode() {
	switch os.Args[1] {
	case "scoop":
		src.ScoopPkgInstall()
	case "choco":
		src.ChocoPkgInstall()
	case "test":
		fmt.Println("true")
	case "newyamlfile":
		src.NewYamlFile()
	case "version":
		fmt.Println(
			"Windows-package-autoinstaller v"+src.Version,
			"\nCreated by "+src.Yellow("Angel(Tom5521)")+"\nUnder the "+src.Red("MIT")+" License",
		)
	case "addCustomBuckets":
		if len(os.Args) >= 3 {
			src.ScoopBucketInstall(os.Args[2])
		}
	}
}
