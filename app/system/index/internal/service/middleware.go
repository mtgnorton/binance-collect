package service

import (
	"gf-admin/utility/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 中间件管理服务
var (
	Middleware = middleware{}
)

type middleware struct {
}

func (s *middleware) Auth(r *ghttp.Request) {

	g.Log("auth").Debug(r.Context(), "是否登录验证中间件开始执行")
	user, err := FrontTokenInstance.GetUser(r.Context())
	g.Log("auth").Debugf(r.Context(), "登录的用户为：%#v", user)

	if err != nil {
		response.JsonErrorLogExit(r, err)
	}

	if user.Id == 0 {
		response.JsonErrorLogExit(r, response.NewError("未登录或会话已过期，请您登录后再继续", g.Map{"user": user}))

	}

	r.Middleware.Next()
}
