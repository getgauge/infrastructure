$GOLANG_BOOTSTRAPPER_VERSION='1.1'

# Create temp directory
New-Item "C:/tmp" -ItemType Directory

# install chocolatey
$chocolateyUseWindowsCompression='false'
$ErrorActionPreference = "Stop"
[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12
Set-ExecutionPolicy Bypass -Scope Process -Force; Invoke-Expression ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))

# install packages
cinst git -y
cinst golang -y
cinst jdk8 -y
cinst maven -y
cinst dotnet4.5 -y
cinst nvm -y
nvm install 8
nvm install 10
nvm install 12
cinst dotnetcore-sdk -y
cinst python3 -y
cinst pip -y
cinst unzip -y
# Refresh envs
refreshenv

# Remove chocolatey from temp location
Remove-Item C:\\Users\\ContainerAdministrator\\AppData\\Local\\Temp\\chocolatey -Force -Recurse


# install gocd bootstrapper
Invoke-WebRequest https://github.com/ketan/gocd-golang-bootstrapper/releases/download/${GOLANG_BOOTSTRAPPER_VERSION}/go-bootstrapper-${GOLANG_BOOTSTRAPPER_VERSION}.windows.amd64.exe -Outfile C:\\go-agent.exe
mkdir C:\go
Add-LocalGroupMember -Group "Administrators" -Member "ContainerAdministrator"
