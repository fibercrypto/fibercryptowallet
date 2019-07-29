# Install QT 5.13.0

sudo apt install python3-requests p7zip-full wget

wget https://git.kaidan.im/lnj/qli-installer/raw/master/qli-installer.py
chmod +x qli-installer.py

./qli-installer.py 5.13.0 linux desktop

export PATH=./5.13.0/gcc_64/bin:$PATH
