package deposit_withdraw

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/dto"
	"gf-admin/app/model"
	"gf-admin/utility/custom_error"
	"math/big"

	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/v2/frame/g"
)

type TransactionSimpleProcessor struct {
}

// 返回站内交易的集合,如果区块中的某个交易解析失败，则直接返回空交易集合,等待后续由管理员或技术人员处理
func (tp *TransactionSimpleProcessor) DistinguishAndParse(ctx context.Context, blockInfo *OriginBlock) ([]*Transaction, error) {

	transactions := make([]*Transaction, 0)
	for _, originTransaction := range blockInfo.Transactions {
		transaction, err := NewTransaction(ctx, &originTransaction)
		isInterior, err := transaction.IsInterior(ctx)

		if err != nil {
			return nil, err
		}
		if isInterior {
			_, err = transaction.SetType(ctx)
			if err != nil {
				return nil, err
			}
			LogInfofDw(ctx, "interior tx %#v", transaction)

			if err != nil {
				return nil, custom_error.Wrap(err, "解析交易失败", g.Map{
					"transaction": originTransaction,
				})
			}
			transactions = append(transactions, transaction)

		}
	}
	return transactions, nil

}

// HandleRecharge 将检测到的充值记录写入到 collects 表中，状态为 model.COLLECT_STATUS_WAIT_FEE
func (tp *TransactionSimpleProcessor) HandleRecharge(ctx context.Context, transaction *Transaction) error {
	LogInfofDw(ctx, "new recharge transaction:%#v", transaction)
	// 判断hash是否重复
	d := dao.Collects.Ctx(ctx)

	idVar, err := d.Where(dao.Collects.Columns().RechargeHash, transaction.Hash).Value(dao.Collects.Columns().Id)
	if idVar.Int() > 0 {
		return custom_error.New("recharge hash already exists", g.Map{
			"transaction": transaction,
		})
	}
	//value为实际归集金额， 实际归集金额应该等于充值金额-手续费,该操作在扫描归集时进行
	_, err = d.OmitEmptyData().Insert(dto.Collects{
		Symbol:          transaction.Symbol,
		RechargeHash:    transaction.Hash,
		Status:          model.COLLECT_STATUS_WAIT_FEE,
		RechargeValue:   transaction.Value.String(),
		UserId:          transaction.UserID,
		UserAddress:     transaction.To,
		ContractAddress: transaction.ContractAddress,
	})
	return err
}

// HandleFee 根据hash更新collects表为 model.COLLECT_STATUS_WAIT_COLLECT  ，queue_task表的状态为 model.QUEUE_TASK_STATUS_SUCCESS
func (tp *TransactionSimpleProcessor) HandleFee(ctx context.Context, transaction *Transaction) error {
	LogInfofDw(ctx, "new fee transaction:%#v", transaction)
	err := dao.Collects.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {

		// 判断是否重复处理
		idVar, err := dao.Collects.Ctx(ctx).
			Where(dao.Collects.Columns().CollectHash, transaction.Hash).
			Where(dao.Collects.Columns().Status, g.Slice{
				model.COLLECT_STATUS_WAIT_COLLECT, model.COLLECT_STATUS_PROCESS_COLLECT, model.COLLECT_STATUS_WAIT_NOTIFY, model.COLLECT_STATUS_PROCESS_NOTIFY, model.COLLECT_STATUS_FINISH_NOTIFY,
			}).
			Value(dao.Collects.Columns().Id)
		if idVar.Int() > 0 {
			return custom_error.New("HandleFee hash has handle", g.Map{
				"transaction": transaction,
			})
		}

		// 根据充值hash找到collect表对应的记录并更新状态为 model.COLLECT_STATUS_WAIT_COLLECT
		_, err = dao.Collects.Ctx(ctx).Update(g.Map{
			dao.Collects.Columns().Status: model.COLLECT_STATUS_WAIT_COLLECT,
		}, g.Map{
			dao.Collects.Columns().HandfeeHash: transaction.Hash,
		})
		if err != nil {
			return custom_error.Wrap(err, "HandleFee 更新充值记录状态失败", g.Map{
				"transaction": transaction,
			})
		}

		// 更新队列表的状态为成功
		err = tp.updateQueueTaskSuccess(ctx, transaction)

		if err != nil {
			return custom_error.Wrap(err, "HandleFee 更新队列任务状态失败", g.Map{
				"transaction": transaction,
			})
		}
		return nil
	})
	return err
}

