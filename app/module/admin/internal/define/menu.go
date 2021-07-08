package define

import (
	"payget/library/query"
	"strings"

	"xorm.io/builder"
)

// ==========================================================================================
// API
// ==========================================================================================
type MenuApiDetailReq struct {
	Id uint `v:"min:1#请选择查看的菜单"`
}

type MenuApiDeleteReq struct {
	Id uint `v:"min:1#请选择删除的菜单"`
}

type MenuApiCreateUpdateBase struct {
	Pid        uint   // 父级ID
	Title      string `v:"required#请输入菜单名称"` // 菜单标题
	Icon       string // 图标
	Path       string // 菜单路径
	Component  string // 菜单组件
	Target     string // 目标
	Permission string // 权限标识
	Type       uint   // 类型：0 菜单 1节点
	Status     uint   // 是否显示：1显示 2不显示
	Note       string // 备注
	Sort       uint   `v:"required#请输入排序"` // 显示顺序
}

type MenuApiCreateReq struct {
	Nodes []string // 权限节点，用来快速生成页面按钮权限
	MenuApiCreateUpdateBase
}

type MenuApiUpdateReq struct {
	MenuApiCreateUpdateBase
	Id uint `v:"min:1#请选择需要修改的菜单"`
}

// ==========================================================================================
// Service
// ==========================================================================================
type MenuServiceCreateUpdateBase struct {
	Pid        uint   `json:"pid"`        // 父级ID
	Title      string `json:"title"`      // 菜单标题
	Icon       string `json:"icon"`       // 图标
	Path       string `json:"path"`       // 菜单路径
	Component  string `json:"component"`  // 菜单组件
	Target     string `json:"target"`     // 目标
	Permission string `json:"permission"` // 权限标识
	Type       uint   `json:"type"`       // 类型：0 菜单 1节点
	Status     uint   `json:"status"`     // 是否显示：1显示 2不显示
	Note       string `json:"note"`       // 备注
	Sort       uint   `json:"sort"`       // 显示顺序
}

type MenuServiceCreateReq struct {
	MenuServiceCreateUpdateBase
	Nodes     []string `json:"nodes"` // 权限节点，用来快速生成页面按钮权限
	CreatedBy uint     `json:"created_by"`
}

type MenuServiceUpdateReq struct {
	MenuServiceCreateUpdateBase
	Id        uint `json:"id"`
	UpdatedBy uint `json:"updated_by"`
}

type MenuServiceCreateRes struct {
	MenuId uint `json:"menu_id"`
}

// 查询
type MenuServiceGetListReq struct {
	query.Params
	Title string `json:"title"`
}

func (q *MenuServiceGetListReq) Build() builder.Cond {
	cond := builder.NewCond()
	if q.Title != "" {
		cond = cond.And(builder.Like{"sys_menu.title", strings.TrimSpace(q.Title)})
	}
	return cond
}
