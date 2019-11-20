# FiberCrypto wallet

FiberCrypto wallet is a cryptocurrency software wallet aimed at:

- Provide easy-to-use interactions to users
- State
- Out-of-the-box support for every SkyFiber token in a single place
- Support other altcoins
- Facilitate exchange of crypto assets
- Buy and sell supported crypto assets using fiat (e.g. USD, GBP, EUR, ...)
- Integrations with trading tools
- Offer basic blockchain-specific tools

## Development

### Project folder structure

Project files are organized as follows:

- `main.go` : Application entry point
- `CHANGELOG.md` : Project changelog
- `Makefile` : Project build rules
- `README.md` : This file.
- `*.qrc` : QML resource index files.
- `qtquickcontrols2.conf` : QT Quick controls configuration file.
- `./resources` : Static resources.
- `./resources/images` : Graphics resources needed by the application.
- `./resources/images/icons` : Project and third-party icons
- `./resources/fonts` : Font files needed to compile the application.
- `./src` : Application source code.
- `./src/ui` : QML definitions for application GUI components.
- `./src/ui/Dialogs` : QML definitions for reusable dialogs.
- `./src/ui/Delegates` : QML specs for partial views.
- `./src/core` : Core go-lang interfaces.
- `./src/main` : Project specific source code.
- `./src/util` : Reusable code.
- `./src/util/logging` : Event logging infrastructure.
- `./src/models` : QT models linking coin-specific models to application GUI.
- `./src/coin` : Source code for altcoin integrations.
- `./src/coin/mocks` : Types implementing `core` interfaces for generic testing scenarios
- `./src/coin/skycoin` : Skycoin wallet integration
- `./src/coin/skycoin/models` : Skycoin implementation of golang core interfaces.
- `./src/coin/skycoin/blockchain` : Skycoin blockchain API.
- `./src/coin/skycoin/sign` : Skycoin sign API.
- `vendor` : Project dependencies managed by `dep`.

### Architecture

FiberCrypto wallet supports multiple altcoins. In order to cope with this complexity GUI code and QT models rely on strict interfaces which shall be implemented to add support for a given coin. Each such integration must have two main components:

- `Models API`: Implements application core interfaces.
- `Sign API` : Implements altcoin transaction and message signing primitives required by application code.
- `Blockchain API` : Provides communication between application and altcoin service nodes to query for data via REST, JSON-RPC and other similar low-level client-server API.
- `Peer-exchange API` (optional): Implements peer-to-peer interactions with altcoin blockchain nodes.

### Build System

The build system is [Qt framework](https://www.qt.io/ "The Qt Company"). The front-end is programmed in [QML](http://doc.qt.io/qt-5/qmlapplications.html "QML Applications"), and the back-end in [Go](https://golang.org/ "The Go Programming Language"), using [therecipe/qt](https://github.com/therecipe/qt/ "therecipe/qt").

#### Requirements

##### Qt version

[Linux/X11 requirements](http://doc.qt.io/qt-5/linux.html)  
[MacOS requirements](http://doc.qt.io/qt-5/macos.html)  
[Windows requirements](http://doc.qt.io/qt-5/windows.html)  

The minimum Qt version required is [Qt 5.12.0 LTS](https://download.qt.io/archive/qt/5.12/5.12.0/ "Qt Archive"). However, is highly recommended using [Qt 5.12.1 LTS](https://download.qt.io/archive/qt/5.12/5.12.1/ "Qt Archive") or any later version of Qt5 due to some bugs:  
- [QTBUG-68156](https://bugreports.qt.io/browse/QTBUG-68156 "Incompatible version of OpenSSL on Ubuntu 18.04")  
- [QTBUG-72811](https://bugreports.qt.io/browse/QTBUG-72811 "[Reg 5.11 -> 5.12] QQC2 buttons not react to click when holding for about a second")

We always recommend using the latest Qt version. See [Qt Archive](https://download.qt.io/archive/qt/ "Qt Archive").

#### Make targets

Common actions are automated with the help of `make`. The following targets have been implemnented:

```
deps                           Add dependencies
run                            Run FiberCrypto Wallet.
install-deps-no-envs           Install therecipe/qt with -tags=no_env set
install-docker-deps            Install docker images for project compilation using docker
install-deps-Linux             Install Linux dependencies
install-deps-Darwin            Install osx dependencies
install-deps-Windows           Install Windowns dependencies
install-deps                   Install dependencies
build-docker                   Build project using docker
build                          Build FiberCrypto Wallet.
clean-Windows                  Clean project FiberCrypto Wallet.
clean                          Clean project FiberCrypto Wallet.
test-sky                       Run Skycoin plugin test suite
test                           Run project test suite
install-linters                Install linters
lint                           Run linters. Use make install-linters first.
```

Type `make help` in your console for details.

## WIP
This is a Work-In-Progress.
