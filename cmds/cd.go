package cmds

import (
	// "fmt"
	"os"
	"errors"
)

// Creating struct for cd command
type Cd_cmd struct {
	Args []string
}

// method to run the command
func (cmd Cd_cmd) Run() error {
	if len(cmd.Args) == 0 {
    	return errors.New("empty string passed to cd")
	}
	
	os.Chdir(cmd.Args[0])
    return nil
}