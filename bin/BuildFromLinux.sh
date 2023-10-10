#!/bin/bash

cd ..

sudo fyne-cross windows -arch=amd64 -env FYNE_THEME=dark


if [ ! -d "builds" ]; then
  mkdir builds
fi

cp -rf ./fyne-cross/bin/windows-amd64/FetchBox.exe ./builds/FetchBox.exe
