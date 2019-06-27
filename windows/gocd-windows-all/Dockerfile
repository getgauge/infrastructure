FROM mcr.microsoft.com/windows/servercore:ltsc2016

ENV TMP c:/tmp
ENV TEMP c:/tmp

COPY install-packages.ps1 C:/

RUN powershell -File C:\install-packages.ps1

CMD C:\\go-agent.exe