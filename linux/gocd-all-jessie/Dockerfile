FROM golang:latest

ENV container docker

ARG GOLANG_BOOTSTRAPPER_VERSION=2.1
ARG UID=1000
ARG GID=1000

ADD ["https://github.com/ketan/gocd-golang-bootstrapper/releases/download/${GOLANG_BOOTSTRAPPER_VERSION}/go-bootstrapper-${GOLANG_BOOTSTRAPPER_VERSION}.linux.amd64", "/usr/local/bin/go-agent"]

RUN (groupadd -g ${GID} go && \ 
useradd -u ${UID} -g go -d /home/go -m go && \
apt-get update && apt-get install -y --no-install-recommends apt-transport-https dirmngr &&\
apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 3FA7E0328081BFF6A14DA29AA6A19B38D3D831EF &&\
echo "deb https://download.mono-project.com/repo/debian stable-stretch main" | tee /etc/apt/sources.list.d/mono-official-stable.list &&\
curl -sL https://deb.nodesource.com/setup_10.x | bash - &&\
apt-get update &&\
apt-get install -y --no-install-recommends mono-complete openjdk-8-jdk-headless git zip unzip nodejs build-essential npm jq maven libssl1.0-dev libreadline-dev zlib1g-dev &&\
echo "Installing rbenv" &&\
su go -c "cd /home/go && git clone git://github.com/sstephenson/rbenv.git .rbenv && git clone git://github.com/sstephenson/ruby-build.git .rbenv/plugins/ruby-build" && echo 'export PATH="/home/go/.rbenv/bin:/home/go/.rbenv/plugins/ruby-build/bin:/home/go/.rbenv/shims/:$PATH"' >> /etc/default/go-agent && echo 'eval "$(rbenv init -)"' >> /etc/default/go-agent && . /etc/default/go-agent &&\
echo "Installing Pypi" &&\
curl https://bootstrap.pypa.io/get-pip.py | python3 &&\
rm -rf /var/lib/apt/lists/*  &&\
chown go /home/go && pip install virtualenv && su go -c ". /etc/default/go-agent && pip install --upgrade --user awscli==1.11.18" && echo "export PATH=$PATH:/home/go/.local/bin" >> /etc/default/go-agent &&\
su go -c ". /etc/default/go-agent && cd /home/go && rbenv install 2.4.0 && rbenv local 2.4.0 && gem install bundler && rbenv global 2.4.0" &&\
su go -c "mkdir -p /home/go/.ssh/ && ssh-keyscan -t rsa github.com >> /home/go/.ssh/known_hosts && git config --global user.email \"gaugeci@gmail.com\" && git config --global user.name gaugeci" &&\
chmod +x /usr/local/bin/go-agent &&\
mkdir go)

CMD /usr/local/bin/go-agent