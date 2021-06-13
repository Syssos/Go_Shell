package main

import (
	"os"
	"fmt"
	"net/http"
)

func main(){
	if len(os.Args) > 1 {
		resp, resErr := http.Get(os.Args[1])
		check(resErr)
		fmt.Println(string([]rune(resp.Status)[0:3]))
	} else {
		fmt.Println("No arguments given")
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}