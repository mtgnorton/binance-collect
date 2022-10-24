package define

import "github.com/gogf/gf/v2/frame/g"

type MemberIndexReq struct {
	g.Meta `path:"/member-html" method:"get" tags:"个人中心" summary:"个人中心"`
}

type MemberIndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
