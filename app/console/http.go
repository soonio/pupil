package console

import (
	"github.com/soonio/pupil/app"
	"github.com/soonio/pupil/pkg/http"
	"github.com/soonio/pupil/route"
	"github.com/urfave/cli/v2"
)

func init() {
	commands = append(commands,
		&cli.Command{Name: "serve", Usage: "start http serve.", Action: func(context *cli.Context) error {
			serve := http.New()
			route.Register(serve)
			return serve.Start(app.Config.Http.Addr)
		}},
	)
}
