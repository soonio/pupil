package http

import (
	"github.com/labstack/echo/v4"
	"sync"
)

var (
	once sync.Once
	e    *echo.Echo
)

func Server() *echo.Echo {
	once.Do(func() {
		e = echo.New()
		e.HideBanner = true
		e.JSONSerializer = &JsonSerializer{} // 使用自定义的JSON编码器
	})
	return e
}
