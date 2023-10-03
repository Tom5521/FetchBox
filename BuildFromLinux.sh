#!/bin/bash


sudo fyne-cross windows -arch=amd64 -env FYNE_THEME=dark


if [ ! -d "builds" ]; then
  mkdir builds
fi

cp -rf ./fyne-cross/bin/windows-amd64/WPA.exe ./builds/WPA.exe
