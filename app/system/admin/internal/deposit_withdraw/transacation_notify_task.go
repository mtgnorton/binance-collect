package deposit_withdraw

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/dto"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"gf-admin/utility/custom_error"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type NotifyData struct {
	Symbol        string
	Value         string
	Type          string
	UserOrOrderId string
	UniqueId      string
}

type NotifyTask struct {
	entity.Notify
	NotifyData NotifyData `json:"notify_data"`
}

// retryInterval 通知重试的时间间隔,key为失败次数,值为时间间隔
var retryInterval = map[int]int{
	1: 4,
	2: 10,
	3: 10,
	4: 60,
	5: 120,
	6: 360,
	7: 900,
}

func (task *NotifyTask) IsRetry(ctx context.Context) bool {
	if task.IsImmediatelyRetry > 0 || task.FailAmount == 0 {
		return true
	}

	if task.NotifyAt == nil {
		return false
	}

	diffMin := (gtime.Timestamp() - task.NotifyAt.Timestamp()) / 60

	needDiffMin := gconv.Int64(retryInterval[task.FailAmount])

	g.Dump(diffMin, needDiffMin)
	return diffMin == needDiffMin
}

func (task *NotifyTask) MarkFail(ctx context.Context, recordErr error) {
	LogErrorfDw(ctx, recordErr)
	err := dao.Notify.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := dao.Notify.Ctx(ctx).Update(g.Map{
			dao.Notify.Columns().Status:             model.NOTIFY_STATUS_FAIL,
			dao.Notify.Columns().FailAmount:         task.FailAmount + 1,
			dao.Notify.Columns().IsImmediatelyRetry: 0,
			dao.Notify.Columns().NotifyAt:           gtime.Now(),
		}, g.Map{
			dao.Notify.Columns().Id: task.Id,
		})
		if err != nil {
			return err
		}

		_, err = dao.NotifyLog.Ctx(ctx).Insert(dto.NotifyLog{
			NotifyId:   task.Id,
			Log:        recordErr.Error(),
			FailAmount: task.FailAmount + 1,
		})
		return err
	})
	if err != nil {
		LogErrorfDw(ctx, custom_error.Wrap(err, "failed to update notify task", g.Map{
			"task": *task,
		}))
	}
}

func (task *NotifyTask) SendAfterSuccess(ctx context.Context) error {

	switch task.Type {
	case model.NOTIFY_TYPE_RECHARGE:
		_, err := dao.Collects.Ctx(ctx).Update(g.Map{
			dao.Collects.Columns().Status: model.COLLECT_STATUS_FINISH_NOTIFY,
		}, g.Map{
			dao.Collects.Columns().Id: task.RelationId,
		})

		return err
	case model.NOTIFY_TYPE_WITHDRAW:
		_, err := dao.Withdraws.Ctx(ctx).Update(g.Map{
			dao.Withdraws.Columns().Status: model.WITHDRAW_STATUS_FINISH_NOTIFY,
		}, g.Map{
			dao.Withdraws.Columns().Id: task.RelationId,
		})

		return err
	}

	_, err := dao.Notify.Ctx(ctx).Update(g.Map{
		dao.Notify.Columns().IsImmediatelyRetry: 0,
		dao.Notify.Columns().NotifyAt:           gtime.Now(),
	}, g.Map{
		dao.Notify.Columns().Id: task.Id,
	})

	return err
}
