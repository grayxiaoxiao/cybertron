package main

import (
	"log"

	"github.com/grayxiaoxiao/cybertron/cmd/cube"

	"github.com/spf13/cobra"
)

var (
    version string = "v.1.1.1"
    rootCmd = &cobra.Command{
        Use:   "cybertron",
        Short: "cybertron: generate go struct from tmpl file/通过tmpl文件构建struct",
        Long:  "cybertron: generate go struct from tmpl file/通过tmpl文件构建struct",
        Version: version,
    }
)

func init() {
    rootCmd.AddCommand(cube.CmdCube)
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}
