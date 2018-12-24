
FROM gocd/gocd-agent-ubuntu-18.04:v18.11.0

LABEL MAINTAINER="getgauge" \
 VERSION="0.2" \
 DESCRIPTION="debian9 gocd agent with openjdk1.11, git, maven 3.6" \
 REPOSITORY="getgauge/gocd-jdk-mvn"

ENV container docker

RUN apt-get update &&\
apt-get install -y --no-install-recommends gnupg2 openjdk-11-jdk maven &&\
rm -rf /var/lib/apt/lists/*