package deposit_withdraw

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/utility/custom_error"
	"math/big"

	"github.com/gogf/gf/v2/util/grand"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/gogf/gf/v2/frame/g"
)

type TransactionTransfer struct {
	Ch chan *TransferTask
}

func NewTransactionTransfer() *TransactionTransfer {
	return &TransactionTransfer{
		Ch: make(chan *TransferTask),
	}
}

func (qc *TransactionTransfer) Run(ctx context.Context) *TransactionTransfer {
	go func() {
		for task := range qc.Ch {
			qc.Consume(ctx, task)
		}
	}()
	return qc
}

func (qc *TransactionTransfer) Consume(ctx context.Context, task *TransferTask) {
	if task.Id == 0 {
		task.Hash = "temp:" + gtime.TimestampMicroStr() + grand.Letters(10)
		id, err := dao.QueueTask.Ctx(ctx).InsertAndGetId(task)
		if err != nil {
			logErrorfDw(ctx, custom_error.Wrap(err, "failed to insert queue task", g.Map{
				"task": task,
			}))
			return
		}
		task.Id = int(id)
	}
	nonce, err := ChainClient.GetLastNonce(ctx, task.From)

	if err != nil {
		task.MarkFail(ctx, custom_error.Wrap(err, "failed to get last nonce", g.Map{
			"task": *task,
		}))
		return
	}

	gasLimit := ChainClient.GetGasLimit(ctx, task.Symbol)

	gasPrice, err := ChainClient.GetGasPrice(ctx)

	if err != nil {
		task.MarkFail(ctx, custom_error.Wrap(err, "failed to get gas price", g.Map{
			"task": *task,
		}))
		return
	}

	value, ok := big.NewInt(0).SetString(task.Value, 10)
	if !ok {
		task.MarkFail(ctx, custom_error.New("failed to get value", g.Map{
			"task": *task,
		}))
		return
	}

	toAddress := common.HexToAddress(task.To)

	legacyTx := &types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &toAddress,
		Value:    value,
		Data:     nil,
	}

	var hash string

	logInfofDw(ctx, "send tx data %#v", *legacyTx)
	hash, err = ChainClient.SendTransaction(ctx, legacyTx, task.PrivateKey, task.ContractAddress)

	if err != nil {

		isNotTry := false
		// 根据错误类型判断是否重试
		gcode := gerror.Code(err)
		if gcode.Code() == model.QUEUE_FAIL_BALANCE_INSUFFICIENT {
			isNotTry = true
		}
		task.MarkFail(ctx, custom_error.Wrap(err, "transfer fail", g.Map{
			"task": *task,
			"tx": g.Map{
				"nonce":           nonce,
				"to":              task.To,
				"value":           value,
				"privateKey":      task.PrivateKey,
				"gasLimit":        gasLimit,
				"gasPrice":        gasPrice,
				"contractAddress": task.ContractAddress,
			},
		}), isNotTry)
		return
	}
	task.Hash = hash
	_, err = dao.QueueTask.Ctx(ctx).Update(g.Map{
		dao.QueueTask.Columns().Hash:     hash,
		dao.QueueTask.Columns().Status:   model.QUEUE_TASK_STATUS_PROCESS,
		dao.QueueTask.Columns().GasLimit: gasLimit,
		dao.QueueTask.Columns().GasPrice: gasPrice,
		dao.QueueTask.Columns().Nonce:    nonce,
		dao.QueueTask.Columns().SendAt:   gtime.Now().String(),
	}, g.Map{
		dao.QueueTask.Columns().Id: task.Id,
	})

	if err != nil {
		logErrorfDw(ctx, custom_error.Wrap(err, "update error", g.Map{
			"task": *task,
		}))
		return
	}
	err = task.SendAfterFunc(ctx)

	if err != nil {
		logErrorfDw(ctx, custom_error.Wrap(err, "transfer SendAfterFunc", g.Map{
			"task": *task,
		}))
		return
	}
}
