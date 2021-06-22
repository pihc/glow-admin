package middleware

import "github.com/gogf/gf/net/ghttp"

var (
	Admin = &middleware{}
)

type middleware struct {
}

//Auth authHook is the HOOK function implements JWT logistics.
func (mid *middleware) Auth(r *ghttp.Request) {
	Auth.MiddlewareFunc()(r)
	r.Middleware.Next()
}
