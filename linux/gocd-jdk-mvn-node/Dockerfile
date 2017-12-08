FROM getgauge/gocd-jdk-mvn

LABEL MAINTAINER="getgauge"
LABEL VERSION="0.1"
LABEL DESCRIPTION="centos gocd agent with openjdk1.8, git, maven 3.5, nodejs"
LABEL REPOSITORY="getgauge/gocd-jdk-mvn-node"

ENV container docker

RUN (curl --silent --location https://rpm.nodesource.com/setup_9.x | bash - &&\
yum -y update && yum install -y nodejs npm)
