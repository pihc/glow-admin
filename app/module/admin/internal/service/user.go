package service

import (
	"context"
	"payget/app/dao"
	"payget/app/model"
	"payget/app/module/admin/internal/define"
	"payget/app/shared"
	"payget/library/query"
	"payget/library/tools"

	"github.com/gogf/gf/util/grand"

	"github.com/gogf/gf/database/gdb"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var User = userService{}

type userService struct {
}

func (s *userService) Login(ctx context.Context, dto model.LoginDTO) (*model.SysUser, error) {
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

	if user.Password != tools.GenPassword(dto.Password, user.Salt) {
		return nil, gerror.New("用户名或密码错误")
	}
	return &user, nil
}

func (s *userService) GetUserInfo(ctx context.Context) (*define.UserInfoRes, error) {
	curUser := shared.Context.Get(ctx).User
	var user model.SysUser
	err := dao.SysUser.Data(g.Map{dao.SysUser.Columns.Id: curUser.Id}).Scan(&user)
	if err != nil {
		return nil, err
	}

	var userVO define.UserInfoRes
	if err := gconv.Struct(user, &userVO); err != nil {
		return nil, err
	}
	// 角色
	roles, err := Role.GetRolesByUserId(curUser.Id)
	if err != nil {
		return nil, err
	}
	userVO.Roles = roles
	// 权限
	perms, err := Menu.GetPermissionList(user.Id)
	if err != nil {
		return nil, err
	}
	userVO.Permissions = perms
	return &userVO, nil
}

func (s *userService) GetList(ctx context.Context, req *define.UserServiceGetListReq) (*query.Result, error) {
	temp := make([]*define.UserServiceGetListRes, 0)
	result, err := query.Page(dao.SysUser.M, req, &temp)
	if err != nil {
		return nil, err
	}
	var userIds []uint
	for _, v := range temp {
		userIds = append(userIds, v.Id)
	}

	roles, err := dao.GetRoleList(userIds)
	if err != nil {
		return nil, err
	}

	// 外挂
	for _, v := range temp {
		data, ok := roles[v.Id]
		if ok {
			v.Roles = data
		}
	}

	return result.WithRecords(temp), nil
}

func (s *userService) Create(ctx context.Context, req *define.UserServiceDoCreateReq) (*define.UserServiceCreateRes, error) {
	curUser := shared.Context.Get(ctx).User
	var user model.SysUser
	if err := gconv.Struct(req, &user); err != nil {
		return nil, err
	}
	user.CreatedBy = curUser.Id
	user.Salt = grand.Letters(6) // 盐
	user.Password = tools.GenPassword(req.Password, user.Salt)

	lastId := 0
	if err := dao.SysRole.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 新增用户
		lastId, err := dao.SysUser.Ctx(ctx).Data(user).InsertAndGetId()
		if err != nil {
			return err
		}

		// 关联角色
		var temp []*model.SysUserRole
		for _, v := range req.RoleIds {
			temp = append(temp, &model.SysUserRole{
				UserId: uint(lastId),
				RoleId: v,
			})
		}
		if _, err := dao.SysUserRole.Ctx(ctx).Data(temp).Insert(); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &define.UserServiceCreateRes{UserId: uint(lastId)}, nil
}

func (s *userService) Update(ctx context.Context, req *define.UserServiceDoUpdateReq) error {
	curUser := shared.Context.Get(ctx).User
	var user model.SysUser
	if err := gconv.Struct(req, &user); err != nil {
		return err
	}
	user.UpdateBy = curUser.Id

	return dao.SysRole.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 编辑用户
		if _, err := dao.SysUser.Ctx(ctx).Data(user).FieldsEx(dao.SysUser.Columns.Id, dao.SysUser.Columns.Password).Where(dao.SysUser.Columns.Id, req.Id).Update(); err != nil {
			return err
		}

		// 删除用户和角色的关联
		if _, err := dao.SysUserRole.Ctx(ctx).Where(dao.SysUserRole.Columns.UserId, user.Id).Delete(); err != nil {
			return err
		}

		// 重新关联角色
		var temp []*model.SysUserRole
		for _, v := range req.RoleIds {
			temp = append(temp, &model.SysUserRole{
				UserId: user.Id,
				RoleId: v,
			})
		}
		if _, err := dao.SysUserRole.Ctx(ctx).Data(temp).Insert(); err != nil {
			return err
		}
		return nil
	})
}

func (s *userService) Delete(ctx context.Context, id uint) error {
	return dao.SysRole.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 删除用户
		if _, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns.Id, id).Delete(); err != nil {
			return err
		}
		// 删除用户和角色的关联
		if _, err := dao.SysUserRole.Ctx(ctx).Where(dao.SysUserRole.Columns.UserId, id).Delete(); err != nil {
			return err
		}
		return nil
	})
}

func (s *userService) ChangeStatus(ctx context.Context, req *define.UserServiceChangeStatusReq) error {
	_, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns.Id, req.Id).Data(req).FieldsEx(dao.SysUser.Columns.Id).Update()
	return err
}

func (s *userService) ResetPassword(ctx context.Context, id uint) error {
	salt := grand.Letters(6) // 盐
	password := tools.GenPassword("111111", salt)
	_, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns.Id, id).Data(g.Map{
		dao.SysUser.Columns.Password: password,
		dao.SysUser.Columns.Salt:     salt,
	}).Update()
	return err
}

func (s *userService) ChangePwd(ctx context.Context, req *define.UserServiceChangePwdReq) error {
	curUser := shared.Context.Get(ctx).User
	var user model.SysUser
	err := dao.SysUser.WherePri(curUser.Id).Scan(&user)
	if err != nil {
		return err
	}

	oldPwd := tools.GenPassword(req.OldPassword, user.Salt)
	if user.Password != oldPwd {
		return gerror.New("密码错误")
	}

	salt := grand.Letters(6)
	newPwd := tools.GenPassword(req.NewPassword, salt)

	_, err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns.Id, user.Id).Data(model.SysUser{
		Password: newPwd,
		Salt:     salt,
		UpdateBy: curUser.Id,
	}).Update()

	return err
}
