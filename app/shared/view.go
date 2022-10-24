package shared

import (
	"context"
	"gf-admin/app/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
)

type sView struct{}

var insView = sView{}

// 视图管理服务
func View() *sView {
	return &insView
}

// 渲染指定模板页面
func (s *sView) RenderTpl(ctx context.Context, tpl string, data ...model.View) {
	var (
		viewObj  = model.View{}
		viewData = make(g.Map)
		request  = g.RequestFromCtx(ctx)
	)
	if len(data) > 0 {
		viewObj = data[0]
	}
	if viewObj.Title == "" {
		viewObj.Title = g.Cfg().MustGet(ctx, `front.title`).String()
	} else {
		viewObj.Title = viewObj.Title + ` - ` + g.Cfg().MustGet(ctx, `front.title`).String()
	}
	if viewObj.Keywords == "" {
		viewObj.Keywords = g.Cfg().MustGet(ctx, `front.keywords`).String()
	}
	if viewObj.Description == "" {
		viewObj.Description = g.Cfg().MustGet(ctx, `front.description`).String()
	}
	// 去掉空数据
	viewData = gconv.Map(viewObj)
	for k, v := range viewData {
		if g.IsEmpty(v) {
			delete(viewData, k)
		}
	}
	// 内置对象
	//viewData["BuildIn"] = &viewBuildIn{httpRequest: request}
	// 内容模板
	if viewData["MainTpl"] == nil {
		viewData["MainTpl"] = s.getDefaultMainTpl(ctx)
		// 如果mainTpl以/开头，则表示是绝对路径，不需要使用layout布局文件
	} else if gstr.HasPrefix(gconv.String(viewData["MainTpl"]), "/") {
		tpl = gstr.TrimLeft(gconv.String(viewData["MainTpl"]), "/")
	}
	// 提示信息
	//if notice, _ := Session().GetNotice(ctx); notice != nil {
	//	_ = Session().RemoveNotice(ctx)
	//	viewData["Notice"] = notice
	//}
	// 渲染模板
	_ = request.Response.WriteTpl(tpl, viewData)
	// 开发模式下，在页面最下面打印所有的模板变量
	//fmt.Printf("%#v", viewData)

	if gmode.IsDevelop() {
		g.Dump("viewData", viewData)
		_ = request.Response.WriteTplContent(`{{dump .}}`, viewData)
	}
}

// 渲染默认模板页面
func (s *sView) Render(ctx context.Context, data ...model.View) {
	s.RenderTpl(ctx, g.Cfg().MustGet(ctx, "viewer.indexLayout").String(), data...)
}

// 跳转中间页面
func (s *sView) Render302(ctx context.Context, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "页面跳转中"
	}
	view.MainTpl = s.getViewFolderName(ctx) + "/pages/302.html"
	s.Render(ctx, view)
}

// 401页面
func (s *sView) Render401(ctx context.Context, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "无访问权限"
	}
	view.MainTpl = s.getViewFolderName(ctx) + "/pages/401.html"
	s.Render(ctx, view)
}

// 403页面
func (s *sView) Render403(ctx context.Context, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "无访问权限"
	}
	view.MainTpl = s.getViewFolderName(ctx) + "/pages/403.html"
	s.Render(ctx, view)
}

// 404页面
func (s *sView) Render404(ctx context.Context, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "资源不存在"
	}
	view.MainTpl = s.getViewFolderName(ctx) + "/pages/404.html"
	s.Render(ctx, view)
}

// 500页面
func (s *sView) Render500(ctx context.Context, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "请求执行错误"
	}
	view.MainTpl = s.getViewFolderName(ctx) + "/pages/500.html"
	s.Render(ctx, view)
}

// 获取视图存储目录
func (s *sView) getViewFolderName(ctx context.Context) string {
	return gstr.Split(g.Cfg().MustGet(ctx, "viewer.indexLayout").String(), "/")[0]
}

// 获取自动设置的MainTpl
func (s *sView) getDefaultMainTpl(ctx context.Context) string {
	var (
		urlPathArray = gstr.SplitAndTrim(g.RequestFromCtx(ctx).URL.Path, "/")
		mainTpl      string
	)
	// 根据path获取模板文件名
	if len(urlPathArray) > 0 && gstr.Contains(urlPathArray[len(urlPathArray)-1], "-html") {
		mainTpl = gstr.Replace(urlPathArray[len(urlPathArray)-1], "-html", "") + ".html"
	}
	// 根据path获取模版文件路径
	if len(urlPathArray) > 1 && mainTpl != "" {
		paths := urlPathArray[0 : len(urlPathArray)-1]
		mainTpl = gstr.Join(paths, "/") + "/" + mainTpl
	}

	if mainTpl == "" {
		mainTpl = "index.html"
	}

	return mainTpl
}
