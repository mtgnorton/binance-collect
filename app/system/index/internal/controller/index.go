package controller

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/app/system/index/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

var Index = index{}

type index struct {
}

func (i *index) IndexHtml(ctx context.Context, req *define.IndexReq) (res *define.IndexRes, err error) {
	user, err := service.FrontTokenInstance.GetUser(ctx)
	if err != nil {
		return
	}
	nodes, err := service.Node.Query(ctx, true)
	if err != nil {
		return
	}

	posts, err := service.Posts.IndexList(ctx)
	if err != nil {
		return
	}

	shared.View().Render(ctx, model.View{
		Title: "首页",
		User:  user,
		Data: g.Map{
			"nodes": nodes,
			"posts": posts,
		},
	})
	return
}
