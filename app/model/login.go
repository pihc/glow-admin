package model

//LoginDTO 登录dto
type LoginDTO struct {
	Username string `json:"username" v:"required|length:4,30#请输入账号|账号长度为:min到:max位"`
	Password string `json:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
}
type DTOMenu struct {
	SysMenu
	Children []*DTOMenu `orm:"-" json:"children"`
}

const (
	MenuTypeMenu     = 0 // 菜单
	MenuTypeBtn      = 1 // 按钮
	MenuStatusShow   = 1 // 展示
	MenuStatusHidden = 0 // 隐藏
)

type RoleWithUserId struct {
	SysRole
	UserId uint `json:"user_id"`
}
