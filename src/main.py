#
# Copyright (c) - All Rights Reserved.
#
# This project is licenced under the MIT License.
#
import tkinter as tk
import yaml,subprocess,sys




class Sh:
    def __init__(self, PowerShell=False, CustomSt=False, Stdin=False, Stdout=False, Stderr=False):
        self.PowerShell = PowerShell
        self.CustomSt = CustomSt
        self.Stdin = Stdin
        self.Stdout = Stdout
        self.Stderr = Stderr

    def Cmd(self, input_str):
        shell = "cmd"
        if self.PowerShell:
            shell = "PowerShell.exe"

        cmd = [shell, "/C", input_str]
        if self.CustomSt:
            if self.Stderr:
                stderr = subprocess.PIPE
            else:
                stderr = None

            if self.Stdin:
                stdin = subprocess.PIPE
            else:
                stdin = None

            if self.Stdout:
                stdout = subprocess.PIPE
            else:
                stdout = None

            process = subprocess.Popen(cmd, stderr=stderr, stdin=stdin, stdout=stdout)
        else:
            process = subprocess.Popen(cmd)

        process.communicate()
        return process.returncode

    def Out(self, input_str):
        cmd = ["cmd", "/C", input_str]
        if self.PowerShell == True:
            cmd[0] = "PowerShell.exe"
        process = subprocess.Popen(cmd, stdout=subprocess.PIPE)
        out, _ = process.communicate()
        return out.decode(), process.returncode

if len(sys.argv) > 1:
    if sys.argv[1] == "term":
        Sh().Cmd(".\\main.exe")
        exit()

    args = " ".join(sys.argv[1:])
    Sh().Cmd(".\\main.exe "+args)
    exit()

def InstallChocoPackages():
    Sh().Cmd(".\\main.exe choco noend")


def InstallScoopPackages():
    Sh().Cmd(".\\main.exe scoop noend")


def NewYmlFile():
    Sh().Cmd(".\\main.exe newyamlfile noend")

def getymldata():
    try: 
        with open("packages.yml", "r") as file:
            yaml_data = yaml.safe_load(file)
        return yaml_data
    except FileNotFoundError:
        NewYmlFile()
        return getymldata()
yaml_data = getymldata()

window = tk.Tk()
window.title("Windows Package AutoInstaller")
window.resizable(False,False)
window.minsize(width=344, height=327)

#red_font = tk.font.Font(family="Arial", size=12, weight="bold", underline=True, foreground="red")


def save_text():
    choco_text = edited_text_choco.get(1.0, tk.END)
    scoop_text = edited_text_scoop.get(1.0,tk.END)
    yaml_data["choco"] = choco_text
    yaml_data["scoop"] = scoop_text
    with open("packages.yml", "w") as yml_file:
        yaml.dump(yaml_data, yml_file)



choco_label = tk.Label(window, text="Choco packages to install:")
choco_label.pack()
edited_text_choco = tk.Text(window, wrap="word", width=40, height=4)
edited_text_choco.insert(tk.END, yaml_data["choco"])
edited_text_choco.pack(padx=10, pady=10)


scoop_label = tk.Label(window,text="Scoop packages to install:")
scoop_label.pack()
edited_text_scoop = tk.Text(window, wrap="word", width=40, height=4)
edited_text_scoop.insert(tk.END, yaml_data["scoop"])
edited_text_scoop.pack(padx=10, pady=10)


save_button = tk.Button(window, text="Save package lists", command=lambda: save_text())
save_button.pack(pady=5)

label = tk.Label(window, text="Select any option")
label.pack()

InsChocopackbt = tk.Button(window, text="Install Choco packages.", command=InstallChocoPackages)
InsChocopackbt.pack()

InstallScooppackbt = tk.Button(window, text="Install Scoop Packages", command=InstallScoopPackages)
InstallScooppackbt.pack()

window.mainloop()
