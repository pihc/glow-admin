// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// SysConfig is the golang structure for table sys_config.
type SysConfig struct {
	Id        uint        `orm:"id,primary" json:"id"`         // 主键ID
	Name      string      `orm:"name"       json:"name"`       // 分组名称
	Sort      int         `orm:"sort"       json:"sort"`       // 排序
	CreatedBy uint        `orm:"created_by" json:"created_by"` // 添加人
	CreatedAt *gtime.Time `orm:"created_at" json:"created_at"` // 添加时间
	UpdatedBy uint        `orm:"updated_by" json:"updated_by"` // 更新人
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updated_at"` // 更新时间
}