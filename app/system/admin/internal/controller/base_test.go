package controller

import (
	"context"
	"fmt"
	"gf-admin/app/shared"
	"gf-admin/app/system/admin/internal/service"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/os/gtimer"
	"time"
)

var ports = garray.NewIntArray(true)
var ctx context.Context = context.Background()

func init() {
	genv.Set("UNDER_TEST", "1")
	for i := 7000; i <= 8000; i++ {
		ports.Append(i)
	}
}

// 新建一个测试服务，参数为需要测试的控制器，返回值为服务器对应的端口
func NewTestServer(bindRouter func(group *ghttp.RouterGroup)) int {

	p, _ := ports.PopRand()
	s := g.Server(p)
	s.Use(ghttp.MiddlewareHandlerResponse)

	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config-unit.toml") //单元测试时，日志输出会作为响应返回，所以暂时建立专门的测试配置文件关闭所有log输出

	prefix, err := g.Cfg().Get(ctx, "server.prefix")
	if err != nil {
		g.Log().Fatalf(ctx, "get server admin prefix error,error info following : %s", err)
	}
	s.Group(prefix.String(), func(group *ghttp.RouterGroup) {
		group.Middleware(
			shared.Middleware.Ctx,
			service.Middleware.OperationLog,
			service.Middleware.ResponseHandler,
		)
		bindRouter(group)
	})
	g.Log().SetStdoutPrint(false)
	g.Log().SetDebug(false)

	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()

	gtimer.AddOnce(ctx, time.Second*2, func(ctx context.Context) {
		s.Shutdown()

	})

	time.Sleep(100 * time.Millisecond)

	return p
}

func NewTestClient(p int) *gclient.Client {
	client := g.Client()
	client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d/admin/api/", p))
	return client
}
