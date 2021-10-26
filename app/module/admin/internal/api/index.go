package api

import (
	"github.com/gogf/gf/util/gconv"
	"glow-admin/app/model"
	"glow-admin/app/module/admin/internal/define"
	"glow-admin/app/module/admin/internal/service"
	"glow-admin/app/shared"
	"glow-admin/library/respond"
	"glow-admin/library/result"
	"glow-admin/library/tools"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Index = indexApi{}

type indexApi struct {
}

func (a *indexApi) Captcha(r *ghttp.Request) respond.Json {
	id, base := tools.GetVerifyImgString()
	return result.Success(g.Map{
		"key":     id,
		"captcha": base,
	}, "")
}

func (a *indexApi) Menus(r *ghttp.Request) respond.Json {
	return result.Response(service.Menu.GetMenuList(shared.Context.Get(r.Context()).User.Id))
}

func (a *indexApi) UserInfo(r *ghttp.Request) respond.Json {
	return result.Response(service.User.GetUserInfo(r.Context()))
}

func (a *indexApi) Login(r *ghttp.Request) respond.Json {
	var (
		apiReq *define.LoginReq
		dto    *model.LoginDTO
	)
	if err := r.Parse(&apiReq); err != nil {
		return result.Error(err)
	}
	if !tools.VerifyString(apiReq.Key, apiReq.Captcha) {
		return result.Errorf("验证码不正确")
	}
	if err := gconv.Struct(apiReq, &dto); err != nil {
		return result.Error(err)
	}
	return result.Response(service.User.Login(r.Context(), dto))
}

func (a *indexApi) Logout(r *ghttp.Request) respond.Json {
	return result.Response("退出成功", nil)
}
