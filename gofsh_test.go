package main

import (
	"os"
	"fmt"
	"strings"
    "os/exec"
	"testing"

	"github.com/Syssos/gofsh/src/color"
	"github.com/Syssos/gofsh/src/filelog"
)

var home string = filelog.GetHomeDir()

func TestNonInteractiveLS(t *testing.T) {
	cases := []struct{
		cmd []string
		want error
	}{
		{[]string{"ls"}, nil},
		{[]string{"ls", "../"}, nil},
		{[]string{"ls", "-la"}, nil},
	}
	for _, i := range cases {
		if len(i.cmd) > 1 {
			out, res := exec.Command(home+"/go/bin/gofsh", i.cmd[0], strings.Join(i.cmd[1:], " ")).Output()
			if res != i.want{
				t.Errorf("Error listing working directory")
			}
			fmt.Println(color.Green + string(out) + color.Reset)
		} else {
			out, res := exec.Command(home+"/go/bin/gofsh", i.cmd[0]).Output()
			if res != i.want{
				t.Errorf("Error listing working directory")
			}
			fmt.Println(color.Green + string(out) + color.Reset)
		}
	}
}

func TestNonInteractivePWD(t *testing.T) {
	cases := []struct{
		cmd string
		want error
	}{
		{"pwd", nil},
	}
	for _, i := range cases {
		out, res := exec.Command(home+"/go/bin/gofsh", i.cmd).Output()
		if res != i.want{
			t.Errorf("Error Printing working directory")
		}
		fmt.Println(color.Green + string(out) + color.Reset)
	}
}

func TestNonInteractiveSite(t *testing.T) {
	cases := []struct{
		cmd []string
		want error
	}{
		{[]string{"site", "https://syssos.app"}, nil},
		{[]string{"site", "https://google.com"}, nil},
	}
	for _, i := range cases {
		if len(i.cmd) > 1 {
			out, res := exec.Command(home+"/go/bin/gofsh", i.cmd[0], i.cmd[1]).Output()
			if res != i.want{
				t.Errorf("Error Printing working directory")
			}
			fmt.Println(color.Green + string(out) + color.Reset)
		} else {
			out, res := exec.Command(home+"/go/bin/gofsh", i.cmd[0]).Output()
			if res != i.want{
				t.Errorf("Error Printing working directory")
			}
			fmt.Println(color.Green + string(out) + color.Reset)
		}
	}
}

func TestNonInteractivePond(t *testing.T) {
	cases := []struct{
		cmd []string
		want error
	}{
		{[]string{"pond", "syssos_pond"}, nil},
	}
	for _, i := range cases {
		if len(i.cmd) > 1 {
			out, res := exec.Command(home+"/go/bin/gofsh", i.cmd[0], i.cmd[1]).Output()
			if _, err := os.Stat("./" + i.cmd[1] + "/server.go"); os.IsNotExist(err) {
				t.Errorf("Error creating pond files")
				
			} else {
				err := os.RemoveAll("./" + i.cmd[1])
				if err != nil {
					t.Errorf("Error removing pond files")
				}
			}
			if res != i.want{
				t.Errorf("Error Printing working directory")
			}
			fmt.Println(color.Green + string(out) + color.Reset)
		} else {
			
		}
	}
}