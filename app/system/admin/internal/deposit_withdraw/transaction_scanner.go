package deposit_withdraw

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"gf-admin/app/shared"
	"gf-admin/utility/custom_error"
	"math/big"
	"time"

	"github.com/gogf/gf/v2/util/grand"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/os/gtimer"

	"github.com/gogf/gf/v2/frame/g"
)

type TransactionScanner struct {
	scanInterval time.Duration //seconds
}

func NewTransactionScanner(ctx context.Context, scanInterval int) *TransactionScanner {
	return &TransactionScanner{
		scanInterval: time.Duration(scanInterval) * time.Second,
	}
}

// 执行扫描任务
func (ts *TransactionScanner) Scan(ctx context.Context, chTransfer chan *TransferTask, chNotify chan *NotifyTask) {
	gtimer.AddSingleton(ctx, ts.scanInterval, func(ctx context.Context) {
		tranferTasks, err := ts.scanTxFee(ctx)
		if err != nil {
			LogErrorfDw(ctx, err)
		}
		for _, task := range tranferTasks {
			chTransfer <- task
		}
		tranferTasks, err = ts.scanTxCollect(ctx)
		if err != nil {
			LogErrorfDw(ctx, err)
		}

		for _, task := range tranferTasks {
			chTransfer <- task
		}

		tranferTasks, err = ts.scanTxWithdraw(ctx)
		if err != nil {
			LogErrorfDw(ctx, err)
		}
		for _, task := range tranferTasks {
			chTransfer <- task
		}

		tranferTasks, err = ts.scanFailTxTransferTask(ctx)
		if err != nil {
			LogErrorfDw(ctx, err)
		}
		for _, task := range tranferTasks {
			chTransfer <- task
		}

		notifyTasks, err := ts.scanNotifyCollect(ctx)
		if err != nil {
			LogErrorfDw(ctx, err)
		}
		for _, task := range notifyTasks {
			chNotify <- task
		}

		notifyTasks, err = ts.scanNotifyWithdraw(ctx)
		if err != nil {
			LogErrorfDw(ctx, err)
		}
		for _, task := range notifyTasks {
			chNotify <- task
		}

		notifyTasks, err = ts.scanFailNotify(ctx)
		if err != nil {
			LogErrorfDw(ctx, err)
		}
		for _, task := range notifyTasks {
			chNotify <- task
		}
	})

}

//连续充值两笔发生短时间手续费向上波动，可能会造成第二笔归集失败
//以下为归集bnb的情况，如果是erc20情况类似
//假设连续充值两笔bnb
//
//第一笔 10bnb
//第二笔 20bnb
//
//第一次scanFee扫描
//假设此时获取到的手续费为1bnb
//此时写入collect表的数据为
//第一笔归集的金额9bnb,手续费1bnb
//第二笔归集的金额19bnb,手续费1bnb
//
//队列实际进行转账
//假设发生了手续费波动
//如果手续费上涨为2bnb
//第一笔进行归集,归集金额9bnb,手续费2bnb
//第二笔进行归集,归集金额19bnb,手续费2bnb,此时会失败,用户此时的余额为19bnb

