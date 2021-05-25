/*
	This package will run an http request to a specific website to ensure it is working how it should
	This will be turned into a way to test websites I create to make sure there is no errors when accessing the webserver.
*/

package isrunning

import (
	"fmt"
	"net/http"
	// "io/ioutil"
	// "log"
	"errors"
	"github.com/Syssos/Go_Shell/color"
)

type Isrunning_cmd struct {
	Args []string
	Link string
	Active bool
}

func (cmd *Isrunning_cmd) Run() error{
	if len(cmd.Args) > 0 {
		if cmd.Args[0] == "status" {
			if cmd.Link != "" {
				switch cmd.Active {
				case true:
					fmt.Println(color.Green + fmt.Sprintf("Site %v is active", cmd.Link) + color.Reset)
				case false:
					fmt.Println(color.Red + fmt.Sprintf("Site %v is not active", cmd.Link) + color.Reset)
				}
			} else {
				return errors.New("Please set link before checking status, for help check usage")
			}
		} else {
			cmd.Link = cmd.Args[0]
			resp, resErr := http.Get(cmd.Args[0])
			if resErr != nil {
				// handle error
				return resErr
			}
		
		    if resp.Status == "200 OK" {
		    	cmd.Active = true
				fmt.Println(color.Green + fmt.Sprintf("200 - Site %v is active", cmd.Link) + color.Reset)
		    } else {
		    	cmd.Active = false
		    	fmt.Println(color.Red + fmt.Sprintf("%v - Site %v could have issues", resp.Status, cmd.Link) + color.Reset)

		    }

		}	    
	 //    body, err := ioutil.ReadAll(resp.Body)
	 //    if err != nil {
		// 	fmt.Println(err)
	 //    }
		// //Convert the body to type string
	
	 //    sb := string(body)
	 //    log.Printf(sb)
	    
	} else {
		fmt.Println(cmd)
		fmt.Println("No Instructions as to what site should do, add arguments, or use help site")
	}
	return nil
}

func (cmd Isrunning_cmd) Usage() {
	usagestr := "\n\tsite - Sends get request to server, if server exists and returns 200, status is considered active\n\n\tUsage:\n\t\tsite status\n\t\tsite <url>\n"
	colorfied := usagestr
	fmt.Println(colorfied)
}
