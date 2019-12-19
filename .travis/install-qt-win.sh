wget "http://download.qt.io/official_releases/online_installers/qt-unified-windows-x86-online.exe" -O qt.exe
./qt.exe --verbose --script util/qt-headless.qs					
export MWDIR="/c/Users/travis/Qt/Tools/mingw730_64"
export QTDIR="/c/Users/travis/Qt/5.12.3/mingw73_64"
export PATH="$MWDIR/bin:$QTDIR/bin:$PATH"
