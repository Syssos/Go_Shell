package cmds

import (
	"fmt"
	"os"
	"errors"
)

// Command structure for pwd command
type Pwd_cmd struct {
	Args []string
}

// method that runs the command
func (cmd Pwd_cmd) Run() error {
	if len(cmd.Args) > 0 {
		fmt.Printf("Can't print directories yet, %v\n", cmd.Args[0])
	}

	cwd, err := os.Getwd()
	if err != nil {
		return errors.New(fmt.Sprintf("%v",err))
	}

	fmt.Println(cwd)

	return nil
}