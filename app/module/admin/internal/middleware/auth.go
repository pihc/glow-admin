package middleware

import (
	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"payget/app/model"
	"payget/app/module/admin/internal/define"
	"payget/app/module/admin/internal/service"
	"payget/app/shared"
	"payget/library/result"
	"time"
)

var (
	// The underlying JWT middleware.
	Auth *jwt.GfJWTMiddleware
)

// Initialization function,
// rewrite this function to customized your own JWT settings.
func init() {
	authMiddleware, err := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "clean code",
		Key:             []byte("secret key"),
		Timeout:         time.Hour * 30,
		MaxRefresh:      time.Hour * 30,
		IdentityKey:     "uid",
		TokenLookup:     "header: Authorization",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		LoginResponse:   LoginResponse,
		RefreshResponse: RefreshResponse,
		LogoutResponse:  LogoutResponse,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
	})
	if err != nil {
		glog.Fatal("JWT Error:" + err.Error())
	}
	Auth = authMiddleware
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler get the identity from JWT and set the identity for every request
// Using this function, by r.GetParam("id") get identity
func IdentityHandler(r *ghttp.Request) interface{} {
	claims := jwt.ExtractClaims(r)
	customCtx := &model.Context{
		Data: make(g.Map),
	}
	shared.Context.Init(r, customCtx)
	customCtx.User = &model.ContextUser{Id: gconv.Uint(claims[Auth.IdentityKey])}
	return claims[Auth.IdentityKey]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(r *ghttp.Request, code int, message string) {
	r.Response.WriteJson(result.ErrorCode(code, "未登录或会话已过期，请您登录后再继续"))
	r.ExitAll()
}

// LoginResponse is used to define customized login-successful callback function.
func LoginResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	r.Response.WriteJson(result.Success(g.Map{
		"access_token": token,
	}, ""))

	r.ExitAll()
}

// RefreshResponse is used to get a new token no matter current token is expired or not.
func RefreshResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	r.Response.WriteJson(g.Map{
		"token":      token,
		"token_type": Auth.TokenHeadName,
		//"expire": expire.Format(time.RFC3339),
	})
	r.ExitAll()
}

// LogoutResponse is used to set token blacklist.
func LogoutResponse(r *ghttp.Request, code int) {
	r.Response.WriteJson(g.Map{
		"code": code,
		"msg":  "success",
	})
	r.ExitAll()
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// if your identityKey is 'id', your user data must have 'id'
// Check error (e) to determine the appropriate error message.
func Authenticator(r *ghttp.Request) (interface{}, error) {

	var (
		apiReq *define.LoginReq
		dto    model.LoginDTO
	)

	if err := r.Parse(&apiReq); err != nil {
		return "", err
	}
	//if !tools.VerifyString(apiReq.Key, apiReq.Captcha) {
	//	return "", gerror.New("验证码不正确")
	//}
	if err := gconv.Struct(apiReq, &dto); err != nil {
		return "", err
	}
	user, err := service.User.Login(r.Context(), dto)
	if err != nil {
		return nil, err
	}

	//if user := service.User.GetUserByUsernamePassword(serviceReq); user != nil {
	//	return user, nil
	//}

	return g.Map{
		"uid": user.Id,
	}, nil
}
