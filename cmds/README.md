# Commands

This package will contain a majority, if not all of the commands utilized by the shell.

While this package will use a filelog.Flog instance from a [Flog](https://github.com/Syssos/Go_Shell/blob/main/filelog/filelog.go) struct, it is still importable and usable in other coding project, for whatever reason a shell would be needed.

## Install

To install this package use the go get command, this should also pull in the filelog package. 

```
go get github.com/Syssos/Go_Shell/filelog
```

## Usage
To start using the shell, in another file, we must start by creating an instance of a [Loop](https://github.com/Syssos/Go_Shell/blob/main/cmds/loop.go). This will require a Flog instance, as well as a cmds.Commands struct instace.

For an example you can look at the main function, within the [main](https://github.com/Syssos/Go_Shell/blob/main/main.go) file at the root of the repo.

In this example we see that we use both a Flog instance
```go
flog := LoggerFromFile() //filelog.Flog{}
loop := cmds.Loop{flog, command_struct}

loopErr := loop.Run()
```
The Flog instance uses a command_struct, which is essentually a struct, of structs, and is what holds the commands the shell has access to.

The loop will run until an unhandled error, or exit is triggered.