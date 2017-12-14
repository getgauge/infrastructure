FROM microsoft/windowsservercore:1709

LABEL MAINTAINER="getgauge" \
 VERSION="0.1" \
 DESCRIPTION="windows gocd agent with windowsservercore" \
 REPOSITORY="getgauge/gocd-agent"

ADD https://download.gocd.org/binaries/17.12.0-5626/win/go-agent-17.12.0-5626-jre-64bit-setup.exe c:\\build\\gocd-agent.exe

RUN C:\\build\\gocd-agent.exe /S
