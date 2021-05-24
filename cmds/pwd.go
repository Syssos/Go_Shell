package cmds

import (
	"fmt"
	"os"
	"errors"
)

// Command structure for pwd command
type Pwd_cmd struct {
	args []string
}

// method that runs the command
func (cmd Pwd_cmd) Run() error {
	if len(cmd.args) > 0 {
		return errors.New(fmt.Sprintf("Can't print directories yet, %v", cmd.args[0]))
	}

	cwd, err := os.Getwd()
	if err != nil {
		return errors.New(fmt.Sprintf("%v",err))
	}

	fmt.Println(cwd)

	return nil
}