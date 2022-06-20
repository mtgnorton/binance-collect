package deposit_withdraw

import (
	"encoding/json"
	"fmt"
	"gf-admin/app/dao"
	"gf-admin/app/dto"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"gf-admin/app/shared"
	"math/big"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/grand"
)

func getBlockInfoByLocalJson(filename string) (*OriginBlock, error) {
	blockJSON := gfile.GetContents(gfile.Pwd() + "/test_data/" + filename)

	response := new(proxyBlockWithTransactions)
	err := json.Unmarshal([]byte(blockJSON), response)

	b := response.toBlock()
	return &b, err
}

func setUserAddress(userAddress string) (func(), error) {
	_, err := gcache.Remove(ctx, model.CACHE_KEY_USER_ADDRESSES)
	if err != nil {
		return nil, err
	}
	d := dao.UserAddresses.Ctx(ctx)

	idVar, err := d.Where(dao.UserAddresses.Columns().Address, userAddress).Value(dao.UserAddresses.Columns().Id)
	if err != nil {
		return nil, err
	}

	if idVar.Int() == 0 {
		id, err := d.InsertAndGetId(dto.UserAddresses{
			Address:        userAddress,
			ExternalUserId: grand.N(10000, 100000),
			PrivateKey:     nil,
			Type:           0,
		})
		if err != nil {
			return nil, err
		}
		return func() {
			_, err = d.Delete(id)
			fmt.Printf("delete user address %s\n", userAddress)
		}, nil
	}
	return nil, nil
}

func setCollectAddress(address string) (func(), error) {

	oldCollectAddress, err := shared.Config.GetString(ctx, model.CONFIG_MODULE_BINNABCE, model.CONFIG_KEY_COLLECT_ADDRESS)
	if err != nil {
		return nil, err
	}
	err = shared.Config.Set(ctx, model.CONFIG_MODULE_BINNABCE, model.CONFIG_KEY_COLLECT_ADDRESS, address)
	if err != nil {
		return nil, err
	}
	return func() {
		_ = shared.Config.Set(ctx, model.CONFIG_MODULE_BINNABCE, model.CONFIG_KEY_COLLECT_ADDRESS, oldCollectAddress)
	}, nil
}
func setFeeWithdrawAddress(address string) (func(), error) {
	oldAddress, err := shared.Config.GetString(ctx, model.CONFIG_MODULE_BINNABCE, model.CONFIG_KEY_FEE_WITHDRAW_ADDRESS)
	if err != nil {
		return nil, err
	}
	err = shared.Config.Set(ctx, model.CONFIG_MODULE_BINNABCE, model.CONFIG_KEY_FEE_WITHDRAW_ADDRESS, address)
	if err != nil {
		return nil, err
	}
	return func() {
		_ = shared.Config.Set(ctx, model.CONFIG_MODULE_BINNABCE, model.CONFIG_KEY_FEE_WITHDRAW_ADDRESS, oldAddress)
	}, nil
}

func generateWithdraw(hash string) (func(), error) {
	d := dao.Withdraws.Ctx(ctx)

	idVar, err := d.Where(dao.Withdraws.Columns().Hash, hash).Value(dao.Withdraws.Columns().Id)
	if err != nil {
		return nil, err
	}

	if idVar.Int() == 0 {
		id, err := d.InsertAndGetId(dto.Withdraws{
			Id:              1,
			UserId:          1,
			ExternalOrderId: 1,
			ExternalUserId:  1,
			Hash:            hash,
			Symbol:          1,
			From:            1,
			To:              1,
			Value:           1,
			Status:          1,
		})
		if err != nil {
			return nil, err
		}
		return func() {
			_, err = d.Delete(id)
		}, nil
	}
	return nil, nil
}

