'''
  Copyright Tom5521(c) - All Rights Reserved.
 
  This project is licensed under the MIT License.
'''

import shutil
import requests
import os

# Create a "tmp" directory if it doesn't exist.
if not os.path.exists("tmp"):
    os.mkdir("tmp")

# Download the opengl32.dll if it doesn't exist in the "tmp" directory.
if not os.path.exists("tmp/opengl32.7z"):
    print("Downloading opengl32.dll...")
    url = 'https://downloads.fdossena.com/geth.php?r=mesa64-latest'
    r = requests.get(url)
    with open('tmp/opengl32.7z', 'wb') as f:
        f.write(r.content)

# Unzip the downloaded 7z file if the "tmp/opengl" directory doesn't exist.
if not os.path.exists("tmp/opengl"):
    print("Unzipping 7z file...")
    os.mkdir("tmp/opengl")
    os.system("7z e -o\"tmp/opengl\" tmp/opengl32.7z")

# Compress the files for Windows.
print("Compressing for Windows...")
if os.path.exists("builds/FetchBox-win64.zip"):
    os.remove("builds/FetchBox-win64.zip")
if not os.path.exists("FetchBox-win64/"):
    os.mkdir("FetchBox-win64/")
shutil.copy2("builds/FetchBox.exe", "FetchBox-win64/")
shutil.copy2("tmp/opengl/opengl32.dll", "FetchBox-win64/")
os.system("zip -r builds/FetchBox-win64.zip FetchBox-win64/")
shutil.rmtree("FetchBox-win64/")



