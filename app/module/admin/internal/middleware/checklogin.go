package middleware

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"glow-admin/app/model"
	"glow-admin/app/shared"
	"glow-admin/library/token"
	"net/http"
)

// 依赖auth
func CheckLogin(r *ghttp.Request) {
	t := r.GetHeader("Authorization")
	if t == "" {
		if t = r.GetString("token"); t == "" {
			r.Response.WriteHeader(http.StatusOK)
			_ = r.Response.WriteJsonExit(g.Map{
				"code": 401,
				"data": nil,
				"msg":  "参数缺失",
			})
		}
	}
	t = gstr.Replace(t, "Bearer ", "")
	data, err := token.Parse(t)
	if err != nil {
		r.Response.WriteHeader(http.StatusOK)
		_ = r.Response.WriteJsonExit(g.Map{
			"code": 401,
			"data": nil,
			"msg":  "Token失效",
		})
	}

	// 写入变量到上下文中
	customCtx := &model.Context{
		Data: make(g.Map),
		User: &data.Data,
	}
	shared.Context.Init(r, customCtx)

	// 生成新的token
	expire := gconv.Int(g.Config().Get("jwt.expire"))
	newT, _ := token.Generate(data.Data, expire)
	r.Response.Writer.Header().Set("Access-Control-Expose-Headers", "Authorization")
	r.Response.Writer.Header().Set("Authorization", "Bearer "+newT)

	r.Middleware.Next()
}
