package define

//LoginReq 登录请求
type LoginReq struct {
	Username string `json:"username" v:"required|length:4,30#请输入账号|账号长度为:min到:max位"`
	Password string `json:"password" v:"required|length:4,30#请输入密码|密码长度不够"`
	Captcha  string `json:"captcha" v:"required#验证码不能为空"`
	Key      string `json:"key" v:"required#验证码KEY不能为空"`
}
