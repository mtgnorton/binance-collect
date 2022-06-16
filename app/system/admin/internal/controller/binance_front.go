package controller

import (
	"context"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
)

var BinanceApi = binanceApi{}

type binanceApi struct{}

func (b *binanceApi) CreateAddress(ctx context.Context, req *define.CreateAddressReq) (res *define.CreateAddressRes, err error) {
	res = &define.CreateAddressRes{}
	res.CreateAddressOutput, err = service.BinanceService.CreateAddress(ctx, req.CreateAddressInput)
	return
}

func (b *binanceApi) ApplyWithdraw(ctx context.Context, req *define.ApplyWithdrawReq) (res *define.ApplyWithdrawRes, err error) {
	res = &define.ApplyWithdrawRes{}
	res.ApplyWithdrawOutput, err = service.BinanceService.ApplyWithdraw(ctx, req.ApplyWithdrawInput)
	return

}
