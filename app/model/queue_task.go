package model

import (
	"gf-admin/app/model/entity"
)

const (
	QUEUE_TASK_STATUS_WAIT    = "wait"
	QUEUE_TASK_STATUS_PROCESS = "process"
	QUEUE_TASK_STATUS_SUCCESS = "success"
	QUEUE_TASK_STATUS_FAIL    = "fail"

	// 任务最大重试次数
	QUEUE_FAIL_MAX_TRY_AMOUNT = 5

	// 队列失败类型
	QUEUE_FAIL_BALANCE_INSUFFICIENT = -32000 //余额不足
)

type FrontTaskAndLog struct {
	Task *entity.QueueTask      `json:"task"`
	Logs []*entity.QueueTaskLog `json:"logs"`
}
