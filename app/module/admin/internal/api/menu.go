package api

import (
	"payget/app/module/admin/internal/define"
	"payget/app/module/admin/internal/service"
	"payget/library/respond"
	"payget/library/result"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/net/ghttp"
)

var Menu = menuApi{}

type menuApi struct {
}

func (a *menuApi) GetList(r *ghttp.Request) respond.Json {
	var (
		data *define.MenuServiceGetListReq
	)
	if err := r.Parse(&data); err != nil {
		return result.Error(err)
	}

	return result.Response(service.Menu.GetList(r.Context(), data))
}

func (a *menuApi) Delete(r *ghttp.Request) respond.Json {
	var (
		data *define.MenuApiDeleteReq
	)
	if err := r.Parse(&data); err != nil {
		return result.Error(err)
	}
	if err := service.Menu.Delete(r.Context(), data.Id); err != nil {
		return result.Error(err)
	}
	return result.Success("", "删除成功")
}

func (a *menuApi) GetDetail(r *ghttp.Request) respond.Json {
	var (
		data *define.MenuApiDetailReq
	)
	if err := r.Parse(&data); err != nil {
		return result.Error(err)
	}
	return result.Response(service.Menu.GetDetail(r.Context(), data.Id))
}

func (a *menuApi) Create(r *ghttp.Request) respond.Json {
	var (
		data             *define.MenuApiDoCreateReq
		serviceCreateReq *define.MenuServiceDoCreateReq
	)
	if err := r.ParseForm(&data); err != nil {
		return result.Error(err)
	}
	if err := gconv.Struct(data, &serviceCreateReq); err != nil {
		return result.Error(err)
	}

	res, err := service.Menu.Create(r.Context(), serviceCreateReq)
	if err != nil {
		return result.Error(err)
	}
	return result.Success(res, "添加成功")
}

func (a *menuApi) Update(r *ghttp.Request) respond.Json {
	var (
		data             *define.MenuApiDoUpdateReq
		serviceUpdateReq *define.MenuServiceDoUpdateReq
	)
	if err := r.ParseForm(&data); err != nil {
		return result.Error(err)
	}
	if err := gconv.Struct(data, &serviceUpdateReq); err != nil {
		return result.Error(err)
	}
	if err := service.Menu.Update(r.Context(), serviceUpdateReq); err != nil {
		return result.Error(err)
	}
	return result.Success("", "编辑成功")
}