func TestTransactionSimpleProcessor_DistinguishAndParse(t *testing.T) {

	tp := NewTransactionProcessor(ctx, model.PROCESSOR_SIMPLE)

	gtest.C(t, func(t *gtest.T) {

		// test recharge transaction
		userAddress := "0x808da0ceea1c25ba09dcba49f9fc604e9ff166d8"

		gc, err := setUserAddress(userAddress)
		t.AssertNil(err)
		if gc != nil {

			defer gc()
		}

		b, err := getBlockInfoByLocalJson("block_15913847.json")
		t.AssertNil(err)
		transactions, err := tp.DistinguishAndParse(ctx, b)

		t.AssertNil(err)
		t.Assert(len(transactions), 1)
		tx := *transactions[0]

		t.Assert(tx.Type, model.TRANSACTION_TYPE_REACHRGE)
		t.Assert(tx.Value.Int64(), 400000000000000000)

		valueDecimal, _ := tx.ValueDecimal.Float64()
		t.Assert(valueDecimal, 0.4)

	})

	// test collect transaction
	gtest.C(t, func(t *gtest.T) {
		userAddress := "0x59c4358b6e4b4bbc6a24fb4883cabec9f957227d"

		collectAddress := "0xf75043489fddd995fc4b0e8d9b002be1696f46d8"

		gc, err := setCollectAddress(collectAddress)
		t.AssertNil(err)
		if gc != nil {
			defer gc()
		}
		gc, err = setUserAddress(userAddress)
		t.AssertNil(err)
		if gc != nil {
			defer gc()
		}
		b, err := getBlockInfoByLocalJson("block_15913847.json")
		t.AssertNil(err)
		transactions, err := tp.DistinguishAndParse(ctx, b)

		t.AssertNil(err)

		tx := *transactions[0]
		t.Assert(tx.Type, model.TRANSACTION_TYPE_COLLECT)

	})
	// test fee transaction
	gtest.C(t, func(t *gtest.T) {
		gc, err := setFeeWithdrawAddress("0x3bbe68d6c846fa419054fbbb7e3abdb67a5bef73")
		t.AssertNil(err)
		if gc != nil {
			defer gc()
		}
		gc, err = setUserAddress("0x45b2884a0e496211878e38ab49a5e8a2dca80f33")
		t.AssertNil(err)
		if gc != nil {
			defer gc()
		}
		b, err := getBlockInfoByLocalJson("block_15913847.json")
		t.AssertNil(err)
		transactions, err := tp.DistinguishAndParse(ctx, b)
		t.AssertNil(err)
		t.Assert(len(transactions), 1)
		tx := *transactions[0]
		g.DumpWithType(tx)
		t.Assert(tx.Type, model.TRANSACTION_TYPE_FEE)

	})
	// test withdraw transaction
	gtest.C(t, func(t *gtest.T) {
		//清空user地址缓存
		gc, err := setUserAddress("")
		t.AssertNil(err)
		if gc != nil {
			defer gc()
		}
		gc, err = setFeeWithdrawAddress("0x1480401a41d458119c96f3fac234c0594b14f1a1")
		t.AssertNil(err)
		if gc != nil {
			defer gc()
		}
		gc, err = generateWithdraw("0x2f5a0e443812e057ea5d3a645527a4777dd172f938c12b2f435b75c946af0b12")

		t.AssertNil(err)
		if gc != nil {
			defer gc()
		}
		b, err := getBlockInfoByLocalJson("block_15913847.json")
		t.AssertNil(err)
		transactions, err := tp.DistinguishAndParse(ctx, b)
		t.AssertNil(err)

		g.DumpWithType(transactions)
		tx := *transactions[0]
		t.Assert(tx.Type, model.TRANSACTION_TYPE_WITHDRAW)

	})

}

var ethLowerTx *Transaction = &Transaction{
	Symbol:          "BNB",
	ContractAddress: "0x",
	Hash:            "0x5263f2bdaa2e7c31bd3233c48e35c79343a9cb5e2cce65e1dc31937947deb10a",
	From:            "0x991195b40a5bdf4725afbd4f10f579bca25308f5",
	UserID:          1,
	To:              "0x81023633832221b512018a21f8a3c6a6fe774913",
	Value:           *big.NewInt(10000000000001),
}
var ethTx *Transaction = &Transaction{
	Symbol:          "BNB",
	ContractAddress: "0x",
	Hash:            "0x5263f2bdaa2e7c31bd3233c48e35c79343a9cb5e2cce65e1dc31937947deb10a",
	From:            "0x991195b40a5bdf4725afbd4f10f579bca25308f5",
	UserID:          1,
	To:              "0x81023633832221b512018a21f8a3c6a6fe774913",
	Value:           *big.NewInt(1000000000000000),
}

//var erc20Tx *Transaction = &Transaction{
//	Symbol: "MTG-USD",
//	Hash:   "0xaf91e8c5c507bca0bee49256aff374c162c96e6f966101aa225149f0c9fc09b1",
//	From:   "0x991195b40a5bdf4725afbd4f10f579bca25308f5",
//	UserID: 1,
//	To:     "0x4538453345313141456631414333353934313435",
//	Value:  *big.NewInt(1000000),
//}

