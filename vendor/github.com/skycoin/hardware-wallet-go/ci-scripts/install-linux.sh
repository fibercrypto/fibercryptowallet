#!/usr/bin/env bash

set -x

sudo apt-get update
# USB HID enumeration libraries
sudo apt install libudev-dev libusb-1.0-0-dev
# SDL for building hardware wallet emulator
sudo apt install libsdl2-dev libsdl2-image-dev libegl1-mesa-dev libgles2-mesa-dev
# OpenGL swrast driver for xvfb
sudo apt install mesa-utils and libgl1-mesa-glx
