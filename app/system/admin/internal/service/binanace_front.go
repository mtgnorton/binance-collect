package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/dto"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/deposit_withdraw"
	"gf-admin/utility/custom_error"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/os/gcache"

	"github.com/gogf/gf/v2/text/gstr"

	"github.com/gogf/gf/v2/frame/g"
)

var BinanceService = binanceService{}

type binanceService struct {
}

func (b *binanceService) CreateAddress(ctx context.Context, in *define.CreateAddressInput) (out *define.CreateAddressOutput, err error) {
	out = &define.CreateAddressOutput{}

	// 如果用户已经存在,直接返回数据库里的地址
	addressVar, err := dao.UserAddresses.Ctx(ctx).Where(g.Map{
		dao.UserAddresses.Columns().ExternalUserId: in.UserId,
	}).Value(dao.UserAddresses.Columns().Address)
	if addressVar.String() != "" {
		out.Address = addressVar.String()
		return
	}

	//生成新的地址
	var privateKey string
	out.Address, privateKey, err = deposit_withdraw.CreateAddress()
	if err != nil {
		return
	}
	_, err = dao.UserAddresses.Ctx(ctx).Insert(dto.UserAddresses{
		ExternalUserId: in.UserId,
		Type:           model.USER_ADDRESS_TYPE_GENERATE,
		Address:        out.Address,
		PrivateKey:     privateKey,
	})
	return
}

func (b *binanceService) ApplyWithdraw(ctx context.Context, in *define.ApplyWithdrawInput) (out *define.ApplyWithdrawOutput, err error) {
	out = &define.ApplyWithdrawOutput{}

	// 判断 提现金额
	if in.Value <= 0 {
		return out, custom_error.New("提现金额必须大于0")
	}
	// 根据 in.ExternalOrderId 判断提现申请是否存在，存在直接返回错误
	idVar, err := dao.Withdraws.Ctx(ctx).Where(g.Map{
		dao.Withdraws.Columns().ExternalOrderId: in.ExternalOrderId,
	}).Value(dao.Withdraws.Columns().Id)
	if idVar.Int() > 0 {
		return out, custom_error.New("提现订单号申请已存在")
	}

	in.Symbol = gstr.ToUpper(in.Symbol)
	// 判断提现货币是否存在
	contract := &entity.Contracts{}
	err = dao.Contracts.Ctx(ctx).Where(dao.Contracts.Columns().Symbol, in.Symbol).Scan(&contract)
	if err != nil || contract.Id == 0 {
		return out, custom_error.New("提现货币查询失败")
	}

	// 判断用户地址是否存在
	idVar, err = dao.UserAddresses.Ctx(ctx).Where(dao.UserAddresses.Columns().Address, in.UserAddress).Value(dao.UserAddresses.Columns().Id)
	if idVar.Int() == 0 {
		return out, custom_error.New("用户地址不存在")
	}

	// 传递过来的value单位是ether,转为 wei
	valueWei, err := deposit_withdraw.ChainClient.EtherToWei(ctx, gconv.String(in.Value), in.Symbol)

	if err != nil {
		return out, custom_error.New("金额转换失败")
	}

	_, err = gcache.Remove(ctx, model.CACHE_KEY_USER_ADDRESSES)
	if err != nil {
		return out, err
	}

	_, err = dao.Withdraws.Ctx(ctx).OmitEmptyData().Insert(dto.Withdraws{
		ExternalOrderId: in.ExternalOrderId,
		ExternalUserId:  in.ExternalUserId,
		Symbol:          in.Symbol,
		Value:           valueWei,
		To:              in.To,
		UserId:          idVar.Int(),
		UserAddress:     in.UserAddress,
		ContractAddress: contract.Address,
		Status:          model.WITHDRAW_STATUS_WAIT,
	})

	return out, err
}
