package service

import (
	"gf-admin/utility/response"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gvalid"
	"reflect"
)

// 中间件管理服务
var (
	Middleware = serviceMiddleware{
		IgnoreAuthUrls: g.Cfg().MustGet(gctx.New(), "casbin.ignoreUrls").Strings(),
	}
)

type serviceMiddleware struct {
	IgnoreAuthUrls []string
}

// 返回处理中间件
func (s *serviceMiddleware) ResponseHandler(r *ghttp.Request) {

	r.Middleware.Next()
	g.Log().Debug(r.Context(), "响应中间件开始执行")
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
	//g.Dump(res, err)
	if err != nil {

		code = gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}

		if v, ok := err.(gvalid.Error); ok {
			msg = v.FirstError().Error()
		} else {
			msg = err.Error()
		}
		if msg == "" {
			msg = code.Message()
		}
	}

	if res == nil || reflect.ValueOf(res).IsNil() {
		response.JsonExit(r, code.Code(), msg, g.Map{})

	}

	//else {
	//	resJson, err := gjson.Encode(res)
	//	if err == nil {
	//		match, err := gregex.Match(`(?i)"message":"(.+)"`, resJson)
	//		if err == nil && len(match) > 1{
	//			msg = string(match[1])
	//		}
	//	}
	//}
	response.JsonExit(r, code.Code(), msg, res)

}

func (s *serviceMiddleware) Auth(r *ghttp.Request) {

	g.Log("auth").Debug(r.Context(), "是否登录验证中间件开始执行")
	administrator, err := AdminTokenInstance.GetAdministrator(r.Context())
	if err != nil {
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), err.Error())
	}

	if administrator.Id == 0 {
		// 根据当前请求方式执行不同的返回数据结构
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), "未登录或会话已过期，请您登录后再继续")
	}

	r.Middleware.Next()
}

func (s *serviceMiddleware) Permission(r *ghttp.Request) {

	administrator, err := AdminTokenInstance.GetAdministrator(r.Context())
	if err != nil {
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), err.Error())
	}

	url := r.Request.URL.Path
	method := r.Request.Method
	prefix, err := g.Cfg().Get(r.Context(), "server.prefix")
	if err != nil {
		g.Log().Fatalf(r.Context(), "get server admin prefix error,error info following : %s", err)
	}

	path := gstr.Replace(url, prefix.String(), "")
	g.Log("auth").Debugf(r.Context(), "权限认证,用户为:%s,path为:%s,method为:%s", administrator.Username, path, method)

	//fmt.Println(s.IgnoreAuthUrls)
	//if garray.NewStrArrayFrom(s.IgnoreAuthUrls, true).ContainsI(path) {
	//	r.Middleware.Next()
	//	return
	//}

	isAllow, err := Enforcer.Auth(administrator.Username, path, method)
	if err != nil || !isAllow {
		g.Dump(administrator.Username, path, method)
		g.Dump(err, isAllow)
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), "没有权限")
	}
	r.Middleware.Next()
}
