package shared

import (
	"gf-admin/app/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var Middleware = middleware{}

type middleware struct {
}

// 自定义上下文对象
func (s *middleware) Ctx(r *ghttp.Request) {
	g.Log().Debug(r.Context(), "ctx中间件开始执行")
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Data: make(g.Map),
	}
	Context.Init(r, customCtx)
	// 执行下一步请求逻辑
	r.Middleware.Next()
}
