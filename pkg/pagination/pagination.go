package pagination

import (
	"net/http"
	"strconv"
)

// Paginator 分页器
type Paginator struct {
	Items any   `json:"items"`             // 数据
	Total int64 `json:"total"`             // 总页数
	Page  int   `json:"page" query:"page"` // 当前页
	Size  int   `json:"size" query:"size"` // 数据条数
}

// New 创建一个分页对象
func New(r *http.Request, max ...int) *Paginator {
	var p = Paginator{Page: 1, Size: 10}

	vs := r.URL.Query()
	if vs.Has("page") {
		page, _ := strconv.ParseInt(vs.Get("page"), 0, 0)
		if page > 0 {
			p.Page = int(page)
		}
	}

	if vs.Has("size") {
		size, _ := strconv.ParseInt(vs.Get("size"), 0, 0)
		if size > 0 {
			p.Size = int(size)
		}
	}

	if len(max) > 0 {
		if p.Size > max[0] {
			p.Size = max[0]
		}
	}

	return &p
}

// Offset 获取数据库查询的偏移量
func (p *Paginator) Offset() int {
	return (p.Page - 1) * p.Size
}

// Load 装载数据
func (p *Paginator) Load(data any) *Paginator {
	p.Items = data
	return p
}