// scanTxFee 将collects表中需要转出手续费的归集放入队列，不需要转出手续费的进程将状态置为待归集
func (ts *TransactionScanner) scanTxFee(ctx context.Context) ([]*TransferTask, error) {

	tasks := make([]*TransferTask, 0)

	var collects []*entity.Collects
	err := dao.Collects.Ctx(ctx).Where(g.Map{
		dao.Collects.Columns().Status: model.COLLECT_STATUS_WAIT_FEE,
	}).Scan(&collects)

	if err != nil {
		return tasks, err
	}

	LogInfofDw(ctx, "scanTxFee begin,tx length: %d", len(collects))
	for _, collect := range collects {
		onceFee, err := ChainClient.GetProbablyOnceGasPrice(ctx, collect.Symbol)
		if err != nil {
			return tasks, custom_error.Wrap(err, "scanTxFee error", g.Map{"collect": collect})
		}
		rechargeValue, ok := big.NewInt(0).SetString(collect.RechargeValue, 10)
		if !ok {

			return tasks, custom_error.New("scanTxFee error", g.Map{"collect": collect})
		}
		actualValue := rechargeValue

		// 当归集法币时，不需要转出手续费，只需要将归集的法币数量设置为rechargeValue-onceFee即可
		if collect.Symbol == model.CONTRACT_DEFAULT_SYMBOL {
			g.Dump(rechargeValue.String(), onceFee.String(), "ffffff")
			actualValue = rechargeValue.Sub(rechargeValue, &onceFee)
			if actualValue.Cmp(big.NewInt(0)) < 0 {
				actualValue = big.NewInt(0)
			}
		}

		settingMinRechargeValue, err := ChainClient.GetMinCollectValue(ctx)
		if err != nil {
			return tasks, custom_error.Wrap(err, "GetMinCollectValue error", g.Map{"collect": collect})
		}

		// 210000000000000 为一次手续费的数量

		//如果充值的金额小于设定的最小充值金额，则不进行归集
		if settingMinRechargeValue.Cmp(actualValue) > 0 {

			_, err = dao.Collects.Ctx(ctx).Update(g.Map{
				dao.Collects.Columns().Status: model.COLLECT_STATUS_FAIl_TOO_LOW_AMOUNT,
				dao.Collects.Columns().Value:  actualValue.String(),
			}, g.Map{
				dao.Collects.Columns().Id: collect.Id,
			})
			if err != nil {
				return tasks, custom_error.Wrap(err, "scanTxFee update collects error", g.Map{"collect": collect})
			}
			return tasks, nil
		}

		// 状态默认为待归集
		status := model.COLLECT_STATUS_WAIT_COLLECT

		var queueTask *TransferTask
		if collect.Symbol != model.CONTRACT_DEFAULT_SYMBOL {

			//  判断用户地址的eth余额是否足够作为手续费，如果足够则不需要转出手续费

			userEthBalance, err := ChainClient.GetBalance(ctx, collect.UserAddress)

			if err != nil {
				return tasks, err
			}

			if userEthBalance.Cmp(&onceFee) < 0 {
				// 封装转出手续费的任务
				feeWithdrawAddress, err := ChainClient.GetFeeWithdrawAddress(ctx)
				if err != nil {

					return tasks, custom_error.Wrap(err, "GetFeeWithdrawAddress error", g.Map{"collect": collect})
				}
				feeWithdrawPrivateKey, err := ChainClient.GetFeeWithdrawPrivateKey(ctx)
				if err != nil {

					return tasks, custom_error.Wrap(err, "GetFeeWithdrawPrivateKey error", g.Map{"collect": collect})
				}
				// 从手续费地址转出手续费到用户地址
				queueTask = &TransferTask{
					QueueTask: entity.QueueTask{
						Symbol:          model.CONTRACT_DEFAULT_SYMBOL,
						ContractAddress: model.CONTRACT_DEAULT_SYMBOL_ADDRESS,
						From:            feeWithdrawAddress,
						To:              collect.UserAddress,
						Value:           onceFee.String(),
						Type:            model.TRANSACTION_TYPE_FEE,
						PrivateKey:      feeWithdrawPrivateKey,
						RelationId:      collect.Id,
					},
				}

				status = model.COLLCT_STATUS_PROCESS_FEE
			}

		}

		// 更新collects表的状态
		updateData := g.Map{
			dao.Collects.Columns().Status: status,
			dao.Collects.Columns().Value:  actualValue.String(),
		}
		_, err = dao.Collects.Ctx(ctx).Update(updateData, g.Map{
			dao.Collects.Columns().Id: collect.Id,
		})

		if err != nil {

			return tasks, custom_error.Wrap(err, "scanTxFee update collects error", g.Map{"collect": collect, "updateData": updateData})
		}
		if queueTask != nil {
			tasks = append(tasks, queueTask)
		}
	}
	return tasks, nil

}

