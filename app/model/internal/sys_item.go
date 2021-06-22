// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// SysItem is the golang structure for table sys_item.
type SysItem struct {
	Id         uint        `orm:"id,primary"  json:"id"`         // 唯一性标识
	Name       string      `orm:"name"        json:"name"`       // 站点名称
	Type       uint        `orm:"type"        json:"type"`       // 站点类型:1普通站点 2其他
	Url        string      `orm:"url"         json:"url"`        // 站点地址
	Image      string      `orm:"image"       json:"image"`      // 站点图片
	IsDomain   uint        `orm:"is_domain"   json:"isDomain"`   // 是否二级域名:1是 2否
	Status     uint        `orm:"status"      json:"status"`     // 状态：1在用 2停用
	Note       string      `orm:"note"        json:"note"`       // 站点备注
	Sort       uint        `orm:"sort"        json:"sort"`       // 显示顺序
	CreateUser uint        `orm:"create_user" json:"createUser"` // 添加人
	CreateTime *gtime.Time `orm:"create_time" json:"createTime"` // 添加时间
	UpdateUser uint        `orm:"update_user" json:"updateUser"` // 更新人
	UpdateTime *gtime.Time `orm:"update_time" json:"updateTime"` // 更新时间
	Mark       uint        `orm:"mark"        json:"mark"`       // 有效标识(1正常 0删除)
}