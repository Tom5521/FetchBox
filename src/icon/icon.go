/*
 * Copyright (c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package icon

import (
	"Windows-package-autoinstaller/src/bundled"
	"os"

	"fyne.io/fyne/v2"
)

var (
	// Dark Icons
	DevICON_Dark      fyne.Resource = bundled.Dev_Dark
	DownloadICON_Dark fyne.Resource = bundled.Download_Dark
	ErrorICON_Dark    fyne.Resource = bundled.Error_Dark
	InstallICON_Dark  fyne.Resource = bundled.Install_Dark
	SaveICON_Dark     fyne.Resource = bundled.Save_Dark
	RestartICON_Dark  fyne.Resource = bundled.Restart_Dark
	InfoICON_Dark     fyne.Resource = bundled.Info_Dark

	// Light Icons
	DevICON_Light      fyne.Resource = bundled.Dev_Light
	DownloadICON_Light fyne.Resource = bundled.Download_Light
	ErrorICON_Light    fyne.Resource = bundled.Error_Light
	InstallICON_Light  fyne.Resource = bundled.Install_Light
	SaveICON_Light     fyne.Resource = bundled.Save_Light
	RestartICON_Light  fyne.Resource = bundled.Restart_Light
	InfoICON_Light     fyne.Resource = bundled.Info_Light

	// Themed Icons
	DevICON      fyne.Resource
	DownloadICON fyne.Resource
	ErrorICON    fyne.Resource
	InstallICON  fyne.Resource
	SaveICON     fyne.Resource
	RestartICON  fyne.Resource
	InfoICON     fyne.Resource

	// No-Theme Icons
	AppICON         fyne.Resource = bundled.App
	PlaceholderICON fyne.Resource = bundled.Placeholder
)

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
