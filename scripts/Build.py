'''
  Copyright Tom5521(c) - All Rights Reserved.
 
  This project is licenced under the MIT License.
'''
import os,platform,shutil



if platform.system() != "Windows":
    os.system("sudo fyne-cross windows -arch=amd64 -env FYNE_THEME=dark")
    if not os.path.exists("builds"):
        os.mkdir("builds")
    shutil.copy("./fyne-cross/bin/windows-amd64/FetchBox.exe","./builds/FetchBox.exe")
elif platform.system() == "Windows":
    os.environ['FYNE_THEME'] = 'dark'
    os.system("fyne package -os windows --src . --exe FetchBox.exe")
