package controller

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/app/system/index/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

var Posts = posts{}

type posts struct {
}

func (p *posts) NewPostHtml(ctx context.Context, req *define.PostsNewHtmlReq) (res *define.PostsNewHtmlRes, err error) {
	user, err := service.FrontTokenInstance.GetUser(ctx)
	if err != nil {
		return
	}
	nodes, err := service.Node.Query(ctx, true)
	shared.View().Render(ctx, model.View{
		Title: "创建新主题",
		User:  user,
		Data: g.Map{
			"nodes": nodes,
		},
	})
	return
}

func (p *posts) Store(ctx context.Context, req *define.PostsStoreReq) (res *define.PostsStoreRes, err error) {
	return service.Posts.Store(ctx, req)
}

func (p *posts) Reply(ctx context.Context, req *define.ReplyStoreReq) (res *define.ReplyStoreRes, err error) {
	return service.Reply.Store(ctx, req)
}
