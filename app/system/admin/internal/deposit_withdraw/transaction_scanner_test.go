package deposit_withdraw

import (
	"gf-admin/app/dao"
	"gf-admin/app/dto"
	"gf-admin/app/model"
	"testing"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/util/grand"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestTransactionScanner_scanFailQueueTask(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		scanner := NewTransactionScanner(ctx, 5)
		_, err := scanner.scanFailTxTransferTask(ctx)
		t.AssertNil(err)
	})
}

func TestTransactionScanner_scanWithdraw(t *testing.T) {

	thirdAddress := "0x991195b40a5bDF4725AfbD4f10F579BCa25308F5"

	gtest.C(t, func(t *gtest.T) {
		//feeAddress, err := ChainClient.GetFeeWithdrawAddress(ctx)

		_, err := dao.Withdraws.Ctx(ctx).OmitEmptyData().Insert(dto.Withdraws{
			UserId:          2,
			ExternalOrderId: grand.N(1, 100000),
			ExternalUserId:  grand.N(1, 100000),
			Symbol:          model.CONTRACT_DEFAULT_SYMBOL,
			ContractAddress: model.CONTRACT_DEAULT_SYMBOL_ADDRESS,
			To:              thirdAddress,
			Value:           grand.N(10000000000, 1000000000000),
			Status:          model.WITHDRAW_STATUS_WAIT,
		})

		t.AssertNil(err)
		scanner := NewTransactionScanner(ctx, 5)
		tasks, err := scanner.scanTxWithdraw(ctx)
		t.AssertNil(err)
		t.AssertGT(len(tasks), 0)
	})
}

func TestTransactionScanner_scanNotifyWithdraw(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		thirdAddress := "0x991195b40a5bDF4725AfbD4f10F579BCa25308F5"

		_, err := dao.Withdraws.Ctx(ctx).OmitEmptyData().Insert(dto.Withdraws{
			UserId:          2,
			ExternalOrderId: grand.N(1, 100000),
			ExternalUserId:  grand.N(1, 100000),
			Symbol:          model.CONTRACT_DEFAULT_SYMBOL,
			ContractAddress: model.CONTRACT_DEAULT_SYMBOL_ADDRESS,
			To:              thirdAddress,
			Value:           grand.N(10000000000, 1000000000000),
			Status:          model.WITHDRAW_STATUS_WAIT_NOTIFY,
		})
		t.AssertNil(err)
		scanner := NewTransactionScanner(ctx, 5)
		tasks, err := scanner.scanNotifyWithdraw(ctx)
		t.AssertNil(err)
		t.AssertGT(len(tasks), 0)
	})
}

func TestTransactionScanner_scanNotifyCollect(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		scanner := NewTransactionScanner(ctx, 5)
		tasks, err := scanner.scanNotifyCollect(ctx)
		g.Dump(*tasks[0])
		t.AssertNil(err)
		t.AssertGT(len(tasks), 0)
	})
}

func TestTransactionScanner_scanFailNotify(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		scanner := NewTransactionScanner(ctx, 5)
		tasks, err := scanner.scanFailNotify(ctx)
		t.AssertNil(err)
		t.AssertGT(len(tasks), 0)
		g.Dump(*tasks[0])
	})
}
func TestTransactionScanner_scanFailTxTransferTask(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		scanner := NewTransactionScanner(ctx, 5)
		tasks, err := scanner.scanFailTxTransferTask(ctx)
		t.AssertNil(err)
		t.AssertGT(len(tasks), 0)
		g.Dump(*tasks[0])
	})
}
