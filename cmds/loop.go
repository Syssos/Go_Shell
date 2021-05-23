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

// Setting up logging/getting current user data
var current_user string = filelog.GetUser()
var current_dir string = filelog.GetCurrentDir()
var flog filelog.Flog = filelog.F_init()

// Setting up commands
var cd Cd_cmd = Cd_cmd{[]string{}}
var pwd Pwd_cmd = Pwd_cmd{[]string{}}
var ls Ls_cmd = Ls_cmd{[]string{}}
var command_struct commands = commands{ls, pwd, cd}

// Commands struct will be a "list" of commands shell can run
type commands struct {
	ls Ls_cmd
	pwd Pwd_cmd
	cd Cd_cmd
}
// used to execute every command in command struct
type exec interface {
	Run() error
}
// responsible for running and logging errors with command
func execute(e exec) {
	err := e.Run()
	if err != nil {
		flog.Errormsg = errors.New(fmt.Sprintf("cmd: %v", err))
		flog.Err()
	}
}

func Loop() error{
	/*
		Logs start/finish of program, runs command loop
		Return: Error if command loop ran into problems
	*/

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
	/*
		Infinite loop, displays input line, everything in loop runs each time enter key is pressed
		Return: Int representation of success or failure, error if error occurs

	*/

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
			flog.Errormsg = cerr
			flog.Err()
		}
	}
	return 1, errors.New("End of loop")
}

func getCommands(cmd string) []string {
	/*


	*/

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
		command_struct.ls.args = args
		execute(command_struct.ls)
		return 0, nil
	case "pwd":
		command_struct.pwd.args = args
		execute(command_struct.pwd)
		return 0, nil
	case "cd":
		command_struct.cd.args = args
		execute(command_struct.cd)
		return 0, nil
	default:
		return 1, errors.New(fmt.Sprintf("Command not found: %v", cmd))
	}
}