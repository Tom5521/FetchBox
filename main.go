/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package main

import (
	"FetchBox/internal/cli"
	"FetchBox/internal/graph"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] != "dev" {
			cli.Init() // Check the cmd args
			return
		}
	}
	graph.Init() // Initialize the graphical mode
}
