package bootstrap

import (
	"github.com/soonio/pupil/app"
	"github.com/soonio/pupil/bootstrap/internal"
	"github.com/soonio/pupil/pkg/http"
	"github.com/soonio/pupil/route"
	"github.com/urfave/cli/v2"
	"os"
)

var Flags = []cli.Flag{
	&cli.StringFlag{Name: "config", Aliases: []string{"c"}, Value: "config.yaml", Usage: "config file"},
}

// Bootstrap 项目初始化
func Bootstrap(context *cli.Context) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	app.Home = dir

	return internal.Viper(context.String("config"))
}

// Http 服务初始化
func Http(c *cli.Context) error {
	serve := http.Server()

	route.Initialize()

	return serve.Start(app.Config.Http.Addr)
}
