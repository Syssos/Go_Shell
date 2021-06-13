#!/bin/bash

DIR="$HOME/.gofsh"

mkdir "$HOME/go"
mkdir "$HOME/go/bin"

if [ -d "$DIR" ]; then
  go test -v ./
else
  chmod u+x "build.sh"
  ./build.sh
  sudo cp "$GOPATH/bin/gofsh" "$HOME/go/bin"
  go test -v ./
fi