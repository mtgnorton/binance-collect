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

func (ua *binanceAdmin) WithdrawList(ctx context.Context, req *define.WithdrawListReq) (res *define.WithdrawListRes, err error) {
	res = new(define.WithdrawListRes)
	res.WithdrawListOutput, err = service.BinanceAdmin.WithdrawList(ctx, req.WithdrawListInput)
	return
}
