package define

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*创建地址*/
type CreateAddressReq struct {
	g.Meta `path:"/dw-create-address" method:"post" summary:"创建地址" tags:"币安出入金"`
	*CreateAddressInput
}
type CreateAddressRes struct {
	*CreateAddressOutput
}

type CreateAddressInput struct {
	UserId int `json:"user_id" v:"required|integer#用户ID不能为空|用户ID必须是整数"`
}
type CreateAddressOutput struct {
	Address string `json:"address"`
}

/*申请提现*/
type ApplyWithdrawReq struct {
	g.Meta `path:"/dw-apply-withdraw" method:"post" summary:"申请提现" tags:"币安出入金"`
	*ApplyWithdrawInput
}
type ApplyWithdrawRes struct {
	*ApplyWithdrawOutput
}

type ApplyWithdrawInput struct {
	ExternalOrderId string  `json:"order_id" dc:"提现唯一标志符" v:"required|length:1,64#外部订单号不能为空|外部订单号长度不能超过64"`
	ExternalUserId  int     `json:"user_id" dc:"该提现申请的用户id" v:"required|integer#外部用户ID不能为空|外部用户ID必须是整数"`
	Value           float64 `json:"value" dc:"提现金额" v:"required|float#金额不能为空|金额必须是数字"`
	UserAddress     string  `json:"user_address" dc:"创建用户时，该系统返回的地址" v:"required|length:42,42#地址不能为空|提现地址包含0x，且总长度为42"`
	To              string  `json:"to_address" dc:"用户提现到的地址" v:"required|length:42,42#提现地址不能为空|提现地址包含0x，且总长度为42"`
	Symbol          string  `json:"symbol" dc:"提现的货币，可选的货币如BNB,BSC-USD，根据实际项目决定" v:"required|length:1,32#币种不能为空|币种长度不能超过32"`
}
type ApplyWithdrawOutput struct {
}
