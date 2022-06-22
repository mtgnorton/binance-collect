package controller

import (
	"context"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
)

var BinanceAdmin = binanceAdmin{}

type binanceAdmin struct {
}

func (ua *binanceAdmin) UserAddressList(ctx context.Context, req *define.UserAddressListReq) (res *define.UserAddressListRes, err error) {
	res = new(define.UserAddressListRes)
	res.UserAddressListOutput, err = service.BinanceAdmin.UserAddressList(ctx, req.UserAddressListInput)
	return
}

func (ua *binanceAdmin) CollectList(ctx context.Context, req *define.CollectListReq) (res *define.CollectListRes, err error) {
	res = new(define.CollectListRes)
	res.CollectListOutput, err = service.BinanceAdmin.CollectList(ctx, req.CollectListInput)
	return
}
func (ua *binanceAdmin) CollectUpdate(ctx context.Context, req *define.CollectUpdateReq) (res *define.CollectUpdateRes, err error) {
	res = new(define.CollectUpdateRes)
	err = service.BinanceAdmin.CollectUpdate(ctx, req.CollectUpdateInput)
	return
}

func (ua *binanceAdmin) CollectDestroy(ctx context.Context, req *define.CollectDestroyReq) (res *define.CollectDestroyRes, err error) {
	res = new(define.CollectDestroyRes)
	err = service.BinanceAdmin.CollectDestroy(ctx, req.CollectDestroyInput)
	return
}

func (ua *binanceAdmin) WithdrawList(ctx context.Context, req *define.WithdrawListReq) (res *define.WithdrawListRes, err error) {
	res = new(define.WithdrawListRes)
	res.WithdrawListOutput, err = service.BinanceAdmin.WithdrawList(ctx, req.WithdrawListInput)
	return
}
func (ba *binanceAdmin) WithdrawUpdate(ctx context.Context, req *define.WithdrawUpdateReq) (res *define.WithdrawUpdateRes, err error) {
	res = new(define.WithdrawUpdateRes)
	err = service.BinanceAdmin.WithdrawUpdate(ctx, req.WithdrawUpdateInput)
	return
}
func (ba *binanceAdmin) WithdrawDestroy(ctx context.Context, req *define.WithdrawDestroyReq) (res *define.WithdrawDestroyRes, err error) {
	res = new(define.WithdrawDestroyRes)
	err = service.BinanceAdmin.WithdrawDestroy(ctx, req.WithdrawDestroyInput)
	return
}

func (ua *binanceAdmin) QueueTaskList(ctx context.Context, req *define.QueueTaskListReq) (res *define.QueueTaskListRes, err error) {
	res = new(define.QueueTaskListRes)
	res.QueueTaskListOutput, err = service.BinanceAdmin.QueueTaskList(ctx, req.QueueTaskListInput)
	return
}
func (ua *binanceAdmin) QueueTaskUpdate(ctx context.Context, req *define.QueueTaskUpdateReq) (res *define.QueueTaskUpdateRes, err error) {
	res = new(define.QueueTaskUpdateRes)
	err = service.BinanceAdmin.QueueTaskUpdate(ctx, req.QueueTaskUpdateInput)
	return
}

func (ua *binanceAdmin) QueueTaskDestroy(ctx context.Context, req *define.QueueTaskDestroyReq) (res *define.QueueTaskDestroyRes, err error) {
	res = new(define.QueueTaskDestroyRes)
	err = service.BinanceAdmin.QueueTaskDestroy(ctx, req.QueueTaskDestroyInput)
	return
}

func (ua *binanceAdmin) NotifyList(ctx context.Context, req *define.NotifyListReq) (res *define.NotifyListRes, err error) {
	res = new(define.NotifyListRes)
	res.NotifyListOutput, err = service.BinanceAdmin.NotifyList(ctx, req.NotifyListInput)
	return
}

func (ba *binanceAdmin) NotifyUpdate(ctx context.Context, req *define.NotifyUpdateReq) (res *define.NotifyUpdateRes, err error) {
	res = new(define.NotifyUpdateRes)
	err = service.BinanceAdmin.NotifyUpdate(ctx, req.NotifyUpdateInput)
	return
}

func (ba *binanceAdmin) NotifyDestroy(ctx context.Context, req *define.NotifyDestroyReq) (res *define.NotifyDestroyRes, err error) {
	res = new(define.NotifyDestroyRes)
	err = service.BinanceAdmin.NotifyDestroy(ctx, req.NotifyDestroyInput)
	return
}

func (ua *binanceAdmin) ContractList(ctx context.Context, req *define.ContractListReq) (res *define.ContractListRes, err error) {
	res = new(define.ContractListRes)
	res.ContractListOutput, err = service.BinanceAdmin.ContractList(ctx, req.ContractListInput)
	return
}

func (ua *binanceAdmin) ContractStore(ctx context.Context, req *define.ContractStoreReq) (res *define.ContractStoreRes, err error) {
	res = new(define.ContractStoreRes)
	err = service.BinanceAdmin.ContractStore(ctx, req.ContractStoreInput)
	return
}

func (ua *binanceAdmin) ContractUpdate(ctx context.Context, req *define.ContractUpdateReq) (res *define.ContractUpdateRes, err error) {
	res = new(define.ContractUpdateRes)
	err = service.BinanceAdmin.ContractUpdate(ctx, req.ContractUpdateInput)
	return
}

func (ua *binanceAdmin) ContractInfo(ctx context.Context, req *define.ContractInfoReq) (res *define.ContractInfoRes, err error) {
	res = new(define.ContractInfoRes)
	res.ContractInfoOutput, err = service.BinanceAdmin.ContractInfo(ctx, req.ContractInfoInput)
	return
}

func (ua *binanceAdmin) ContractDestroy(ctx context.Context, req *define.ContractDestroyReq) (res *define.ContractDestroyRes, err error) {
	res = new(define.ContractDestroyRes)
	err = service.BinanceAdmin.ContractDestroy(ctx, req.ContractDestroyInput)
	return
}

func (ua *binanceAdmin) LoseBlockList(ctx context.Context, req *define.LoseBlockListReq) (res *define.LoseBlockListRes, err error) {
	res = new(define.LoseBlockListRes)
	res.LoseBlockListOutput, err = service.BinanceAdmin.LoseBlockList(ctx, req.LoseBlockListInput)
	return
}

func (ua *binanceAdmin) LoseBlockStore(ctx context.Context, req *define.LoseBlockStoreReq) (res *define.LoseBlockStoreRes, err error) {
	res = new(define.LoseBlockStoreRes)
	err = service.BinanceAdmin.LoseBlockStore(ctx, req.LoseBlockStoreInput)
	return
}

func (ua *binanceAdmin) LoseBlockDestroy(ctx context.Context, req *define.LoseBlockDestroyReq) (res *define.LoseBlockDestroyRes, err error) {
	res = new(define.LoseBlockDestroyRes)
	err = service.BinanceAdmin.LoseBlockDestroy(ctx, req.LoseBlockDestroyInput)
	return
}
