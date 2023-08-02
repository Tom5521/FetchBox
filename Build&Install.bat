@echo off
cd .\src\
echo Compiling the executable...
go build -o ../InstallApps.exe main.go
echo Executable Compiled!!!

pause