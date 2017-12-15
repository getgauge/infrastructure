FROM microsoft/nanoserver

LABEL MAINTAINER="getgauge" \
 VERSION="0.1" \
 DESCRIPTION="windows gocd agent with nanoserver, OpenJDK1.8, Git 2.15.1.2" \
 REPOSITORY="getgauge/gocd-agent"

RUN powershell -C \
 mkdir c:\\build; \
 Invoke-WebRequest http://cdn.azul.com/zulu/bin/zulu8.21.0.1-jdk8.0.131-win_x64.zip -OutFile C:\\build\\openjdk.zip; \
 Invoke-WebRequest https://github.com/git-for-windows/git/releases/download/v2.15.1.windows.2/MinGit-2.15.1.2-64-bit.zip -OutFile C:\\build\\git.zip;\
 Invoke-WebRequest https://download.gocd.org/binaries/17.12.0-5626/generic/go-agent-17.12.0-5626.zip -OutFile c:\\build\\gocd-agent.zip;\
 Expand-Archive C:\\build\\openjdk.zip C:\\Java; \
 Rename-Item C:\\Java\\zulu8.21.0.1-jdk8.0.131-win_x64 OpenJDK; \
 Expand-Archive C:\\build\\git.zip C:\\git; \
 Expand-Archive c:\\build\gocd-agent.zip C:\\; \
 Rename-Item C:\\go-agent-17.12.0 gocd-agent;

RUN setx /M PATH "%PATH%;C:\Git\cmd;C:\Java\OpenJDK\bin"

COPY bootstrap.ps1 c:\\gocd-agent\\bootstrap.ps1
WORKDIR c:\\gocd-agent
CMD ["powershell", "-File", "c:\\gocd-agent\\bootstrap.ps1"]

