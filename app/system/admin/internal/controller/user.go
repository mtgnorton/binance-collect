package controller

import (
	"context"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
)

var User = user{}

type user struct {
}

func (u *user) List(ctx context.Context, req *define.UserListReq) (res *define.UserListRes, err error) {
	res = &define.UserListRes{}
	res.UserListOutput, err = service.User.List(ctx, req.UserListInput)
	return
}

func (u *user) ToggleDestroy(ctx context.Context, req *define.UserToggleDestroyReq) (res *define.UserToggleDestroyRes, err error) {
	err = service.User.ToggleDestroy(ctx, req.UserToggleDestroyInput)
	return
}
