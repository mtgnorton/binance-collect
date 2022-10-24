package index

import (
	"context"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/controller"
	"gf-admin/app/system/index/internal/service"
	"gf-admin/utility/response"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gsession"

	"github.com/gogf/gf/v2/net/goai"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
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
	uploadPath := g.Cfg().MustGet(ctx, "upload.FrontPath").String()
	if uploadPath == "" {
		g.Log().Fatal(ctx, "文件上传配置路径不能为空")
	}
	if !gfile.IsDir(uploadPath) {
		err := gfile.Mkdir(uploadPath)
		if err != nil {
			return
		}
	}
	s.AddStaticPath("/upload", uploadPath)

	// HOOK, 开发阶段禁止浏览器缓存,方便调试
	if gmode.IsDevelop() {
		s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
			r.Response.Header().Set("Cache-Control", "no-store")
		})
	}

	prefix, err := g.Cfg().Get(ctx, "front.prefix")
	if err != nil {
		g.Log().Fatalf(ctx, "get server front prefix error,error info following : %s", err)
	}

	service.FrontTokenInstance.Init(ctx)

	s.Group(prefix.String(), func(group *ghttp.RouterGroup) {
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(
				shared.Middleware.Ctx,
				shared.Middleware.ResponseHandler,
				shared.Middleware.TokenInCookieToHeader,
			)

			//通过header将对应用户注入到上下文对象中

			service.FrontTokenInstance.LoadConfig().Middleware(group)
			//无需验证登录的路由
			group.Bind(controller.NoAuth)

			group.Bind(controller.Index)

			group.Group("/", func(group *ghttp.RouterGroup) {

				group.Middleware(service.Middleware.Auth)

				// 需要登录验证的路由
				group.Bind(controller.Personal)

				group.Bind(controller.Posts)

				group.Bind(controller.Member)

			})

		})
	})
	enhanceOpenAPIDoc(s)
	sessionConfig(s)

	s.SetPort(g.Cfg().MustGet(ctx, "front.Port").Int())
	s.SetServerRoot(g.Cfg().MustGet(ctx, "front.ServerRoot").String())
	s.Run()
}

func sessionConfig(s *ghttp.Server) {

	err := s.SetConfigWithMap(g.Map{
		"SessionStorage": gsession.NewStorageRedis(g.Redis("session")),
	})
	if err != nil {
		g.Log().Fatalf(gctx.New(), "init session driver error, %s", err)
	}
}

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	openapi.Components = goai.Components{
		SecuritySchemes: goai.SecuritySchemes{
			"ApiKeyAuth": goai.SecuritySchemeRef{
				Ref: "", // 暂时还不知道该值是干什么用的
				Value: &goai.SecurityScheme{
					Type: "apiKey",
					In:   "header",
					Name: "Authorization",
				},
			},
		},
	}
	openapi.Security = &goai.SecurityRequirements{
		goai.SecurityRequirement{"ApiKeyAuth": []string{}},
	}
	// API description.
	openapi.Info.Title = `gf-admin`
	openapi.Info.Description = `前台接口文档`
}
