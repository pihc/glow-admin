package service

import (
	"context"
	"glow-admin/app/dao"
	"glow-admin/app/model"
	"glow-admin/app/module/admin/internal/define"
	"glow-admin/app/shared"
	"glow-admin/library/query"
	"glow-admin/library/token"
	"glow-admin/library/tools"

	"github.com/gogf/gf/util/grand"

	"github.com/gogf/gf/database/gdb"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var User = userService{}

type userService struct {
}

// 登录
func (s *userService) Login(ctx context.Context, req *model.LoginDTO) (g.Map, error) {
	user, err := s.InfoByUsername(ctx, req.Username)
	if err != nil || user == nil {
		return nil, gerror.New("用户名或密码错误")
	}
	pwd := tools.GenPassword(req.Password, user.Salt)
	if user.Password != pwd {
		return nil, gerror.New("用户名或密码错误")
	}
	// 生成token
	expire := gconv.Int(g.Config().Get("jwt.expire"))
	tk, err := token.Generate(model.ContextUser{
		Id: user.Id,
	}, expire)

	return g.Map{
		"access_token": tk,
	}, nil
}

// 获取用户信息通过账号
func (s *userService) InfoByUsername(ctx context.Context, username string) (*model.SysUser, error) {
	user := (*model.SysUser)(nil)
	if err := dao.SysUser.Ctx(ctx).Where(g.Map{dao.SysUser.C.Username: username}).Scan(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserInfo(ctx context.Context) (*define.UserInfoRes, error) {
	curUser := shared.Context.Get(ctx).User
	var user model.SysUser
	err := dao.SysUser.WherePri(curUser.Id).Scan(&user)
	if err != nil {
		return nil, err
	}

	var userVO define.UserInfoRes
	if err := gconv.Struct(user, &userVO); err != nil {
		return nil, err
	}
	// 角色
	roles, err := Role.GetRolesByUserId(user.Id)
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

func (s *userService) Create(ctx context.Context, req *define.UserServiceCreateReq) (*define.UserServiceCreateRes, error) {
	if req.CreatedBy == 0 {
		req.CreatedBy = shared.Context.Get(ctx).User.Id
	}
	userMap := gconv.Map(req)
	userMap["salt"] = grand.Letters(6) // 盐
	userMap["password"] = tools.GenPassword(req.Password, userMap["salt"].(string))

	lastId := 0
	if err := dao.SysRole.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 新增用户
		lastId, err := dao.SysUser.Ctx(ctx).Data(userMap).InsertAndGetId()
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

func (s *userService) Update(ctx context.Context, req *define.UserServiceUpdateReq) error {
	if req.UpdatedBy == 0 {
		req.UpdatedBy = shared.Context.Get(ctx).User.Id
	}

	return dao.SysRole.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 编辑用户
		if _, err := dao.SysUser.Ctx(ctx).Data(req).FieldsEx(dao.SysUser.C.Id, dao.SysUser.C.Password).Where(dao.SysUser.C.Id, req.Id).Update(); err != nil {
			return err
		}

		// 删除用户和角色的关联
		if _, err := dao.SysUserRole.Ctx(ctx).Where(dao.SysUserRole.C.UserId, req.Id).Delete(); err != nil {
			return err
		}

		// 重新关联角色
		var temp []*model.SysUserRole
		for _, v := range req.RoleIds {
			temp = append(temp, &model.SysUserRole{
				UserId: req.Id,
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
		if _, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.C.Id, id).Delete(); err != nil {
			return err
		}
		// 删除用户和角色的关联
		if _, err := dao.SysUserRole.Ctx(ctx).Where(dao.SysUserRole.C.UserId, id).Delete(); err != nil {
			return err
		}
		return nil
	})
}

func (s *userService) ChangeStatus(ctx context.Context, req *define.UserServiceChangeStatusReq) error {
	_, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.C.Id, req.Id).Data(req).FieldsEx(dao.SysUser.C.Id).Update()
	return err
}

func (s *userService) ResetPassword(ctx context.Context, id uint) error {
	salt := grand.Letters(6) // 盐
	password := tools.GenPassword("111111", salt)
	_, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.C.Id, id).Data(g.Map{
		dao.SysUser.C.Password: password,
		dao.SysUser.C.Salt:     salt,
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

	_, err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.C.Id, user.Id).Data(g.Map{
		dao.SysUser.C.Password:  newPwd,
		dao.SysUser.C.Salt:      salt,
		dao.SysUser.C.UpdatedBy: curUser.Id,
	}).Update()

	return err
}
