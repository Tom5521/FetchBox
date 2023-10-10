#!/bin/bash

cd ..

ReGenerateBundle()
{
  fyne bundle --pkg icon --name "$1" -o ./src/icon/bundled.go "$2"
}

AppendBundle ()
{
  fyne bundle --pkg icon --name "$1" -o ./src/icon/bundled.go -append "$2"
}



ReGenerateBundle Placeholder ./Assets/Placeholder.png
AppendBundle App ./Assets/Icon.png

# Dark Icons

AppendBundle Save_Dark ./Assets/Icons-Dark/Save.png
AppendBundle Dev_Dark ./Assets/Icons-Dark/Dev.png
AppendBundle Install_Dark ./Assets/Icons-Dark/Install.png
AppendBundle Info_Dark ./Assets/Icons-Dark/Info.png
AppendBundle Error_Dark ./Assets/Icons-Dark/Error.png
AppendBundle Restart_Dark ./Assets/Icons-Dark/Restart.png
AppendBundle Download_Dark ./Assets/Icons-Dark/Download.png
AppendBundle Uninstall_Dark ./Assets/Icons-Dark/Uninstall.png

# Light Icons

AppendBundle Save_Light ./Assets/Icons-Light/Save.png
AppendBundle Dev_Light ./Assets/Icons-Light/Dev.png
AppendBundle Install_Light ./Assets/Icons-Light/Install.png
AppendBundle Info_Light ./Assets/Icons-Light/Info.png
AppendBundle Error_Light ./Assets/Icons-Light/Error.png
AppendBundle Restart_Light ./Assets/Icons-Light/Restart.png
AppendBundle Download_Light ./Assets/Icons-Light/Download.png
AppendBundle Uninstall_Light ./Assets/Icons-Light/Uninstall.png
