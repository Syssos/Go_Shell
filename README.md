# Go Shell [![Build Status](https://travis-ci.com/Syssos/Go_Shell.svg?branch=main)](https://travis-ci.com/github/Syssos/Go_Shell)  [![Cody Code](https://syssos.app/static/images/index/cody_code.svg)](https://syssos.app)

<div style="text-align:center" alt="Go Shell img"><img src=".https://raw.githubusercontent.com/Syssos/Go_Shell/main/settings/images/Go_Shell.png" /></div>
<!-- ![Go Shell img](https://raw.githubusercontent.com/Syssos/Go_Shell/main/settings/images/Go_Shell.png) -->

Go Shell is an interactive shell, aimed at automating tasks often used for creating, or testing web based applications or programs. Simple linux commands are included to help navigate through directories.

### Dependancies for package:

	- [Go](https://golang.org/)
	- TOML parsing package ([komkom/toml](https://github.com/komkom/toml))

## Install

All of the packages used in this repository were designed with external usablility in mind. Each package directory will contain a readme with all of the information you'll need to get started. The directions below should get you started with using the shell as.

These instructions are for targeted towards linux user's.

To start things off this will require some knowledge of working with Go modules, and packages. As well as how to build and install programs. If you have any questions about Go or how to do these things I just so happen to have a repository [here](https://github.com/Syssos/Learning_Go) I created as I learned Go, leaving notes behind that may be useful to you.

To get the files we need clone the package to a location you can use.

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

## Usage
![Go Shell Example](https://raw.githubusercontent.com/Syssos/Go_Shell/main/settings/images/Shell.PNG)

### help
The help command is designed to work with every command the loop has access to. Each command has a Usage() method which is called when help is ran. To use this feature use the command help followed by the command you need help with.

```
$ help ls
```

### exit
This exits the shell properly and allows for logging of the user quiting.

```
$ exit
```

## Basic Commads

### ls
This command should work just about the same as it does an a native linux system, minus advanced functionality. While it takes a location to list it cannot except flags at the moment.

```
$ ls
$ ls ../
$ ls /Go_Shell
```
### pwd
This command will print the working directory. Its a pretty straight forward command and doesn't take arguments.

```
$ pwd
```
### cd
This command changes the working directory much like cd in linux, like ls this will take a location to change to but will not accept flags.

```
$ cd ../
$ cd Go_Shell/
```

## Website Oriented Commands

The commands under this section will be related to http request based commands.

### site

The site command is a command that allows for the user to see a status code for a specific url. The purpose of this is mainly as a check to ensure the web application or program we are working with is returning an "OK" status.

```
$ site https://github.com/Syssos/Go_Shell
```

After the a url is entered the site command will "remember" the url, if at any point in that shell instance you want the status of that site, the word status can be used to indicate it.

```
$ site status
```
the output for both of these commands should be

```
200 - Site https://github.com/Syssos/Go_Shell is active
```


## Logging

This program will log errors to a file called "GoShellLogfile.txt". Due to the state of this shell the logging file location is undecided, for know the shell will live int the ```/tmp``` directory. Be aware that this file location may change in future version as there could be a need for it in a new location to better suit personal use cases.

![Logging Example](https://raw.githubusercontent.com/Syssos/Go_Shell/main/settings/images/Error_Log_Example.PNG)

## Travis

Travis-CI is a bonus feature of this repository. Due to the nature of this project the code can change at a fast pace. To prevent errors from occuring to someone who clones the repo, the travis build state is indicated at the top of this readme. I have multiple machines and need to share code between them. Travis allows me to do that and ensure the code is working at the same time.
