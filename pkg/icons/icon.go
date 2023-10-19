/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package icon

import (
	"os"

	"fyne.io/fyne/v2"
)

var (
	// Dark Icons
	DevICON_Dark       fyne.Resource = Dev_Dark
	DownloadICON_Dark  fyne.Resource = Download_Dark
	ErrorICON_Dark     fyne.Resource = Error_Dark
	InstallICON_Dark   fyne.Resource = Install_Dark
	SaveICON_Dark      fyne.Resource = Save_Dark
	RestartICON_Dark   fyne.Resource = Restart_Dark
	InfoICON_Dark      fyne.Resource = Info_Dark
	UninstallICON_Dark fyne.Resource = Uninstall_Dark

	// Light Icons
	DevICON_Light       fyne.Resource = Dev_Light
	DownloadICON_Light  fyne.Resource = Download_Light
	ErrorICON_Light     fyne.Resource = Error_Light
	InstallICON_Light   fyne.Resource = Install_Light
	SaveICON_Light      fyne.Resource = Save_Light
	RestartICON_Light   fyne.Resource = Restart_Light
	InfoICON_Light      fyne.Resource = Info_Light
	UninstallICON_Light fyne.Resource = Uninstall_Light

	// Themed Icons
	DevICON       fyne.Resource
	DownloadICON  fyne.Resource
	ErrorICON     fyne.Resource
	InstallICON   fyne.Resource
	SaveICON      fyne.Resource
	RestartICON   fyne.Resource
	InfoICON      fyne.Resource
	UninstallICON fyne.Resource

	// No-Theme Icons
	AppICON         fyne.Resource = App
	PlaceholderICON fyne.Resource = Placeholder
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
		UninstallICON = UninstallICON_Light
	} else {
		DevICON = DevICON_Dark
		InstallICON = InstallICON_Dark
		DownloadICON = DownloadICON_Dark
		ErrorICON = ErrorICON_Dark
		SaveICON = SaveICON_Dark
		RestartICON = RestartICON_Dark
		InfoICON = InfoICON_Dark
		UninstallICON = UninstallICON_Dark
	}
}
