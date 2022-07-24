package route

import (
	"github.com/labstack/echo/v4"
	"github.com/soonio/pupil/app/api"
)

func Register(e *echo.Echo) {
	// 路由采用显示路由，所有的路由都标识清楚，便于查找
	e.GET("/version", api.Home.Version)

	e.POST("validate/default", api.Home.ValidateDefaultError)
	e.POST("/validate/multi-lang", api.Home.ValidateMultiLangError)

	e.GET("/dict/:key", api.Dict.Get)
	e.GET("/dict", api.Dict.List)
	e.POST("/dict", api.Dict.Save)

}
