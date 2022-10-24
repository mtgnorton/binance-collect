package define

import (
	"gf-admin/app/model"

	"github.com/gogf/gf/v2/frame/g"
)

type IndexReq struct {
	g.Meta `path:"/" method:"get" tags:"首页" summary:"首页"`
}

type IndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type PostsNewHtmlReq struct {
	g.Meta `path:"/posts-new-html" method:"get" tags:"主题相关" summary:"发表主题页面"`
}

type PostsNewHtmlRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type PostsListReq struct {
	g.Meta `path:"/posts" method:"get" tags:"主题相关" summary:"主题列表"`
}

type PostsListRes struct {
	List []model.PostInfoInList `json:"list"`
}

type PostsStoreReq struct {
	g.Meta  `path:"/posts-store" method:"post" summary:"创建主题" tags:"主题相关"`
	NodeId  uint   `v:"required#节点id不能为空" json:"node_id"`
	Title   string `v:"required#标题不能为空" json:"title"`
	Content string `v:"required#内容不能为空" json:"content"`
}

type PostsStoreRes struct {
}

type PostsUpdateReq struct {
	g.Meta  `path:"/posts-update" method:"post" summary:"更新主题" tags:"主题相关"`
	Id      uint   `v:"required#主题id不能为空" json:"id"`
	Title   string `v:"required#标题不能为空" json:"title"`
	Content string `v:"required#内容不能为空" json:"content"`
}

type PostsUpdateRes struct {
}

type PostsMoveReq struct {
	g.Meta `path:"/posts-move" method:"post" summary:"移动主题" tags:"主题相关"`
	Id     uint `v:"required#主题id不能为空" json:"id"`
	NodeId uint `v:"required#新节点id不能为空" json:"node_id"`
}

type PostsMoveRes struct {
}

type PostsDetailReq struct {
	g.Meta `path:"/posts-detail-html/{post_id}" method:"get" summary:"主题详情" tags:"主题相关"`
	PostId uint `v:"required#主题id不能为空" json:"post_id"`
}

type PostsDetailRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
