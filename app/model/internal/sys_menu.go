// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// SysMenu is the golang structure for table sys_menu.
type SysMenu struct {
	Id         uint        `orm:"id,primary" json:"id"`         // 主键ID
	Pid        uint        `orm:"pid"        json:"pid"`        // 父级ID
	Title      string      `orm:"title"      json:"title"`      // 菜单标题
	Icon       string      `orm:"icon"       json:"icon"`       // 图标
	Path       string      `orm:"path"       json:"path"`       // 菜单路径
	Component  string      `orm:"component"  json:"component"`  // 菜单组件
	Target     string      `orm:"target"     json:"target"`     // 目标
	Permission string      `orm:"permission" json:"permission"` // 权限标识
	Type       uint        `orm:"type"       json:"type"`       // 类型：1目录 2菜单 3节点
	Status     uint        `orm:"status"     json:"status"`     // 是否显示：1显示 2不显示
	Note       string      `orm:"note"       json:"note"`       // 备注
	Sort       uint        `orm:"sort"       json:"sort"`       // 显示顺序
	CreatedBy  uint        `orm:"created_by" json:"created_by"` // 添加人
	CreatedAt  *gtime.Time `orm:"created_at" json:"created_at"` // 添加时间
	UpdateBy   uint        `orm:"update_by"  json:"update_by"`  // 更新人
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updated_at"` // 更新时间
}
