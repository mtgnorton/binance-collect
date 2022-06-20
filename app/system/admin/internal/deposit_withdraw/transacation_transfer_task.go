package deposit_withdraw

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"gf-admin/utility/custom_error"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type TransferTask struct {
	entity.QueueTask
}

// 将一个任务标记为失败，并且可以将失败次数设置为最大，在某些情况下应该设为不再重试，如余额不足
func (task *TransferTask) MarkFail(ctx context.Context, recordErr error, isNotTry ...bool) {

	LogErrorfDw(ctx, recordErr)

	failAmount := 1

	err := dao.QueueTask.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {

		// 记录失败原因
		oldFailAmount, err := dao.QueueTask.Ctx(ctx).Value(dao.QueueTask.Columns().FailAmount, g.Map{dao.QueueTask.Columns().Id: task.Id})
		if err != nil {
			return err
		}

		if len(isNotTry) == 1 && isNotTry[0] {
			failAmount = model.QUEUE_FAIL_MAX_TRY_AMOUNT - oldFailAmount.Int()
		}

		// 更新queueTask失败次数
		_, err = dao.QueueTask.Ctx(ctx).Update(g.Map{
			dao.QueueTask.Columns().FailAmount: &gdb.Counter{
				Field: dao.QueueTask.Columns().FailAmount,
				Value: float64(failAmount),
			},
			dao.QueueTask.Columns().Status: model.QUEUE_TASK_STATUS_FAIL,
		}, g.Map{
			dao.QueueTask.Columns().Id: task.Id,
		})

		if err != nil {
			return err
		}

		_, err = dao.QueueTaskLog.Ctx(ctx).Insert(entity.QueueTaskLog{
			QueueTaskId: task.Id,
			Log:         recordErr.Error(),
			FailAmount:  oldFailAmount.Int() + 1,
		})

		return err
	})
	if err != nil {
		LogErrorfDw(ctx, custom_error.Wrap(err, "failed to update transfer log", g.Map{
			"task": *task,
		}))
	}
}

func (task *TransferTask) SendAfterFunc(ctx context.Context) error {
	switch task.Type {
	// 将转出手续费的hash更新到collects表
	case model.TRANSACTION_TYPE_FEE:
		_, err := dao.Collects.Ctx(ctx).Update(g.Map{
			dao.Collects.Columns().HandfeeHash: task.Hash,
		}, g.Map{
			dao.Collects.Columns().Id: task.RelationId,
		})

		return err
	// 将归集的hash更新到collects表
	case model.TRANSACTION_TYPE_COLLECT:
		_, err := dao.Collects.Ctx(ctx).Update(g.Map{
			dao.Collects.Columns().CollectHash: task.Hash,
		}, g.Map{
			dao.Collects.Columns().Id: task.RelationId,
		})

		return err

	case model.TRANSACTION_TYPE_WITHDRAW:
		_, err := dao.Withdraws.Ctx(ctx).Update(g.Map{
			dao.Withdraws.Columns().Hash: task.Hash,
		}, g.Map{
			dao.Withdraws.Columns().Id: task.RelationId,
		})

		return err
	}
	return nil
}
