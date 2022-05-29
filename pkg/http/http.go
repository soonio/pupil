package http

import (
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	var e = echo.New()
	e.HideBanner = true
	e.JSONSerializer = &JsonSerializer{} // 使用自定义的JSON编码器
	return e
}
