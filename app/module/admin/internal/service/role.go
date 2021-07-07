package service

import (
	"context"
	"payget/app/dao"
	"payget/app/model"
	"payget/app/module/admin/internal/define"
	"payget/app/shared"
	"payget/library/query"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/errors/gerror"

	"github.com/gogf/gf/frame/g"
	"xorm.io/builder"
)

var Role = roleService{}

type roleService struct {
}

func (s *roleService) GetList(ctx context.Context, req *define.RoleServiceGetListReq) (*query.Result, error) {
	roles := make([]*model.SysRole, 0)
	result, err := query.Page(dao.SysRole.M, req, &roles)
	if err != nil {
		return nil, err
	}
	return result.WithRecords(roles), nil
}

func (s *roleService) All() ([]*model.SysRole, error) {
	var temp []*model.SysRole
	err := dao.SysRole.Scan(&temp)
	if err != nil {
		return nil, err
	}

	return temp, nil
}

func (s *roleService) Create(ctx context.Context, req *define.RoleServiceDoCreateReq) (*define.RoleServiceCreateRes, error) {
	curUser := shared.Context.Get(ctx).User
	roleMap := gconv.Map(req)
	roleMap["created_by"] = curUser.Id
	result, err := dao.SysRole.Ctx(ctx).Data(roleMap).Insert()
	if err != nil {
		return nil, err
	}
	n, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &define.RoleServiceCreateRes{RoleId: uint(n)}, nil
}

func (s *roleService) Update(ctx context.Context, req *define.RoleServiceDoUpdateReq) error {
	curUser := shared.Context.Get(ctx).User
	roleMap := gconv.Map(req)
	roleMap["updated_by"] = curUser.Id
	_, err := dao.SysRole.
		Ctx(ctx).
		Data(roleMap).
		FieldsEx(dao.SysRole.Columns.Id).
		Where(dao.SysRole.Columns.Id, req.Id).
		Update()
	return err
}

func (s *roleService) Delete(ctx context.Context, id uint) error {
	_, err := dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns.Id, id).Delete()
	return err
}

func (s *roleService) GetRolesByUserId(userID uint) (list []*model.SysRole, err error) {
	err = dao.SysUserRole.As("ur").Data(g.Map{
		"ur.user_id": userID,
	}).InnerJoin("sys_user u", "ur.user_id=u.id").
		InnerJoin("sys_role r", "ur.role_id=r.id").Scan(&list)
	return
}

func (s *roleService) SetMenus(ctx context.Context, req *define.RoleServiceSetMenusReq) error {
	role, err := dao.SysRole.FindOne(req.RoleId)
	if err != nil {
		return err
	}
	if role == nil {
		return gerror.New("角色不存在")
	}

	var temps []model.SysRoleMenu
	for _, v := range req.MenuIds {
		temps = append(temps, model.SysRoleMenu{
			RoleId: req.RoleId,
			MenuId: v,
		})
	}

	return dao.SysRole.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 删除角色和菜单之间的关联
		if _, err = dao.SysRoleMenu.Ctx(ctx).Where(dao.SysRoleMenu.Columns.RoleId, req.RoleId).Delete(); err != nil {
			return err
		}
		// 保存角色所拥有的菜单
		if _, err = dao.SysRoleMenu.Ctx(ctx).Data(temps).Insert(); err != nil {
			return err
		}
		return nil
	})
}

func (s *roleService) GetMenus(roleId uint) ([]*define.RoleServiceMenuListRes, error) {
	// 获取所有菜单
	var list []*define.RoleServiceMenuListRes
	cond := builder.NewCond()
	cond = cond.And(
		builder.Eq{"status": 1},
	)
	if err := query.All(dao.SysMenu.OrderAsc("sort"), cond, &list); err != nil {
		return nil, err
	}

	// 获取角色拥有权限的菜单
	menuIds, err := dao.SysRoleMenu.Where(dao.SysRoleMenu.Columns.RoleId, roleId).Fields(dao.SysRoleMenu.Columns.MenuId).Array()
	if err != nil {
		return nil, err
	}

	for _, v := range list {
		for _, menuId := range menuIds {
			if menuId.Uint() == v.Id {
				v.Checked = true
				v.Open = true
				break
			}
		}
	}

	return list, nil
}
