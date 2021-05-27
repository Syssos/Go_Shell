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


// Setting up logging
// var Flog filelog.Flog = filelog.F_init()

// Commands struct will be a "list" of commands shell can run
type Commands struct {
	Ls Ls_cmd
	Pwd Pwd_cmd
	Cd Cd_cmd
	Site Site_cmd
}

// Structure for Loop, Responsible for keeping track of commands and logger
type Loop struct {
	Command_struct Commands
	Flog filelog.Flog
}

// Start of command line process, responsible for greeting and salute
func (l *Loop) Run() error{
	/*
		Logs start/finish of program, runs command loop
		Return: Error if command loop ran into problems
	*/

	// Logging when the program is starting/Printing greeting message
	l.Flog.Greet()
	suc, err := l.cmdLoop()
	
	if err != nil {
		fmt.Println(err)
	} else if suc == 0 {
		// Logging when the user ends the program (only works with "exit")
		l.Flog.Salute()
	}

	return nil
}

// Inifinite loop, runs until told by "exit" command
func (l *Loop)cmdLoop() (int, error){
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
		
			_, helperr := l.helpCommand(parsed_input[1])
			if helperr != nil {
				l.Flog.Errormsg = helperr
				l.Flog.Err()
			}
		} else {
		
			_, cerr := l.runCommand(parsed_input[0], parsed_input[1:])
			if cerr != nil {
				l.Flog.Errormsg = cerr
				l.Flog.Err()
			}
		}	
	}

	return 1, errors.New("End of loop")
}

// Runs the command based on cmd string
func (l *Loop) runCommand(cmd string, args []string) (int, error) {
	/*
		Switch statement responsible for executing command in command struct
		Return: int representation of err or success, Error if error occurs

	*/

	switch cmd {
	case "ls":
		l.Command_struct.Ls.Args = args
		lsErr := execute(l.Command_struct.Ls)
		l.hasError(lsErr)
		return 0, nil
	case "pwd":
		l.Command_struct.Pwd.Args = args
		pwdErr := execute(l.Command_struct.Pwd)
		l.hasError(pwdErr)
		return 0, nil
	case "cd":
		l.Command_struct.Cd.Args = args
		cdErr := execute(l.Command_struct.Cd)
		l.hasError(cdErr)
		return 0, nil
	case "site":
		l.Command_struct.Site.Args = args
		siteErr := execute(&l.Command_struct.Site)
		l.hasError(siteErr)
		return 0, nil
	default:
		return 1, errors.New(fmt.Sprintf("Command not found: %v", cmd))
	}
}

// Prints help messages for commands
func (l *Loop)helpCommand(cmd string) (int, error){
	/*
		Switch statement responsible for printing command usage
		Return: int representation of err or success, Error if error occurs

	*/

	switch cmd {
	case "ls":
		PrintUsage(l.Command_struct.Ls)
		return 0, nil
	case "pwd":
		PrintUsage(l.Command_struct.Pwd)
		return 0, nil
	case "cd":
		PrintUsage(l.Command_struct.Cd)
		return 0, nil
	case "site":
		PrintUsage(l.Command_struct.Site)
		return 0, nil
	default:
		return 1, errors.New(fmt.Sprintf("Command not found: %v", cmd))
	}
}

func (l *Loop) hasError(err error) {
	if err != nil {
		l.Flog.Errormsg = err
		l.Flog.Err()	
	}
}

// used to execute every command in command struct
type exec interface {
	Run() error
}

// responsible for site and logging errors with command
func execute(e exec) error {
	err := e.Run()
	if err != nil {
		return errors.New(fmt.Sprintf("%v", err))
	}
	return nil
}

// Interface used to print usage for command
type info interface {
	Usage()
}

// Function that prints usage
func PrintUsage(i info) {
	i.Usage()
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
