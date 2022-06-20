package deposit_withdraw

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/utility/custom_error"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"
)

type TransactionNotifier struct {
	Ch chan *NotifyTask
}

func NewTransactionNotifier() *TransactionNotifier {
	return &TransactionNotifier{
		Ch: make(chan *NotifyTask, 100),
	}
}
func (qc *TransactionNotifier) Run(ctx context.Context) {
	go func() {
		for task := range qc.Ch {
			if task.IsRetry(ctx) {
				qc.Consume(ctx, task)
			}
		}
	}()
}

func (tn *TransactionNotifier) Consume(ctx context.Context, task *NotifyTask) {
	// 如果task.Id == 0 需要将task插入数据库
	if task.Id == 0 {
		id, err := dao.Notify.Ctx(ctx).InsertAndGetId(task)
		if err != nil {
			LogErrorfDw(ctx, custom_error.Wrap(err, "failed to insert notify", g.Map{
				"task": *task,
			}))
			return
		}
		task.Id = uint(id)
	}
	resp, err := tn.Send(ctx, gconv.Map(task.NotifyData), task.NotifyAddress)
	if err != nil {
		task.MarkFail(ctx, err)
		return
	}
	if resp == model.NOTIFY_SEND_SUCCESS {
		_, err = dao.Notify.Ctx(ctx).Update(g.Map{
			dao.Notify.Columns().Status: model.NOTIFY_STATUS_FINISH,
		}, g.Map{
			dao.Notify.Columns().Id: task.Id,
		})
		if err != nil {
			task.MarkFail(ctx, custom_error.Wrap(err, "failed to update notify", g.Map{
				"task": *task,
			}))
			return
		}
	} else {
		task.MarkFail(ctx, custom_error.New("failed to send notify,resp is :"+resp, g.Map{
			"task": *task,
		}))
		return
	}
	err = task.SendAfterFunc(ctx)

	if err != nil {
		LogErrorfDw(ctx, custom_error.Wrap(err, "notify SendAfterFunc", g.Map{
			"task": *task,
		}))
		return
	}
}

func (tn *TransactionNotifier) Send(ctx context.Context, data map[string]interface{}, link string) (string, error) {
	values := url.Values{}
	for k, v := range data {
		values.Add(k, v.(string))
	}
	args := values.Encode()
	resp, err := http.Post(link,
		"application/x-www-form-urlencoded", strings.NewReader(args))

	LogInfofDw(ctx, "send notify,link is :"+link+",data is :"+args)
	if err != nil {
		return "", custom_error.Wrap(err, "failed to send notify", g.Map{
			"data": data,
			"link": link,
		})
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", custom_error.Wrap(err, "failed read resp", g.Map{
			"data": data,
			"link": link,
		})
	}
	return string(content), nil
}
