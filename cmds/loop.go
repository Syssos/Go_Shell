package cmds

import (
	"github.com/Syssos/Go_Shell/color"
	"os"
	"os/user"
	"fmt"
	"bufio"
	"path"
	"errors"
	"strings"
)

var current_user string = getUser()
var current_dir string = getCurrentDir()
var greetingMessage string = fmt.Sprintf("Hello there, General %v. Welcome to the shit show", current_user)
var saluteMessage string = "Later homeboy"

func getUser() string{
	use, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	return use.Name
}

type messages struct {
	Greeting, Salutation, Errormsg, LogFile string
}

func (m messages) Greet() {

	m.Log(fmt.Sprintf("User %v started instance", current_user))
	fmt.Println(m.Greeting, "\n")
}

func (m messages) Salute() {

	m.Log(fmt.Sprintf("User %v exited instance", current_user))
	fmt.Println(m.Salutation)
}

func (m messages) Err() {

	m.Log(m.Errormsg)
	fmt.Println(m.Errormsg)
}

func (m messages) Log(msg string) {
	f, err := os.OpenFile(m.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
	    fmt.Println(err)
	}
	f.Write([]byte(msg+"\n"))
	f.Close()
}


var mes messages = messages{
	greetingMessage, 
	saluteMessage, 
	"",
	current_dir + "/logfile.txt",
}

func Loop() error{
	mes.Greet()
	suc, err := cmdLoop()
	if err != nil {
		fmt.Println(err)
	} else if suc == 0 {
		mes.Salute()
	}

	return nil
}

func getCurrentDir() string{
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return cwd
}

func cmdLoop() (int, error){

	for ;; {

		input := bufio.NewReader(os.Stdin)
		cwd := getCurrentDir()

		fmt.Printf("%v - %v: ", color.Red + current_user + color.Reset, color.Blue + path.Base(cwd) + color.Reset)
		in, _ := input.ReadString('\n')

		parsed_input := getCommands(in)

		if parsed_input[0] == "exit" {
			return 0, nil
		}
		_, cerr := runCommand(parsed_input[0], parsed_input[1:])
		if cerr != nil {
			mes.Errormsg = fmt.Sprintf("%v", cerr)
			mes.Err()
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