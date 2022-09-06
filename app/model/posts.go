package model

import "github.com/gogf/gf/os/gtime"

type PostWithoutContent struct {
	Id               uint        `json:"id"                ` //
	NodeId           uint        `json:"node_id"           ` // 节点id
	UserId           uint        `json:"user_id"           ` // 用户id
	Username         string      `json:"username"          ` // 用户名
	Title            string      `json:"title"             ` // 标题
	TopEndTime       *gtime.Time `json:"top_end_time"      ` // 置顶截止时间,为空说明没有置顶
	CharacterAmount  uint        `json:"character_amount"  ` // 字符长度
	VisitsAmount     uint        `json:"visits_amount"     ` // 访问次数
	CollectionAmount uint        `json:"collection_amount" ` // 收藏次数
	ReplyAmount      uint        `json:"reply_amount"      ` // 回复次数
	ThanksAmount     uint        `json:"thanks_amount"     ` // 感谢次数
	ShieldedAmount   uint        `json:"shielded_amount"   ` // 被屏蔽次数
	Weight           int         `json:"weight"            ` // 权重
	ReplyLastTime    *gtime.Time `json:"reply_last_time"   ` // 最后回复时间
	CreatedAt        *gtime.Time `json:"created_at"        ` // 主题创建时间
	UpdatedAt        *gtime.Time `json:"updated_at"        ` // 主题更新时间
	DeletedAt        *gtime.Time `json:"deleted_at"        ` // 删除时间

}