// 将collects表中需要归集的交易放入队列
func (ts *TransactionScanner) scanTxCollect(ctx context.Context) ([]*TransferTask, error) {
	tasks := make([]*TransferTask, 0)
	var collects []*entity.Collects
	err := dao.Collects.Ctx(ctx).Where(g.Map{
		dao.Collects.Columns().Status: model.COLLECT_STATUS_WAIT_COLLECT,
	}).Scan(&collects)
	if err != nil {
		return tasks, err
	}
	LogInfofDw(ctx, "scanTxCollect begin,tx length: %d", len(collects))

	collectAddress, err := ChainClient.GetCollectAddress(ctx)
	if err != nil {
		return tasks, err
	}

	// 判断手续费是否充足 允许20%的误差，因为手续费可能波动
	for _, collect := range collects {
		onceFee, err := ChainClient.GetProbablyOnceGasPrice(ctx, collect.Symbol)
		if err != nil {
			return tasks, custom_error.Wrap(err, "scanTxCollect error", g.Map{"collect": collect})
		}

		balance, err := ChainClient.GetBalance(ctx, collect.UserAddress)
		if err != nil {
			return tasks, custom_error.Wrap(err, "scanTxCollect get balance error", g.Map{
				"collect": collect,
			})
		}

		allowMinNonceFee := big.NewInt(0).Div(onceFee.Mul(&onceFee, big.NewInt(8)), big.NewInt(10))
		if balance.Cmp(allowMinNonceFee) < 0 { //手续费不足
			_, err := dao.Collects.Ctx(ctx).Update(g.Map{
				dao.Collects.Columns().Status: model.COLLECT_STATUS_WAIT_FEE,
			}, g.Map{
				dao.Collects.Columns().Id: collect.Id,
			})
			if err != nil {
				return tasks, custom_error.Wrap(err, "手续费不足，充值归集状态改为待转出手续费", g.Map{
					"collect": collect,
				})
			}
			return tasks, custom_error.New("手续费不足，充值归集状态改为待转出手续费", g.Map{
				"collect":          collect,
				"allowMinNonceFee": allowMinNonceFee.String(),
				"balance":          balance.String(),
			})
		}

		// 获取用户私钥
		userAddresses, err := ChainClient.GetUserAddresses(ctx)
		if err != nil {
			return tasks, custom_error.Wrap(err, "", g.Map{
				"collect": collect,
			})
		}
		userAddressInfo, ok := userAddresses[collect.UserAddress]
		if !ok {
			return tasks, custom_error.New("获取用户私钥失败", g.Map{
				"collect": collect,
			})
		}
		// 封装归集任务
		queueTask := &TransferTask{
			QueueTask: entity.QueueTask{
				Symbol:          collect.Symbol,
				ContractAddress: collect.ContractAddress,
				From:            collect.UserAddress,
				To:              collectAddress,
				Value:           collect.Value,
				Type:            model.TRANSACTION_TYPE_COLLECT,
				PrivateKey:      userAddressInfo.PrivateKey,
				RelationId:      collect.Id,
			},
		}
		// 更新collects表的状态
		updateData := g.Map{
			dao.Collects.Columns().Status: model.COLLECT_STATUS_PROCESS_COLLECT,
		}
		_, err = dao.Collects.Ctx(ctx).Update(updateData, g.Map{
			dao.Collects.Columns().Id: collect.Id,
		})

		if err != nil {

			return tasks, custom_error.Wrap(err, "scanTxCollect update collects error", g.Map{"collect": collect, "updateData": updateData})
		}

		tasks = append(tasks, queueTask)

	}
	return tasks, nil
}

// 检测失败的队列任务进行重试，失败次数需要小于5次，并且失败时间超过3min

