# Windows-package-autoinstaller

This is a little program why use for installing packages automatically in my windows making a frontend for scoop and choco package managers to which I plan to add many more functions

## Features

#### Install choco and scoop packages

You only need to fill the text box with the package names and press *Install [pkg manager] packages*

#### Save package lists

This program saves the pkg list in a yaml file, you can edit it without oppening the program!

#### Request admin permissions

If you want to install choco packages you need administrator permissions, this program check if choco/sudo or gsudo is installed to exec the cmds as administrator without restarting the program

#### Cool Interface

I use Fyne framework which in its defaul aspect in itself is very nice.

#### Easy to use

The installation of choco and scoop package managers is fully automated.

#### CLI mode

If you don't need a graphic interface, you can use it only in cmd or powershell!

## CLI mode

The syntax of the CLI mode is something like this:


```
[binary] [parameter]
```
### Parameters

- `choco`: Used to install the choco packages specified in the yaml configuration file.
- `scoop`: It serves for the same purpose as the previous one but for scoop.
- `Newyamlfile`: Used to create a new yaml file (overwrite the old config file if it exists)
- `AddCustomBuckets`: Used to add a custom bucket for scoop package manager, the bucket name is specified in a third argument
- `Version`: Print the program version in the cmd
