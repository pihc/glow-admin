package service

import (
	"github.com/gogf/gf/frame/g"
	"payget/app/dao"
	"payget/app/model"
	"payget/library/query"
	"xorm.io/builder"
)

var Menu = menuService{}

type menuService struct {
}

func (m *menuService) GetMenuList(userID uint) (list []*model.DTOMenu, err error) {
	if userID == 1 {
		return m.GetChildMenuAll(0)
	}
	return m.GetChildrenMenuByPid(userID, 0)
}
func (m *menuService) GetChildMenuAll(pid uint) (list []*model.DTOMenu, err error) {
	cond := builder.NewCond()
	cond = cond.And(
		builder.Eq{"pid": pid},
		builder.Eq{"status": 1},
		builder.Eq{"type": 0},
	)
	err = query.All(dao.SysMenu.OrderAsc("sort"), cond, &list)
	if err != nil {
		return
	}
	for _, v := range list {
		var child []*model.DTOMenu
		child, err = m.GetChildMenuAll(v.Id)
		if err != nil {
			return nil, err
		}
		v.Children = child
	}
	return
}

func (m *menuService) GetChildrenMenuByPid(userID, pid uint) (list []*model.DTOMenu, err error) {
	list, err = dao.GetPermissionsListByUserId(userID, pid)
	if err != nil {
		return
	}
	for _, v := range list {
		var child []*model.DTOMenu
		child, err = m.GetChildrenMenuByPid(userID, v.Id)
		if err != nil {
			return
		}
		v.Children = child
	}
	return
}
func (m *menuService) GetPermissionList(userID uint) ([]string, error) {
	var (
		perm  []*model.SysMenu
		perms []string
	)
	if userID == 1 {
		// 超级管理员
		if err := dao.SysMenu.Data(g.Map{
			dao.SysMenu.Columns.Type:   1,
			dao.SysMenu.Columns.Status: 1,
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
