package api

import (
	"fmt"
	"github.com/soonio/pupil/pkg/erro"
	"net/http"

	"github.com/soonio/pupil/app/types"

	"github.com/labstack/echo/v4"
)

type home struct{}

var Home = new(home)

func (h *home) Version(c echo.Context) error {
	return c.JSON(200, 1)
}

// ValidateMultiLangError
// @tags    验证器错误消息
// @Summary 获取多语言错误(当前仅支持英文|中文)
// @Param   token           header   string         true "认证Token"
// @Param   Accept-Language header   string         true "示例值: en | zh"
// @Param   k               formData string         true "键"
// @Param   v               formData string         true "值"
// @Success 200             {object} api.Response{} ""
// @Router  /validate/multi-lang [POST]
func (h *home) ValidateMultiLangError(c echo.Context) error {
	p := &types.AdminDictSaveRequest{}
	_ = c.Bind(p)

	if err := c.Validate(p); err != nil {
		// 对于错误复杂的处理方式，可以封装在Failure方法中
		//return Failure(c, http.StatusUnprocessableEntity, err.(*validator.Error).Lang(validator.Language(c)).Error())
		return ParameterError(c, err)
	}

	return Success(c, p)
}

// ValidateDefaultError
// @tags    验证器错误消息
// @Summary 获取默认错误
// @Param   token header   string         true "认证Token"
// @Param   k     formData string         true "键"
// @Param   v     formData string         true "值"
// @Success 200   {object} api.Response{} ""
// @Router  /validate/default [POST]
func (h *home) ValidateDefaultError(c echo.Context) error {
	p := &types.AdminDictSaveRequest{}
	_ = c.Bind(p)

	if err := c.Validate(p); err != nil {
		return Failure(c, http.StatusUnprocessableEntity, err.Error())
	}

	return Success(c, p)

}

// Error
// @tags    错误
// @Summary 使用自定义错误
// @Success 200   {object} api.Response{} ""
// @Router  /erro [GET]
func (h *home) Error(c echo.Context) error {
	err := erro.New("some error.")

	if err != nil {
		if err, ok := err.(*erro.Erro); ok {
			fmt.Println("error:", err.Error())
			fmt.Println("path:", err.On())
			return c.JSON(http.StatusOK, &Response{
				StatusOperateFailure,
				"错误消息",
				map[string]string{
					"msg":    err.Error(),
					"caller": err.On(),
				}})
		}
		return Failure(c, http.StatusInternalServerError, err.Error())
	}

	return Success(c)
}
