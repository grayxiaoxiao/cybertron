package test

import (
	"github.com/grayxiaoxiao/cybertron/cmd/cube/transform"
	"testing"
)

func Test_CubeNewWithoutArgs(t *testing.T) {
	newCmd := transform.CmdNew
	newCmd.Execute()
}

func Test_CubeNewWithPath(t *testing.T) {
	newCmd := transform.CmdNew
	newCmd.SetArgs([]string{"models/bumblebee/blank"})
	newCmd.Execute()
}

func Test_CubeNewWithFields(t *testing.T) {
	newCmd := transform.CmdNew
	newCmd.SetArgs([]string{"models/bumblebee/customer", "id:int64", "name:string", "address:string", "status:int", "description", "serial_number:string", "settlement_type:int"})
	newCmd.Execute()
}
