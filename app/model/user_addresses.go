package model

import "github.com/gogf/gf/v2/os/gtime"

const (
	USER_ADDRESS_TYPE_GENERATE = 0
	USER_ADDRESS_TYPE_IMPORT   = 1
)

type FrontUserAddress struct {
	Id             uint        `json:"id"               ` // ID
	Address        string      `json:"address"          ` // 以太坊地址
	ExternalUserId string      `json:"external_user_id" ` // 外部用户id
	CreateAt       *gtime.Time `json:"create_at"        ` // 创建时间
	UpdateAt       *gtime.Time `json:"update_at"        ` // 更新时间
}
