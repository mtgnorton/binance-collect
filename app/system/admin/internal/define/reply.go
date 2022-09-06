package define

import (
	"gf-admin/app/model"
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/os/gtime"
)

type ReplyListInput struct {
	Username  string     `dc:"用户名查询"`
	BeginTime gtime.Time `dc:"创建开始时间"`
	EndTime   gtime.Time `dc:"创建结束时间"`
	IsDestroy bool       `dc:"是否删除"`
	model.PageSizeInput
	model.OrderFieldDirectionInput
}

type ReplyListOutput struct {
	List []entity.Replies `json:"list"`
	model.PageSizeOutput
}

type ReplyListReq struct {
	g.Meta `path:"/reply-list" method:"get" summary:"回复列表" tags:"回复管理"`
	*ReplyListInput
}

type ReplyListRes struct {
	*ReplyListOutput
}

type ReplyToggleDestroyInput struct {
	Id        uint `json:"id"`
	IsDestroy bool `json:"is_destroy"`
}
type ReplyToggleDestroyReq struct {
	g.Meta `path:"/reply-toggle-destroy" method:"delete" summary:"删除回复" tags:"回复管理"`
	*ReplyToggleDestroyInput
}

type ReplyToggleDestroyRes struct {
}
