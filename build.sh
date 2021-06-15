#!/bin/bash

# Create gofsh folder at /home/usr/gofsh
DIR="/etc/gofsh"

if [ -d "$DIR" ]; then
  echo "Directory $DIR exists, is the application installed?"
  exit 1
else
  echo "Createing directory $DIR"
  sudo mkdir "$DIR"
  echo ""
  sudo cp -r ./etc/config/* "$DIR"
fi

# check for go folder
GODIR="$HOME/go/bin"
if [ ! -d "$GODIR" ]; then
  mkdir "$GODIR"
  echo ""
fi

# Install commands for shell
echo "Building pond command..."
go build "./src/commands/pond/pond.go"
echo "Building r2h command..."
go build "./src/commands/r2h/r2h.go"
echo "Building site command..."
go build "./src/commands/site/site.go"

mv "pond" "r2h" "site" "$HOME/go/bin"

# Install needed packages
go get github.com/komkom/toml

# Install goshell
echo ""
echo "Installing gofsh command"
go install .
echo "Moving gofsh binary to /usr/bin"
sudo cp "$HOME/go/bin/gofsh" "/usr/bin"