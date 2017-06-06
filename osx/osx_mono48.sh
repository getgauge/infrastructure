xcode-select --install
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

brew update
brew install git
brew install Caskroom/cask/java
brew install go
brew install maven

curl -O https://download.mono-project.com/archive/4.8.1/macos-10-universal/MonoFramework-MDK-4.8.1.macos10.xamarin.universal.pkg
sudo installer -pkg MonoFramework-MDK-4.8.1.macos10.xamarin.universal.pkg -target /

brew install ruby-build
brew install rbenv
rbenv install 2.3.0
rbenv global 2.3.0
echo 'export PATH="$HOME/.rbenv/shims:$PATH"' >> ~/.bash_profile
source ~/.bash_profile
gem install bundler

brew tap cosmo0920/mingw_w64
brew install mingw-w64
curl -O http://crossgcc.rts-software.org/download/gcc-4.8.0-qt-4.8.4-win32/gcc-4.8.0-qt-4.8.4-for-mingw32.dmg
sudo hdiutil attach gcc-4.8.0-qt-4.8.4-for-mingw32.dmg
sudo installer -package /Volumes/gcc-4.8.0-qt-4.8.4-for-mingw32/gcc-4.8.0-qt-4.8.4-for-mingw32.pkg -target /
cd $HOME
hdiutil unmount /Volumes/gcc-4.8.0-qt-4.8.4-for-mingw32
echo 'export PATH="$PATH:/usr/local/gcc-4.8.0-qt-4.8.4-for-mingw32/win32-gcc/bin/"' >> ~/.bash_profile

curl -O http://s.sudre.free.fr/Software/files/Packages.dmg
sudo hdiutil attach Packages.dmg 
sudo installer -package /Volumes/Packages\ */Install\ Packages.pkg -target /
cd $HOME
hdiutil unmount /Volumes/Packages\ *

curl -O https://download.gocd.io/binaries/17.4.0-4892/generic/go-agent-17.4.0-4892.zip
unzip go-agent-17.4.0-4892.zip
echo "agent.auto.register.key=$AGENT_AUTO_REGISTER_KEY" > go-agent-17.4.0/config/autoregister.properties
echo "agent.auto.register.resources=FT,UT,darwin,installers" >> go-agent-17.4.0/config/autoregister.properties
echo "agent.auto.register.hostname=$AGENT_NAME" >> go-agent-17.4.0/config/autoregister.properties