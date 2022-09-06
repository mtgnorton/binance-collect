package service

import (
	"context"
	"fmt"
	"gf-admin/app/dao"
	"gf-admin/app/system/admin/internal/define"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var Node = node{}

type node struct {
}

func (n *node) List(ctx context.Context, in *define.NodeListInput) (output *define.NodeListOutput, err error) {
	output = &define.NodeListOutput{}
	d := dao.Nodes.Ctx(ctx)
	if in.Name != "" {
		d = d.WhereLike(dao.Nodes.Columns().Name, fmt.Sprintf("%%%s%%", in.Name))

	}
	if in.IsIndex != "" {
		d = d.Where(dao.Nodes.Columns().IsIndex, in.IsIndex)
	}

	err = d.Scan(&output.List)

	return
}

func (n *node) Store(ctx context.Context, in *define.NodeStoreInput) (err error) {
	d := dao.Nodes.Ctx(ctx)
	exist, err := d.Where(dao.Nodes.Columns().Name, in.Name).Count()
	if err != nil {
		return err
	}
	if exist > 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "节点名称已存在")
	}
	_, err = d.Insert(in)
	return err
}

func (n *node) Update(ctx context.Context, in *define.NodeUpdateInput) (err error) {
	d := dao.Nodes.Ctx(ctx)
	err = n.ExistById(ctx, in.Id)
	if err != nil {
		return
	}
	nameExist, err := d.WhereNot(dao.Nodes.Columns().Id, in.Id).Where(dao.Nodes.Columns().Name, in.Name).Count()
	if err != nil {
		return err
	}
	if nameExist > 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "节点名称已存在")
	}

	_, err = d.WherePri(in.Id).Update(in)
	return err
}

func (n *node) Destroy(ctx context.Context, Id uint) (err error) {
	err = n.ExistById(ctx, Id)
	if err != nil {
		return err
	}
	d := dao.Nodes.Ctx(ctx)
	_, err = d.WherePri(Id).Delete()
	return err
}

func (n *node) ExistById(ctx context.Context, Id uint) (err error) {
	d := dao.Nodes.Ctx(ctx)
	exist, err := d.WherePri(Id).Count()
	if err != nil {
		return
	}
	if exist == 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "节点不存在")
	}
	return
}
