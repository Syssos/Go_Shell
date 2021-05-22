package cmds

import (
	"fmt"
	"os"
	"errors"
)

func Pwd(args []string) error{

	cwd, err := os.Getwd()
	if err != nil {
		return errors.New(fmt.Sprintf("%v",err))
	}

	fmt.Println(cwd)

	return nil
}