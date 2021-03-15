package cube

import (
	"github.com/grayxiaoxiao/cybertron/cmd/cube/transform"

	"github.com/spf13/cobra"
)

var CmdCube = &cobra.Command{
    Use:   "cube",
    Short: "generate the struct files/构建struct",
    Long:  "generate the struct files/构建struct",
    Run:   run,
}

func init() {
    CmdCube.AddCommand(transform.CmdNew)
}

func run(cmd *cobra.Command, args []string) {
}
