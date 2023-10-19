/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package main

import (
	"fmt"
	"os"

	"FetchBox/cmd/core"
	"FetchBox/internal/graph"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] != "dev" {
			TermMode() // Check the cmd args
			return
		}
	}
	graph.Init() // Initialize the graphical mode
}

func TermMode() {
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
			"\nCreated by "+core.Yellow("Angel(Tom5521)")+"\nUnder the "+core.Red("MIT")+" License",
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