func Test_Processor_Scanner_Queue(t *testing.T) {
	// 测试eth 充值->转出手续费->归集
	gtest.C(t, func(t *gtest.T) {

		collectModel := dao.Collects.Ctx(ctx)
		condition := g.Map{
			dao.Collects.Columns().RechargeHash: ethLowerTx.Hash,
		}
		processor := NewTransactionProcessor(ctx, model.PROCESSOR_SIMPLE)
		scanner := NewTransactionScanner(ctx, 5)

		// 测试充值金额过小
		//充值检测
		_, err := collectModel.Delete(condition)
		t.AssertNil(err)
		err = processor.HandleRecharge(ctx, ethLowerTx)
		t.AssertNil(err)
		var collect entity.Collects
		err = collectModel.Where(condition).Scan(&collect)
		t.AssertNil(err)

		t.AssertGT(collect.Id, 0)
		t.Assert(collect.Status, model.COLLECT_STATUS_WAIT_FEE)

		// 充值扫描
		tasks, err := scanner.scanTxFee(ctx)
		t.AssertNil(err)
		t.Assert(len(tasks), 0)

		err = collectModel.Where(condition).Scan(&collect)
		t.AssertNil(err)
		t.Assert(collect.Status, model.COLLECT_STATUS_FAIl_TOO_LOW_AMOUNT)

		// 测试正常充值
		//充值检测
		_, err = collectModel.Delete(condition)
		t.AssertNil(err)
		err = processor.HandleRecharge(ctx, ethTx)
		t.AssertNil(err)
		err = collectModel.Where(condition).Scan(&collect)
		t.AssertNil(err)

		t.AssertGT(collect.Id, 0)
		t.Assert(collect.Status, model.COLLECT_STATUS_WAIT_FEE)
		t.AssertNil(err)

		// 充值扫描
		tasks, err = scanner.scanTxFee(ctx)
		t.AssertNil(err)
		t.Assert(len(tasks), 0)

		err = collectModel.Where(condition).Scan(&collect)
		t.AssertNil(err)
		t.Assert(collect.Status, model.COLLECT_STATUS_WAIT_COLLECT)

		// 归集扫描
		tasks, err = scanner.scanTxCollect(ctx)
		if err != nil {
			LogErrorfDw(ctx, err)
		}
		t.AssertNil(err)
		t.Assert(len(tasks), 1)
		err = collectModel.Where(condition).Scan(&collect)
		t.AssertNil(err)
		t.Assert(collect.Status, model.COLLECT_STATUS_PROCESS_COLLECT)

		//测试交易队列归集交易
		queue := NewTransactionTransfer()

		_, err = dao.QueueTask.Ctx(ctx).Where(dao.QueueTask.Columns().From, ethTx.To).Delete()

		t.AssertNil(err)

		queue.Consume(ctx, tasks[0])

		var databaseTask *entity.QueueTask
		err = dao.QueueTask.Ctx(ctx).Where(dao.QueueTask.Columns().From, ethTx.To).Scan(&databaseTask)
		t.AssertNil(err)
		t.Assert(databaseTask.Status, model.QUEUE_TASK_STATUS_PROCESS)

		//检测归集交易
		collectTx := &Transaction{
			Symbol:          "BNB",
			ContractAddress: "0x",
			Hash:            databaseTask.Hash,
			From:            databaseTask.From,
			UserID:          1,
			To:              databaseTask.To,
		}
		err = processor.HandleCollect(ctx, collectTx)
		t.AssertNil(err)
		// 队列中的任务状态应该为 model.QUEUE_TASK_STATUS_SUCCESS
		err = dao.QueueTask.Ctx(ctx).Where(dao.QueueTask.Columns().From, ethTx.To).Scan(&databaseTask)
		t.AssertNil(err)
		t.Assert(databaseTask.Status, model.QUEUE_TASK_STATUS_SUCCESS)

	})
}

func TestTransactionSimpleProcessor_Quo(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// big.Int 类型做除法会省略小数,所以使用big.Float
		n1 := big.NewInt(10)
		n2 := big.NewInt(2)
		r1 := n1.Quo(n1, n2)
		fmt.Println(r1.String())

		//

	})
}
