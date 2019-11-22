#!/usr/bin/env bash

set -x

# Install gimme
curl -sL -o "$HOME/bin/gimme" https://raw.githubusercontent.com/travis-ci/gimme/master/gimme
chmod +x "$HOME/bin/gimme"

