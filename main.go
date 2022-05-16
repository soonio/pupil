package main

import (
	"fmt"
	"github.com/soonio/pupil/bootstrap"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func main() {
	terminal := &cli.App{
		Name:   "pupil",
		Usage:  "小学生爱学习~",
		Flags:  bootstrap.Flags,
		Before: bootstrap.Bootstrap,
		Commands: []*cli.Command{
			{Name: "serve", Usage: "start http serve.", Action: bootstrap.Http},
		},
	}

	sort.Sort(cli.FlagsByName(terminal.Flags))
	sort.Sort(cli.CommandsByName(terminal.Commands))

	err := terminal.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
	}
}
