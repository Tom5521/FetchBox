/*
 * Copyright (c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package icon

import (
	"errors"
	"os"

	"fyne.io/fyne/v2"
)

var (
	// Dark Icons
	DevICON_Dark      fyne.Resource
	DownloadICON_Dark fyne.Resource
	ErrorICON_Dark    fyne.Resource
	InstallICON_Dark  fyne.Resource
	SaveICON_Dark     fyne.Resource
	RestartICON_Dark  fyne.Resource
	InfoICON_Dark     fyne.Resource

	// Light Icons
	DevICON_Light      fyne.Resource
	DownloadICON_Light fyne.Resource
	ErrorICON_Light    fyne.Resource
	InstallICON_Light  fyne.Resource
	SaveICON_Light     fyne.Resource
	RestartICON_Light  fyne.Resource
	InfoICON_Light     fyne.Resource

	// Themed Icons
	DevICON      fyne.Resource
	DownloadICON fyne.Resource
	ErrorICON    fyne.Resource
	InstallICON  fyne.Resource
	SaveICON     fyne.Resource
	RestartICON  fyne.Resource
	InfoICON     fyne.Resource

	// No-Theme Icons
	AppICON         fyne.Resource
	PlaceholderICON fyne.Resource
)

func LoadResource(app fyne.App, Filename string) (fyne.Resource, error) {
	ret, err := fyne.LoadResourceFromPath(Filename)
	if err != nil {
		return PlaceholderICON, errors.New("Error loading " + Filename + "resource")
	}
	return ret, nil
}

func LoadIcons(app fyne.App, errWin func(fyne.App, error, fyne.Window)) {
	var err error

	PlaceholderICON, err = LoadResource(app, "./Assets/Placeholder.png")
	if err != nil {
		errWin(app, err, nil)
	}
	AppICON, err = LoadResource(app, "./Assets/Icon.png")
	if err != nil {
		errWin(app, err, nil)
	}

	icons := []string{"Dev.png", "Install.png", "Download.png", "Error.png", "Save.png", "Restart.png", "Info.png"}

	// Load dark icons
	for _, name := range icons {
		path := "./Assets/Icons-Dark/" + name
		icon, err := LoadResource(app, path)
		if err != nil {
			errWin(app, err, nil)
		}
		switch name {
		case "Dev.png":
			DevICON_Dark = icon
		case "Install.png":
			InstallICON_Dark = icon
		case "Download.png":
			DownloadICON_Dark = icon
		case "Error.png":
			ErrorICON_Dark = icon
		case "Save.png":
			SaveICON_Dark = icon
		case "Restart.png":
			RestartICON_Dark = icon
		case "Info.png":
			InfoICON_Dark = icon
		}
	}

	// Load light icons
	for _, name := range icons {
		path := "./Assets/Icons-Light/" + name
		icon, err := LoadResource(app, path)
		if err != nil {
			errWin(app, err, nil)
		}
		switch name {
		case "Dev.png":
			DevICON_Light = icon
		case "Install.png":
			InstallICON_Light = icon
		case "Download.png":
			DownloadICON_Light = icon
		case "Error.png":
			ErrorICON_Light = icon
		case "Save.png":
			SaveICON_Light = icon
		case "Restart.png":
			RestartICON_Light = icon
		case "Info.png":
			InfoICON_Light = icon
		}
	}
}

func SetThemeIcons(app fyne.App, errWin func(fyne.App, error, fyne.Window)) {
	if os.Getenv("FYNE_THEME") == "light" {
		DevICON = DevICON_Light
		InstallICON = InstallICON_Light
		DownloadICON = DownloadICON_Light
		ErrorICON = ErrorICON_Light
		SaveICON = SaveICON_Light
		RestartICON = RestartICON_Light
		InfoICON = InfoICON_Light
	} else {
		DevICON = DevICON_Dark
		InstallICON = InstallICON_Dark
		DownloadICON = DownloadICON_Dark
		ErrorICON = ErrorICON_Dark
		SaveICON = SaveICON_Dark
		RestartICON = RestartICON_Dark
		InfoICON = InfoICON_Dark
	}
}
