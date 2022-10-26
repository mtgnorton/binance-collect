// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Messages is the golang structure of table forum_messages for DAO operations like Where/Data.
type Messages struct {
	g.Meta          `orm:"table:forum_messages, do:true"`
	Id              interface{} //
	UserId          interface{} // 用户id
	Username        interface{} // 用户名
	RepliedUserId   interface{} // 被回复用户id,用户a向用户b回复，用户b为 被回复用户id
	RepliedUsername interface{} // 被回复用户名
	PostId          interface{} // 关联主题id
	ReplyId         interface{} // 关联回复id
	IsRead          interface{} // 是否已读，否: 0,是: 1
	CreatedAt       *gtime.Time // 创建时间
	DeletedAt       *gtime.Time // 删除时间
}