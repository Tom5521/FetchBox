@echo off
cd .\src\
echo Compiling the golang binary...
go build -o ..\builds\main.exe ..\main.go
echo Go Binary Compiled!!!
echo Packing python files with the golang binary...
pyinstaller --onefile --distpath ..\builds --add-binary "../builds/main.exe;." --name InstallApps.exe main.py
rm src\InstallApps.exe.spec
rm src\build
echo Packing Complete!!!


pause
