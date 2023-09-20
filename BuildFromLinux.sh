#!/bin/bash


sudo fyne-cross windows -arch=amd64 -env FYNE_THEME=dark


if [ ! -d "builds" ]; then
  mkdir builds
fi

sudo mv ./fyne-cross/bin/windows-amd64/Windows-package-autoinstaller.exe ./builds/WPA.exe
