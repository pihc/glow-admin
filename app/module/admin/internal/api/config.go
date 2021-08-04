package api

import (
	"glow-admin/app/module/admin/internal/define"
	"glow-admin/app/module/admin/internal/service"
	"glow-admin/library/respond"
	"glow-admin/library/result"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Config = configApi{}

type configApi struct {
}

// 配置分组分页
func (a *configApi) GetList(r *ghttp.Request) respond.Json {
	var (
		data *define.ConfigServiceGetListReq
	)
	if err := r.Parse(&data); err != nil {
		return result.Error(err)
	}
	return result.Response(service.Config.GetList(r.Context(), data))
}

// 配置分组列表
func (a *configApi) GetAll(r *ghttp.Request) respond.Json {
	return result.Response(service.Config.GetAll())
}

// 配置分组明细
func (a *configApi) GetDetail(r *ghttp.Request) respond.Json {
	var (
		data *define.ConfigApiDetailReq
	)
	if err := r.Parse(&data); err != nil {
		return result.Error(err)
	}
	return result.Response(service.Config.GetDetail(r.Context(), data.Id))
}

// 配置分组新增
func (a *configApi) Create(r *ghttp.Request) respond.Json {
	var (
		data             *define.ConfigApiCreateReq
		serviceCreateReq *define.ConfigServiceCreateReq
	)
	if err := r.ParseForm(&data); err != nil {
		return result.Error(err)
	}
	if err := gconv.Struct(data, &serviceCreateReq); err != nil {
		return result.Error(err)
	}

	if res, err := service.Config.Create(r.Context(), serviceCreateReq); err != nil {
		return result.Error(err)
	} else {
		return result.Success(res, "添加成功")
	}
}

// 配置分组编辑
func (a *configApi) Update(r *ghttp.Request) respond.Json {
	var (
		data             *define.ConfigApiUpdateReq
		serviceUpdateReq *define.ConfigServiceUpdateReq
	)
	if err := r.ParseForm(&data); err != nil {
		return result.Error(err)
	}
	if err := gconv.Struct(data, &serviceUpdateReq); err != nil {
		return result.Error(err)
	}
	if err := service.Config.Update(r.Context(), serviceUpdateReq); err != nil {
		return result.Error(err)
	}
	return result.Success("", "编辑成功")
}

// 配置分组删除
func (a *configApi) Delete(r *ghttp.Request) respond.Json {
	var (
		data *define.ConfigApiDeleteReq
	)
	if err := r.Parse(&data); err != nil {
		return result.Error(err)
	}
	if err := service.Config.Delete(r.Context(), data.Id); err != nil {
		return result.Error(err)
	}
	return result.Success("", "删除成功")
}
