package deposit_withdraw

import (
	"context"
	"gf-admin/app/model"
	"time"

	"github.com/gogf/gf/v2/os/gtimer"
)

type Manager struct {
	transactionProcessor      TransactionProcessorAbstract
	transactionScanner        TransactionScannerAbstract
	transactionTransfer       *TransactionTransfer
	transactionNotifier       *TransactionNotifier
	detectIntervalMillisecond time.Duration //检测时间间隔,默认1s检测一次
}

func NewManager(ctx context.Context, options ...func(manager *Manager)) *Manager {
	var manager = &Manager{
		transactionProcessor:      NewTransactionProcessor(ctx, model.PROCESSOR_SIMPLE),
		transactionScanner:        NewTransactionScanner(ctx, 5),
		transactionTransfer:       NewTransactionTransfer(),
		transactionNotifier:       NewTransactionNotifier(),
		detectIntervalMillisecond: time.Duration(1000) * time.Millisecond,
	}
	for _, option := range options {
		option(manager)
	}

	return manager
}

func (m *Manager) Run(ctx context.Context) {

	m.run(ctx)
	gtimer.AddSingleton(ctx, m.detectIntervalMillisecond, func(ctx context.Context) {
		m.run(ctx)
	})

	m.transactionTransfer.Run(ctx)
	m.transactionNotifier.Run(ctx)
	m.transactionScanner.Scan(ctx, m.transactionTransfer.Ch, m.transactionNotifier.Ch)
}

// 检测是否有新的交易
// 将新的交易封装成transaction
// 根据交易类型进行不同的处理
func (m *Manager) run(ctx context.Context) {

	blockInfo, err := m.Detect(ctx)
	if err != nil {
		logErrorfDw(ctx, err)
	}
	// 说明此时没有新的区块产生
	if blockInfo == nil {
		return
	}
	go func() {
		transactions, err := m.transactionProcessor.DistinguishAndParse(ctx, blockInfo)
		if err != nil {
			logErrorfDw(ctx, err)
		}
		m.Dispatch(ctx, transactions)
	}()

}

// Dispatch 交易调度
func (m *Manager) Dispatch(ctx context.Context, transactions []*Transaction) {
	if len(transactions) == 0 {
		return
	}
	var err error
	for _, transaction := range transactions {
		switch transaction.Type {
		case model.TRANSACTION_TYPE_REACHRGE:
			err = m.transactionProcessor.HandleRecharge(ctx, transaction)

		case model.TRANSACTION_TYPE_FEE:
			err = m.transactionProcessor.HandleFee(ctx, transaction)
		case model.TRANSACTION_TYPE_COLLECT:
			err = m.transactionProcessor.HandleCollect(ctx, transaction)
		case model.TRANSACTION_TYPE_WITHDRAW:
			err = m.transactionProcessor.HandleWithdraw(ctx, transaction)
		}
		if err != nil {
			logErrorfDw(ctx, err)
		}
	}
}

// Detect 获取检测区块
func (m *Manager) Detect(ctx context.Context) (*OriginBlock, error) {

	newestNumber, detectNumber, err := ChainClient.GetDetectNumber(ctx)
	if err != nil {
		return nil, err
	}
	logInfofDw(ctx, "block detect ,newest block is %d,detect block is %d", newestNumber, detectNumber)
	if detectNumber == 0 {
		return nil, err
	}
	blockInfo, err := ChainClient.GetBlockInfoByNumber(ctx, detectNumber)

	return blockInfo, err
}
