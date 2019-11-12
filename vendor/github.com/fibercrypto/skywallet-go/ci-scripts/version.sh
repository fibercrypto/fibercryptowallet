#!/usr/bin/env bash
version=$(git describe --tags --exact-match HEAD 2> /dev/null)
if [ $? -ne 0 ]
then
    version=$(cat ./VERSION 2> /dev/null)
    if [ $? -ne 0 ]
    then
        version='v0.0.0'
    fi
fi
echo $version
