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
type UserApiCreateUpdateBase struct {
	Nickname string `v:"required#请输入昵称"`    // 昵称
	Username string `v:"required#请输入登录用户名"` // 登录用户名
	Password string `v:"required#请输入登录密码"`  // 登录密码
	Avatar   string // 头像
	Mobile   string // 手机号码
	Email    string // 邮箱地址
	Intro    string // 个人简介
	Note     string // 备注
	Status   uint   // 状态：1正常 2禁用
	RoleIds  []uint `v:"required#请选择角色"`
}

type UserApiDoCreateReq struct {
	UserApiCreateUpdateBase
}

type UserApiDoUpdateReq struct {
	UserApiCreateUpdateBase
	Id uint `v:"min:1#请选择需要修改的用户"`
}

type UserApiChangeStatusReq struct {
	Id     uint `v:"min:1#请选择需要修改的用户"`
	Status uint
}

// ==========================================================================================
// Service
// ==========================================================================================
type UserServiceDoCreateReq struct {
	UserServiceCreateUpdateBase
}

type UserServiceDoUpdateReq struct {
	UserServiceCreateUpdateBase
	Id uint
}

type UserServiceCreateUpdateBase struct {
	Nickname string // 昵称
	Username string // 登录用户名
	Password string // 登录密码
	Mobile   string // 手机号码
	Email    string // 邮箱地址
	Intro    string // 个人简介
	Note     string // 备注
	Status   uint   // 状态：1正常 2禁用
	RoleIds  []uint
}

type UserServiceCreateRes struct {
	UserId uint `json:"user_id"`
}

type UserInfoRes struct {
	model.SysUser
	Roles []*model.SysRole `json:"roles"`
	//Authorities    []*model.SysMenu `json:"authorities"`
	Permissions []string `json:"permissions"`
}

// 查询请求
type UserServiceGetListReq struct {
	query.Params
	Nickname string `json:"nickname"`
	Username string `json:"username"`
}

func (q *UserServiceGetListReq) Build() builder.Cond {
	cond := builder.NewCond()
	if q.Nickname != "" {
		cond = cond.And(builder.Like{"sys_user.nickname", strings.TrimSpace(q.Nickname)})
	}
	if q.Username != "" {
		cond = cond.And(builder.Like{"sys_user.username", strings.TrimSpace(q.Username)})
	}
	return cond
}

type UserServiceGetListRes struct {
	model.SysUser
	Roles []*model.SysRole `json:"roles"`
}

type UserServiceChangeStatusReq struct {
	Id     uint
	Status uint
}