//  HandleCollect 根据hash更新collects表为 model.COLLECT_STATUS_WAIT_NOTIFY，queue_task表的状态为 model.QUEUE_TASK_STATUS_SUCCESS
func (tp *TransactionSimpleProcessor) HandleCollect(ctx context.Context, transaction *Transaction) error {

	LogInfofDw(ctx, "new collect transaction:%#v", transaction)
	err := dao.Collects.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {

		// 判断是否重复处理
		idVar, err := dao.Collects.Ctx(ctx).
			Where(dao.Collects.Columns().CollectHash, transaction.Hash).
			Where(dao.Collects.Columns().Status, g.Slice{
				model.COLLECT_STATUS_WAIT_NOTIFY, model.COLLECT_STATUS_PROCESS_NOTIFY, model.COLLECT_STATUS_FINISH_NOTIFY,
			}).
			Value(dao.Collects.Columns().Id)
		if idVar.Int() > 0 {
			return custom_error.New("HandleCollect has handle", g.Map{
				"transaction": transaction,
			})
		}

		// 根据归集hash找到collect表对应的记录并更新状态为 model.COLLECT_STATUS_WAIT_NOTIFY
		_, err = dao.Collects.Ctx(ctx).Update(g.Map{
			dao.Collects.Columns().Status: model.COLLECT_STATUS_WAIT_NOTIFY,
		}, g.Map{
			dao.Collects.Columns().CollectHash: transaction.Hash,
		})
		if err != nil {
			return custom_error.Wrap(err, "HandleCollect 更新充值记录状态失败", g.Map{
				"transaction": transaction,
			})
		}
		// 更新队列表的状态为成功
		err = tp.updateQueueTaskSuccess(ctx, transaction)

		if err != nil {
			return custom_error.Wrap(err, "HandleCollect 更新队列任务状态失败", g.Map{
				"transaction": transaction,
			})
		}
		return nil
	})
	return err
}

// HandleWithdraw 根据hash更新提现表的状态为 model.WITHDRAW_STATUS_WAIT_NOTIFY, queue_task表的状态为 model.QUEUE_TASK_STATUS_SUCCESS
func (tp *TransactionSimpleProcessor) HandleWithdraw(ctx context.Context, transaction *Transaction) error {
	LogInfofDw(ctx, "new withdraw transaction:%#v", transaction)
	err := dao.Withdraws.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {

		// 判断hash是否重复处理
		idVar, err := dao.Withdraws.Ctx(ctx).
			Where(dao.Withdraws.Columns().Hash, transaction.Hash).
			Where(dao.Withdraws.Columns().Status, g.Slice{
				model.WITHDRAW_STATUS_WAIT_NOTIFY, model.WITHDRAW_STATUS_PROCESS_NOTIFY, model.WITHDRAW_STATUS_FINISH_NOTIFY,
			}).
			Value(dao.Withdraws.Columns().Id)
		if idVar.Int() > 0 {
			return custom_error.New("withdraw hash has handle", g.Map{
				"transaction": transaction,
			})
		}

		_, err = dao.Withdraws.Ctx(ctx).Update(g.Map{
			dao.Withdraws.Columns().Status: model.WITHDRAW_STATUS_WAIT_NOTIFY,
		}, g.Map{
			dao.Withdraws.Columns().Hash: transaction.Hash,
		})
		if err != nil {
			return custom_error.Wrap(err, "HandleWithdraw 更新记录状态失败", g.Map{
				"transaction": transaction,
			})
		}
		// 更新队列表的状态为成功
		err = tp.updateQueueTaskSuccess(ctx, transaction)
		if err != nil {
			return custom_error.Wrap(err, "HandleWithdraw 更新队列任务状态失败", g.Map{
				"transaction": transaction,
			})
		}

		return nil
	})
	return err
}

// updateQueueTaskSuccess 更新队列任务的状态为 model.QUEUE_TASK_STATUS_SUCCESS
func (tp *TransactionSimpleProcessor) updateQueueTaskSuccess(ctx context.Context, transaction *Transaction) error {

	// 计算本次交易实际花费的手续费，单位为ether
	feeWeiBigInt := transaction.GasPrice.Mul(&transaction.GasPrice, big.NewInt(int64(transaction.GasUsed)))

	feeWeiBigFloat := new(big.Float).SetInt(feeWeiBigInt)

	feeEther := feeWeiBigFloat.Quo(feeWeiBigFloat, big.NewFloat(1e18))

	_, err := dao.QueueTask.Ctx(ctx).Update(g.Map{
		dao.QueueTask.Columns().Status:         model.QUEUE_TASK_STATUS_SUCCESS,
		dao.QueueTask.Columns().FinishAt:       gtime.Now().String(),
		dao.QueueTask.Columns().ActualGasUsed:  transaction.GasUsed,
		dao.QueueTask.Columns().ActualGasPrice: transaction.GasPrice.String(),
		dao.QueueTask.Columns().ActualFee:      feeEther,
	}, g.Map{
		dao.QueueTask.Columns().Hash: transaction.Hash,
	})

	return err
}
