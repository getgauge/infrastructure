FROM getgauge/gocd-jdk-mvn

LABEL MAINTAINER="getgauge"
LABEL VERSION="0.1"
LABEL DESCRIPTION="centos gocd agent with openjdk1.8, git, maven 3.5, mono"
LABEL REPOSITORY="getgauge/gocd-jdk-mvn-mono"

ENV container docker

RUN (rpm --import "http://keyserver.ubuntu.com/pks/lookup?op=get&search=0x3FA7E0328081BFF6A14DA29AA6A19B38D3D831EF" &&\
yum-config-manager --add-repo http://download.mono-project.com/repo/centos/ &&\
yum -y update && yum install -y mono-complete)
