package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model/entity"
)

var Node = node{}

type node struct {
}

func (n *node) Query(ctx context.Context, isIndex bool) (nodes []entity.Nodes, err error) {
	d := dao.Nodes.Ctx(ctx)

	if isIndex {
		d.Where(dao.Nodes.Columns().IsIndex, 1)
	}

	err = d.Order(dao.Nodes.Columns().Sort + " asc").
		Order(dao.Nodes.Columns().Id + " asc").
		Scan(&nodes)

	return
}
