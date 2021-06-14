# Gofsh [![Build Status](https://travis-ci.com/Syssos/gofsh.svg?branch=main)](https://travis-ci.com/github/Syssos/gofsh)  [![Cody Code](https://syssos.app/static/images/index/cody_code.svg)](https://syssos.app)

<p align="center">
  <img src="https://syssos.app/static/images/index/gofsh.png" alt="Go Shell img"/>
</p>

An interactive or non-interactive go based shell, aimed at automating tasks often used for creating, or testing web based applications. 

### Dependancies for module:
- [Go](https://golang.org/)
- TOML parsing package, used for logger settings([komkom/toml](https://github.com/komkom/toml))
- HTTP request multiplexer, used in pond generated web apps ([gorilla/mux](https://github.com/gorilla/mux))

## Install

These instructions are targeted towards linux user's.

To get the files you will need, clone the repository to a location on your machine.

```bash
git clone https://github.com/Syssos/gofsh.git
```
Once the repository is cloned cd into it. 

When inside of the gofsh directory we should be able to install and build the program using the build.sh script.

``` bash
cd gofsh
chmod u+x build.sh && ./build.sh
```

Installing the shell will cause ```go get``` to grab any associated packages and install them in the ```$GOPATH/src``` folder.

The build shell script will also place the gofsh bin file in a location executable.

## Usage
Gofsh will handle interactive and non-interactive usage. If you which to run the interactive shell, then doing so can be started with the gofsh command.

```bash
gofsh
```

If you would like to use the shell in non-interactive mode, then the commands can be passed to gofsh as arguments

```bash
gofsh site https://cody.syssos.app
```

## Logging
The logging is handled by a [Flog](https://github.com/Syssos/gofsh/tree/main/src/filelog#filelog) instance. 

Gofsh uses a Flog instance generated with settings saved in a TOML file. These settings are stored and configured in the [``` $HOME/.gofsh/config/LogSettings.toml ```](https://github.com/Syssos/gofsh/blob/main/etc/config/config/LogSettings.toml)

The logger will use the built in "time" package from go. To find more configuration settings that can be used for this logger, [this](https://www.geeksforgeeks.org/time-formatting-in-golang/) link should help find more information.

## Travis
Travis-CI is a continuous integration repository "extension". The badge at the top of the page represents the current status of the gofsh program.

Due to the current state of the program, major changes can be introduced at any time. To prevent errors from occuring to someone who clones the repo, the travis build state is indicated at the top of this readme.

## Testing 

There are 2 testing methods accounted for in this repository. The first is a generic all around test for the gofsh functionality, this is done with the [gofsh_test.go](https://github.com/Syssos/gofsh/blob/main/gofsh_test.go) file. Due to travis being used for this project, the [test.sh](https://github.com/Syssos/gofsh/blob/main/test.sh) script is included for travis to run upon the latest commit.
