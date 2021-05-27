/*
	This package will run an http request to a specific website to ensure it is working how it should
	This will be turned into a way to test websites I create to make sure there is no errors when accessing the webserver.

	This package will later be converted into, or used in, another package that allows for website route enumeration via a "common routes" txt file.
*/

package cmds

import (
	"fmt"
	"net/http"
	"errors"
	"github.com/Syssos/Go_Shell/color"
)

// Struct for keeping track of the site command data while program is running
type Site_cmd struct {
	Args []string
	Link string
	Active bool
}

// Method for handling the "site" command call from "cmds/loop.go"
func (cmd *Site_cmd) Run() error{

	if len(cmd.Args) > 0 {
		
		if cmd.Args[0] == "status" {
		
			statusError := cmd.Status()
			if statusError != nil {
				return statusError
			}
		
		} else {
		
			cmd.Link = cmd.Args[0]
			
			statusError := cmd.Status()
			if statusError != nil {
				return statusError
			}
		}	    
	} else {
	
		return errors.New("Incorrect use of site arguments check the usage by running 'help site'")
	}
	
	return nil
}

// Prints usage of "site" command
func (cmd Site_cmd) Usage() {

	usagestr := "\n\tsite - Sends get request to server, if server returns 200, status is considered active. As of right now anything other then 200 will be seen as an error, including redirects such as 301.\n\n\tUsage:\n\t\tsite status\n\t\tsite <url>\n"
	colorfied := usagestr
	fmt.Println(colorfied)
}

// Fucntion used for checking the 
func (cmd *Site_cmd) Status() error {

	if cmd.Link != "" {
	
		resp, resErr := http.Get(cmd.Link)
		if resErr != nil {
			return resErr
		}

		// If the first 3 characters in the response status are 200, note resp.Status is type string
		if string([]rune(resp.Status)[0:3]) == "200" {
	
	    	cmd.Active = true
			fmt.Println(color.Green + fmt.Sprintf("200 - Site %v is active", cmd.Link) + color.Reset)
	    } else {
			
	    	cmd.Active = false
	    	fmt.Println(color.Red + fmt.Sprintf("%v - Site %v could have issues", resp.Status, cmd.Link) + color.Reset)
	    }
	
	} else {
	
		return errors.New("Please set link before checking status, for help check usage")
	}
	
	return nil
}