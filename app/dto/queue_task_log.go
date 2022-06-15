// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// QueueTaskLog is the golang structure of table queue_task_log for DAO operations like Where/Data.
type QueueTaskLog struct {
	g.Meta      `orm:"table:queue_task_log, do:true"`
	Id          interface{} // ID
	QueueTaskId interface{} // 队列任务id
	Log         interface{} // 错误日志
	FailAmount  interface{} // 第几次失败
	CreateAt    *gtime.Time // 创建时间
	UpdateAt    *gtime.Time // 更新时间
}