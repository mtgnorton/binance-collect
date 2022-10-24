package controller

import (
	"context"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/app/system/index/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

var Personal = personal{}

type personal struct {
}

func (p *personal) Logout(ctx context.Context, req *define.AuthLogoutReq) (res *define.AuthLogoutRes, err error) {
	res = &define.AuthLogoutRes{}
	err = service.Auth.Logout(ctx)

	g.RequestFromCtx(ctx).Response.RedirectTo("/")

	return
}

func (p *personal) Info(ctx context.Context, req *define.AuthInfoReq) (res *define.AuthInfoRes, err error) {

	res = &define.AuthInfoRes{}
	res.AuthInfoOutput, err = service.Auth.Info(ctx)

	return
}
