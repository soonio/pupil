package console

import (
	"github.com/urfave/cli/v2"
)

var commands []*cli.Command

// Register 把命令注册到命令管理器中
func Register(terminal *cli.App) {
	terminal.Commands = append(terminal.Commands, commands...)
}
