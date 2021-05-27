package main

import (
	
	"fmt"

	"github.com/Syssos/Go_Shell/cmds"
)

func main() {


	// Setting up commands from cmds directory
	cd := cmds.Cd_cmd{[]string{}}
	pwd := cmds.Pwd_cmd{[]string{}}
	ls := cmds.Ls_cmd{[]string{}}
	site := cmds.Site_cmd{[]string{}, "", false}

	// Creating command struct to hold available commands
	command_struct := cmds.Commands{ls, pwd, cd, site}

	loop := cmds.Loop{command_struct}

	err := loop.Run()
	if err != nil{
		fmt.Println(err)
	}
}