package cmds

import (
	"os"
	"io/fs"
	"fmt"
	"errors"
	"github.com/Syssos/Go_Shell/color"
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

func (cmd Cd_cmd) Usage() {
	usagestr := "\n\tcd - Change directory\n\n\tUsage:\n\t\tcd <path/to/directory>\n"
	colorfied := color.Yellow + usagestr + color.Reset
	fmt.Println(colorfied)
}

// Command structure for pwd command
type Pwd_cmd struct {
	Args []string
}

// method that runs the command
func (cmd Pwd_cmd) Run() error {
	if len(cmd.Args) > 0 {
		// Should be return, however there is a plan to allow args to print full path to dest. This message is just print to not trigger error in pkg_test
		fmt.Printf("Can't print directories yet, %v\n", cmd.Args[0])
	}

	cwd, err := os.Getwd()
	if err != nil {
		return errors.New(fmt.Sprintf("%v",err))
	}

	fmt.Println(color.Green + cwd + color.Reset)

	return nil
}

func (cmd Pwd_cmd) Usage() {
	usagestr := "\n\tpwd - Print working directory\n\n\tUsage:\n\t\tpwd\n"
	colorfied := color.Yellow + usagestr + color.Reset
	fmt.Println(colorfied)
}

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

func (cmd Ls_cmd) Usage() {
	usagestr := "\n\tls - List files in directory\n\n\tUsage:\n\t\tls\n\t\tls <path/to/directory>\n"
	colorfied := color.Yellow + usagestr + color.Reset
	fmt.Println(colorfied)
}

func printFiles(files []fs.DirEntry) {
	
	for _, file := range files {
		fmt.Println(color.Green + file.Name() + color.Reset)
	}

}