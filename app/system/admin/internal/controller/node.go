package controller

import (
	"context"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
)

var Node = node{}

type node struct {
}

func (n *node) List(ctx context.Context, req *define.NodeListReq) (res *define.NodeListRes, err error) {
	res = &define.NodeListRes{}
	res.NodeListOutput, err = service.Node.List(ctx, &req.NodeListInput)
	return
}

func (n *node) Store(ctx context.Context, req *define.NodeStoreReq) (err error) {
	err = service.Node.Store(ctx, req.NodeStoreInput)
	return
}

func (n *node) Update(ctx context.Context, req *define.NodeUpdateReq) (err error) {
	err = service.Node.Update(ctx, req.NodeUpdateInput)
	return
}

func (n *node) Destroy(ctx context.Context, req *define.NodeDestroyReq) (err error) {
	err = service.Node.Destroy(ctx, req.Id)
	return
}
