/*
	Filelog is a package dedicated to logging files to a specific output file.

	The Flog struct is what gives a developer the ability to utilize this package, Multiple instances of this package can be created and used to log different insidents
	for different cases.

	As the cmds package is developed this pacackage will be used increasingly more. The ability to log command specific errors, to an appropriate log file, will become
	 more of a neccesicity
*/

package filelog

import (
	
	"os"
	"fmt"
	"time"
	"os/user"

	"github.com/Syssos/gofsh/src/color"
)

// Flog structure to keep track of needed material
type Flog struct {

	Greeting    string
	Salutation  string
	LogFile     string
	DtFormat    string
	DtLocation  *time.Location
	Errormsg    error
}

// Prints greeting message and logs when the shell instance was started
func (m Flog) Greet() {
	
	dt := time.Now().In(m.DtLocation)

	m.Log(fmt.Sprintf("%v - User %v started instance",  dt.Format(m.DtFormat), GetUser()))

	fmt.Println(m.Greeting, "\n")
}

// Prints salute message and logs when the shell instance is closed
func (m Flog) Salute() {

	dt := time.Now().In(m.DtLocation)

	m.Log(fmt.Sprintf("%v - User %v exited instance",  dt.Format(m.DtFormat), GetUser()))

	fmt.Println(m.Salutation)
}

func (m Flog) Logmsg(msg string) {

	dt := time.Now().In(m.DtLocation)

	m.Log(fmt.Sprintf("%v - %v",  dt.Format(m.DtFormat), msg))
}

// Prints error message and logs it
func (m Flog) Err() {

	dt := time.Now().In(m.DtLocation)

	m.Log(fmt.Sprintf("%v - %v",  dt.Format(m.DtFormat), m.Errormsg))

	fmt.Println(color.Red + fmt.Sprintf("%v", m.Errormsg) + color.Reset)
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