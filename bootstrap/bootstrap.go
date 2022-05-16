package bootstrap

import (
	"github.com/soonio/pupil/app"
	"github.com/soonio/pupil/bootstrap/internal"
	"github.com/soonio/pupil/pkg/http"
	"github.com/urfave/cli/v2"
	"os"
)

var Flags = []cli.Flag{
	&cli.StringFlag{Name: "config", Aliases: []string{"c"}, Value: "config.yaml", Usage: "config file"},
}

func Bootstrap(context *cli.Context) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	app.Home = dir

	return internal.Viper(context.String("config"))
}

func Http(c *cli.Context) error {
	serve := http.Server()
	return serve.Start(app.Config.Http.Addr)
}
