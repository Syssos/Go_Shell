package cmds

import (
	// "fmt"
	"os"
	"errors"
)

func Cd(args []string) error {
	if len(args) > 0 {
		os.Chdir(args[0])
    	return nil
	}
    return errors.New("empty string passed to cd")

 //    _, err := os.Getwd()
 //    if err != nil {
 //    	return errors.New(fmt.Sprintf("%v", err))
 //    }

}