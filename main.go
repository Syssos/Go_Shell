package main

import (
	
	"fmt"
	"encoding/json"
	"bytes"
	"io/ioutil"

	"github.com/Syssos/Go_Shell/cmds"
	"github.com/Syssos/Go_Shell/filelog"
	"github.com/komkom/toml"
)

func main() {
	
	// Setting up commands from cmds directory
	cd := cmds.Cd_cmd{[]string{}}
	pwd := cmds.Pwd_cmd{[]string{}}
	ls := cmds.Ls_cmd{[]string{}}
	site := cmds.Site_cmd{[]string{}, "", false}

	flog := LoggerFromFile()
	command_struct := cmds.Commands{ls, pwd, cd, site}
	loop := cmds.Loop{command_struct, flog}

	loopErr := loop.Run()
	if loopErr != nil{
		flog.Errormsg = loopErr
		flog.Err()
	}
}

func LoggerFromFile() filelog.Flog {

	file, openErr := ioutil.ReadFile("settings/cmds.toml")
	if openErr != nil {
		fmt.Println(openErr)
	}

	doc := string(file)
	dec := json.NewDecoder(toml.New(bytes.NewBufferString(doc)))
	st := struct {
	  Logger struct {
	    Greeting string `json: "Greeting"`
	    Salute string `json: "Salute"`
	    LogFile string `json: "LogFile"`
	    Errormsg bool `json: "Errormsg"`
	  } `json: "Logger"`
	}{}

	err := dec.Decode(&st)
	if err != nil {
	  panic(err)
	}

	flog := filelog.Flog{ st.Logger.Greeting, st.Logger.Salute, st.Logger.LogFile, nil}
	return flog
}