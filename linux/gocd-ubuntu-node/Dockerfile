FROM gocd/gocd-agent-ubuntu-18.04:v18.11.0

LABEL description="debian image with node and dependencies for headless chrome"\
    repository="getgauge/gocd-ubuntu-node"


ENV container docker

# requirements for chrome headless
# https://github.com/Googlechrome/puppeteer/issues/290#issuecomment-322838700
RUN apt-get update && apt-get install -y gnupg &&\
curl -sL https://deb.nodesource.com/setup_11.x | bash - &&\
apt-get update && apt-get install -y nodejs make g++ &&\
apt-get update &&\
apt-get install -y git-all gconf-service libasound2 libatk1.0-0 libc6 libcairo2 libcups2 libdbus-1-3 \
  libexpat1 libfontconfig1 libgcc1 libgconf-2-4 libgdk-pixbuf2.0-0 libglib2.0-0 libgtk-3-0 \
  libnspr4 libpango-1.0-0 libpangocairo-1.0-0 libstdc++6 libx11-6 libx11-xcb1 libxcb1 libxcomposite1 \
  libxcursor1 libxdamage1 libxext6 libxfixes3 libxi6 libxrandr2 libxrender1 libxss1 libxtst6 ca-certificates \
  fonts-liberation libappindicator1 libnss3 lsb-release xdg-utils wget &&\
apt autoremove -y &&\
rm -rf /var/lib/apt/lists/*