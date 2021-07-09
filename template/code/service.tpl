package service

import (
	"context"
	"payget/app/dao"
	"payget/app/model"
	"payget/app/module/admin/internal/define"
	"payget/app/shared"
	"payget/library/query"
)

var ${.table.short_name|UcFirst} = ${.table.short_name}Service{}

type ${.table.short_name}Service struct {
}

func (s *${.table.short_name}Service) GetList(ctx context.Context, req *define.${.table.short_name|UcFirst}ServiceGetListReq) (*query.Result, error) {
	temp := make([]*model.${.table.name|CaseCamel}, 0)
	result, err := query.Page(dao.${.table.name|CaseCamel}.M, req, &temp)
	if err != nil {
		return nil, err
	}
	return result.WithRecords(temp), nil
}

func (s *${.table.short_name}Service) All() ([]*model.${.table.name|CaseCamel}, error) {
	var temp []*model.${.table.name|CaseCamel}
	err := dao.${.table.name|CaseCamel}.Scan(&temp)
	if err != nil {
		return nil, err
	}

	return temp, nil
}

func (s *${.table.short_name}Service) Create(ctx context.Context, req *define.${.table.short_name|UcFirst}ServiceCreateReq) (*define.${.table.short_name|UcFirst}ServiceCreateRes, error) {
	if req.CreatedBy == 0 {
		req.CreatedBy = shared.Context.Get(ctx).User.Id
	}
	lastId, err := dao.${.table.name|CaseCamel}.Ctx(ctx).Data(req).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	return &define.${.table.short_name|UcFirst}ServiceCreateRes{${.table.short_name|UcFirst}Id: uint(lastId)}, nil
}

func (s *${.table.short_name}Service) Update(ctx context.Context, req *define.${.table.short_name|UcFirst}ServiceUpdateReq) error {
	if req.UpdatedBy == 0 {
		req.UpdatedBy = shared.Context.Get(ctx).User.Id
	}
	_, err := dao.${.table.name|CaseCamel}.Ctx(ctx).Data(req).FieldsEx(dao.${.table.name|CaseCamel}.Columns.Id).Where(dao.${.table.name|CaseCamel}.Columns.Id, req.Id).Update()
	return err
}

func (s *${.table.short_name}Service) Delete(ctx context.Context, id uint) error {
	_, err := dao.${.table.name|CaseCamel}.Ctx(ctx).Where(dao.${.table.name|CaseCamel}.Columns.Id, id).Delete()
	return err
}
