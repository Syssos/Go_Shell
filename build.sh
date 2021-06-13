#!/bin/bash

# Create gofsh folder at /home/usr/gofsh
DIR="$HOME/.gofsh"
PONDDIR="$HOME/.gofsh/pond_templates"

if [ -d "$DIR" ]; then
  echo "Directory $DIR exists, is the application installed?"
  exit 1
else
  echo "Createing directory $DIR"
  mkdir "$DIR"
  echo "Createing directory $PONDDIR"
  mkdir "$PONDDIR"
  echo ""
  echo "Coping files from etc/config/pond_templates to $PONDDIR"
  cp -r "./etc/config/pond_templates" "$DIR"
fi

# Install commands for shell
echo ""
echo "Building pond command..."
go build "./src/commands/pond/pond.go"
echo "Building r2h command..."
go build "./src/commands/r2h/r2h.go"
echo "Building site command..."
go build "./src/commands/site/site.go"

mv "pond" "r2h" "site" "$HOME/go/bin"

# Install goshell
echo ""
echo "Installing gofsh command"
go install .
echo "Moving gofsh binary to /usr/bin"
sudo cp "$HOME/go/bin/gofsh" "/usr/bin"