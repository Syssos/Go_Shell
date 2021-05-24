package main

import (
	"testing"
	"fmt"
	"github.com/Syssos/Go_Shell/cmds"
)

// Setting up commands from cmds directory
var cd cmds.Cd_cmd = cmds.Cd_cmd{[]string{}}
var pwd cmds.Pwd_cmd = cmds.Pwd_cmd{[]string{}}
var ls cmds.Ls_cmd = cmds.Ls_cmd{[]string{}}

// Creating command struct to hold available commands
var command_struct commands = commands{ls, pwd, cd}

// Commands struct will be a "list" of commands shell can run
type commands struct {
	ls cmds.Ls_cmd
	pwd cmds.Pwd_cmd
	cd cmds.Cd_cmd
}

// used to execute every command in command struct
type exec interface {
	Run() error
}
// responsible for running and logging errors with command
func execute(e exec) error {
	err := e.Run()
	if err != nil {
		return err
	}
	return nil
}

func TestPwd(t *testing.T) {

	cases := []struct{
		args []string
		want error
	}{
		{[]string{}, nil},
		{[]string{"./"}, nil},
	}
	for _, i := range cases {
		command_struct.pwd.Args = i.args
		res := execute(command_struct.pwd)
		if res != i.want{
			t.Errorf("Error Printing working directory")
		}
	}
}

func TestLs(t *testing.T) {
	cases := []struct{
		args []string
		want error
	}{
		{[]string{}, nil},
		{[]string{"./"}, nil},
	}
	for _, i := range cases {
		command_struct.ls.Args = i.args
		res := execute(command_struct.ls)
		if res != i.want{
			t.Errorf("Error Printing working directory")
		}
	}
}

func TestCd(t *testing.T) {
	cases := []struct{
		args []string
		want error
	}{
		{[]string{"./"}, nil},
		{[]string{"../"}, nil},
	}

	for _, i := range cases {
		command_struct.cd.Args = i.args
		res := execute(command_struct.cd)
		fmt.Println(res)
		if res != i.want{
			t.Errorf("Error Printing working directory")
		}
	}
}