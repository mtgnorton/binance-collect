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

func (n *node) Store(ctx context.Context, req *define.NodeStoreReq) (res *define.NodeStoreRes, err error) {
	err = service.Node.Store(ctx, req.NodeStoreInput)
	return
}

func (n *node) Update(ctx context.Context, req *define.NodeUpdateReq) (res *define.NodeUpdateRes, err error) {
	err = service.Node.Update(ctx, req.NodeUpdateInput)
	return
}

func (n *node) Destroy(ctx context.Context, req *define.NodeDestroyReq) (res *define.NodeDestroyRes, err error) {
	err = service.Node.Destroy(ctx, req.Id)
	return
}

func (n *node) Info(ctx context.Context, req *define.NodeInfoReq) (res *define.NodeInfoRes, err error) {
	res = &define.NodeInfoRes{}
	res.NodeInfoOutput, err = service.Node.Info(ctx, req.Id)
	return
}
