package service

import (
	"github.com/gogf/gf/frame/g"
	"payget/app/dao"
	"payget/app/model"
)

var Role = roleService{}

type roleService struct {
}

func (r *roleService) GetRolesByUserId(userID uint) (list []*model.SysRole, err error) {
	err = dao.SysUserRole.As("ur").Data(g.Map{
		"ur.user_id": userID,
	}).InnerJoin("sys_user u", "ur.user_id=u.id").
		InnerJoin("sys_role r", "ur.role_id=r.id").Scan(&list)
	return
}
