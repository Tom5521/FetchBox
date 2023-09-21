@echo off

set FYNE_THEME=dark

fyne package -os windows --src . --exe WPA.exe

pause
