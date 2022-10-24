package shared

import (
	"bytes"
	"fmt"
	"gf-admin/app/model"
	"gf-admin/utility/response"
	"reflect"
	"strings"

	"github.com/gogf/gf/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gutil"
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

func (s *middleware) TokenInCookieToHeader(r *ghttp.Request) {
	v := r.Cookie.Get("gf-token")
	g.Dump("gf-token", v.String())
	if v.String() != "" {
		r.Header.Set("Authorization", "Bearer "+v.String())
	}
	r.Middleware.Next()
}

// 返回处理中间件
func (s *middleware) ResponseHandler(r *ghttp.Request) {

	buffers := bytes.NewBuffer([]byte(""))
	g.DumpTo(buffers, r.GetMap(), gutil.DumpOption{})

	g.Log().Infof(r.Context(), "请求的url为：%s,客户端端传递过来的参数如下", r.URL.Path)
	g.Log().Infof(r.Context(), "%s", buffers)

	r.Middleware.Next()

	//系统运行时错误
	if err := r.GetError(); err != nil {
		r.Response.Status = 200
		r.Response.ClearBuffer()
		fmt.Println(333)

		response.JsonErrorLogExit(r, err)
	}

	//如果已经有返回内容，那么该中间件什么也不做
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		res  interface{}
		code gcode.Code = gcode.CodeOK
	)

	res = r.GetHandlerResponse()

	if msg == "" {
		if strings.Contains(r.URL.Path, "-update") {
			msg = "更新成功"
		} else if strings.Contains(r.URL.Path, "-delete") {
			msg = "删除成功"
		} else if strings.Contains(r.URL.Path, "-store") {
			msg = "保存成功"
		} else if strings.Contains(r.URL.Path, "-info") || strings.Contains(r.URL.Path, "-list") {
			msg = "获取成功"
		}
	}

	if res == nil || reflect.ValueOf(res).IsNil() {
		response.JsonExit(r, code.Code(), msg, g.Map{})
	}

	response.JsonExit(r, code.Code(), msg, res)

}
