package api

import "github.com/labstack/echo/v4"

type home struct{}

var Home = new(home)

func (h *home) Version(c echo.Context) error {
	return c.JSON(200, 1)
}
