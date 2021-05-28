# Go Shell [![Build Status](https://travis-ci.com/Syssos/Go_Shell.svg?branch=main)](https://travis-ci.com/github/Syssos/Go_Shell)  [![Cody Code](https://syssos.app/static/images/index/cody_code.svg)](https://syssos.app)

<p align="center">
  <img src="https://raw.githubusercontent.com/Syssos/Go_Shell/main/settings/images/Go_Shell.png" alt="Go Shell img"/>
</p>

Go Shell is an interactive shell, aimed at automating tasks often used for creating, or testing web based applications. 

### Dependancies for package:
- [Go](https://golang.org/)
- TOML parsing package ([komkom/toml](https://github.com/komkom/toml))

Simple linux commands are included to help navigate through directories.

## Table of Contents
- [Overview](#go-shell)
- [Table of Contents](#table-of-contents)
- [Install](#install)
- [Usage](#usage)
  * [Core Commands](#core-commands)
    - [help](#help)
    - [exit](#exit)
    - [ls](#ls)
    - [pwd](#pwd)
    - [cd](#cd)
  * [Website Oriented Commands](#website-oriented-commands)
    - [site](#site)
- [Logging](#logging)
- [Travis](#travis)

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
<p align="center">
  <img src="https://raw.githubusercontent.com/Syssos/Go_Shell/main/settings/images/Shell.PNG" alt="Go Shell img"/>
</p>

### Core Commands

Each command structure has a Usage() method which is called when the "help" command is ran. To use this feature, enter ```help``` followed by the command you need help with.

### Help
The help command is designed to work with every command the loop is able to run.

```
$ help ls
```

### Exit
This exits the shell properly and allows for logging of the user quiting.

```
$ exit
```
### ls
This command should work just about the same as it does an a native linux system, minus flag functionality. While it takes a location to list it cannot except other flags at the moment.

```
$ ls
$ ls ../
$ ls Go_Shell/
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

The site command is a command that allows for the user to see a status code for a specific url. 

The purpose of this is mainly as a check to ensure the web application or program we are working with is returning an "OK" status.

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
<p align="center">
  <img src="https://raw.githubusercontent.com/Syssos/Go_Shell/main/settings/images/Error_Log_Example.PNG" alt="Go Shell img"/>
</p>

Example held in [settings/example_log.txt](https://github.com/Syssos/Go_Shell/blob/main/settings/example_log.txt)

This program utilizes toml formated data saved in a [.toml](https://github.com/Syssos/Go_Shell/blob/main/settings/cmds.toml).

If you are unfamiliar with the TOML format, more information can be found on their [github](https://github.com/toml-lang/toml) page.

The settings in the file mentioned above are only used for the filelog package. The "Greeting" and "Salute" variables are likely to get removed for lack of relevance to the logger. Other settings refer to datatime settings and location settings.


## Travis

Travis-CI is a bonus feature of this repository. Due to the nature of this project the code can change at a fast pace. To prevent errors from occuring to someone who clones the repo, the travis build state is indicated at the top of this readme. I have multiple machines and need to share code between them. Travis allows me to do that and ensure the code is working at the same time.
