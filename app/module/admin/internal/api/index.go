package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"payget/app/module/admin/internal/service"
	"payget/app/shared"
	"payget/library/respond"
	"payget/library/result"
	"payget/library/tools"
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
