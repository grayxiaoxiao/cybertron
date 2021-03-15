package test

import (
	"cybertron/cmd/cube/transform"
	"testing"
)

func Test_CubeNewCommand(t *testing.T) {
	newCmd := transform.CmdNew
	newCmd.SetArgs([]string{"models/bumblebee/customer", "id:int64", "name:string", "address:string", "status:int", "description", "serial_number:string", "settlement_type:int"})
	newCmd.Execute()
}
