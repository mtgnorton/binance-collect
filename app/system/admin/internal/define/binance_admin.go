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

type CollectUpdateReq struct {
	g.Meta `path:"/binance-collect-update" method:"put" summary:"归集更新" tags:"币安后台接口"`
	*CollectUpdateInput
}
type CollectUpdateRes struct {
}

type CollectUpdateInput struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type CollectDestroyReq struct {
	g.Meta `path:"/binance-collect-destroy" method:"delete" summary:"归集删除" tags:"币安后台接口"`
	*CollectDestroyInput
}
type CollectDestroyRes struct {
}

type CollectDestroyInput struct {
	Id int `json:"id"`
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
	UserAddress     string `dc:"提现地址查询" json:"user_address"`
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

type WithdrawUpdateReq struct {
	g.Meta `path:"/binance-withdraw-update" method:"put" summary:"提现更新" tags:"币安后台接口"`
	*WithdrawUpdateInput
}
type WithdrawUpdateRes struct {
}
type WithdrawUpdateInput struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type WithdrawDestroyReq struct {
	g.Meta `path:"/binance-withdraw-destroy" method:"delete" summary:"提现删除" tags:"币安后台接口"`
	*WithdrawDestroyInput
}

type WithdrawDestroyRes struct {
}
type WithdrawDestroyInput struct {
	Id int `json:"id"`
}

type QueueTaskListReq struct {
	g.Meta `path:"/binance-queue-task-list" method:"get" summary:"队列任务列表" tags:"币安后台接口"`
	*QueueTaskListInput
}

type QueueTaskListRes struct {
	*QueueTaskListOutput
}

type QueueTaskListInput struct {
	Hash   string `dc:"hash查询" json:"hash"`
	From   string `dc:"开始时间查询" json:"from"`
	To     string `dc:"结束时间查询" json:"to"`
	Status string `dc:"状态查询" json:"status"`
	model.PageSizeInput
}

type QueueTaskListOutput struct {
	List  []*model.FrontTaskAndLog `json:"list"`
	Page  int                      `json:"page"`
	Size  int                      `json:"size"`
	Total int                      `json:"total"`
}

type QueueTaskUpdateReq struct {
	g.Meta `path:"/binance-queue-task-update" method:"put" summary:"队列任务更新" tags:"币安后台接口"`
	*QueueTaskUpdateInput
}
type QueueTaskUpdateRes struct {
}
type QueueTaskUpdateInput struct {
	Id         int    `dc:"队列任务id" json:"id"`
	Status     string `dc:"状态" json:"status"`
	FailAmount int    `dc:"失败次数" json:"fail_amount"`
}

type QueueTaskDestroyReq struct {
	g.Meta `path:"/binance-queue-task-destroy" method:"delete" summary:"队列任务删除" tags:"币安后台接口"`
	*QueueTaskDestroyInput
}
type QueueTaskDestroyRes struct {
}

type QueueTaskDestroyInput struct {
	Id int `dc:"队列任务id" json:"id"`
}

type NotifyListReq struct {
	g.Meta `path:"/binance-notify-list" method:"get" summary:"通知列表" tags:"币安后台接口"`
	*NotifyListInput
}

type NotifyListRes struct {
	*NotifyListOutput
}

type NotifyListOutput struct {
	List  []*model.FrontNotify `json:"list"`
	Page  int                  `json:"page"`
	Size  int                  `json:"size"`
	Total int                  `json:"total"`
}

type NotifyListInput struct {
	UniqueId string `dc:"唯一id查询" json:"unique_id"`
	model.PageSizeInput
}

type NotifyUpdateReq struct {
	g.Meta `path:"/binance-notify-update" method:"put" summary:"通知更新" tags:"币安后台接口"`
	*NotifyUpdateInput
}
type NotifyUpdateRes struct {
}
type NotifyUpdateInput struct {
	Id                 int    `dc:"通知id" json:"id"`
	Status             string `dc:"状态" json:"status"`
	IsImmediatelyRetry bool   `dc:"是否立即重试" json:"is_immediately_retry"`
}

type NotifyDestroyReq struct {
	g.Meta `path:"/binance-notify-destroy" method:"delete" summary:"通知删除" tags:"币安后台接口"`
	*NotifyDestroyInput
}
type NotifyDestroyRes struct {
}
type NotifyDestroyInput struct {
	Id int `dc:"通知id" json:"id"`
}

type ContractListReq struct {
	g.Meta `path:"/binance-contract-list" method:"get" summary:"合约列表" tags:"币安后台接口"`
	*ContractListInput
}

type ContractListRes struct {
	*ContractListOutput
}

type ContractListOutput struct {
	List  []*entity.Contracts `json:"list"`
	Page  int                 `json:"page"`
	Size  int                 `json:"size"`
	Total int                 `json:"total"`
}

type ContractListInput struct {
	model.PageSizeInput
}

type ContractStoreReq struct {
	g.Meta `path:"/binance-contract-store" method:"post" summary:"合约存储" tags:"币安后台接口"`
	*ContractStoreInput
}
type ContractStoreRes struct {
}

type ContractStoreInput struct {
	Symbol        string `json:"symbol "      v:"required#货币必填"      `              // 货币类型
	Address       string `json:"address"   v:"required#地址必填"    `                   // 合约地址
	Decimals      int    `json:"decimals"  v:"required|integer#小数位数不能为空|小数位数必须是整数"` // 小数位数
	IsCollectOpen int    `json:"is_collect_open" `                                  // 是否开启,1是 0否
}

type ContractUpdateReq struct {
	g.Meta `path:"/binance-contract-update" method:"put" summary:"合约更新" tags:"币安后台接口"`
	*ContractUpdateInput
}
type ContractUpdateRes struct {
}
type ContractUpdateInput struct {
	Id            int    `json:"id" v:"required|integer#id不能为空|id必须是整数"`            // id
	Symbol        string `json:"symbol "      v:"required#货币必填"      `              // 货币类型
	Address       string `json:"address"   v:"required#地址必填"    `                   // 合约地址
	Decimals      int    `json:"decimals"  v:"required|integer#小数位数不能为空|小数位数必须是整数"` // 小数位数
	IsCollectOpen int    `json:"is_collect_open" `                                  // 是否开启,1是 0否
}

type ContractInfoReq struct {
	g.Meta `path:"/binance-contract-info" method:"get" summary:"合约信息" tags:"币安后台接口"`
	*ContractInfoInput
}
type ContractInfoRes struct {
	*ContractInfoOutput
}
type ContractInfoOutput struct {
	Contract *entity.Contracts `json:"contract"`
}

type ContractInfoInput struct {
	Id uint `json:"id" v:"required|integer#id不能为空|id必须是整数"`
}

type ContractDestroyReq struct {
	g.Meta `path:"/binance-contract-destroy" method:"delete" summary:"合约删除" tags:"币安后台接口"`
	*ContractDestroyInput
}
type ContractDestroyRes struct {
}
type ContractDestroyInput struct {
	Id uint `json:"id" v:"required|integer#id不能为空|id必须是整数"`
}

type LoseBlockListReq struct {
	g.Meta `path:"/binance-lose-block-list" method:"get" summary:"丢失区块列表" tags:"币安后台接口"`
	*LoseBlockListInput
}

type LoseBlockListRes struct {
	*LoseBlockListOutput
}
type LoseBlockListOutput struct {
	List  []*entity.LoseBlocks `json:"list"`
	Page  int                  `json:"page"`
	Size  int                  `json:"size"`
	Total int                  `json:"total"`
}

type LoseBlockListInput struct {
	model.PageSizeInput
}

type LoseBlockStoreReq struct {
	g.Meta `path:"/binance-lose-block-store" method:"post" summary:"丢失区块存储" tags:"币安后台接口"`
	*LoseBlockStoreInput
}

type LoseBlockStoreRes struct {
}
type LoseBlockStoreInput struct {
	Number int `json:"number" v:"required|integer#区块号不能为空|区块号必须是整数"` // 区块号
}

type LoseBlockDestroyReq struct {
	g.Meta `path:"/binance-lose-block-destroy" method:"delete" summary:"丢失区块删除" tags:"币安后台接口"`
	*LoseBlockDestroyInput
}

type LoseBlockDestroyRes struct {
}

type LoseBlockDestroyInput struct {
	Id uint `json:"id" v:"required|integer#id不能为空|id必须是整数"`
}
