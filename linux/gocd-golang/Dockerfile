FROM gocd/gocd-agent-ubuntu-16.04:v18.2.0

LABEL MAINTAINER="getgauge" \
 VERSION="0.3" \
 DESCRIPTION="ubuntu gocd agent with golang" \
 REPOSITORY="getgauge/gocd-node-xvfb"

ENV container docker

RUN apt-get update && apt-get install -y software-properties-common &&\
    add-apt-repository ppa:longsleep/golang-backports &&\
    apt-get update &&\
    apt-get install -y golang-go

