FROM gocd/gocd-agent-centos-7:v17.12.0

ENV container docker
ENV PATH="/home/go/.rbenv/bin:/home/go/.rbenv/plugins/ruby-build/bin:/home/go/.rbenv/shims/:/usr/local/go/bin:/opt/apache-maven-3.5.4/bin:/home/go/.local/bin:$PATH"
ENV JAVA_HOME="/usr/lib/jvm/java-openjdk"
ENV M2_HOME="/opt/apache-maven-3.5.4"

RUN (yum install -y -q yum-utils &&\
rpm --import "http://keyserver.ubuntu.com/pks/lookup?op=get&search=0x3FA7E0328081BFF6A14DA29AA6A19B38D3D831EF" &&\
yum-config-manager --add-repo http://download.mono-project.com/repo/centos/ && yum -y update &&\
yum install -y -q https://centos7.iuscommunity.org/ius-release.rpm &&\
yum install -y -q epel-release  java-1.8.0-openjdk-devel hg mono-complete zip &&\
yum install -y -q zlib dpkg zlib-devel gcc-c++ patch readline readline-devel libyaml-devel libffi-devel openssl-devel make bzip2 autoconf automake libtool bison curl sqlite-devel &&\
su go -c "cd /home/go && git clone git://github.com/sstephenson/rbenv.git .rbenv && git clone git://github.com/sstephenson/ruby-build.git .rbenv/plugins/ruby-build" &&\
curl -o go.zip https://storage.googleapis.com/golang/go1.10.3.linux-amd64.tar.gz && tar -C /usr/local -xzf go.zip && chown go -R /usr/local/go/ &&\
yum install -y -q nodejs npm python36 python36-setuptools jq fakeroot rpm-build rpmdevtools libxml2-devel &&\
yum erase -y git && yum install -y -q git2u-core &&\
rpm -Uvh https://packages.microsoft.com/config/rhel/7/packages-microsoft-prod.rpm &&\
yum install -y dotnet-sdk-2.2 &&\
yum install -y hub &&\
curl https://bootstrap.pypa.io/get-pip.py | python3 &&\
curl -o mvn.tar.gz http://www-us.apache.org/dist/maven/maven-3/3.5.4/binaries/apache-maven-3.5.4-bin.tar.gz && tar -C /opt -xzf mvn.tar.gz &&\
rm go.zip mvn.tar.gz &&\
echo -e 'eval "$(rbenv init -)"' >> /etc/default/go-agent &&\
chown go /home/go && pip install virtualenv && su go -c "pip install --upgrade --user awscli==1.16.193" &&\
su go -c "curl -sL https://github.com/shyiko/jabba/raw/master/install.sh | bash && . ~/.jabba/jabba.sh" &&\
su go -c "curl https://raw.githubusercontent.com/creationix/nvm/v0.34.0/install.sh | bash && . ~/.bashrc && nvm install --lts && nvm use --lts" &&\
su go -c "source /etc/default/go-agent && cd /home/go && rbenv install 2.6.0 && rbenv local 2.6.0 && gem install rake && gem install parser && gem install json && gem install bundler && rbenv global 2.6.0" &&\
su go -c "mkdir -p /home/go/.ssh/ && ssh-keyscan -t rsa github.com >> /home/go/.ssh/known_hosts && git config --global user.email \"gaugeci@gmail.com\" && git config --global user.name gaugeci" &&\
echo "export GO_SERVER_SYSTEM_PROPERTIES=\"$GO_SERVER_SYSTEM_PROPERTIES -Dplugin.cd.go.contrib.elastic-agent.docker.log.level=debug\"" >> /etc/default/go-agent &&\
source /etc/default/go-agent)
