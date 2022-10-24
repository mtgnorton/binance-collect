package model

import (
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/util/gmeta"
)

type Node struct {
	gmeta.Meta `orm:"table:forum_nodes"`
	entity.Nodes
}
