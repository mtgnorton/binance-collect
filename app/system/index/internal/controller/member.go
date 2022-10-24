package controller

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/app/system/index/internal/service"
)

var Member = member{}

type member struct {
}

func (m *member) Index(ctx context.Context, req *define.MemberIndexReq) (res *define.MemberIndexRes, err error) {
	user, err := service.FrontTokenInstance.GetUser(ctx)
	if err != nil {
		return
	}
	shared.View().Render(ctx, model.View{
		Title: "个人中心",
		User:  user,
	})
	return
}
