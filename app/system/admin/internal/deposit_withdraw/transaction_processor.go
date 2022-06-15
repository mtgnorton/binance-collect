package deposit_withdraw

import (
	"context"
	"gf-admin/app/model"
)

// 初始化一个交易处理器
func NewTransactionProcessor(ctx context.Context, processorType string, options ...func(processor TransactionProcessorAbstract)) (tp TransactionProcessorAbstract) {

	switch processorType {
	case model.PROCESSOR_SIMPLE:
		tp = &TransactionSimpleProcessor{}
	}
	for _, option := range options {
		option(tp)
	}

	return tp
}
