#!/bin/bash

DIR="$HOME/.gofsh"

mkdir "$HOME/go"
mkdir "$HOME/go/bin"

if [ -d "$DIR" ]; then
  go test -v ./
  ./clean.sh
else
  chmod u+x "build.sh"
  chmod u+x "clean.sh"
  ./build.sh
  sudo cp "$GOPATH/bin/gofsh" "$HOME/go/bin"
  go test -v ./
  ./clean.sh
fi