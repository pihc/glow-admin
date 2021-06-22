package service

import (
	"context"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"payget/app/dao"
	"payget/app/model"
	"payget/app/module/admin/internal/define"
	"payget/app/shared"
	"payget/library/tools"
)

var User = userService{}

type userService struct {
}

func (u *userService) Login(ctx context.Context, dto model.LoginDTO) (*model.SysUser, error) {
	record, err := dao.SysUser.Where(g.Map{
		dao.SysUser.Columns.Username: dto.Username,
	}).One()
	if err != nil {
		return nil, err
	}
	var user model.SysUser
	if err := record.Struct(&user); err != nil {
		return nil, err
	}
	if user.Password != tools.GenPassword(dto.Password) {
		return nil, gerror.New("用户名或密码错误")
	}
	return &user, nil
}

func (u *userService) GetUserInfo(ctx context.Context) (*define.RspUserInfo, error) {
	curUser := shared.Context.Get(ctx).User
	var user model.SysUser
	err := dao.SysUser.Data(g.Map{dao.SysUser.Columns.Id: curUser.Id}).Scan(&user)
	if err != nil {
		return nil, err
	}
	// 拷贝属性
	var userVO define.RspUserInfo
	if err := gconv.Struct(user, &userVO); err != nil {
		return nil, err
	}
	roles, err := Role.GetRolesByUserId(curUser.Id)
	if err != nil {
		return nil, err
	}
	userVO.Roles = roles

	// 获取权限菜单
	//menus,err := Menu.GetMenuList(curUser.Id)
	//if err!=nil{
	//	return nil,err
	//}
	//userVO.men
	//userInfoVo.setAuthorities(menuList)
	perms, err := Menu.GetPermissionList(user.Id)
	if err != nil {
		return nil, err
	}
	userVO.PermissionList = perms
	return &userVO, nil
}
