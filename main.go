package main

import (
	
	"fmt"
	"time"
	"bytes"
	"io/ioutil"
	"encoding/json"

	"github.com/komkom/toml"
	"github.com/Syssos/Go_Shell/cmds"
	"github.com/Syssos/Go_Shell/filelog"
)
// Creates filelog.Flog and cmds.Loop instance, then runs Loop for command line interpretor
func main() {
	
	cd     := cmds.Cd_cmd{[]string{}}
	pwd    := cmds.Pwd_cmd{[]string{}}
	ls     := cmds.Ls_cmd{[]string{}}
	site   := cmds.Site_cmd{[]string{}, "", false}
	hostme := cmds.HostMe_cmd{[]string{}}
	pond   := cmds.Pond_cmd{[]string{}}

	command_struct := cmds.Commands{ls, pwd, cd, site, hostme, pond}
	
	// Creating Flog instance from settings/cmds.toml
	flog := LoggerFromFile()

	loop := cmds.Loop{flog, command_struct}

	loopErr := loop.Run()
	if loopErr != nil{
		flog.Errormsg = loopErr
		flog.Err()
	}
}

// Generating Flog instance from settings retrieved from file
func LoggerFromFile() filelog.Flog {

	file, openErr := ioutil.ReadFile("settings/cmds.toml")
	if openErr != nil {
		fmt.Println(openErr)
	}

	doc := string(file)
	// Decodes toml to *json.Decoder
	dec := json.NewDecoder(toml.New(bytes.NewBufferString(doc)))
	
	st  := struct {
	  Logger struct {
	    Greeting string `json: "Greeting"`
	    Salute string `json: "Salute"`
	    LogFile string `json: "LogFile"`
	    DtFormat string `json: "DtFormat"`
	    DtTimeZone string `json: "DtTimeZone"`
	    DtOffset int `json: "DtOffset"`
	    Errormsg bool `json: "Errormsg"`
	  } `json: "Logger"`
	}{}

	err := dec.Decode(&st)
	if err != nil {
	  panic(err)
	}

	// Setting the error logs timestamp, timezone
	location := time.FixedZone(st.Logger.DtTimeZone, st.Logger.DtOffset*60*60)
	
	flog := filelog.Flog{ st.Logger.Greeting, st.Logger.Salute, st.Logger.LogFile, st.Logger.DtFormat, location, nil}
	return flog
}