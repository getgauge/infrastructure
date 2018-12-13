xcode-select --install
command -v brew >/dev/null 2>&1 || {/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"}

brew update
brew install git || brew upgrade git
brew install jq || brew upgrade jq
brew cask reinstall java
brew install go || brew upgrade go
brew install maven || brew upgrade maven
brew install nodejs || brew upgrade nodejs
brew cask install dotnet-sdk || brew cask upgrade dotnet-sdk

# Install jabba (java version manager) to handle java versions
curl -sL https://github.com/shyiko/jabba/raw/master/install.sh | sh
echo "source ~/.jabba/jabba.sh" >> ~/.bash_profile
source ~/.bash_profile

# Install java 10 explicitly
jabba install openjdk@1.10-0


command -v pip >/dev/null 2>&1 || {
    brew install python3
    echo 'alias pip="pip3"' >> ~/.bash_profile
    echo 'alias python="python3"' >> ~/.bash_profile
}

brew install ruby-build || brew upgrade ruby-build
brew install rbenv || brew upgrade rbenv
command -v rbenv >/dev/null 2>&1 || {
    rbenv install 2.3.0
    rbenv global 2.3.0
    echo 'export PATH="$HOME/.rbenv/shims:$PATH"' >> ~/.bash_profile
    source ~/.bash_profile
    gem install bundler
}

brew tap cosmo0920/mingw_w64
brew install mingw-w64 || brew upgrade mingw-w64

if [ ! -f gcc-4.8.0-qt-4.8.4-for-mingw32.dmg ]; then
    curl -O http://crossgcc.rts-software.org/download/gcc-4.8.0-qt-4.8.4-win32/gcc-4.8.0-qt-4.8.4-for-mingw32.dmg
    sudo hdiutil attach gcc-4.8.0-qt-4.8.4-for-mingw32.dmg
    sudo installer -package /Volumes/gcc-4.8.0-qt-4.8.4-for-mingw32/gcc-4.8.0-qt-4.8.4-for-mingw32.pkg -target /
    cd $HOME
    hdiutil unmount /Volumes/gcc-4.8.0-qt-4.8.4-for-mingw32
    echo 'export PATH="$PATH:/usr/local/gcc-4.8.0-qt-4.8.4-for-mingw32/win32-gcc/bin/"' >> ~/.bash_profile
fi

if [ ! -f Packages.dmg ]; then
    curl -O http://s.sudre.free.fr/Software/files/Packages.dmg
    sudo hdiutil attach Packages.dmg 
    sudo installer -package /Volumes/Packages\ */Install\ Packages.pkg -target /
    cd $HOME
    hdiutil unmount /Volumes/Packages\ *
fi

if [ ! -f go-agent-17.4.0-4892.zip ]; then
    curl -O https://download.gocd.io/binaries/17.4.0-4892/generic/go-agent-17.4.0-4892.zip
    unzip go-agent-17.4.0-4892.zip
fi
