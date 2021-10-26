// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// SysRoleDao is the manager for logic model data accessing and custom defined data operations functions management.
type SysRoleDao struct {
	gmvc.M                // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	C      sysRoleColumns // C is the short type for Columns, which contains all the column names of Table for convenient usage.
	DB     gdb.DB         // DB is the raw underlying database management object.
	Table  string         // Table is the underlying table name of the DAO.
}

// SysRoleColumns defines and stores column names for table sys_role.
type sysRoleColumns struct {
	Id        string // 主键ID
	Name      string // 角色名称
	Code      string // 角色标签
	Note      string // 备注
	Sort      string // 排序
	CreatedBy string // 添加人
	CreatedAt string // 添加时间
	UpdatedBy string // 更新人
	UpdatedAt string // 更新时间
}

// NewSysRoleDao creates and returns a new DAO object for table data access.
func NewSysRoleDao() *SysRoleDao {
	columns := sysRoleColumns{
		Id:        "id",
		Name:      "name",
		Code:      "code",
		Note:      "note",
		Sort:      "sort",
		CreatedBy: "created_by",
		CreatedAt: "created_at",
		UpdatedBy: "updated_by",
		UpdatedAt: "updated_at",
	}
	return &SysRoleDao{
		C:     columns,
		M:     g.DB("default").Model("sys_role").Safe(),
		DB:    g.DB("default"),
		Table: "sys_role",
	}
}
