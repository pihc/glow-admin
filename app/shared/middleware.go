package shared

import (
	"github.com/gogf/gf/net/ghttp"
)

var (
	Middleware = middleware{}
)

type middleware struct {
}

func (mid *middleware) Cors(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
