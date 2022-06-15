package deposit_withdraw

import (
	"context"
)

type TransactionProcessorAbstract interface {
	DistinguishAndParse(context.Context, *OriginBlock) ([]*Transaction, error) //判断交易类型
	HandleCollect(context.Context, *Transaction) error                         // 处理归集
	HandleWithdraw(context.Context, *Transaction) error                        // 处理提现
	HandleRecharge(context.Context, *Transaction) error                        // 处理充值
	HandleFee(context.Context, *Transaction) error                             // 处理手续费
}
