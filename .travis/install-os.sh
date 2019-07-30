# Install QT 5.13.0

brew update
brew install p7zip | true

brew install pyenv | true

# Add pyenv initializer to shell startup script
echo 'eval "$(pyenv init -)"' >> ~/.bash_profile
echo "~/5.13.0/gcc_64/bin"      >> ~/.bash_profile

# Reload your profile
source ~/.bash_profile

# Install python 3.6.8
pyenv install 3.6.8
pyenv global 3.6.8

pip install requests

(cd $HOME && wget https://git.kaidan.im/lnj/qli-installer/raw/master/qli-installer.py && chmod +x qli-installer.py && ./qli-installer.py 5.13.0 mac desktop)

