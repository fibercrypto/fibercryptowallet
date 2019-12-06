#!/usr/bin/env bash

set -x

sudo apt-get update
# PIP
sudo apt-get install python3-pip
# OpenGL swrast driver for xvfb
sudo apt install mesa-utils and libgl1-mesa-glx
# Hardware wallet for running tests against emulator
mkdir -p tmp/hardware-wallet
git clone --depth=1 --single-branch --branch develop https://github.com/skycoin/hardware-wallet.git tmp/hardware-wallet
git -C tmp/hardware-wallet submodule update --init --recursive
( cd ./tmp/hardware-wallet && sh "ci-scripts/install-linux.sh" )
export PATH="/usr/local/bin:$(pwd)/tmp/hardware-wallet/protoc/bin:$PATH"
