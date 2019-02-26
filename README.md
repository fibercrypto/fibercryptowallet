# Wallet GUI client

## Build System

The build system is [Qt framework](https://www.qt.io/ "The Qt Company"). The front-end is programmed in [QML](http://doc.qt.io/qt-5/qmlapplications.html "QML Applications"), and the back-end in [Python](https://www.python.org/ "Python"), using [PySide2](https://wiki.qt.io/Qt_for_Python "Qt for Python").

### Requirements

#### Qt version

[Linux/X11 requirements](http://doc.qt.io/qt-5/linux.html)  
[MacOS requirements](http://doc.qt.io/qt-5/macos.html)  
[Windows requirements](http://doc.qt.io/qt-5/windows.html)  

The minimum Qt version required is [Qt 5.12.0 LTS](https://download.qt.io/archive/qt/5.12/5.12.0/ "Qt Archive"). However, is highly recommended using [Qt 5.12.1 LTS](https://download.qt.io/archive/qt/5.12/5.12.1/ "Qt Archive") due to some bugs:
- [QTBUG-68156](https://bugreports.qt.io/browse/QTBUG-68156 "Incompatible version of OpenSSL on Ubuntu 18.04")  
- [QTBUG-72811](https://bugreports.qt.io/browse/QTBUG-72811 "[Reg 5.11 -> 5.12] QQC2 buttons not react to click when holding for about a second")

#### PySide2

From Qt 5.12.0 onwards, the Qt Company officially supports Python as programming language for desktop development.  
The wheel of [PySide2](https://wiki.qt.io/Qt_for_Python "Qt for Python") must match the selected Qt version, i.e. if you use [Qt 5.12.0 LTS](https://download.qt.io/archive/qt/5.12/5.12.0/ "Qt Archive") you must use it with PySide2 5.12.0, and if you use [Qt 5.12.1 LTS](https://download.qt.io/archive/qt/5.12/5.12.1/ "Qt Archive") you must use PySide2 5.12.1.

## WIP
This is a Work-In-Progress.