# gmusic
Simple music player in go that uses raylib for front-end because every GUI framework needs to suck for some reason
# installation
## Linux
```bash
# Arch (-based) Linux
sudo pacman -Sy --needed git go
# Debian (-based) Linux
sudo apt update
sudo apt install git golang-go

git clone https://github.com/griesson-h/gmusic.git
cd gmusic
go build
```
## Windows
```bash
winget install Git.Git Golang.Go # reboot may be required
wget https://github.com/Vuniverse0/mingwInstaller/releases/download/1.2.1/mingwInstaller.exe
mingwInstaller.exe
# follow the installation process
git clone https://github.com/griesson-h/gmusic.git
cd gmusic
go build
```