//todo  如果队列任务的状态为process,并且超过一定时间，说明区块检测的时候该交易没有检测到，需要重新检测
func (ts *TransactionScanner) scanFailTxTransferTask(ctx context.Context) ([]*TransferTask, error) {
	tasks := make([]*TransferTask, 0)
	err := dao.QueueTask.Ctx(ctx).
		WhereLT(dao.QueueTask.Columns().UpdateAt, time.Now().Add(-3*time.Minute)).
		Where(dao.QueueTask.Columns().Status, model.QUEUE_TASK_STATUS_FAIL).
		Scan(&tasks)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

// 扫描待提现的记录放入队列,更新记录的状态为 model.WITHDRAW_STATUS_PROCESS
func (ts *TransactionScanner) scanTxWithdraw(ctx context.Context) ([]*TransferTask, error) {
	tasks := make([]*TransferTask, 0)
	var withdraws []*entity.Withdraws
	err := dao.Withdraws.Ctx(ctx).Where(g.Map{
		dao.Withdraws.Columns().Status: model.WITHDRAW_STATUS_WAIT,
	}).Scan(&withdraws)
	if err != nil {
		return tasks, err
	}
	LogInfofDw(ctx, "scanTxWithdraw begin,tx length: %d", len(withdraws))

	feeWithdrawAddress, err := ChainClient.GetFeeWithdrawAddress(ctx)

	if err != nil {
		return tasks, err
	}
	feeWithdrawPrivateKey, err := ChainClient.GetFeeWithdrawPrivateKey(ctx)
	if err != nil {
		return tasks, err
	}

	for _, withdraw := range withdraws {
		// 提现不需要判断提现地址的手续费，如果不足在队列进行实际交易时直接设为失败且不灾再重试
		queueTask := &TransferTask{
			QueueTask: entity.QueueTask{
				Symbol:          withdraw.Symbol,
				ContractAddress: withdraw.ContractAddress,
				From:            feeWithdrawAddress,
				To:              withdraw.To,
				Value:           withdraw.Value,
				Type:            model.TRANSACTION_TYPE_WITHDRAW,
				PrivateKey:      feeWithdrawPrivateKey,
				RelationId:      withdraw.Id,
			},
		}

		// 更新withdraws表的状态为处理中
		updateData := g.Map{
			dao.Withdraws.Columns().Status: model.WITHDRAW_STATUS_PROCESS,
			dao.Withdraws.Columns().From:   feeWithdrawAddress,
		}
		_, err = dao.Withdraws.Ctx(ctx).Update(updateData, g.Map{
			dao.Withdraws.Columns().Id: withdraw.Id,
		})

		if err != nil {

			return tasks, custom_error.Wrap(err, "scanTxWithdraw update withdraw error", g.Map{"withdraw": withdraw, "updateData": updateData})
		}

		tasks = append(tasks, queueTask)

	}

	return tasks, nil
}

// scanNotifyCollect 扫描collect表中状态为 model.COLLECT_STATUS_WAIT_NOTIFY 的记录
func (ts *TransactionScanner) scanNotifyCollect(ctx context.Context) ([]*NotifyTask, error) {
	tasks := make([]*NotifyTask, 0)
	var collects []*entity.Collects
	err := dao.Collects.Ctx(ctx).Where(g.Map{
		dao.Collects.Columns().Status: model.COLLECT_STATUS_WAIT_NOTIFY,
	}).Scan(&collects)
	if err != nil {
		return tasks, err
	}
	LogInfofDw(ctx, "scanNotifyCollect begin,tx length: %d", len(collects))

	notifyAddress, err := shared.Config.GetString(ctx, model.CONFIG_MODULE_BINNABCE, model.CONFIG_KEY_NOTIFY_ADDRESS)
	if err != nil {
		return tasks, err
	}
	userAddresses, err := ChainClient.GetUserAddresses(ctx)
	if err != nil {
		return tasks, err
	}
	for _, collect := range collects {
		userInfo, ok := userAddresses[collect.UserAddress]
		if !ok {

			return tasks, custom_error.New("scanNotifyCollect user address not found", g.Map{
				"collect": collect,
			})
		}
		externalUserId := userInfo.ExternalUserId

		if err != nil {
			return tasks, err
		}

		// 金额单位需要转换成ether
		valueEther, err := ChainClient.WeiToEther(ctx, collect.Value, collect.Symbol)
		if err != nil {
			return tasks, err
		}
		uniqueId := grand.S(20)
		task := &NotifyTask{
			Notify: entity.Notify{
				Type:          model.NOTIFY_TYPE_RECHARGE,
				RelationId:    collect.Id,
				NotifyAddress: notifyAddress,
				Status:        model.NOTIFY_STATUS_WAIT,
				UniqueId:      uniqueId,
			},
			NotifyData: NotifyData{
				Symbol:        collect.Symbol,
				Value:         valueEther,
				Type:          model.NOTIFY_TYPE_RECHARGE,
				UserOrOrderId: externalUserId,
				UniqueId:      uniqueId,
			},
		}
		_, err = dao.Collects.Ctx(ctx).Update(g.Map{
			dao.Collects.Columns().Status: model.COLLECT_STATUS_PROCESS_NOTIFY,
		}, g.Map{
			dao.Collects.Columns().Id: collect.Id,
		})

		if err != nil {
			return tasks, custom_error.Wrap(err, "scanNotifyCollect update collect error", g.Map{"collect": collect})
		}
		tasks = append(tasks, task)
	}
	return tasks, nil

}

// scanNotifyWithdraw 扫描withdraw表中状态为 model.WITHDRAW_STATUS_WAIT_NOTIFY 的记录
func (ts *TransactionScanner) scanNotifyWithdraw(ctx context.Context) ([]*NotifyTask, error) {
	tasks := make([]*NotifyTask, 0)
	var withdraws []*entity.Withdraws
	err := dao.Withdraws.Ctx(ctx).Where(g.Map{
		dao.Withdraws.Columns().Status: model.WITHDRAW_STATUS_WAIT_NOTIFY,
	}).Scan(&withdraws)
	if err != nil {
		return tasks, err
	}
	LogInfofDw(ctx, "scanNotifyWithdraw begin,tx length: %d", len(withdraws))

	notifyAddress, err := shared.Config.GetString(ctx, model.CONFIG_MODULE_BINNABCE, model.CONFIG_KEY_NOTIFY_ADDRESS)
	if err != nil {
		return tasks, err
	}
	for _, withdraw := range withdraws {

		// 金额单位需要转换成ether
		valueEther, err := ChainClient.WeiToEther(ctx, withdraw.Value, withdraw.Symbol)
		if err != nil {
			return tasks, err
		}
		uniqueId := grand.S(20)

		task := &NotifyTask{
			Notify: entity.Notify{
				Type:          model.NOTIFY_TYPE_WITHDRAW,
				RelationId:    withdraw.Id,
				NotifyAddress: notifyAddress,
				Status:        model.NOTIFY_STATUS_WAIT,
				UniqueId:      uniqueId,
			},
			NotifyData: NotifyData{
				Symbol:        withdraw.Symbol,
				Value:         valueEther,
				Type:          model.NOTIFY_TYPE_WITHDRAW,
				UserOrOrderId: withdraw.ExternalOrderId,
				UniqueId:      uniqueId,
			},
		}
		_, err = dao.Withdraws.Ctx(ctx).Update(g.Map{
			dao.Withdraws.Columns().Status: model.WITHDRAW_STATUS_PROCESS_NOTIFY,
		}, g.Map{
			dao.Withdraws.Columns().Id: withdraw.Id,
		})

		if err != nil {
			return tasks, custom_error.Wrap(err, "scanNotifyWithdraw update withdraw error", g.Map{"withdraw": withdraw})
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// scanFailNotify 扫描notify表中状态为 model.NOTIFY_STATUS_FAIL 的记录，且失败次数小于 model.NOTIFY_MAX_RETRY_AMOUNT 次
func (ts *TransactionScanner) scanFailNotify(ctx context.Context) ([]*NotifyTask, error) {
	var notifies []*entity.Notify
	var tasks []*NotifyTask

	err := dao.Notify.Ctx(ctx).
		Where(dao.Notify.Columns().Status, model.NOTIFY_STATUS_FAIL).
		WhereLT(dao.Notify.Columns().FailAmount, model.NOTIFY_MAX_RETRY_AMOUNT).Scan(&notifies)
	if err != nil {
		return nil, err
	}

	LogInfofDw(ctx, "scanFailNotify begin,notifies length: %d", len(notifies))

	for _, notify := range notifies {
		if notify.NotifyData == "" {
			continue
		}
		notifyData := NotifyData{}
		err = gconv.Scan(notify.NotifyData, &notifyData)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, &NotifyTask{
			Notify:     *notify,
			NotifyData: notifyData,
		})
	}
	return tasks, err
}
