package model

import (
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/util/gmeta"
)

const (
	REPLY_STATUS_NORMAL   = "normal"
	REPLY_STATUS_NO_AUDIT = "no_audit"
	REPLY_STATUS_SHIELD   = "shield"
)

type Replies struct {
	gmeta.Meta `orm:"table:forum_replies"`
	entity.Replies
}
