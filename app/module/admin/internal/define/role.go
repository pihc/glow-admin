package define

import (
	"glow-admin/app/model"
	"glow-admin/library/query"
	"strings"

	"xorm.io/builder"
)

// ==========================================================================================
// API
// ==========================================================================================

// API删除
type RoleApiDeleteReq struct {
	Id uint `v:"min:1#请选择需要删除的角色"`
}

type RoleApiGetMenusReq struct {
	Id uint `v:"min:1#请选择角色"`
}

// API创建/修改基类
type RoleApiCreateUpdateBase struct {
	Name string `v:"required#请输入角色名称"`
	Code string `v:"required#请输入角色标识"`
	Note string
}

// API创建
type RoleApiCreateReq struct {
	RoleApiCreateUpdateBase
}

// API修改
type RoleApiUpdateReq struct {
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

// Service创建
type RoleServiceCreateReq struct {
	RoleServiceCreateUpdateBase
	CreatedBy uint `json:"created_by"`
}

// Service修改
type RoleServiceUpdateReq struct {
	RoleServiceCreateUpdateBase
	Id        uint `json:"id"`
	UpdatedBy uint `json:"updated_by"`
}

// Service创建/修改基类
type RoleServiceCreateUpdateBase struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Note string `json:"note"`
}

// Service创建返回结果
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
	RoleId  uint   `json:"role_id"`
	MenuIds []uint `json:"menu_ids"`
}
