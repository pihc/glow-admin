package query

import (
	"fmt"
	"math"

	"github.com/gogf/gf/database/gdb"
)

//Query 分页查询条件
type Query interface {
	GetPageIndex() int
	GetPageSize() int
	Build(*gdb.Model) *gdb.Model
	GetOrder() string
}

//Params 分页参数
type Params struct {
	Page  int    `json:"page" form:"page"`
	Limit int    `json:"limit" form:"limit"`
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}

//GetPageIndex 获取当前页码
func (p *Params) GetPageIndex() int {
	if p.Page == 0 {
		return 1
	}
	return p.Page
}

//GetPageSize 获取分页大小
func (p *Params) GetPageSize() int {
	if p.Limit == 0 {
		return 20
	}
	if p.Limit > 100 {
		return 100
	}
	return p.Limit
}

//GetOrder 获取排序
func (p *Params) GetOrder() string {
	if len(p.Sort) > 0 && len(p.Order) > 0 {
		return fmt.Sprintf("%s %s", p.Sort, p.Order)
	}
	return ""
}

//Data 分页结果
type Result struct {
	Records interface{} `json:"records"`
	Current int         `json:"current"`
	Pages   int         `json:"pages"`
	Size    int         `json:"size"`
	Total   int         `json:"total"`
}

func (r *Result) WithRecords(data interface{}) *Result {
	r.Records = data
	return r
}

func NewResult(page, limit, total int) *Result {
	pages := int(math.Ceil(float64(total) / float64(limit)))
	return &Result{
		Records: nil,
		Current: page,
		Pages:   pages,
		Size:    limit,
		Total:   total,
	}
}

//Page 分页查询
func Page(m *gdb.Model, query Query, bean interface{}) (*Result, error) {
	m = query.Build(m)
	pageIndex := query.GetPageIndex()
	pageSize := query.GetPageSize()
	listM := m.Page(pageIndex, pageSize)
	order := query.GetOrder()
	if len(order) > 0 {
		listM = listM.Order(order)
	}
	total, err := m.Fields("1").Count()
	if err != nil {
		return nil, err
	}
	err = listM.Scan(bean)
	if err != nil {
		return nil, err
	}

	return NewResult(pageIndex, pageSize, total).WithRecords(bean), nil
}

//All
func All(m *gdb.Model, query Query, bean interface{}) error {
	err := query.Build(m).Scan(bean)
	if err != nil {
		return err
	}
	return nil
}
