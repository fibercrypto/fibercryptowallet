#!/usr/bin/env bash

set -x

brew update
# USB HID enumeration libraries
# SDL for building hardware wallet emulator
brew install sdl2_image sdl2 mesa mesalib-glw

brew update
# USB HID enumeration libraries
brew install libusb
# SDL for building hardware wallet emulator
brew install sdl2_image sdl2 mesa mesalib-glw
