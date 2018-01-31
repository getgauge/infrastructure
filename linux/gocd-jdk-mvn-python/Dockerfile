FROM getgauge/gocd-jdk-mvn

LABEL MAINTAINER="getgauge" \
 VERSION="0.1"\
 DESCRIPTION="centos gocd agent with openjdk1.8, git, maven 3.5, python"\
 REPOSITORY="getgauge/gocd-jdk-mvn-python"

ENV container docker

RUN (yum -y update && yum install -y epel-release && \
su go -c "yum install -y python34 python34-setuptools && easy_install-3.4 pip && ln -sf /usr/bin/python3 /usr/bin/python")
