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
	// "errors"
)

type Isrunning_cmd struct {
	Args []string
	Link string
	Active bool
}

func (cmd Isrunning_cmd) Run() error{
	if len(cmd.Args) > 0 {

		resp, resErr := http.Get(cmd.Args[0])
		if resErr != nil {
			// handle error
			fmt.Println(resErr)
		}
	
	    if resp.Status == "200 OK" {
	    	fmt.Println("yeah")
	    }
	    /*
	    body, err := ioutil.ReadAll(resp.Body)
	    if err != nil {
			fmt.Println(err)
	    }
		//Convert the body to type string
	
	    sb := string(body)
	    log.Printf(sb)
	    */
	} else {
		fmt.Println("No Instructions as to what site should do, add arguments, or use help site")
	}
	return nil
}

func (cmd Isrunning_cmd) Usage() {
	usagestr := "\n\tsite - Currently not configure for use, more for an example\n\n\tUsage:\n\t\tsite <url>\n"
	colorfied := usagestr
	fmt.Println(colorfied)
}
