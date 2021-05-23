package cmds

import (
	// "fmt"
	"os"
	"errors"
)

type Cd_cmd struct {
	args []string
}

func (cmd Cd_cmd) Run() error {
	if len(cmd.args) > 0 {
		os.Chdir(cmd.args[0])
    	return nil
	}
	
    return errors.New("empty string passed to cd")
}

func Cd(args []string) error {
	if len(args) > 0 {
		os.Chdir(args[0])
    	return nil
	}
	
    return errors.New("empty string passed to cd")
}