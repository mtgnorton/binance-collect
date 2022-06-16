package service

import (
	"context"
	"fmt"
	"gf-admin/app/dao"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/utility/custom_error"
)

var BinanceAdmin = binanceAdmin{}

type binanceAdmin struct {
}

func (ba *binanceAdmin) UserAddressList(ctx context.Context, in *define.UserAddressListInput) (out *define.UserAddressListOutput, err error) {
	out = &define.UserAddressListOutput{}

	d := dao.UserAddresses.Ctx(ctx)

	if in.ExternalUserId != 0 {
		d = d.Where(dao.UserAddresses.Columns().ExternalUserId, in.ExternalUserId)
	}

	if in.Address != "" {
		d = d.WhereLike(dao.UserAddresses.Columns().Address, fmt.Sprintf("%%%s%%", in.Address))
	}

	out.Page = in.Page
	out.Size = in.Size
	out.Total, err = d.Count()
	if err != nil {
		return out, custom_error.New(err.Error())
	}

	d = d.Order(dao.UserAddresses.Columns().Id, "desc").Page(in.Page, in.Size)
	err = d.Scan(&out.List)

	return

}

func (ba *binanceAdmin) CollectList(ctx context.Context, in *define.CollectListInput) (out *define.CollectListOutput, err error) {
	out = &define.CollectListOutput{}

	d := dao.Collects.Ctx(ctx)

	if in.RechargeHash != "" {
		d = d.Where(dao.Collects.Columns().RechargeHash, in.RechargeHash)
	}

	if in.HandfeeHash != "" {
		d = d.Where(dao.Collects.Columns().HandfeeHash, in.HandfeeHash)
	}
	if in.CollectHash != "" {
		d = d.Where(dao.Collects.Columns().CollectHash, in.CollectHash)
	}

	if in.UserAddress != "" {
		d = d.Where(dao.Collects.Columns().UserAddress, in.UserAddress)
	}
	if in.Status != "" {
		d = d.Where(dao.Collects.Columns().Status, in.Status)
	}
	out.Page = in.Page
	out.Size = in.Size
	out.Total, err = d.Count()
	if err != nil {
		return out, custom_error.New(err.Error())
	}
	d = d.Order(dao.Collects.Columns().Id, "desc").Page(in.Page, in.Size)
	err = d.Scan(&out.List)
	return
}

func (ba *binanceAdmin) WithdrawList(ctx context.Context, in *define.WithdrawListInput) (out *define.WithdrawListOutput, err error) {
	out = &define.WithdrawListOutput{}

	d := dao.Withdraws.Ctx(ctx)

	if in.Hash != "" {
		d = d.Where(dao.Withdraws.Columns().Hash, in.Hash)
	}

	if in.To != "" {
		d = d.Where(dao.Withdraws.Columns().To, in.To)
	}

	if in.ExternalOrderId != "" {
		d = d.Where(dao.Withdraws.Columns().ExternalOrderId, in.ExternalOrderId)
	}

	if in.ExternalUserId != 0 {
		d = d.Where(dao.Withdraws.Columns().ExternalUserId, in.ExternalUserId)
	}
	if in.Status != "" {
		d = d.Where(dao.Withdraws.Columns().Status, in.Status)
	}
	out.Page = in.Page
	out.Size = in.Size
	out.Total, err = d.Count()
	if err != nil {
		return out, custom_error.New(err.Error())
	}
	d = d.Order(dao.Withdraws.Columns().Id, "desc").Page(in.Page, in.Size)
	err = d.Scan(&out.List)
	return
}
