#!/bin/bash

DIR="$HOME/.gofsh"

if [ -d "$DIR" ]; then
  go test -v ./
  ./clean.sh
else
  chmod u+x "build.sh"
  chmod u+x "clean.sh"
  ./build.sh
  go test -v ./
  ./clean.sh
fi