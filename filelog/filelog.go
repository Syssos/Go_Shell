package filelog

import (
	"os"
	"os/user"
	"fmt"
	"time"
)

// Getting current user for messages
var current_user string = GetUser()

// Creating messages to display to the user
var greetingMessage string = fmt.Sprintf("Hello there, General %v. Welcome to the Go Shell", current_user)
var saluteMessage string = "Later homeboy"
var logfile string = "/tmp/GoShellLogfile.txt"

// Flog structure to keep track of needed material
type Flog struct {
	Greeting, Salutation, LogFile string
	Errormsg error
}

// Prints greeting message and logs when the shell instance was started
func (m Flog) Greet() {
	dt := time.Now()
	m.Log(fmt.Sprintf("%v - User %v started instance",  dt.Format("01-02-2006 15:04:05"), current_user))
	fmt.Println(m.Greeting, "\n")
}

// Prints salute message and logs when the shell instance is closed
func (m Flog) Salute() {
	dt := time.Now()
	m.Log(fmt.Sprintf("%v - User %v exited instance",  dt.Format("01-02-2006 15:04:05"), current_user))
	fmt.Println(m.Salutation)
}

// Prints error message and logs it
func (m Flog) Err() {
	dt := time.Now()
	m.Log(fmt.Sprintf("%v - %v",  dt.Format("01-02-2006 15:04:05"), m.Errormsg))
	fmt.Println(m.Errormsg)
}

// Command that opens log file and writes content to it.
func (m Flog) Log(msg string) {
	f, err := os.OpenFile(m.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
	    fmt.Println(err)
	}
	f.Write([]byte(msg+"\n"))
	f.Close()
}

// Command to get the current users name
func GetUser() string{
	use, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	return use.Name
}

// Creates instance of Flog to use in code, multiple Flog instances can be used to log to multiple files
func F_init() Flog {
	var mes Flog = Flog{
		greetingMessage, 
		saluteMessage, 
		logfile,
		nil,
	}
	return mes
}

// Gets the current user direectory 
func GetHomeDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	return dirname
}

// Gets the current working directory for the current directory the user is in when the program is ran.
func GetCurrentDir() string{
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return cwd
}