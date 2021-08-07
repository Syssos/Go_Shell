# Gofsh [![Build Status](https://travis-ci.com/Syssos/gofsh.svg?branch=main)](https://travis-ci.com/github/Syssos/gofsh)

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

### Configuration
When the gofsh shell is first installed the default logger settings will have the logfile stored in /tmp. If you would like to have a more permanent log file, it is recommended that the settings file is updated.

Gofsh uses settings generated with information saved in a TOML file. These settings are stored and configured in the file [``` /etc/gofsh/config/LogSettings.toml ```](https://github.com/Syssos/gofsh/blob/main/etc/config/config/LogSettings.toml).

Changing the LogFile value to the path and filename you would like to use will change where the log file is saved.

Note that the filename does not need to be an existing file, but the path to the file must be valid.

```bash
/path/must/exist/youdecide.txt
```
my log file location

```bash
/home/syssos/Documents/gofsh_logs/gofshlog.txt
```

The logger will use the built in "time" package from Go, to generate the timestamp. To find more information on configuration settings that can be used for this logger, [this](https://yourbasic.org/golang/format-parse-string-time-date-example/) link should help.

## Uninstalling
The go shell will place files in the following locations.

```
github.com/Syssos/gofsh/etc/config/*                 -> /etc/gofsh/config
github.com/Syssos/gofsh/bin/* (built from build.sh)  -> $HOME/go/bin

gofsh binary                                         -> /usr/bin
```
The [clean.sh](https://github.com/Syssos/gofsh/blob/main/clean.sh) will remove the files from the computer, however will not remove any files created from the shell.

## Usage
Gofsh will handle interactive and non-interactive usage. If you wish to use the non-interactive shell, then doing so can be done by passing gofsh the command and the arguments.

```bash
gofsh
```

If you would like to use the shell in non-interactive mode, then the commands can be passed to gofsh as arguments

```bash
gofsh site https://cody.syssos.app
```
## Testing 

There are 2 testing methods accounted for in this repository. The first is a generic all around test for the gofsh functionality, this is done with the [gofsh_test.go](https://github.com/Syssos/gofsh/blob/main/gofsh_test.go) file. Due to travis being used for this project, the [test.sh](https://github.com/Syssos/gofsh/blob/main/test.sh) script is included for travis to run upon the latest commit.
