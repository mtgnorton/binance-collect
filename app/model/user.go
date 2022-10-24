package model

import "github.com/gogf/gf/os/gtime"

const (
	USER_STATUS_NORMAL        = ""
	USER_STATUS_DISABLE_LOGIN = "disable_login"
	USER_STATUS_DISABLE_POST  = "disable_post"
	USER_STATUS_DISABLE_REPLY = "disable_reply"
)

type BalanceChangeType string

const (
	BALANCE_CHANGE_TYPE_LOGIN           BalanceChangeType = "login"
	BALANCE_CHANGE_TYPE_REGSITER        BalanceChangeType = "register"
	BALANCE_CHANGE_TYPE_ESTABLISH_POST  BalanceChangeType = "establish_post"
	BALANCE_CHANGE_TYPE_ESTABLISH_REPLY BalanceChangeType = "establish_reply"
	BALANCE_CHANGE_TYPE_THANK_REPLY     BalanceChangeType = "thanks_reply"
	BALANCE_CHANGE_TYPE_THANK_POST      BalanceChangeType = "thanks_post"
	BALANCE_CHANGE_TYPE_ACTIVITY        BalanceChangeType = "activity"
)

type UserSummary struct {
	Id       uint   `json:"id" gconv:"id"`
	Username string `json:"username" gconv:"username"`
	Email    string `json:"email" gconv:"email"`
}

type UserInfoWithoutPass struct {
	Id                  uint        `json:"id"                     ` //
	Username            string      `json:"username"               ` // 用户名
	Email               string      `json:"email"                  ` // email
	Description         string      `json:"description"            ` // 简介
	Avatar              string      `json:"avatar"                 ` // 头像地址
	Status              string      `json:"status"                 ` // 状态
	PostsAmount         uint        `json:"posts_amount"           ` // 创建主题次数
	ReplyAmount         uint        `json:"reply_amount"           ` // 回复次数
	ShieldedAmount      uint        `json:"shielded_amount"        ` // 被屏蔽次数
	FollowByOtherAmount uint        `json:"follow_by_other_amount" ` // 被关注次数
	TodayActivity       uint        `json:"today_activity"         ` // 今日活跃度
	Remark              string      `json:"remark"                 ` // 备注
	LastLoginIp         string      `json:"last_login_ip"          ` // 最后登陆IP
	LastLoginTime       *gtime.Time `json:"last_login_time"        ` // 最后登陆时间
	CreatedAt           string      `json:"created_at"             ` // 注册时间
	UpdatedAt           gtime.Time  `json:"updated_at"             ` // 更新时间
	DeletedAt           *gtime.Time `json:"deleted_at"             ` // 删除时间
}
