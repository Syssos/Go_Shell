/*
	The cmds package is responsible for holding all of the basic shell commands. To add to the commands the shell has access to they will for now be part of the cmds package.

	The Loop will consist of 2 main components, a logger and a command struct. The command struct will be created out of structures that represent commands, this allows
	for arguments to be passed easily and an easy command execution method. The logger used for this command loop is another structure known as Flog. These two components
	components will not need to change much as the shell is running.
*/

package cmds

import (
	
	"os"
	"fmt"
	"path"
	"bufio"
	"errors"
	"strings"

	"github.com/Syssos/Go_Shell/color"
	"github.com/Syssos/Go_Shell/filelog"
)

// Holds structures for each command loop can run for easier access
type Commands struct {
	
	Ls   Ls_cmd
	Pwd  Pwd_cmd
	Cd   Cd_cmd
	Site Site_cmd
}

// Responsible for keeping track of commands and logger
type Loop struct {

	Flog           filelog.Flog
	Command_struct Commands
}
// Where command loop starts, greeting and salute messages are displayed here
func (l *Loop) Run() error{

	l.Flog.Greet()
	suc, err := l.cmdLoop()
	if err != nil {

		l.Flog.Errormsg = err
		l.Flog.Err()

	} else if suc == 0 {
		
		l.Flog.Salute()
	}

	return nil
}
// Inifinite loop, runs until told by "exit" command
func (l *Loop)cmdLoop() (int, error){

	for ;; {

		cwd   := filelog.GetCurrentDir()
		cuser := filelog.GetUser()
		
		// Printing what user see's before their input on command line
		fmt.Printf("%v - %v: ", color.Cyan + cuser + color.Reset, color.Blue + path.Base(cwd) + color.Reset)

		input := bufio.NewReader(os.Stdin)
		in, _ := input.ReadString('\n')
		
		parsed_input := createCmdSlice(in)
		
		if len(parsed_input) > 0 {
			if parsed_input[0] == "exit" {
			
				return 0, nil

			} else if parsed_input[0] == "help" {
			
				if len(parsed_input) > 1 {
					_, helpErr := l.helpCommand(parsed_input[1])
					if helpErr != nil {
						l.Flog.Errormsg = helpErr
						l.Flog.Err()
					}
				} else {

					l.Flog.Errormsg = errors.New("No command specified to help with")
					l.Flog.Err()
				}
				
			} else {
			
				_, cerr := l.runCommand(parsed_input[0], parsed_input[1:])
				if cerr != nil {
					l.Flog.Errormsg = cerr
					l.Flog.Err()
				}
			}	
		} else {
			// Blank input was passed, restarting loop to collect input
		}
	}

	return 1, errors.New("End of loop")
}
// Runs the command based on cmd string
func (l *Loop) runCommand(cmd string, args []string) (int, error) {

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
// Used to check for errors in runCommand
func (l *Loop) hasError(err error) {

	if err != nil {
		l.Flog.Errormsg = err
		l.Flog.Err()	
	}
}

// Interface used to print usage for command
type info interface {
	Usage()
}

// Function that prints usage, utilizes info
func PrintUsage(i info) {
	i.Usage()
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

// creates parsed slice from string based off spaces
func createCmdSlice(cmd string) []string {

	commands := []string{}

	words := strings.Fields(cmd)
	for _, word := range words {
		commands = append(commands, word)
	}

	return commands
}
