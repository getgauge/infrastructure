FROM gocd/gocd-agent-ubuntu-16.04:v18.2.0

LABEL MAINTAINER="getgauge" \
 VERSION="0.3" \
 DESCRIPTION="ubuntu gocd agent with golang" \
 REPOSITORY="getgauge/gocd-node-xvfb"

ENV container docker

RUN apt update && apt install -y software-properties-common &&\
    add-apt-repository ppa:gophers/archive && apt-get update &&\
    apt-get install -y golang-1.10-go zip unzip &&\
    echo "export PATH=/usr/lib/go-1.10/bin:$PATH" >> /etc/default/go-agent

