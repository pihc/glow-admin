package service

import (
	"context"
	"glow-admin/app/dao"
	"glow-admin/app/model"
	"glow-admin/app/module/admin/internal/define"
	"glow-admin/app/shared"
	"glow-admin/library/query"
)

var Config = configService{}

type configService struct {
}

// 配置分组分页
func (s *configService) GetList(ctx context.Context, req *define.ConfigServiceGetListReq) (*query.Result, error) {
	temp := make([]*model.SysConfig, 0)
	result, err := query.Page(dao.SysConfig.M, req, &temp)
	if err != nil {
		return nil, err
	}
	return result.WithRecords(temp), nil
}

// 配置分组列表
func (s *configService) GetAll() ([]*model.SysConfig, error) {
	var temp []*model.SysConfig
	err := dao.SysConfig.Scan(&temp)
	if err != nil {
		return nil, err
	}

	return temp, nil
}

// 配置分组明细
func (s *configService) GetDetail(ctx context.Context, id uint) (*model.SysConfig, error) {
	var config model.SysConfig
	if err := dao.SysConfig.Ctx(ctx).Where(dao.SysConfig.Columns.Id, id).Scan(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

// 配置分组新增
func (s *configService) Create(ctx context.Context, req *define.ConfigServiceCreateReq) (*define.ConfigServiceCreateRes, error) {
	if req.CreatedBy == 0 {
		req.CreatedBy = shared.Context.Get(ctx).User.Id
	}
	lastId, err := dao.SysConfig.Ctx(ctx).Data(req).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	return &define.ConfigServiceCreateRes{ConfigId: uint(lastId)}, nil
}

// 配置分组编辑
func (s *configService) Update(ctx context.Context, req *define.ConfigServiceUpdateReq) error {
	if req.UpdatedBy == 0 {
		req.UpdatedBy = shared.Context.Get(ctx).User.Id
	}
	_, err := dao.SysConfig.Ctx(ctx).Data(req).FieldsEx(dao.SysConfig.Columns.Id).Where(dao.SysConfig.Columns.Id, req.Id).Update()
	return err
}

// 配置分组删除
func (s *configService) Delete(ctx context.Context, id uint) error {
	_, err := dao.SysConfig.Ctx(ctx).Where(dao.SysConfig.Columns.Id, id).Delete()
	return err
}
