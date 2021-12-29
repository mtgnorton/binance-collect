package cmd

import (
	"context"
	"gf-admin/app/system/admin/internal/api"
	"gf-admin/app/system/admin/internal/service"
	"gf-admin/utility/response"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gsession"
	"github.com/gogf/gf/v2/util/gmode"
)

func Run(ctx context.Context) {
	var (
		s   = g.Server()
		oai = s.GetOpenApi()
	)

	// OpenApi自定义信息
	oai.Info.Title = `API Reference`
	oai.Config.CommonResponse = response.JsonRes{}
	oai.Config.CommonResponseDataField = `Data`

	// 静态目录设置
	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	if uploadPath == "" {
		g.Log().Fatal(ctx, "文件上传配置路径不能为空")
	}
	if !gfile.IsDir(uploadPath) {
		gfile.Mkdir(uploadPath)
	}
	s.AddStaticPath("/upload", uploadPath)

	// HOOK, 开发阶段禁止浏览器缓存,方便调试
	if gmode.IsDevelop() {
		s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
			r.Response.Header().Set("Cache-Control", "no-store")
		})
	}

	// 前台系统路由注册
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware.Ctx,
			service.Middleware.ResponseHandler,
		)
		group.Bind(
			api.Administrator,
		)
		// 权限控制路由
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(service.Middleware.Auth)

		})
	})
	// 自定义丰富文档
	enhanceOpenAPIDoc(s)
	sessionConfig(s)
	// 启动Http Server
	s.Run()
	return

}
func sessionConfig(s *ghttp.Server) {

	s.SetConfigWithMap(g.Map{
		"SessionStorage": gsession.NewStorageRedis(g.Redis("session")),
	})
}
func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info.Title = `gf-admin`
	openapi.Info.Description = `后台接口文档`
}
