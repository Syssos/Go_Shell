package cmds

import (
	"fmt"
	"os"
	"errors"
	"github.com/Syssos/Go_Shell/color"
)

// Creating struct for ls command
type Ls_cmd struct {
	args []string
}

// Method to run the ls command
func (cmd Ls_cmd) Run() error {
	for _, arg := range cmd.args {
		if string([]rune(arg)[0]) == "-" {
			fmt.Printf("%v: is a flag statement\n", string([]rune(arg)[1:]))
		}
	}
	files, err := os.ReadDir(".")

	if err != nil {
		return errors.New(fmt.Sprintf("%v",err))
	}

	for _, file := range files {
		fmt.Println(color.Purple + file.Name() + color.Reset)
	}

	return nil
}