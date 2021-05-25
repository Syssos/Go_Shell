package cmds

import (
	"github.com/Syssos/Go_Shell/color"
	"github.com/Syssos/Go_Shell/filelog"
	"github.com/Syssos/Go_Shell/microservices/isrunning"
	"os"
	"fmt"
	"bufio"
	"path"
	"errors"
	"strings"
)

// Setting up logging
var flog filelog.Flog = filelog.F_init()

// to change log file
// flog.LogeFile = "path/to/file"

// Setting up commands from cmds directory
var cd Cd_cmd = Cd_cmd{[]string{}}
var pwd Pwd_cmd = Pwd_cmd{[]string{}}
var ls Ls_cmd = Ls_cmd{[]string{}}

// Setting up command from another package then cmds
var running isrunning.Isrunning_cmd = isrunning.Isrunning_cmd{[]string{}, "", false}

// Creating command struct to hold available commands
var command_struct commands = commands{ls, pwd, cd, running}

// Commands struct will be a "list" of commands shell can run
type commands struct {
	ls Ls_cmd
	pwd Pwd_cmd
	cd Cd_cmd
	running isrunning.Isrunning_cmd
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

// Interface used to print usage for command
type info interface {
	Usage()
}

// Function that prints usage
func PrintUsage(i info) {
	i.Usage()
}

// Start of command line process, responsible for greeting and salute
func Loop() error{
	/*
		Logs start/finish of program, runs command loop
		Return: Error if command loop ran into problems
	*/

	// Logging when the program is starting/Printing greeting message
	flog.Greet()
	suc, err := cmdLoop()
	if err != nil {
		fmt.Println(err)
	} else if suc == 0 {
		// Logging when the user ends the program (only works with "exit")
		flog.Salute()
	}

	return nil
}

// Inifinite loop, runs until told by "exit" command
func cmdLoop() (int, error){
	/*
		Infinite loop, displays input line, everything in loop runs each time enter key is pressed
		Return: Int representation of success or failure, error if error occurs

	*/

	for ;; {

		cwd := filelog.GetCurrentDir()
		cuser := filelog.GetUser()
		fmt.Printf("%v - %v: ", color.Cyan + cuser + color.Reset, color.Blue + path.Base(cwd) + color.Reset)

		input := bufio.NewReader(os.Stdin)
		in, _ := input.ReadString('\n')
		parsed_input := createCmdSlice(in)
		if parsed_input[0] == "exit" {
			return 0, nil
		} else if parsed_input[0] == "help" {
			_, helperr := helpCommand(parsed_input[1])
			if helperr != nil {
				flog.Errormsg = helperr
				flog.Err()
			}
		} else {
			_, cerr := runCommand(parsed_input[0], parsed_input[1:])
			if cerr != nil {
				flog.Errormsg = cerr
				flog.Err()
			}
		}	
	}
	return 1, errors.New("End of loop")
}

// creates parsed slice from string 
func createCmdSlice(cmd string) []string {
	/*
		Parses input from user and turns it into string slice
		Return: returns string slice of commands and flags

	*/

	commands := []string{}

	words := strings.Fields(cmd)
	for _, word := range words {
		commands = append(commands, word)
	}

	return commands
}

// Runs the command based on cmd string
func runCommand(cmd string, args []string) (int, error) {
	/*
		Switch statement responsible for executing command in command struct
		Return: int representation of err or success, Error if error occurs

	*/

	switch cmd {
	case "ls":
		command_struct.ls.Args = args
		execute(command_struct.ls)
		return 0, nil
	case "pwd":
		command_struct.pwd.Args = args
		execute(command_struct.pwd)
		return 0, nil
	case "cd":
		command_struct.cd.Args = args
		execute(command_struct.cd)
		return 0, nil
	case "cody":
		if len(args) > 0 {
			command_struct.running.Args = args
			command_struct.running.Link = args[0]
			execute(command_struct.running)
			return 0, nil
		}
		fmt.Println("Added link")
		return 0, nil
	default:
		return 1, errors.New(fmt.Sprintf("Command not found: %v", cmd))
	}
}

// Prints help messages for commands
func helpCommand(cmd string) (int, error){
	/*
		Switch statement responsible for printing command usage
		Return: int representation of err or success, Error if error occurs

	*/

	switch cmd {
	case "ls":
		PrintUsage(command_struct.ls)
		return 0, nil
	case "pwd":
		PrintUsage(command_struct.pwd)
		return 0, nil
	case "cd":
		PrintUsage(command_struct.cd)
		return 0, nil
	case "cody":
		PrintUsage(command_struct.running)
		return 0, nil
	default:
		return 1, errors.New(fmt.Sprintf("Command not found: %v", cmd))
	}
}
