package api

import (
	"net/http"

	"github.com/soonio/pupil/app/logic"
	"github.com/soonio/pupil/app/types"
	"github.com/soonio/pupil/pkg/pagination"

	"github.com/labstack/echo/v4"
)

var Dict = new(dictApi)

type dictApi struct{}

// Get
// @tags     字典管理
// @Summary  获取字典配置
// @Param    token  header    string                     true  "认证Token"
// @Param    key    path      int                        true  "键"
// @Success  200    {object}  api.Response{data=string}  "值"
// @Router   /dict/{key} [GET]
func (i *dictApi) Get(c echo.Context) error {
	v, err := logic.Dict.Get(c.Param("key"))
	if err != nil {
		return ding(err)
	}
	return Success(c, v)
}

// List
// @tags     字典管理
// @Summary  获取列表
// @Param    token  header    string                                     true   "认证Token"
// @Param    page   query     int                                        false  "页码，默认1"
// @Param    size   query     int                                        false  "分页大小，参考返回的分页大小(默认10)"
// @Success  200    {object}  api.Response{data=types.FeedbackResponse}  ""
// @Router   /dict [GET]
func (i *dictApi) List(c echo.Context) error {
	var paginator = pagination.New(c.Request())
	err := logic.Dict.List(paginator)
	if err != nil {
		return ding(err)
	}
	return Success(c, paginator)
}

// Save
// @tags     字典管理
// @Summary  创建字典
// @Param    token  header    string          true  "认证Token"
// @Param    k      formData  string          true  "键"
// @Param    v      formData  string          true  "值"
// @Success  200    {object}  api.Response{}  ""
// @Router   /dict [POST]
func (i *dictApi) Save(c echo.Context) error {
	p := &types.AdminDictSaveRequest{}
	_ = c.Bind(p)
	if err := c.Validate(p); err != nil {
		return Failure(c, http.StatusUnprocessableEntity, err.Error())
	}

	if err := logic.Dict.Save(p.K, p.V); err != nil {
		return ding(err)
	}
	return Success(c)
}
