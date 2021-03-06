package api

import (
	"net/http"

	"github.com/soonio/pupil/pkg/validator"

	"github.com/labstack/echo/v4"
)

// 业务状态码定义需要从600之后开始
const (
	StatusOK             = 0
	StatusOperateFailure = 655
)

var TextCode = map[int]string{
	StatusOK:             "success",
	StatusOperateFailure: "操作失败",
}

func StatusText(code int) string {
	return TextCode[code]
}

// Response 接口响应消息的格式
type Response struct {
	Code int    `json:"code"`           // 业务状态码
	Msg  string `json:"msg"`            // 业务消息
	Data any    `json:"data,omitempty"` // 数据
}

// Success 成功时响应消息
func Success(c echo.Context, data ...any) error {
	var d any
	if len(data) > 0 {
		d = data[0]
	}
	return c.JSON(http.StatusOK, &Response{0, TextCode[0], d})
}

// Failure 失败时响应消息
func Failure(c echo.Context, httpCode int, msg ...string) error {
	var m string
	if len(msg) > 0 {
		m = msg[0]
	} else {
		if httpCode >= 600 {
			m = StatusText(httpCode)
		} else {
			m = http.StatusText(httpCode)
		}
	}
	if httpCode < 600 {
		return c.JSON(httpCode, map[string]string{"message": m})
	} else {
		return c.JSON(http.StatusOK, &Response{httpCode, m, nil})
	}
}

// ParameterError 参数错误时的便捷响应方法
func ParameterError(c echo.Context, err error) error {
	if e, ok := err.(*validator.Error); ok {
		return Failure(c, http.StatusUnprocessableEntity, e.LangInHeader(c.Request()).Error())
	}
	return Failure(c, http.StatusUnprocessableEntity, err.Error())
}
