package validator

import "github.com/labstack/echo/v4"

// 主要针对echo框架做的适配多语言

type language struct {
	AcceptLanguage string `header:"Accept-Language"` // 语言
}

// Language 获取请求头中的默认语言
func Language(c echo.Context) string {
	var l = &language{}
	_ = c.Echo().Binder.(*echo.DefaultBinder).BindHeaders(c, l)
	return l.AcceptLanguage
}
