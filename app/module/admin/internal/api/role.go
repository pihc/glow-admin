package api

import (
	"glow-admin/app/module/admin/internal/define"
	"glow-admin/app/module/admin/internal/service"
	"glow-admin/library/respond"
	"glow-admin/library/result"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Role = roleApi{}

type roleApi struct {
}

func (a *roleApi) GetList(r *ghttp.Request) respond.Json {
	var (
		data *define.RoleServiceGetListReq
	)
	if err := r.Parse(&data); err != nil {
		return result.Error(err)
	}
	return result.Response(service.Role.GetList(r.Context(), data))
}

func (a *roleApi) All(r *ghttp.Request) respond.Json {
	return result.Response(service.Role.All())
}

func (a *roleApi) Create(r *ghttp.Request) respond.Json {
	var (
		data             *define.RoleApiCreateReq
		serviceCreateReq *define.RoleServiceCreateReq
	)
	if err := r.ParseForm(&data); err != nil {
		return result.Error(err)
	}
	if err := gconv.Struct(data, &serviceCreateReq); err != nil {
		return result.Error(err)
	}

	if res, err := service.Role.Create(r.Context(), serviceCreateReq); err != nil {
		return result.Error(err)
	} else {
		return result.Success(res, "添加成功")
	}
}

func (a *roleApi) Update(r *ghttp.Request) respond.Json {
	var (
		data             *define.RoleApiUpdateReq
		serviceUpdateReq *define.RoleServiceUpdateReq
	)
	if err := r.ParseForm(&data); err != nil {
		return result.Error(err)
	}
	if err := gconv.Struct(data, &serviceUpdateReq); err != nil {
		return result.Error(err)
	}
	if err := service.Role.Update(r.Context(), serviceUpdateReq); err != nil {
		return result.Error(err)
	}
	return result.Success("", "编辑成功")
}

func (a *roleApi) Delete(r *ghttp.Request) respond.Json {
	var (
		data *define.RoleApiDeleteReq
	)
	if err := r.Parse(&data); err != nil {
		return result.Error(err)
	}
	if err := service.Role.Delete(r.Context(), data.Id); err != nil {
		return result.Error(err)
	}
	return result.Success("", "删除成功")
}

func (a *roleApi) GetMenus(r *ghttp.Request) respond.Json {
	var (
		data *define.RoleApiGetMenusReq
	)
	if err := r.Parse(&data); err != nil {
		return result.Error(err)
	}
	return result.Response(service.Role.GetMenus(data.Id))
}

func (a *roleApi) SetMenus(r *ghttp.Request) respond.Json {
	var (
		data       *define.RoleApiSetMenusReq
		serviceReq *define.RoleServiceSetMenusReq
	)
	if err := r.ParseForm(&data); err != nil {
		return result.Error(err)
	}
	if err := gconv.Struct(data, &serviceReq); err != nil {
		return result.Error(err)
	}
	if err := service.Role.SetMenus(r.Context(), serviceReq); err != nil {
		return result.Error(err)
	}
	return result.Success(nil, "设置成功")
}
