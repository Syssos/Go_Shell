package cmds

import (
	"github.com/Syssos/Go_Shell/color"
	"github.com/Syssos/Go_Shell/filelog"
	"os"
	"fmt"
	"bufio"
	"path"
	"errors"
	"strings"
)

var current_user string = filelog.GetUser()
var current_dir string = filelog.GetCurrentDir()
var flog filelog.Flog = filelog.F_init()

func Loop() error{
	
	flog.Greet()
	suc, err := cmdLoop()
	if err != nil {
		fmt.Println(err)
	} else if suc == 0 {
		flog.Salute()
	}

	return nil
}

func cmdLoop() (int, error){

	for ;; {

		input := bufio.NewReader(os.Stdin)
		cwd := filelog.GetCurrentDir()

		fmt.Printf("%v - %v: ", color.Red + current_user + color.Reset, color.Blue + path.Base(cwd) + color.Reset)
		in, _ := input.ReadString('\n')

		parsed_input := getCommands(in)

		if parsed_input[0] == "exit" {
			return 0, nil
		}
		_, cerr := runCommand(parsed_input[0], parsed_input[1:])
		if cerr != nil {
			flog.Errormsg = fmt.Sprintf("%v", cerr)
			flog.Err()
		}
	}
	return 1, errors.New("End of loop")
}

func getCommands(cmd string) []string {
	commands := []string{}

	words := strings.Fields(cmd)
	for _, word := range words {
		commands = append(commands, word)
	}

	return commands
}

func runCommand(cmd string, args []string) (int, error) {
	switch cmd {
	case "ls":
		lsErr := Ls(args)
		if lsErr != nil {
			return 1, errors.New(fmt.Sprintf("ls: %v", lsErr))
		}
		return 0, nil
	case "pwd":
		pwdErr := Pwd(args)
		if pwdErr != nil {
			return 1, errors.New(fmt.Sprintf("pwd: %v", pwdErr))
		}
		return 0, nil
	case "cd":
		cdErr := Cd(args)
		if cdErr != nil {
			return 1, errors.New(fmt.Sprintf("pwd: %v", cdErr))
		}
		return 0, nil
	default:
		return 1, errors.New(fmt.Sprintf("Command not found: %v", cmd))
	}
}