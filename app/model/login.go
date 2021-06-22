package model

//LoginDTO 登录dto
type LoginDTO struct {
	Username string `json:"username" v:"required|length:4,30#请输入账号|账号长度为:min到:max位"`
	Password string `json:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
}
