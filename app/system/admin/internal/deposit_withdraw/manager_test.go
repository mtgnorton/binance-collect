package deposit_withdraw

import (
	"testing"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/test/gtest"
)

var manager *Manager

func InitManager() {

	manager = NewManager(ctx)

	_, _ = g.DB().Exec(ctx, "truncate table collects;")
	_, _ = g.DB().Exec(ctx, "truncate table queue_task;")
	_, _ = g.DB().Exec(ctx, "truncate table queue_task_log;")
	_, _ = g.DB().Exec(ctx, "truncate table withdraws;")
	_, _ = g.DB().Exec(ctx, "truncate table notify;")
	_, _ = g.DB().Exec(ctx, "truncate table notify_log;")
}

func Test_Manager_Detect(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		InitManager()
		manager.Run(ctx)
	})
}

// 测试某一个区块
func Test_Manager_Single_Block(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		InitManager()

		m := manager
		detectNumber := 20148458
		blockInfo, err := ChainClient.GetBlockInfoByNumber(ctx, detectNumber)
		t.AssertNil(err)
		transactions, err := m.transactionProcessor.DistinguishAndParse(ctx, blockInfo)

		t.AssertNil(err)

		m.Dispatch(ctx, transactions)

	})
}
