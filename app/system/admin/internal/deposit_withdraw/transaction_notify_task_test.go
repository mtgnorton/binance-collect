package deposit_withdraw

import (
	"encoding/json"
	"gf-admin/app/model/entity"
	"testing"
	"time"

	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestTransactionNotifier_Consume(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		scanner := NewTransactionScanner(ctx, 5)
		tasks, err := scanner.scanNotifyCollect(ctx)

		t.AssertNil(err)
		t.AssertGT(len(tasks), 0)

		notifier := NewTransactionNotifier()
		g.Dump(json.Marshal(*tasks[0]))
		g.Dump(*tasks[0])
		notifier.Consume(ctx, tasks[0])
	})

}

// retryInterval 通知重试的时间间隔,key为失败次数,值为时间间隔
//var retryInterval = map[int]int{
//	1: 4,
//	2: 10,
//	3: 10,
//	4: 60,
//	5: 120,
//	6: 360,
//	7: 900,
//}
func TestNotifyTask_IsRetry(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		task := &NotifyTask{
			Notify: entity.Notify{
				Id:                 0,
				Type:               "",
				RelationId:         0,
				NotifyData:         "",
				NotifyAddress:      "",
				FailAmount:         1,
				Status:             "",
				IsImmediatelyRetry: 0,
				CreateAt:           nil,
				NotifyAt:           gtime.Now().Add(time.Minute * -1),
			},
			NotifyData: NotifyData{},
		}

		r1 := task.IsRetry(ctx)
		t.Assert(r1, false)

		task.NotifyAt = gtime.Now().Add(time.Minute * -4)

		r2 := task.IsRetry(ctx)
		t.Assert(r2, true)

		task.FailAmount = 2
		task.NotifyAt = gtime.Now().Add(time.Minute * -10)

		r3 := task.IsRetry(ctx)
		t.Assert(r3, true)

		task.FailAmount = 3
		task.NotifyAt = gtime.Now().Add(time.Minute * -10)

		r4 := task.IsRetry(ctx)
		t.Assert(r4, true)
	})
}

func Test_NewNotifyResponseServer(t *testing.T) {

	s := g.Server()
	s.SetPort(8787)

	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("success")
	})
	s.Run()
}
