#!/bin/bash


export FYNE_THEME=dark
#sudo fyne-cross windows -arch=amd64 -app-id com.Tom5521.WPA


if [ ! -d "builds" ]; then
  mkdir builds
fi


unzip ./fyne-cross/dist/windows-amd64/Windows-package-autoinstaller.exe.zip
mv Windows-package-autoinstaller.exe ./builds/WPA.exe
