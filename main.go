package main

import (
	
	"fmt"

	"github.com/Syssos/Go_Shell/cmds"
)


func main() {
	err := cmds.Loop()
	if err != nil{
		fmt.Println(err)
	}
}