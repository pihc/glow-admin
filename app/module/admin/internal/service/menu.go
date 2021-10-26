package service

import (
	"context"
	"glow-admin/app/dao"
	"glow-admin/app/model"
	"glow-admin/app/module/admin/internal/define"
	"glow-admin/app/shared"
	"glow-admin/library/query"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var Menu = menuService{
	permissionMap: map[string]string{"index": "查询", "add": "添加", "edit": "修改", "delete": "删除"},
}

type menuService struct {
	permissionMap map[string]string
}

func (s *menuService) GetList(ctx context.Context, req *define.MenuServiceGetListReq) ([]*model.SysMenu, error) {
	temp := make([]*model.SysMenu, 0)
	if err := query.All(dao.SysMenu.M, req, &temp); err != nil {
		return nil, err
	}
	return temp, nil
}

func (s *menuService) Delete(ctx context.Context, id uint) error {
	_, err := dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.C.Id, id).Delete()
	return err
}

func (s *menuService) GetDetail(ctx context.Context, id uint) (*model.SysMenu, error) {
	var menu model.SysMenu
	if err := dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.C.Id, id).Scan(&menu); err != nil {
		return nil, err
	}
	return &menu, nil
}

func (s *menuService) Create(ctx context.Context, req *define.MenuServiceCreateReq) (*define.MenuServiceCreateRes, error) {
	if req.CreatedBy == 0 {
		req.CreatedBy = shared.Context.Get(ctx).User.Id
	}
	// 添加菜单
	lastId, err := dao.SysMenu.Ctx(ctx).Data(req).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	// 添加权限
	if len(req.Nodes) > 0 {
		var menus []model.SysMenu
		for k, v := range req.Nodes {
			menus = append(menus, model.SysMenu{
				Pid:        uint(lastId),
				Title:      s.permissionMap[v] + req.Title,
				Target:     "_self",
				Permission: strings.Replace(strings.TrimLeft(req.Path, "/"), "/", ":", -1) + ":" + v,
				Type:       model.MenuTypeBtn,
				Status:     1,
				Note:       "",
				Sort:       uint((k + 1) * 5),
				CreatedBy:  req.CreatedBy,
			})
		}
		if _, err := dao.SysMenu.Ctx(ctx).Data(menus).Insert(); err != nil {
			return nil, err
		}
	}

	return &define.MenuServiceCreateRes{MenuId: uint(lastId)}, nil
}

func (s *menuService) Update(ctx context.Context, req *define.MenuServiceUpdateReq) error {
	curUser := shared.Context.Get(ctx).User
	menuMap := gconv.Map(req)
	menuMap["updated_by"] = curUser.Id
	_, err := dao.SysMenu.Ctx(ctx).Data(menuMap).FieldsEx(dao.SysMenu.C.Id).Where(dao.SysMenu.C.Id, req.Id).Update()
	return err
}

func (s *menuService) GetMenuList(userID uint) (list []*model.DTOMenu, err error) {
	if userID == 1 {
		return s.GetChildMenuAll(0)
	}
	return s.GetChildrenMenuByPid(userID, 0)
}

func (s *menuService) GetChildMenuAll(pid uint) (list []*model.DTOMenu, err error) {
	//cond := builder.NewCond()
	//cond = cond.And(
	//	builder.Eq{"pid": pid},
	//	builder.Eq{"status": 1},
	//	builder.Eq{"type": 0},
	//)
	if err = dao.SysMenu.OrderAsc("sort").Where("pid = ? AND status = 1 AND type = 0", pid).Scan(&list); err != nil {
		return
	}
	//err = query.All(dao.SysMenu.OrderAsc("sort"), cond, &list)
	//if err != nil {
	//	return
	//}
	for _, v := range list {
		var child []*model.DTOMenu
		child, err = s.GetChildMenuAll(v.Id)
		if err != nil {
			return nil, err
		}
		v.Children = child
	}
	return
}

func (s *menuService) GetChildrenMenuByPid(userID, pid uint) (list []*model.DTOMenu, err error) {
	list, err = dao.GetPermissionsListByUserId(userID, pid)
	if err != nil {
		return
	}
	for _, v := range list {
		var child []*model.DTOMenu
		child, err = s.GetChildrenMenuByPid(userID, v.Id)
		if err != nil {
			return
		}
		v.Children = child
	}
	return
}

func (s *menuService) GetPermissionList(userID uint) ([]string, error) {
	var (
		perm  []*model.SysMenu
		perms []string
	)
	if userID == 1 {
		// 超级管理员
		if err := dao.SysMenu.Data(g.Map{
			dao.SysMenu.C.Type:   model.MenuTypeBtn,
			dao.SysMenu.C.Status: model.MenuStatusShow,
		}).Scan(&perm); err != nil {
			return nil, err
		}
		for _, v := range perm {
			perms = append(perms, v.Permission)
		}
	} else {
		perm, err := dao.GetPermissionList(userID)
		if err != nil {
			return nil, err
		}
		for _, v := range perm {
			perms = append(perms, v.Permission)
		}
	}
	return perms, nil
}
