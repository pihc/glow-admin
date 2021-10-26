package service

import (
	"context"
	"glow-admin/app/dao"
	"glow-admin/app/model"
	"glow-admin/app/module/admin/internal/define"
	"glow-admin/app/shared"
	"glow-admin/library/query"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"

	"github.com/gogf/gf/frame/g"
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
	return nil, nil
}

func (s *roleService) All() ([]*model.SysRole, error) {
	var temp []*model.SysRole
	err := dao.SysRole.Scan(&temp)
	if err != nil {
		return nil, err
	}

	return temp, nil
}

func (s *roleService) Create(ctx context.Context, req *define.RoleServiceCreateReq) (*define.RoleServiceCreateRes, error) {
	if req.CreatedBy == 0 {
		req.CreatedBy = shared.Context.Get(ctx).User.Id
	}
	lastId, err := dao.SysRole.Ctx(ctx).Data(req).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	return &define.RoleServiceCreateRes{RoleId: uint(lastId)}, nil
}

func (s *roleService) Update(ctx context.Context, req *define.RoleServiceUpdateReq) error {
	if req.UpdatedBy == 0 {
		req.UpdatedBy = shared.Context.Get(ctx).User.Id
	}
	_, err := dao.SysRole.Ctx(ctx).Data(req).FieldsEx(dao.SysRole.C.Id).Where(dao.SysRole.C.Id, req.Id).Update()
	return err
}

func (s *roleService) Delete(ctx context.Context, id uint) error {
	_, err := dao.SysRole.Ctx(ctx).Where(dao.SysRole.C.Id, id).Delete()
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
		if _, err = dao.SysRoleMenu.Ctx(ctx).Where(dao.SysRoleMenu.C.RoleId, req.RoleId).Delete(); err != nil {
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
	if err := dao.SysMenu.OrderAsc("sort").Where("status = 1").Scan(&list); err != nil {
		return nil, err
	}
	//cond := builder.NewCond()
	//cond = cond.And(
	//	builder.Eq{"status": 1},
	//)
	//if err := query.All(, cond, &list); err != nil {
	//	return nil, err
	//}

	// 获取角色拥有权限的菜单
	menuIds, err := dao.SysRoleMenu.Where(dao.SysRoleMenu.C.RoleId, roleId).Fields(dao.SysRoleMenu.C.MenuId).Array()
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
