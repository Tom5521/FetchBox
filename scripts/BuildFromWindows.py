import os
from os import system as sys 


os.chdir("..")
os.environ['FYNE_THEME'] = 'dark'
sys("fyne package -os windows --src . --exe FetchBox.exe")

input()
