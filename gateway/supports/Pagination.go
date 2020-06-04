package supports

import (
	"errors"

	"github.com/kataras/iris/v12"
)
type TableVO struct {
	Code int64 `json:"code"`
	Count int64 `json:"count"`
	Data interface{} `json:"data"`
}

func NewTableVO(code int64,count int64, data interface{}) *TableVO {
	return &TableVO{
		Code:code,
		Count:count,
		Data:data,
	}
}

// bootstraptable 分页参数
type Pagination struct {
	PageNumber int //当前看的是第几页
	PageSize   int //每页显示多少条数据

	// 用于分页设置的参数
	Start int
	Limit int

	// 时间范围
	StartDate string
	EndDate   string

	Uid int64 // 公用的特殊参数
}

func NewPagination(ctx iris.Context) (*Pagination, error) {
	start, err1 := ctx.URLParamInt("page")
	size, err2 := ctx.URLParamInt("limit")
	if err1 != nil || err2 != nil {
		return nil, errors.New("请求的分页参数解析错误")
	}

	page := Pagination{
		PageNumber: start,
		PageSize:   size,
	}
	page.pageSetting()
	return &page, nil
}

// 设置分页参数
func (p *Pagination) pageSetting() {
	if p.PageNumber < 1 {
		p.PageNumber = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 1
	}

	p.Start = (p.PageNumber - 1) * p.PageSize
	p.Limit = p.PageSize
}
