package define

import "github.com/gogf/gf/v2/frame/g"

type ReplyStoreReq struct {
	g.Meta  `path:"/comments-store" method:"post" tags:"新增评论" summary:"评论"`
	PostId  uint   `v:"required#主题id不能为空" json:"post_id"`
	Content string `v:"required#内容不能为空" json:"content"`
}

type ReplyStoreRes struct {
}
