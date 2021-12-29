package service

import (
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/utility/response"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 中间件管理服务
var (
	Middleware = serviceMiddleware{}
)

type serviceMiddleware struct {
}

// 返回处理中间件
func (s *serviceMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// 如果已经有返回内容，那么该中间件什么也不做
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		err  error
		msg  string
		res  interface{}
		code gcode.Code = gcode.CodeOK
	)
	res, err = r.GetHandlerResponse()
	if err != nil {

		code = gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}

		msg = err.Error()
		if msg == "" {
			msg = code.Message()
		}

	}

	response.JsonExit(r, code.Code(), msg, res)

}

// 自定义上下文对象
func (s *serviceMiddleware) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &define.Context{
		Session: r.Session,
		Data:    make(g.Map),
	}
	Context.Init(r, customCtx)
	if administrator := Session.GetAdministrator(r.Context()); administrator.Id > 0 {
		customCtx.Administrator = &define.ContextAdministrator{
			Id:       administrator.Id,
			Username: administrator.Username,
			Nickname: administrator.Nickname,
			Avatar:   administrator.Avatar,
		}
	}

	// 执行下一步请求逻辑
	r.Middleware.Next()
}

// 前台系统权限控制，用户必须登录才能访问
func (s *serviceMiddleware) Auth(r *ghttp.Request) {
	administrator := Session.GetAdministrator(r.Context())
	if administrator.Id == 0 {
		// 根据当前请求方式执行不同的返回数据结构
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), "未登录或会话已过期，请您登录后再继续")
	}
	r.Middleware.Next()
}
