package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/dto"
	"gf-admin/utility/response"
	"strings"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

// 中间件管理服务
var (
	Middleware = serviceMiddleware{}
)

type serviceMiddleware struct {
	IgnoreAuthUrls []string
}

func (s *serviceMiddleware) GetIgnoreAuthUrls() []string {
	if s.IgnoreAuthUrls == nil {
		s.IgnoreAuthUrls = g.Cfg().MustGet(context.TODO(), "casbin.ignoreUrls").Strings()
	}
	return s.IgnoreAuthUrls
}

func (s *serviceMiddleware) OperationLog(r *ghttp.Request) {

	r.Middleware.Next()

	if r.Method == "GET" {
		return
	}

	path := r.URL.Path
	pathName, err := dao.AdminMenu.Ctx(r.Context()).Where(g.Map{
		dao.AdminMenu.Columns.Identification: strings.Replace(path, g.Cfg().MustGet(r.Context(), "server.Prefix").String(), "", -1),
	}).Value(dao.AdminMenu.Columns.Name)
	if err != nil {
		response.JsonErrorLogExit(r, err)
	}
	params := r.GetMap()

	responseContent := r.GetHandlerResponse()

	if err != nil {
		responseContent = err.Error()
	}

	adminID, err := AdminTokenInstance.GetAdministratorId(r.Context())
	if err != nil {
		response.JsonErrorLogExit(r, err)

	}

	_, err = dao.AdminLog.Ctx(r.Context()).Insert(dto.AdminLog{
		AdministratorId: adminID,
		Path:            path,
		Method:          r.Method,
		PathName:        pathName,
		Params:          params,
		Response:        responseContent,
	})
	if err != nil {
		response.JsonErrorLogExit(r, err)
	}
}

func (s *serviceMiddleware) Auth(r *ghttp.Request) {

	g.Log("auth").Debug(r.Context(), "是否登录验证中间件开始执行")
	administrator, err := AdminTokenInstance.GetAdministrator(r.Context())
	if err != nil {
		response.JsonErrorLogExit(r, err)
	}

	if administrator.Id == 0 {
		response.JsonErrorLogExit(r, response.NewError("未登录或会话已过期，请您登录后再继续", g.Map{"administrator": administrator}))

	}

	r.Middleware.Next()
}

func (s *serviceMiddleware) Permission(r *ghttp.Request) {

	administrator, err := AdminTokenInstance.GetAdministrator(r.Context())
	if err != nil {
		response.JsonErrorLogExit(r, response.WrapError(err, "没有权限", g.Map{"administrator": administrator}))

	}

	url := r.Request.URL.Path
	method := r.Request.Method
	prefix, err := g.Cfg().Get(r.Context(), "server.Prefix")
	if err != nil {
		response.JsonErrorLogExit(r, err)
	}

	path := gstr.Replace(url, prefix.String(), "")
	g.Log("auth").Infof(r.Context(), "权限认证,用户为:%s,path为:%s,method为:%s", administrator.Username, path, method)

	if garray.NewStrArrayFrom(s.GetIgnoreAuthUrls(), true).ContainsI(path) {
		r.Middleware.Next()
		return
	}

	isAllow, err := Enforcer.Auth(administrator.Username, path, method)
	if err != nil {

		response.JsonErrorLogExit(r, response.WrapError(err, "没有权限", g.Map{"administrator": administrator, "request path": path}))
	}
	if !isAllow {
		response.JsonErrorLogExit(r, response.NewError("没有权限", g.Map{"administrator": administrator, "request path": path}))
	}
	r.Middleware.Next()
}
