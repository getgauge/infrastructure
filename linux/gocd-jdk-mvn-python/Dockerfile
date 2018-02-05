FROM getgauge/gocd-jdk-mvn

LABEL MAINTAINER="getgauge" \
 VERSION="0.1"\
 DESCRIPTION="centos gocd agent with openjdk1.8, git, maven 3.5, python"\
 REPOSITORY="getgauge/gocd-jdk-mvn-python"

ENV container docker

RUN (yum -y update && yum install -y epel-release && yum install -y python34 python34-setuptools && \
easy_install-3.4 pip && \
echo "export VIRTUALENVWRAPPER_PYTHON=/usr/bin/python3.4" >> /etc/bashrc && \
echo "export PATH=/home/go/.local/bin:$PATH" >> /etc/bashrc )
