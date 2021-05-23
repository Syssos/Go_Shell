package main

import (
	"testing"
	"github.com/Syssos/Go_Shell/cmds"
)

func TestPwd(t *testing.T) {
	cases := []struct{
		args []string
		want error
	}{
		{[]string{}, nil},
		{[]string{"./"}, nil},
	}
	for _, i := range cases {
		res := cmds.Pwd(i.args)
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
		res := cmds.Ls(i.args)
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
		res := cmds.Cd(i.args)
		if res != i.want{
			t.Errorf("Error Printing working directory")
		}
	}
}