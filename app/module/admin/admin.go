package admin

import (
	"glow-admin/app/module/admin/internal/api"
	"glow-admin/app/module/admin/internal/middleware"
	"glow-admin/library/respond"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Init() {
	s := g.Server()
	// 前台系统路由注册
	s.Group("/admin", func(group *ghttp.RouterGroup) {
		group.GET("/captcha", respond.Convert(api.Index.Captcha))
		group.POST("/login", respond.Convert(api.Index.Login))
		group.POST("/logout", respond.Convert(api.Index.Logout))
		group.GET("/test", respond.Convert(api.User.Test))

		group.Middleware(middleware.CheckLogin).Group("/", func(group *ghttp.RouterGroup) {
			group.GET("/menus", respond.Convert(api.Index.Menus))
			group.GET("/user_info", respond.Convert(api.Index.UserInfo))

			// 后台用户
			group.GET("/user/index", respond.Convert(api.User.GetList))
			group.POST("/user/edit", respond.Convert(api.User.Update))
			group.POST("/user/add", respond.Convert(api.User.Create))
			group.POST("/user/delete/:id", respond.Convert(api.User.Delete))
			group.POST("/user/change_status", respond.Convert(api.User.ChangeStatus))
			group.POST("/user/reset_pwd/:id", respond.Convert(api.User.ResetPassword))
			group.POST("/user/change_pwd", respond.Convert(api.User.ChangePwd))

			// 角色
			group.GET("/role/index", respond.Convert(api.Role.GetList))
			group.POST("/role/edit", respond.Convert(api.Role.Update))
			group.POST("/role/add", respond.Convert(api.Role.Create))
			group.POST("/role/delete/:id", respond.Convert(api.Role.Delete))
			group.GET("/role/menus/:id", respond.Convert(api.Role.GetMenus))
			group.POST("/role/menus", respond.Convert(api.Role.SetMenus))
			group.GET("/role/all", respond.Convert(api.Role.All))

			// 菜单
			group.GET("/menu/index", respond.Convert(api.Menu.GetList))
			group.POST("/menu/delete/:id", respond.Convert(api.Menu.Delete))
			group.GET("/menu/info/:id", respond.Convert(api.Menu.GetDetail))
			group.POST("/menu/add", respond.Convert(api.Menu.Create))
			group.POST("/menu/edit", respond.Convert(api.Menu.Update))

			//group.GET("/native", respond.Convert(api.Hello.Native))
			//group.GET("/string", respond.Convert(api.Hello.String))
			//group.GET("/json", respond.Convert(api.Hello.Json))
			//group.GET("/xml", respond.Convert(api.Hello.Xml))
			//group.GET("/file", respond.Convert(api.Hello.File))
		})
	})
}
