package define

import (
	"gf-admin/app/model"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/frame/g"
)

type PostListInput struct {
	NodeId    int    `json:"node_id"`
	IsDestroy bool   `json:"is_destroy"`
	Username  string `json:"username"`
	Title     string `json:"title"`
	BeginTime string `json:"begin_time"`
	EndTime   string `json:"end_time"`
	model.PageSizeInput
	model.OrderFieldDirectionInput
}

type PostListOutput struct {
	List []model.PostWithoutContent `json:"posts"`
	model.PageSizeOutput
}

type PostListReq struct {
	g.Meta `path:"/posts-list" method:"get" summary:"主题列表" tags:"主题管理"`
	*PostListInput
}

type PostListRes struct {
	*PostListOutput
}

type PostToggleDestroyInput struct {
	Id        uint `json:"id"`
	IsDestroy bool `json:"is_destroy"`
}
type PostToggleDestroyReq struct {
	g.Meta `path:"/posts-toggle-destroy" method:"delete" summary:"删除主题" tags:"主题管理"`
	*PostToggleDestroyInput
}

type PostToggleDestroyRes struct {
}

type PostToggleTopInput struct {
	Id      uint       `json:"id"`
	EndTime gtime.Time `json:"end_time" dc:"置顶截止时间,为空时，代表取消置顶"`
}

type PostToggleTopReq struct {
	g.Meta `path:"/posts-toggle-top" method:"put" summary:"置顶主题" tags:"主题管理"`
	*PostToggleTopInput
}

type PostToggleTopRes struct {
}
