package admin

import (
	"payget/app/module/admin/internal/api"
	"payget/app/module/admin/internal/middleware"
	"payget/library/respond"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Init() {
	s := g.Server()
	// 前台系统路由注册
	s.Group("/admin", func(group *ghttp.RouterGroup) {
		group.GET("/captcha", respond.Convert(api.Index.Captcha))
		//group.POST("/login", respond.Convert(api.Index.Login))
		group.POST("/login", middleware.Auth.LoginHandler)
		group.POST("/refresh_token", middleware.Auth.RefreshHandler)
		group.POST("/logout", middleware.Auth.LogoutHandler)

		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.Admin.Auth)
			group.GET("/menus", respond.Convert(api.Index.Menus))
			group.GET("/user_info", respond.Convert(api.Index.UserInfo))

			// 后台用户
			group.GET("/user/index", respond.Convert(api.User.GetList))
			group.POST("/user/edit", respond.Convert(api.User.Update))
			group.POST("/user/add", respond.Convert(api.User.Create))
			group.POST("/user/delete/:id", respond.Convert(api.User.Delete))
			group.POST("/user/change_status", respond.Convert(api.User.ChangeStatus))
			group.POST("/user/reset_pwd/:id", respond.Convert(api.User.ResetPassword))

			// 角色
			group.GET("/role/index", respond.Convert(api.Role.GetList))
			group.POST("/role/edit", respond.Convert(api.Role.Update))
			group.POST("/role/add", respond.Convert(api.Role.Create))
			group.POST("/role/delete/:id", respond.Convert(api.Role.Delete))
			group.GET("/role/menus/:id", respond.Convert(api.Role.GetMenus))
			group.POST("/role/menus", respond.Convert(api.Role.SetMenus))
			group.GET("/role/all", respond.Convert(api.Role.All))

			//group.GET("/native", respond.Convert(api.Hello.Native))
			//group.GET("/string", respond.Convert(api.Hello.String))
			//group.GET("/json", respond.Convert(api.Hello.Json))
			//group.GET("/xml", respond.Convert(api.Hello.Xml))
			//group.GET("/file", respond.Convert(api.Hello.File))
		})
	})
}
