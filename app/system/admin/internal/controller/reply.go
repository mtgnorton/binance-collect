package controller

import (
	"context"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
)

var Reply = reply{}

type reply struct {
}

func (r *reply) List(ctx context.Context, req *define.ReplyListReq) (res *define.ReplyListRes, err error) {
	res = &define.ReplyListRes{}
	res.ReplyListOutput, err = service.Reply.List(ctx, req.ReplyListInput)
	return
}

func (r *reply) ToggleDestroy(ctx context.Context, req *define.ReplyToggleDestroyReq) (res *define.ReplyToggleDestroyRes, err error) {
	err = service.Reply.ToggleDestroy(ctx, req.ReplyToggleDestroyInput)
	return
}
