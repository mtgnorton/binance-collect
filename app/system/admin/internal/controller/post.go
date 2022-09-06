package controller

import (
	"context"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
)

var Post = post{}

type post struct {
}

func (p *post) List(ctx context.Context, req *define.PostListReq) (res *define.PostListRes, err error) {
	res = &define.PostListRes{}
	res.PostListOutput, err = service.Post.List(ctx, req.PostListInput)
	return
}

func (p *post) ToggleTop(ctx context.Context, req *define.PostToggleTopReq) (res *define.PostToggleTopRes, err error) {
	err = service.Post.ToggleTop(ctx, req.PostToggleTopInput)
	return
}

func (p *post) ToggleDelete(ctx context.Context, req *define.PostToggleDestroyReq) (res *define.PostToggleDestroyRes, err error) {
	err = service.Post.ToggleDestroy(ctx, req.PostToggleDestroyInput)
	return
}
