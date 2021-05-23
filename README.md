# Go Shell [![Build Status](https://travis-ci.com/Syssos/Go_Shell.svg?branch=main)](https://travis-ci.com/github/Syssos/Go_Shell)  [![Cody Code](https://syssos.app/static/images/index/cody_code.svg)](https://syssos.app)

![Go Shell img](https://github.com/Syssos/Go_Shell/blob/main/Go_Shell.png?raw=true)

This is a simple shell built in go designed to give you basic functionality such as changing directory and listing files. This is intended to be used in future projects, if need arise for a shell. Over time this shell will be worked on to account for tasks I am use often such as running nmap scans and sorting output.

## Install

If you wish to use any of the packages, the deirectory will have a readme containing all of the information you'll need to get started. To use the command line tool as I do I will explain how to do so below.

These instructions are for targeted towards linux user's.

### Building

To start things of this will require some knowledge of working with Go modules, and packages. As well as how to build and install programs. If you have any questions about Go or how to do these things I just so happen to have a repository [here](https://github.com/Syssos/Learning_Go) I created as I learned Go, leaving notes behind that may be useful to you.

Clone the package to a location you can use.

``` 
git clone https://github.com/Syssos/Go_Shell.git
```
Once the repository is cloned cd into it.
```
cd Go_Shell
```
Before we continue I need to note there is no go.mod file included in this repository so one will need to be created. The command I used to create my module is below.

```
go mod init github.com/Syssos/Go_Shell
```
If this is not the module you choose you will need to either run the ``` go get ``` command for the packages included, or change the import statement in the nested packages.

Once the go.mod file is created we can install it.

```
go install .
```

This will place a bin file in "$PATH/bin" called "Go_Shell", if the module name was used above. Running this command will give you the Go Shell.

For me this looks like
```
~/go/bin/Go_Shell
```

Alternativly you can install the "cmds" package with ``` go get ``` and create a script that utilizes the Loop function much like [main.go](https://github.com/Syssos/Go_Shell/blob/main/main.go) does.

## Main

This file will call the Loop function from the cmds package. This loop is what is responsible for all of the commands. The plans to re-use this code mean I need all of the code in one package, more or less, that I can grab and use in another project.

## Logging

This program will log errors to a file called "GoShellLogfile.txt". Due to the state of this shell the logging file location is undecided, for know the shell will live int the ```/tmp``` directory. Be aware that this file location may change in future version as there could be a need for it in a new location to better suit personal use cases.

## Travis

Very unhappy with this, .org worked the first time I used it no issues. the .com is the worst thing I've ever seen.
