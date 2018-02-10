xcode-select --install
command -v brew >/dev/null 2>&1 || {/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"}

brew update
brew install git
brew install jq
brew install Caskroom/cask/java
brew install go
brew install maven
brew install nodejs
brew cask install dotnet-sdk

command -v pip >/dev/null 2>&1 || {
    brew install python3
    echo 'alias pip="pip3"' >> ~/.bash_profile
    echo 'alias python="python3"' >> ~/.bash_profile
}

brew install ruby-build
brew install rbenv
command -v rbenv >/dev/null 2>&1 || {
    rbenv install 2.3.0
    rbenv global 2.3.0
    echo 'export PATH="$HOME/.rbenv/shims:$PATH"' >> ~/.bash_profile
    source ~/.bash_profile
    gem install bundler
}

brew tap cosmo0920/mingw_w64
brew install mingw-w64

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