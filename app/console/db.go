package console

import (
	"errors"
	"fmt"
	"github.com/soonio/pupil/app"
	"github.com/urfave/cli/v2"
)

// 生成db相关的命令参数，使用migrate命令执行
// cmd | xargs migrate

func init() {
	commands = append(commands,
		&cli.Command{
			Name:  "db:create",
			Usage: "create table sql",
			Action: func(c *cli.Context) error {
				if c.Args().Len() == 0 {
					return errors.New("缺少数据表迁移名称")
				}
				_, err := fmt.Printf(
					"create -ext sql -dir database/migrations -seq %s_table\n",
					c.Args().First())
				return err
			}},
		&cli.Command{
			Name:  "db:up",
			Usage: "up",
			Action: func(c *cli.Context) error {
				if c.Args().Len() == 0 {
					return errors.New("缺少迁移的步长")
				}
				_, err := cmd("up", c.Args().First())
				return err
			}},
		&cli.Command{
			Name:  "db:down",
			Usage: "down",
			Action: func(c *cli.Context) error {
				if c.Args().Len() == 0 {
					return errors.New("缺少迁移的步长")
				}
				_, err := cmd("down", c.Args().First())
				return err
			}},
	)
}

func cmd(method string, step string) (n int, err error) {
	return fmt.Printf(
		"-path database/migrations -database \"mysql://%s:%s@tcp(%s)/%s\" %s %s\n",
		app.Config.DB.Username,
		app.Config.DB.Password,
		app.Config.DB.Path,
		app.Config.DB.Dbname,
		method,
		step)
}
