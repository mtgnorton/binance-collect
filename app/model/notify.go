package model

import "gf-admin/app/model/entity"

const (
	NOTIFY_TYPE_RECHARGE = "recharge"
	NOTIFY_TYPE_WITHDRAW = "withdraw"

	NOTIFY_STATUS_WAIT   = "wait"
	NOTIFY_STATUS_FINISH = "finish"
	NOTIFY_STATUS_FAIL   = "fail"

	NOTIFY_SEND_SUCCESS = "success"

	NOTIFY_MAX_RETRY_AMOUNT = 8
)

type FrontNotify struct {
	Notify *entity.Notify      `json:"notify"`
	Logs   []*entity.NotifyLog `json:"logs"`
}
