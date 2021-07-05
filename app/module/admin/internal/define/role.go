package define

import (
	"payget/app/model"
	"payget/library/query"
	"strings"

	"xorm.io/builder"
)

// ==========================================================================================
// API
// ==========================================================================================

type RoleApiCreateUpdateBase struct {
	Name string `v:"required#请输入角色名称"`
	Code string `v:"required#请输入角色标识"`
	Note string
}

type RoleApiDoCreateReq struct {
	RoleApiCreateUpdateBase
}

type RoleApiDoUpdateReq struct {
	RoleApiCreateUpdateBase
	Id uint `v:"min:1#请选择需要修改的角色"`
}

type RoleApiSetMenusReq struct {
	RoleId  uint   `v:"required#请选择角色"`
	MenuIds []uint `v:"required#请选择菜单和权限"`
}

// ==========================================================================================
// Service
// ==========================================================================================

type RoleServiceGetListReq struct {
	query.Params
	Name string `json:"name"`
}

func (q *RoleServiceGetListReq) Build() builder.Cond {
	cond := builder.NewCond()
	if q.Name != "" {
		cond = cond.And(builder.Like{"role.name", strings.TrimSpace(q.Name)})
	}
	return cond
}

type RoleServiceDoCreateReq struct {
	RoleServiceCreateUpdateBase
}

type RoleServiceDoUpdateReq struct {
	RoleServiceCreateUpdateBase
	Id uint
}

type RoleServiceCreateUpdateBase struct {
	Name string
	Code string
	Note string
}

type RoleServiceCreateRes struct {
	RoleId uint `json:"role_id"`
}

// 全部菜单（角色拥有的菜单会打上标记）
type RoleServiceMenuListRes struct {
	model.SysMenu
	Checked bool `json:"checked" `
	Open    bool `json:"open"`
}

type RoleServiceSetMenusReq struct {
	RoleId  uint
	MenuIds []uint
}
