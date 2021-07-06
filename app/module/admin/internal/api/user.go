package api

import (
	"payget/app/module/admin/internal/define"
	"payget/app/module/admin/internal/service"
	"payget/library/respond"
	"payget/library/result"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var User = userApi{}

type userApi struct {
}

func (a *userApi) GetList(r *ghttp.Request) respond.Json {
	var (
		data *define.UserServiceGetListReq
	)
	if err := r.Parse(&data); err != nil {
		result.Error(err)
	}
	return result.Response(service.User.GetList(r.Context(), data))
}

func (a *userApi) Create(r *ghttp.Request) respond.Json {
	var (
		data             *define.UserApiDoCreateReq
		serviceCreateReq *define.UserServiceDoCreateReq
	)
	if err := r.ParseForm(&data); err != nil {
		return result.Error(err)
	}
	if err := gconv.Struct(data, &serviceCreateReq); err != nil {
		return result.Error(err)
	}

	res, err := service.User.Create(r.Context(), serviceCreateReq)
	if err != nil {
		return result.Error(err)
	}
	return result.Success(res, "")
}

func (a *userApi) Update(r *ghttp.Request) respond.Json {
	var (
		data             *define.UserApiDoUpdateReq
		serviceUpdateReq *define.UserServiceDoUpdateReq
	)
	if err := r.ParseForm(&data); err != nil {
		return result.Error(err)
	}
	if err := gconv.Struct(data, &serviceUpdateReq); err != nil {
		return result.Error(err)
	}
	if err := service.User.Update(r.Context(), serviceUpdateReq); err != nil {
		return result.Error(err)
	}
	return result.Success("", "编辑成功")
}

func (a *userApi) Delete(r *ghttp.Request) respond.Json {
	var (
		data *define.UserApiDeleteReq
	)
	if err := r.Parse(&data); err != nil {
		result.Error(err)
	}
	if err := service.User.Delete(r.Context(), data.Id); err != nil {
		return result.Error(err)
	}
	return result.Success("", "删除成功")
}

func (a *userApi) ChangeStatus(r *ghttp.Request) respond.Json {
	var (
		data             *define.UserApiChangeStatusReq
		serviceUpdateReq *define.UserServiceChangeStatusReq
	)
	if err := r.ParseForm(&data); err != nil {
		return result.Error(err)
	}
	if err := gconv.Struct(data, &serviceUpdateReq); err != nil {
		return result.Error(err)
	}
	if err := service.User.ChangeStatus(r.Context(), serviceUpdateReq); err != nil {
		return result.Error(err)
	}
	return result.Success("", "编辑成功")
}

func (a *userApi) ResetPassword(r *ghttp.Request) respond.Json {
	var (
		data *define.UserApiResetPwdReq
	)
	if err := r.Parse(&data); err != nil {
		result.Error(err)
	}
	if err := service.User.ResetPassword(r.Context(), data.Id); err != nil {
		return result.Error(err)
	}
	return result.Success("", "重置成功")
}

func (a *userApi) ChangePwd(r *ghttp.Request) respond.Json {
	var (
		data             *define.UserApiChangePwdReq
		serviceUpdateReq *define.UserServiceChangePwdReq
	)
	if err := r.Parse(&data); err != nil {
		result.Error(err)
	}
	if err := gconv.Struct(data, &serviceUpdateReq); err != nil {
		return result.Error(err)
	}
	if err := service.User.ChangePwd(r.Context(), serviceUpdateReq); err != nil {
		return result.Error(err)
	}
	return result.Success("", "修改成功")
}
