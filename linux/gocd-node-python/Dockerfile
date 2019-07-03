FROM getgauge/gocd-node

LABEL MAINTAINER="getgauge" \
 VERSION="0.1"\
 DESCRIPTION="centos gocd agent with openjdk1.8, git, nodejs, python"\
 REPOSITORY="getgauge/gocd-node-python"

ENV container docker

RUN (yum -y update && yum install -y epel-release && yum install -y python36 python36-setuptools && \
easy_install-3.6 pip && \
ln -sf /usr/bin/python3 /usr/bin/python && \
echo "export PATH=/home/go/.local/bin:$PATH" >> /etc/bashrc )
