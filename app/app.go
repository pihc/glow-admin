package app

import (
	"payget/app/module/admin"
	"payget/app/shared"

	"github.com/gogf/gf/frame/g"
)

func Run() {
	s := g.Server()
	s.Use(shared.Middleware.Cors, shared.Middleware.ErrorHandler)
	admin.Init()
	s.Run()
}
