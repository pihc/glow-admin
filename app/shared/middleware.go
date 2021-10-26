package shared

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
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

// 全局统一错误处理
func (mid *middleware) ErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		// server错误日志会自动记录
		r.Response.ClearBuffer()
		_ = r.Response.WriteJson(g.Map{
			"code": http.StatusInternalServerError,
			"data": nil,
			"msg":  "哎哟我去，服务器居然开小差了，请稍后再试吧！",
		})
		r.Response.WriteHeader(http.StatusOK)
	}
}
