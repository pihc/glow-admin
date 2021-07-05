// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// SysRole is the golang structure for table sys_role.
type SysRole struct {
	Id        uint        `orm:"id,primary" json:"id"`         // 主键ID
	Name      string      `orm:"name"       json:"name"`       // 角色名称
	Code      string      `orm:"code"       json:"code"`       // 角色标签
	Status    uint        `orm:"status"     json:"status"`     // 状态：1正常 2禁用
	Note      string      `orm:"note"       json:"note"`       // 备注
	Sort      uint        `orm:"sort"       json:"sort"`       // 排序
	CreatedBy uint        `orm:"created_by" json:"created_by"` // 添加人
	CreatedAt *gtime.Time `orm:"created_at" json:"created_at"` // 添加时间
	UpdateBy  uint        `orm:"update_by"  json:"update_by"`  // 更新人
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updated_at"` // 更新时间
}
