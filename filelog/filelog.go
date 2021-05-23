package filelog

import (
	"os"
	"os/user"
	"fmt"
)

var current_user string = GetUser()
var current_dir string = GetCurrentDir()
var homeDir string = GetHomeDir()
var greetingMessage string = fmt.Sprintf("Hello there, General %v. Welcome to the shit show", current_user)
var saluteMessage string = "Later homeboy"

type Flog struct {
	Greeting, Salutation, Errormsg, LogFile string
}

func (m Flog) Greet() {

	m.Log(fmt.Sprintf("User %v started instance", current_user))
	fmt.Println(m.Greeting, "\n")
}

func (m Flog) Salute() {

	m.Log(fmt.Sprintf("User %v exited instance", current_user))
	fmt.Println(m.Salutation)
}

func (m Flog) Err() {

	m.Log(m.Errormsg)
	fmt.Println(m.Errormsg)
}

func (m Flog) Log(msg string) {
	f, err := os.OpenFile(m.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
	    fmt.Println(err)
	}
	f.Write([]byte(msg+"\n"))
	f.Close()
}

func GetUser() string{
	use, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	return use.Name
}

func F_init() Flog {
	var mes Flog = Flog{
		greetingMessage, 
		saluteMessage, 
		"",
		"/tmp/GoShellLogfile.txt",
	}

	return mes
}

func GetHomeDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	return dirname
}

func GetCurrentDir() string{
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return cwd
}