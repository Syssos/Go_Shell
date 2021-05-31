# Go Shell [![Build Status](https://travis-ci.com/Syssos/Go_Shell.svg?branch=main)](https://travis-ci.com/github/Syssos/Go_Shell)  [![Cody Code](https://syssos.app/static/images/index/cody_code.svg)](https://syssos.app)

<p align="center">
  <img src="https://raw.githubusercontent.com/Syssos/Go_Shell/main/settings/images/Go_Shell.png" alt="Go Shell img"/>
</p>

Go Shell is an interactive shell, aimed at automating tasks often used for creating, or testing web based applications. 

### Dependancies for module:
- [Go](https://golang.org/)
- TOML parsing package ([komkom/toml](https://github.com/komkom/toml))
- HTTP request multiplexer ([gorilla/mux](https://github.com/gorilla/mux))
- GofshTemplates ([README](https://github.com/Syssos/Go_Shell/tree/main/settings/GofshTemplates/README.md))

While the usefulness of this package comes from the web specific commands, simple linux style commands are included to help with those commands.

## Install

All of the packages used in this repository were designed with external usablility in mind. Each package directory will contain a readme with all of the information you'll need to get started. The directions below should get you started with using the shell as.

These instructions are for targeted towards linux user's.

To start things off this will require some knowledge of working with Go modules, and packages. As well as how to build and install programs. If you have any questions about Go or how to do these things I just so happen to have a repository, [here](https://github.com/Syssos/Learning_Go), I created as I learned Go, leaving notes behind that may be useful to you.

To get the files we need clone the package to a location you can use.

``` 
git clone https://github.com/Syssos/Go_Shell.git
```
Once the repository is cloned cd into it. 

When inside of the Go_Shell directory we should be able to install and build the program with no issues.
```
cd Go_Shell
```
Thanks to the help of [travis](#travis) we should see that the package is in working condition. 

Installing the shell will cause ```go get``` to grab any accociated packages and install them in the ```$GOPATH/src``` folder.

```
go install .
```

This will place a compiled binary file in "$PATH/bin" called "Go_Shell". Running this command will give you the Go Shell.

For me that looks something like
```
~/go/bin/Go_Shell
```
Alternativly you can install the "cmds" package with go get and create a script that utilizes the cmds, and filelog packages much like [main.go](https://github.com/Syssos/Go_Shell/blob/main/main.go) does.

## Usage
As the project grows there will be more and more commands to keep track of. To help with keeping this readme to the point, information on the commands for this shell can be found in the readme in the [cmds](https://github.com/Syssos/Go_Shell/tree/main/cmds) directory.

## Logging
<p align="center">
  <img src="https://raw.githubusercontent.com/Syssos/Go_Shell/main/settings/images/Error_Log_Example.PNG" alt="Go Shell img"/>
  <p align="center">Example held in <a href="https://github.com/Syssos/Go_Shell/blob/main/settings/example_log.txt">settings/example_log.txt</a></p>
</p>

The logging is handled by a [Flog](https://github.com/Syssos/Go_Shell/tree/main/filelog#filelog) instance. This instance is generated with settings saved in a TOML file. To learn more about the settings file check out the [readme](https://github.com/Syssos/Go_Shell/tree/main/settings) for it

## Travis
Travis-CI is a continuous integration repository "extention". The badge at the top of the page represents the current status of the Go Shell program. This is impo

Due to the current state of the program changes can be major changes can be introduced at anytime. To prevent errors from occuring to someone who clones the repo, the travis build state is indicated at the top of this readme. 

## Testing 

As mentioned above, travis is used to ensure that the package is contnuously running. It does this based off of the result from the `pkg_test.go` file.

This file will contain all of the test's needed to ensure the program is working correctly. As commands get added to the repository that are more complex in design, the use of extra _test.go files may become apparent.