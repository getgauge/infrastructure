FROM gocd/gocd-agent-centos-7:v18.10.0

ARG GIT_VERSION=2.15.1
ARG MAVEN_VERSION=3.5.4

LABEL MAINTAINER="getgauge" \
 VERSION="0.1" \
 DESCRIPTION="centos gocd agent with openjdk1.8, git, maven 3.5" \
 REPOSITORY="getgauge/gocd-jdk-mvn"

ENV container docker

RUN (yum install -y yum-utils &&\
yum install -y dh-autoreconf curl-devel expat-devel gettext-devel openssl-devel \
 perl-devel zlib-devel asciidoc xmlto docbook2X getopt &&\
yum groupinstall -y "Development Tools" &&\
curl -o git.tar.gz https://mirrors.edge.kernel.org/pub/software/scm/git/git-$GIT_VERSION.tar.gz &&\
tar -zxf git.tar.gz && cd git-$GIT_VERSION &&\
make configure && ./configure --prefix=/usr && make all && make install &&\
cd .. && rm -rf git-$GIT_VERSION git.tar.gz &&\
yum remove -y dh-autoreconf curl-devel expat-devel gettext-devel openssl-devel\
 perl-devel zlib-devel asciidoc xmlto docbook2X getopt &&\
yum groupremove -y "Development Tools" &&\
yum install -y epel-release java-10-openjdk-devel zip &&\
curl -o mvn.tar.gz http://www-eu.apache.org/dist/maven/maven-3/$MAVEN_VERSION/binaries/apache-maven-$MAVEN_VERSION-bin.tar.gz &&\
tar -C /opt -xzf mvn.tar.gz && rm mvn.tar.gz &&\
echo -e "export JAVA_HOME=/usr/lib/jvm/java-openjdk\nexport M2_HOME=/opt/apache-maven-$MAVEN_VERSION\nexport PATH=/opt/apache-maven-$MAVEN_VERSION/bin:${PATH}" >> /etc/default/go-agent && source /etc/default/go-agent &&\
su go -c "curl -sL https://github.com/shyiko/jabba/raw/master/install.sh | bash && . ~/.jabba/jabba.sh" &&\
chown go /home/go && su go -c "source /etc/default/go-agent" && echo "export PATH=$PATH:/home/go/.local/bin" >> /etc/default/go-agent &&\
su go -c "mkdir -p /home/go/.ssh/ && ssh-keyscan -t rsa github.com >> /home/go/.ssh/known_hosts && git config --global user.email \"gaugeci@gmail.com\" && git config --global user.name gaugeci" &&\
echo "export GO_SERVER_SYSTEM_PROPERTIES=\"$GO_SERVER_SYSTEM_PROPERTIES -Dplugin.cd.go.contrib.elastic-agent.docker.log.level=debug\"" >> /etc/default/go-agent)
