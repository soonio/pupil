package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/soonio/pupil/bootstrap"
	"github.com/soonio/pupil/cli/migrate"
	"github.com/soonio/pupil/cli/tools"
	"github.com/soonio/pupil/pkg/utils"

	"github.com/urfave/cli/v2"
)

var commands []*cli.Command

func main() {
	terminal := &cli.App{
		Name:  "pupil",
		Usage: "pupil-cli",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "config", Aliases: []string{"c"}, Value: "config.yaml", Usage: "设置配置文件"},
		},
		Before: func(context *cli.Context) error {
			bootstrap.Bootstrap(context.String("config"))
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "version",
				Usage: "show version",
				Action: func(context *cli.Context) error {
					v := utils.Version()
					fmt.Printf(`Version:    %s
Build Date: %s
Go Version: %s
OS/Arch:    %s
Git:
  Branch:   %s
  CommitId: %s
`, v.Version, v.BuildDate, v.GoVersion, v.Platform, v.Git.Branch, v.Git.Commit)
					return nil
				},
			},
			{
				Name:   "migrate:version",
				Usage:  "版本信息",
				Action: migrate.Version,
			},
			{
				Name:   "migrate:create",
				Usage:  "创建迁移文件",
				Action: migrate.Create,
			},
			{
				Name:   "migrate:up",
				Usage:  "数据库迁移升级",
				Action: migrate.Up,
			},
			{
				Name:   "migrate:down",
				Usage:  "数据库迁移降级",
				Action: migrate.Down,
			},
			{
				Name:   "migrate:step",
				Usage:  "数据库迁移到指定步长",
				Action: migrate.Steps,
				Flags:  []cli.Flag{&cli.IntFlag{Name: "number", Aliases: []string{"n"}, Usage: "步长，正数向前迁移指定步长，负数向后退回指定步长", Required: true}},
			},
			{
				Name:   "jwt",
				Usage:  "生成接口可用token",
				Action: tools.GenJwt,
				Flags:  []cli.Flag{&cli.IntFlag{Name: "id", Aliases: []string{"i"}, Usage: "用户ID", Value: 1}},
			},
		},
	}

	terminal.Commands = append(terminal.Commands, commands...)

	sort.Sort(cli.FlagsByName(terminal.Flags))
	sort.Sort(cli.CommandsByName(terminal.Commands))

	err := terminal.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
