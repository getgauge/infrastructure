FROM getgauge/gocd-jdk-mvn

LABEL MAINTAINER="getgauge" \
 VERSION="0.2"\
 DESCRIPTION="Ubuntu gocd agent with openjdk1.11, git, maven 3.5, python"\
 REPOSITORY="getgauge/gocd-jdk-mvn-python"

ENV container docker

RUN (apt-get update &&\
apt-get install -y --no-install-recommends python3.7 python3-pip python3-setuptools &&\
rm -rf /var/lib/apt/lists/* &&\
ln -sf /usr/bin/python3 /usr/bin/python && \
ln -sf /usr/bin/pip3 /usr/bin/pip && \
echo "export PATH=/home/go/.local/bin:$PATH" >> /etc/bashrc )
