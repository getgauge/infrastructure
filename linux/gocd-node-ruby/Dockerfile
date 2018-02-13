FROM getgauge/gocd-node

LABEL MAINTAINER="getgauge" \
 VERSION="0.1"\
 DESCRIPTION="centos gocd agent with openjdk1.8, git, nodejs, ruby"\
 REPOSITORY="getgauge/gocd-node-ruby"

ENV container docker

RUN (su go -c "cd /home/go && git clone git://github.com/sstephenson/rbenv.git .rbenv && \
git clone git://github.com/sstephenson/ruby-build.git .rbenv/plugins/ruby-build" && \
echo -e 'export PATH="/home/go/.rbenv/bin:/home/go/.rbenv/plugins/ruby-build/bin:/home/go/.rbenv/shims/:$PATH"' >> /etc/default/go-agent && \
echo 'eval "$(rbenv init -)"' >> /etc/default/go-agent && source /etc/default/go-agent )
