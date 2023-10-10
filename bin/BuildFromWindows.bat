@echo off

cd ..

set FYNE_THEME=dark

fyne package -os windows --src . --exe FetchBox.exe

pause
