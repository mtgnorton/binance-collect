package model

import (
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/util/gmeta"

	"github.com/gogf/gf/v2/os/gtime"
)

type TimePeriod int

const (
	Day1  TimePeriod = 1
	Day3  TimePeriod = 3
	Day7  TimePeriod = 7
	Month TimePeriod = 30
	Year  TimePeriod = 365
)

type Post struct {
	gmeta.Meta `orm:"table:forum_posts"`
	entity.Posts
}

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
type PostInfoInList struct {
	Id     uint `json:"id"                ` //
	NodeId uint `json:"node_id"  
` // 节点id
	NodeName          string      `json:"node_name"         `      // 节点名称
	UserId            uint        `json:"user_id"           `      // 用户id
	Username          string      `json:"username"          `      // 用户名
	Title             string      `json:"title"             `      // 标题
	ReplyAmount       uint        `json:"reply_amount"      `      // 回复次数
	LastChangeTime    string      `json:"last_change_time"   `     // 最后变动时间
	ReplyLastUserId   uint        `json:"reply_last_user_id"   `   // 最后回复用户id
	ReplyLastUsername string      `json:"reply_last_username"    ` // 最后回复用户名
	CreatedAt         *gtime.Time `json:"create_at"         `      // 主题创建时间
	ReplyLastTime     *gtime.Time `json:"create_at"         `      // 主题创建时间
}

type PostPageInput struct {
	Period        TimePeriod `json:"period" dc:"时间段，1:1天，3:3天，7:7天，30:30天，365:365天"`
	NodeId        uint       `json:"node_id" dc:"节点id"`
	IsIndex       bool       `json:"is_index" dc:"是否首页"`
	UserId        uint       `json:"user_id" dc:"用户id"`
	FilterKeyword string     `json:"filter_keyword" dc:"过滤关键字"`
	PageSizeInput
}

type PostPageList struct {
	PageSizeOutput
	List []*PostInfoInList
}

type PostWithComments struct {
	Post
	Replies []*Replies `orm:"with:posts_id=id"`
	Node    Node       `orm:"with:id=node_id"`
}
