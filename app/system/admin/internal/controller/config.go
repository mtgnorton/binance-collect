package controller

import (
	"context"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
)

var Config = configApi{}

type configApi struct {
}

func (c *configApi) List(ctx context.Context, req *define.ConfigListReq) (resp *define.ConfigListRes, err error) {
	resp = new(define.ConfigListRes)
	resp.ConfigListOutput, err = service.Config.GetModules(ctx, req.ConfigListInput)
	return
}

func (c *configApi) Update(ctx context.Context, req *define.ConfigUpdateReq) (resp *define.ConfigUpdateRes, err error) {
	resp = new(define.ConfigUpdateRes)
	err = service.Config.Update(ctx, req.ConfigUpdateInput)
	return
}
