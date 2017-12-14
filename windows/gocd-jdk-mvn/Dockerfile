FROM microsoft/nanoserver

LABEL MAINTAINER="getgauge" \
 VERSION="0.1" \
 DESCRIPTION="centos gocd agent with openjdk1.8, git, maven 3.5" \
 REPOSITORY="getgauge/gocd-jdk-mvn"

ADD http://cdn.azul.com/zulu/bin/zulu8.21.0.1-jdk8.0.131-win_x64.zip C:\\build\\openjdk.zip
ADD https://github.com/git-for-windows/git/releases/download/v2.15.1.windows.2/MinGit-2.15.1.2-64-bit.zip C:\\build\\git.zip
ADD https://download.gocd.org/binaries/17.12.0-5626/win/go-agent-17.12.0-5626-jre-64bit-setup.exe C:\\build\\gocd-agent.exe

RUN powershell -Command \
 Expand-Archive C:\build\openjdk.zip $Env:ProgramFiles\Java; \
 Get-ChildItem -Path $Env:ProgramFiles\Java -Filter "zulu*" | ForEach-Object {$_ | Rename-Item -NewName "OpenJDK" }; \
 & Expand-Archive C:\build\git.zip $Env:ProgramFiles\Git; \
 & C:\\build\\gocd-agent.exe /S; \
 setx /M PATH "%PATH%;C:\Program Files\Git\cmd;C:\Program Files\Java\OpenJDK\bin"