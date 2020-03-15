#!/usr/bin/env bash

set -x

brew update
# USB HID enumeration libraries
# SDL for building hardware wallet emulator
brew install sdl2_image sdl2 mesa mesalib-glw

brew update
# USB HID enumeration libraries
brew install libusb
# Install OS-specific test and build dependencies for hardware-wallet
mkdir -p tmp/hardware-wallet
git clone --depth=1 --single-branch --branch develop https://github.com/skycoin/hardware-wallet.git tmp/hardware-wallet
git -C tmp/hardware-wallet submodule update --init --recursive
( cd ./tmp/hardware-wallet && sh "ci-scripts/install-osx.sh" )
# Include paths for brew packages e.g. SDL2
export SDL_INCLUDE="$(brew --prefix sdl2)/include/SDL2";