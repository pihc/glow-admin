// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// SysMenuDao is the manager for logic model data accessing and custom defined data operations functions management.
type SysMenuDao struct {
	gmvc.M                // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	C      sysMenuColumns // C is the short type for Columns, which contains all the column names of Table for convenient usage.
	DB     gdb.DB         // DB is the raw underlying database management object.
	Table  string         // Table is the underlying table name of the DAO.
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
	Type       string // 类型：0菜单/1按钮
	Status     string // 是否显示：0禁用/1正常
	Note       string // 备注
	Sort       string // 显示顺序
	CreatedBy  string // 添加人
	CreatedAt  string // 创建时间
	UpdatedBy  string // 更新人
	UpdatedAt  string // 更新时间
	Visible    string //
	CreateBy   string // 创建者
}

// NewSysMenuDao creates and returns a new DAO object for table data access.
func NewSysMenuDao() *SysMenuDao {
	columns := sysMenuColumns{
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
		UpdatedBy:  "updated_by",
		UpdatedAt:  "updated_at",
		Visible:    "visible",
		CreateBy:   "create_by",
	}
	return &SysMenuDao{
		C:     columns,
		M:     g.DB("default").Model("sys_menu").Safe(),
		DB:    g.DB("default"),
		Table: "sys_menu",
	}
}
