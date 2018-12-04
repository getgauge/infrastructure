FROM getgauge/gocd-jdk-mvn

LABEL MAINTAINER="getgauge" \
 VERSION="0.2"\
 DESCRIPTION="ubuntu gocd agent with openjdk1.11, git, maven 3.5, nodejs"\
 REPOSITORY="getgauge/gocd-jdk-mvn-node"

ENV container docker

RUN apt-get update && apt-get install -y gnupg &&\
curl -sL https://deb.nodesource.com/setup_11.x | bash - &&\
apt-get update && apt-get install -y nodejs make g++ &&\
apt autoremove -y &&\
rm -rf /var/lib/apt/lists/*

