FROM microsoft/nanoserver:1709

LABEL MAINTAINER="getgauge" \
 VERSION="0.1" \
 DESCRIPTION="windows gocd agent with nanoserver" \
 REPOSITORY="getgauge/gocd-agent"

ADD https://bintray.com/gocd-contrib/gocd_golang_goagent/download_file?file_path=goagent%2F1.92%2Fgocd-golang-agent_windows_amd64_1.92.exe C:\\gocd-agent\\agent.exe
COPY bootstrap.ps1 c:\\gocd-agent\\bootstrap.ps1
CMD [ "powershell", "-File", "c:\\gocd-agent\\bootstrap.ps1"]

