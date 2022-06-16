package define

import (
	"gf-admin/app/model"
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type UserAddressListReq struct {
	g.Meta `path:"/binance-user-address-list" method:"get" summary:"用户地址列表" tags:"币安后台接口"`
	*UserAddressListInput
}

type UserAddressListRes struct {
	*UserAddressListOutput
}

type UserAddressListInput struct {
	Address        string `dc:"地址查询"`
	ExternalUserId int    `dc:"外部用户id查询"`
	model.PageSizeInput
}

type UserAddressListOutput struct {
	List  []*model.FrontUserAddress `json:"list"`
	Page  int                       `json:"page"`
	Size  int                       `json:"size"`
	Total int                       `json:"total"`
}

type CollectListReq struct {
	g.Meta `path:"/binance-collect-list" method:"get" summary:"归集列表" tags:"币安后台接口"`
	*CollectListInput
}

type CollectListRes struct {
	*CollectListOutput
}

type CollectListInput struct {
	RechargeHash string `dc:"充值hash查询" json:"recharge_hash"`
	HandfeeHash  string `dc:"手续费hash查询" json:"handfee_hash"`
	CollectHash  string `dc:"归集hash查询" json:"collect_hash"`
	UserAddress  string `dc:"用户地址查询" json:"user_address"`
	Status       string `dc:"状态查询" json:"status"`
	model.PageSizeInput
}

type CollectListOutput struct {
	List  []*entity.Collects `json:"list"`
	Page  int                `json:"page"`
	Size  int                `json:"size"`
	Total int                `json:"total"`
}

type WithdrawListReq struct {
	g.Meta `path:"/binance-withdraw-list" method:"get" summary:"提现列表" tags:"币安后台接口"`
	*WithdrawListInput
}

type WithdrawListRes struct {
	*WithdrawListOutput
}
type WithdrawListInput struct {
	To              string `dc:"提现地址查询" json:"to"`
	ExternalOrderId string `dc:"外部订单id查询" json:"external_order_id"`
	ExternalUserId  int    `dc:"外部用户id查询" json:"external_user_id"`
	Status          string `dc:"状态查询" json:"status"`
	Hash            string `dc:"hash查询" json:"hash"`
	model.PageSizeInput
}

type WithdrawListOutput struct {
	List  []*entity.Withdraws `json:"list"`
	Page  int                 `json:"page"`
	Size  int                 `json:"size"`
	Total int                 `json:"total"`
}
