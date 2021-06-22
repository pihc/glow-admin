package admin

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"payget/app/module/admin/internal/api"
	"payget/app/module/admin/internal/middleware"
	"payget/library/respond"
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
			//group.GET("/native", respond.Convert(api.Hello.Native))
			//group.GET("/string", respond.Convert(api.Hello.String))
			//group.GET("/json", respond.Convert(api.Hello.Json))
			//group.GET("/xml", respond.Convert(api.Hello.Xml))
			//group.GET("/file", respond.Convert(api.Hello.File))
		})
	})
}
