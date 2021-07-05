// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// SysMenuDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type SysMenuDao struct {
	gmvc.M                 // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB         // DB is the raw underlying database management object.
	Table   string         // Table is the table name of the DAO.
	Columns sysMenuColumns // Columns contains all the columns of Table that for convenient usage.
}

// SysMenuColumns defines and stores column names for table sys_menu.
type sysMenuColumns struct {
	Id         string // 主键ID
	Pid        string // 父级ID
	Title      string // 菜单标题
	Icon       string // 图标
	Path       string // 菜单路径
	Component  string // 菜单组件
	Target     string // 目标
	Permission string // 权限标识
	Type       string // 类型：1目录 2菜单 3节点
	Status     string // 是否显示：1显示 2不显示
	Note       string // 备注
	Sort       string // 显示顺序
	CreatedBy  string // 添加人
	CreatedAt  string // 添加时间
	UpdateBy   string // 更新人
	UpdatedAt  string // 更新时间
}

func NewSysMenuDao() *SysMenuDao {
	return &SysMenuDao{
		M:     g.DB("default").Model("sys_menu").Safe(),
		DB:    g.DB("default"),
		Table: "sys_menu",
		Columns: sysMenuColumns{
			Id:         "id",
			Pid:        "pid",
			Title:      "title",
			Icon:       "icon",
			Path:       "path",
			Component:  "component",
			Target:     "target",
			Permission: "permission",
			Type:       "type",
			Status:     "status",
			Note:       "note",
			Sort:       "sort",
			CreatedBy:  "created_by",
			CreatedAt:  "created_at",
			UpdateBy:   "update_by",
			UpdatedAt:  "updated_at",
		},
	}
}
