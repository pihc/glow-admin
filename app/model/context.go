package model

import (
	"github.com/gogf/gf/frame/g"
)

const (
	// 上下文变量存储键名，前后端系统共享
	ContextKey = "ContextKey"
)

// 请求上下文结构
type Context struct {
	User *ContextUser // 上下文用户信息
	Data g.Map        // 自定KV变量，业务模块根据需要设置，不固定
}

// 请求上下文中的用户信息
type ContextUser struct {
	Id uint // 用户ID
}
