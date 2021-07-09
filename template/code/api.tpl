package api

import (
	"github.com/gogf/gf/net/ghttp"
    "github.com/gogf/gf/util/gconv"
    "payget/app/module/admin/internal/define"
    "payget/app/module/admin/internal/service"
    "payget/library/respond"
    "payget/library/result"
)

var ${.table.short_name|UcFirst} = ${.table.short_name}Api{}

type ${.table.short_name}Api struct {
}

func (a *${.table.short_name}Api) GetList(r *ghttp.Request) respond.Json {
	var (
		data *define.${.table.short_name|UcFirst}ServiceGetListReq
	)
	if err := r.Parse(&data); err != nil {
		return result.Error(err)
	}
	return result.Response(service.${.table.short_name|UcFirst}.GetList(r.Context(), data))
}

func (a *${.table.short_name}Api) All(r *ghttp.Request) respond.Json {
	return result.Response(service.${.table.short_name|UcFirst}.All())
}

func (a *${.table.short_name}Api) Create(r *ghttp.Request) respond.Json {
	var (
		data             *define.${.table.short_name|UcFirst}ApiCreateReq
		serviceCreateReq *define.${.table.short_name|UcFirst}ServiceCreateReq
	)
	if err := r.ParseForm(&data); err != nil {
		return result.Error(err)
	}
	if err := gconv.Struct(data, &serviceCreateReq); err != nil {
		return result.Error(err)
	}

	if res, err := service.${.table.short_name|UcFirst}.Create(r.Context(), serviceCreateReq); err != nil {
		return result.Error(err)
	} else {
		return result.Success(res, "添加成功")
	}
}

func (a *${.table.short_name}Api) Update(r *ghttp.Request) respond.Json {
	var (
		data             *define.${.table.short_name|UcFirst}ApiUpdateReq
		serviceUpdateReq *define.${.table.short_name|UcFirst}ServiceUpdateReq
	)
	if err := r.ParseForm(&data); err != nil {
		return result.Error(err)
	}
	if err := gconv.Struct(data, &serviceUpdateReq); err != nil {
		return result.Error(err)
	}
	if err := service.${.table.short_name|UcFirst}.Update(r.Context(), serviceUpdateReq); err != nil {
		return result.Error(err)
	}
	return result.Success("", "编辑成功")
}

func (a *${.table.short_name}Api) Delete(r *ghttp.Request) respond.Json {
	var (
		data *define.${.table.short_name|UcFirst}ApiDeleteReq
	)
	if err := r.Parse(&data); err != nil {
		return result.Error(err)
	}
	if err := service.${.table.short_name|UcFirst}.Delete(r.Context(), data.Id); err != nil {
		return result.Error(err)
	}
	return result.Success("", "删除成功")
}
