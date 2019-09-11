FROM gocd/gocd-agent-ubuntu-16.04:v18.2.0

LABEL MAINTAINER="getgauge" \
 VERSION="0.3" \
 DESCRIPTION="ubuntu gocd agent with nodejs, git, xvfb" \
 REPOSITORY="getgauge/gocd-node-xvfb"

ENV container docker

RUN apt update && apt install apt-transport-https &&\
    apt-key adv --keyserver hkp://pool.sks-keyservers.net --recv-keys 023EDB0B &&\
    echo deb https://dl.bintray.com/gauge/gauge-deb nightly main | tee -a /etc/apt/sources.list &&\
    curl -sL https://deb.nodesource.com/setup_10.x | bash - &&\
    apt update &&\
    apt install -y nodejs libgtk2.0-0 libnotify-bin libgconf-2-4 libnss3 xvfb libxss1 libasound2