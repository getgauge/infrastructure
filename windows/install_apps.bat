echo "Install apps"
cinst git -y
cinst golang -y
cinst jdk8 -y
cinst maven -y
cinst ruby -version 2.4 -y
cinst dotnet4.5 -y
cinst nvm -y
cinst nodejs.install --version 9.4.0 -y
cinst dotnetcore-sdk -y
cinst python3 -y
cinst pip -y
:: Refresh envs
refreshenv
echo "Installation Finished"