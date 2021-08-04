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
type UserApiResetPwdReq struct {
	Id uint `v:"min:1#请选择需要修改的用户"`
}

type UserApiDeleteReq struct {
	Id uint `v:"min:2#请选择id>1的用户进行删除"`
}

type UserApiChangePwdReq struct {
	OldPassword string `v:"required#请输入旧密码"`
	NewPassword string `v:"required#请输入新密码"`
}

type UserApiCreateUpdateBase struct {
	Nickname string `v:"required#请输入昵称"`    // 昵称
	Username string `v:"required#请输入登录用户名"` // 登录用户名
	Password string `v:"required#请输入登录密码"`  // 登录密码
	Avatar   string // 头像
	Mobile   string // 手机号码
	Email    string // 邮箱地址
	Intro    string // 个人简介
	Note     string // 备注
	Status   uint   `v:"required|between:1,2#请选择状态"` // 状态：1正常 2禁用
	RoleIds  []uint `v:"required#请选择角色"`
}

type UserApiCreateReq struct {
	UserApiCreateUpdateBase
}

type UserApiUpdateReq struct {
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
type UserServiceCreateReq struct {
	UserServiceCreateUpdateBase
	CreatedBy uint `json:"created_by"`
}

type UserServiceUpdateReq struct {
	UserServiceCreateUpdateBase
	Id        uint `json:"id"`
	UpdatedBy uint `json:"updated_by"`
}

type UserServiceCreateUpdateBase struct {
	Nickname string `json:"nickname"` // 昵称
	Username string `json:"username"` // 登录用户名
	Password string `json:"password"` // 登录密码
	Mobile   string `json:"mobile"`   // 手机号码
	Email    string `json:"email"`    // 邮箱地址
	Intro    string `json:"intro"`    // 个人简介
	Note     string `json:"note"`     // 备注
	Status   uint   `json:"status"`   // 状态：1正常 2禁用
	RoleIds  []uint `json:"role_ids"`
}

type UserServiceCreateRes struct {
	UserId uint `json:"user_id"`
}

type UserInfoRes struct {
	model.SysUser
	Roles       []*model.SysRole `json:"roles"`
	Permissions []string         `json:"permissions"`
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
	Id     uint `json:"id"`
	Status uint `json:"status"`
}

type UserServiceChangePwdReq struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
