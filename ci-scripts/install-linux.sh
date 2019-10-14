#!/usr/bin/env bash

set -x

sudo apt-get update
# SDL for building hardware wallet emulator
sudo apt install libsdl2-dev libsdl2-image-dev libegl1-mesa-dev libgles2-mesa-dev
# PIP
sudo apt-get install python3-pip
# OpenGL swrast driver for xvfb
sudo apt install mesa-utils and libgl1-mesa-glx
