#!/usr/bin/env bash

set -e -o pipefail

go build -o $GOPATH/bin/skycoin-hw-cli .
