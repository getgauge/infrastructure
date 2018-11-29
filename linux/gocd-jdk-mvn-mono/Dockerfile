FROM getgauge/gocd-jdk-mvn

LABEL MAINTAINER="getgauge" \
 VERSION="0.2" \
 DESCRIPTION="ubuntu gocd agent with openjdk1.11, git, maven 3.5, mono" \
 REPOSITORY="getgauge/gocd-jdk-mvn-mono" 

ENV container docker

RUN apt-get update && apt-get install -y gnupg &&\
apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 3FA7E0328081BFF6A14DA29AA6A19B38D3D831EF &&\
echo "deb https://download.mono-project.com/repo/ubuntu stable-bionic main" | tee /etc/apt/sources.list.d/mono-official-stable.list &&\
apt update &&\
apt install -y mono-devel &&\
apt autoremove -y &&\
rm -rf /var/lib/apt/lists/*
