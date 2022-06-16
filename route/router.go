package route

import (
	"github.com/labstack/echo/v4"
	"github.com/soonio/pupil/app/api"
)

func Register(e *echo.Echo) {
	e.GET("/version", api.Home.Version)

	e.GET("/dict/:key", api.Dict.Get)
	e.GET("/dict", api.Dict.List)
	e.POST("/dict", api.Dict.Save)
}
