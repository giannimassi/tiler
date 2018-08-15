#!/usr/bin/env bash

DIR=$1

if [[ -d "${DIR}" ]] ; then
    mkdir -p $1/build/mac
    mkdir -p $1/build/windows
    mkdir -p $1/build/linux

    export GOARCH=amd64

    export GOOS=darwin
    go build -o $1/build/mac/tiler

    export GOOS=windows
    go build -o $1/build/windows/tiler.exe

    export GOOS=linux
    go build -o $1/build/linux/tiler
else
    echo "$1 is not a valid directory"
fi