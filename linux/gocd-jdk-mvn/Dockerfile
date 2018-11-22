FROM maven:3.6.0-jdk-11-slim

ARG GOCD_BOOTSTRAPPER_VERSION=2.3    

LABEL MAINTAINER="getgauge" \
 VERSION="0.2" \
 DESCRIPTION="gocd agent with openjdk1.10, git, maven 3.6" \
 REPOSITORY="getgauge/gocd-jdk-mvn"

ENV container docker

RUN (adduser --disabled-password  --gecos '' go &&\
apt-get update && apt-get install --no-install-recommends -y wget ssh git &&\
rm -rf /var/lib/apt/lists/* &&\
wget -q -O /usr/local/bin/bootstrapper https://github.com/ketan/gocd-golang-bootstrapper/releases/download/$GOCD_BOOTSTRAPPER_VERSION/go-bootstrapper-$GOCD_BOOTSTRAPPER_VERSION.linux.amd64 &&\
chown go /home/go &&\
echo "export PATH=$PATH:/home/go/.local/bin" >> /etc/default/go-agent &&\
su go -c "mkdir -p /home/go/.ssh/ && ssh-keyscan -t rsa github.com >> /home/go/.ssh/known_hosts && git config --global user.email \"gaugeci@gmail.com\" && git config --global user.name gaugeci" &&\
echo "export GO_SERVER_SYSTEM_PROPERTIES=\"$GO_SERVER_SYSTEM_PROPERTIES -Dplugin.cd.go.contrib.elastic-agent.docker.log.level=debug\"" >> /etc/default/go-agent)

CMD /usr/local/bin/bootstrapper