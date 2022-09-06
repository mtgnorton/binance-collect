package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/utility/custom_error"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

var Reply = reply{}

type reply struct {
}

func (r *reply) List(ctx context.Context, in *define.ReplyListInput) (out *define.ReplyListOutput, err error) {
	out = &define.ReplyListOutput{}
	d := dao.Replies.Ctx(ctx)
	if in.Username != "" {
		d = d.Where(dao.Replies.Columns().Username, in.Username)
	}
	if !in.BeginTime.IsZero() {
		d = d.WhereGTE(dao.Replies.Columns().CreatedAt, in.BeginTime)
	}

	if !in.EndTime.IsZero() {
		d = d.WhereLTE(dao.Replies.Columns().CreatedAt, in.EndTime)
	}
	if in.IsDestroy {
		d = d.WhereNot(dao.Replies.Columns().DeletedAt, nil)
	} else {
		d = d.Where(dao.Replies.Columns().DeletedAt, nil)
	}

	out.Page = in.Page
	out.Size = in.Size
	out.Total, err = d.Count()
	if err != nil {
		return out, custom_error.New(err.Error())
	}
	d = d.Page(in.Page, in.Size)
	if in.OrderField != "" && in.OrderDirection != "" {
		d = d.Order(in.OrderField, in.OrderDirection)
	}
	err = d.Scan(&out.List)
	return
}

func (r *reply) ToggleDestroy(ctx context.Context, in *define.ReplyToggleDestroyInput) (err error) {
	exist, err := dao.Replies.Ctx(ctx).WherePri(in.Id).Count()
	if err != nil {
		return err
	}
	if exist == 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "主题不存在")
	}
	if err != nil {
		return err
	}

	updateValue := gtime.Now()

	if !in.IsDestroy {
		updateValue = nil
	}
	_, err = dao.Replies.Ctx(ctx).Where(dao.Replies.Columns().Id, in.Id).Update(g.Map{
		dao.Replies.Columns().DeletedAt: updateValue,
	})

	return
}
