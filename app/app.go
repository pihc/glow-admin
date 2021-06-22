package app

import (
	"github.com/gogf/gf/frame/g"
	"payget/app/module/admin"
	"payget/app/shared"
)

func Run() {
	s := g.Server()
	s.Use(shared.Middleware.Cors)
	admin.Init()
	s.Run()
}
