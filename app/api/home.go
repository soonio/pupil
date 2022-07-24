package api

import (
	"github.com/labstack/echo/v4"
	"github.com/soonio/pupil/app/types"
	"github.com/soonio/pupil/pkg/validator"
	"net/http"
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
		return Failure(c, http.StatusUnprocessableEntity, err.(*validator.Error).Lang(validator.Language(c)).Error())
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
