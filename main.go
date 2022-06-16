package main

import (
	"flag"

	"github.com/soonio/pupil/app"
	"github.com/soonio/pupil/bootstrap"
	"github.com/soonio/pupil/pkg/http"
	"github.com/soonio/pupil/pkg/validator"
	"github.com/soonio/pupil/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title                 pupil
// @version               1.0
// @schemes               https
// @host                  https://pupil.localhost.com
func main() {

	var c = flag.String("c", "config.yaml", "the config file")
	flag.Parse()

	bootstrap.Bootstrap(*c)

	var e = echo.New()
	e.Validator = validator.New()
	e.JSONSerializer = &http.JsonSerializer{}
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Recover())

	route.Register(e)

	err := e.Start(app.Config.Http.Addr)
	if err != nil {
		panic(err)
	}
}
