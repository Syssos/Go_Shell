package cmds

import (
	"fmt"
	"os"
	"io/fs"
	"errors"
	"github.com/Syssos/Go_Shell/color"
)

// Creating struct for ls command
type Ls_cmd struct {
	Args []string
}

// Method to run the ls command
func (cmd Ls_cmd) Run() error {

	if len(cmd.Args) > 1{
		for _, arg := range cmd.Args {
			if string([]rune(arg)[0]) == "-" {
				fmt.Printf("%v: is a flag statement\n", string([]rune(arg)[1:]))
			}
		}
	} else if len(cmd.Args) == 1 {

		files, err := os.ReadDir(cmd.Args[0])
		
		if err != nil {
			return errors.New(fmt.Sprintf("%v",err))
		}

		printFiles(files)

		return nil
	}
		
	files, err2 := os.ReadDir(".")

	if err2 != nil {
		return errors.New(fmt.Sprintf("%v",err2))
	}

	printFiles(files)

	return nil
}

func printFiles(files []fs.DirEntry) {
	
	for _, file := range files {
		fmt.Println(color.Purple + file.Name() + color.Reset)
	}

}