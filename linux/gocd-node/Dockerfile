FROM gocd/gocd-agent-centos-7:v18.2.0

ARG GIT_VERSION=2.16.3

LABEL MAINTAINER="getgauge" \
 VERSION="0.2" \
 DESCRIPTION="centos gocd agent with nodejs, git" \
 REPOSITORY="getgauge/gocd-node"

ENV container docker

RUN (yum install yum-utils &&\
yum install -y dh-autoreconf curl-devel expat-devel gettext-devel openssl-devel \
 perl-devel zlib-devel asciidoc xmlto docbook2X getopt libXScrnSaver &&\
yum groupinstall -y "Development Tools" &&\
curl -o git.tar.gz https://mirrors.edge.kernel.org/pub/software/scm/git/git-$GIT_VERSION.tar.gz &&\
tar -zxf git.tar.gz && cd git-$GIT_VERSION &&\
make configure && ./configure --prefix=/usr && make all && make install &&\
cd .. && rm -rf git-$GIT_VERSION git.tar.gz &&\
yum remove -y dh-autoreconf curl-devel expat-devel gettext-devel openssl-devel\
 perl-devel zlib-devel asciidoc xmlto docbook2X getopt &&\
yum groupremove -y "Development Tools" &&\
curl --silent --location https://rpm.nodesource.com/setup_9.x | bash - && yum -y update &&\
yum install -y nodejs npm &&\
echo -e "export JAVA_HOME=/usr/lib/jvm/java-openjdk\nexport M2_HOME=/opt/apache-maven-3.5.0\nexport PATH=/opt/apache-maven-3.5.0/bin:${PATH}" >> /etc/default/go-agent && source /etc/default/go-agent &&\
chown go /home/go && su go -c "source /etc/default/go-agent" && echo "export PATH=$PATH:/home/go/.local/bin" >> /etc/default/go-agent &&\
su go -c "mkdir -p /home/go/.ssh/ && ssh-keyscan -t rsa github.com >> /home/go/.ssh/known_hosts && git config --global user.email \"gaugeci@gmail.com\" && git config --global user.name gaugeci" &&\
echo "export GO_SERVER_SYSTEM_PROPERTIES=\"$GO_SERVER_SYSTEM_PROPERTIES -Dplugin.cd.go.contrib.elastic-agent.docker.log.level=debug\"" >> /etc/default/go-agent)
