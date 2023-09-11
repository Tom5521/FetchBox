#!/bin/bash


sudo fyne-cross windows -arch=amd64 -app-id com.Tom5521.WPA -env FYNE_THEME=dark -app-version 2.0


if [ ! -d "builds" ]; then
  mkdir builds
fi

sudo mv ./fyne-cross/bin/windows-amd64/Windows-package-autoinstaller.exe ./builds/WPA.exe
