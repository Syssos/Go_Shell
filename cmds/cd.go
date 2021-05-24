package cmds

import (
	// "fmt"
	"os"
	"errors"
)

// Creating struct for cd command
type Cd_cmd struct {
	args []string
}

// method to run the command
func (cmd Cd_cmd) Run() error {
	if len(cmd.args) > 0 {
		os.Chdir(cmd.args[0])
    	return nil
	}
	
    return errors.New("empty string passed to cd")
}