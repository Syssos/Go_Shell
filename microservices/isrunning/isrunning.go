package isrunning

import (
	"fmt"
)

// This package will run an http request to a specific website to ensure it is working how it should
// This will be turned into a way to test websites I create to make sure there is no errors when accessing the webserver.

type Isrunning_cmd struct {
	Args []string
	Link string
	Active bool
}

func (cmd Isrunning_cmd) Run() error{
	fmt.Println(cmd.Link)
	return nil
}

func (cmd Isrunning_cmd) Usage() {
	fmt.Println("Made it to usage")
}