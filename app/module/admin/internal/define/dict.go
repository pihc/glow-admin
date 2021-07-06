package define

// ==========================================================================================
// API
// ==========================================================================================
type DictApiDeleteReq struct {
	Id uint `v:"min:1#请选择需要删除的用户"`
}

type DictApiCreateUpdateBase struct {
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

type DictApiDoCreateReq struct {
	DictApiCreateUpdateBase
}

type DictApiDoUpdateReq struct {
	DictApiCreateUpdateBase
	Id uint `v:"min:1#请选择需要修改的用户"`
}

// ==========================================================================================
// Service
// ==========================================================================================
type DictServiceDoCreateReq struct {
	DictServiceCreateUpdateBase
}

type DictServiceDoUpdateReq struct {
	DictServiceCreateUpdateBase
	Id uint
}

type DictServiceCreateUpdateBase struct {
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

type DictServiceCreateRes struct {
	DictId uint `json:"dict_id"`
}
