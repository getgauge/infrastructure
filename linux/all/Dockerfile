FROM gocd/gocd-agent-centos-7:v17.12.0

ENV container docker

RUN (yum install -y -q yum-utils &&\
rpm --import "http://keyserver.ubuntu.com/pks/lookup?op=get&search=0x3FA7E0328081BFF6A14DA29AA6A19B38D3D831EF" &&\
yum-config-manager --add-repo http://download.mono-project.com/repo/centos/ && yum -y update &&\
yum install -y -q epel-release java-1.8.0-openjdk-devel hg mono-complete zip &&\
yum install -y -q git-core zlib zlib-devel gcc-c++ patch readline readline-devel libyaml-devel libffi-devel openssl-devel make bzip2 autoconf automake libtool bison curl sqlite-devel &&\
su go -c "cd /home/go && git clone git://github.com/sstephenson/rbenv.git .rbenv && git clone git://github.com/sstephenson/ruby-build.git .rbenv/plugins/ruby-build" && echo -e 'export PATH="/home/go/.rbenv/bin:/home/go/.rbenv/plugins/ruby-build/bin:/home/go/.rbenv/shims/:$PATH"' >> /etc/default/go-agent && echo 'eval "$(rbenv init -)"' >> /etc/default/go-agent && source /etc/default/go-agent &&\
curl -o go.zip https://storage.googleapis.com/golang/go1.10.3.linux-amd64.tar.gz && tar -C /usr/local -xzf go.zip && echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/default/go-agent && source /etc/default/go-agent && chown go -R /usr/local/go/ &&\
yum install -y -q dh-autoreconf curl-devel expat-devel gettext-devel openssl-devel perl-devel zlib-devel asciidoc xmlto docbook2X getopt && yum groupinstall -y "Development Tools" &&\
curl -o git.tar.gz https://www.kernel.org/pub/software/scm/git/git-2.9.3.tar.gz && tar -zxf git.tar.gz && cd git-2.9.3 && make configure && ./configure --prefix=/usr && make all && make install &&\
yum install -y -q nodejs npm python34 jq fakeroot rpm-build rpmdevtools libxml2-devel &&\
curl https://bootstrap.pypa.io/get-pip.py | python3 &&\
curl -o mvn.tar.gz http://redrockdigimark.com/apachemirror/maven/maven-3/3.5.0/binaries/apache-maven-3.5.0-bin.tar.gz && tar -C /opt -xzf mvn.tar.gz &&\
echo -e "export JAVA_HOME=/usr/lib/jvm/java-openjdk\nexport M2_HOME=/opt/apache-maven-3.5.0\nexport PATH=/opt/apache-maven-3.5.0/bin:${PATH}" >> /etc/default/go-agent && source /etc/default/go-agent &&\
chown go /home/go && pip install virtualenv && su go -c "source /etc/default/go-agent && pip install --upgrade --user awscli==1.11.18" && echo "export PATH=$PATH:/home/go/.local/bin" >> /etc/default/go-agent &&\
su go -c "source /etc/default/go-agent && cd /home/go && rbenv install 2.3.0 && rbenv local 2.3.0 && gem install rake && gem install parser && gem install json && gem install bundler && rbenv install 2.4.0 && rbenv local 2.4.0 && gem install bundler && rbenv global 2.3.0" &&\
su go -c "mkdir -p /home/go/.ssh/ && ssh-keyscan -t rsa github.com >> /home/go/.ssh/known_hosts && git config --global user.email \"gaugeci@gmail.com\" && git config --global user.name gaugeci" &&\
echo "export GO_SERVER_SYSTEM_PROPERTIES=\"$GO_SERVER_SYSTEM_PROPERTIES -Dplugin.cd.go.contrib.elastic-agent.docker.log.level=debug\"" >> /etc/default/go-agent)
