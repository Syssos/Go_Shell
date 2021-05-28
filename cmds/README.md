# cmds package

This package will contain a majority, if not all of the commands utilized by the shell.

While this package will use a filelog.Flog instance from a [Flog](https://github.com/Syssos/Go_Shell/blob/main/filelog/filelog.go) struct, it is still importable and usable in other coding project, for whatever reason a shell would be needed.

## Install

To install this package use the go get command

```bash
go get github.com/Syssos/Go_Shell/cmds
```

This package will also require the filelog package as it is a dependancy for the loop struct. To install that package you can use go get as well.

```bash
go get github.com/Syssos/Go_Shell/filelog
```

## Usage
Every command in this file can be imported and ran in other projects if need be.

This can be done by initualizing an instance of the command struct, and then calling the `Run()` method.

an example of this, using the cmds.pwd command

```go
package main

import (
	"fmt"

	"github.com/Syssos/Go_Shell/cmds"
)

func main () {
	pwd_command := cmds.Pwd_cmd{[string]{}}
	// pwd_command.Args = []string{"arguments", "command", "can", "handle"}

	pwd_command.Run()
}
```

This method should work for every command in the `cmds` package, as they are all designed to be called based off of the Run() method.

Keep in mind that the `cmd.Args` value needs to be a slice of commands that are of type string. The loop is responsible for parsing, so a string line as shown below, will not work.

```go
pwd_command.Args = "arguments command can handle" //will not work, mismatched type string and []string
```