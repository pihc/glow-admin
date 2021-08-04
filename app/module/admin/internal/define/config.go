package define

import (
	"glow-admin/library/query"
	"strings"

	"xorm.io/builder"
)

// ==========================================================================================
// API
// ==========================================================================================

// API配置分组明细
type ConfigApiDetailReq struct {
	Id uint `v:"min:1#请选择查看的配置分组"`
}

// API配置分组删除
type ConfigApiDeleteReq struct {
	Id uint `v:"min:1#请选择需要删除的配置分组"`
}

// API配置分组创建
type ConfigApiCreateReq struct {
	ConfigApiCreateUpdateBase
}

// API配置分组修改
type ConfigApiUpdateReq struct {
	ConfigApiCreateUpdateBase
	Id uint `v:"min:1#请选择需要修改的配置分组"`
}

// API配置分组创建/修改基类
type ConfigApiCreateUpdateBase struct {
	Name string `v:"required#请输入分组名称"` //分组名称
	Sort int    `v:"required#请输入排序"`   //排序
}

// ==========================================================================================
// Service
// ==========================================================================================

// Service配置分组查询
type ConfigServiceGetListReq struct {
	query.Params
	Name string `json:"short_name"`
}

func (q *ConfigServiceGetListReq) Build() builder.Cond {
	cond := builder.NewCond()
	if q.Name != "" {
		cond = cond.And(builder.Like{"role.short_name", strings.TrimSpace(q.Name)})
	}
	return cond
}

// Service配置分组创建
type ConfigServiceCreateReq struct {
	ConfigServiceCreateUpdateBase
	CreatedBy uint `json:"created_by"`
}

// Service配置分组修改
type ConfigServiceUpdateReq struct {
	ConfigServiceCreateUpdateBase
	Id        uint `json:"id"`
	UpdatedBy uint `json:"updated_by"`
}

// Service配置分组创建/修改基类
type ConfigServiceCreateUpdateBase struct {
	Name string `json:""` //分组名称
	Sort int    `json:""` //排序
}

// Service配置分组创建返回结果
type ConfigServiceCreateRes struct {
	ConfigId uint `json:"config_id"`
}
