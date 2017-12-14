FROM microsoft/nanoserver

ARG GIT_VERSION=2.15.1.2
ARG NODE_VERSION=9.2.0

LABEL MAINTAINER="getgauge" \
 VERSION="0.1" \
 DESCRIPTION="windows gocd agent with nodejs, git" \
 REPOSITORY="getgauge/gocd-node"

RUN powershell -Command Expand-Archive C:\build\node.zip C:\; \
 Invoke-WebRequest https://nodejs.org/dist/v${NODE_VERSION}/node-v${NODE_VERSION}-win-x64.zip -OutFile node.zip; \ 
 Invoke-WebRequest https://github.com/git-for-windows/git/releases/download/v2.15.1.windows.2/MinGit-2.15.1.2-64-bit.zip -OutFile git.zip; \ 
 Invoke-WebRequest https://download.gocd.org/binaries/17.12.0-5626/win/go-agent-17.12.0-5626-jre-64bit-setup.exe -OutFile gocd-agent.exe; \ 
 Expand-Archive C:\\build\\node.zip C:\node; \
 Expand-Archive C:\\build\\git.zip $Env:ProgramFiles\Git; \
 & C:\\build\\gocd-agent.exe /S; \
 setx /M PATH "%PATH%;C:\node"