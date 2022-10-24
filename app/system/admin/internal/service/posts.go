package service

import (
	"context"
	"fmt"
	"gf-admin/app/dao"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/utility/response"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var Post = post{}

type post struct {
}

func (p *post) List(ctx context.Context, in *define.PostListInput) (out *define.PostListOutput, err error) {
	out = &define.PostListOutput{}
	d := dao.Posts.Ctx(ctx)
	if in.NodeId != 0 {
		d = d.Where(dao.Posts.Columns().NodeId, in.NodeId)
	}
	if in.IsDestroy {
		d = d.WhereNot(dao.Posts.Columns().DeletedAt, nil)
	} else {
		d = d.Where(dao.Posts.Columns().DeletedAt, nil)

	}
	if in.Username != "" {
		d = d.Where(dao.Posts.Columns().Username, in.Username)
	}

	if in.Title != "" {
		d = d.WhereLike(dao.Posts.Columns().Title, fmt.Sprintf("%%%s%%", in.Title))
	}

	if in.BeginTime != "" {
		d = d.WhereGTE(dao.Posts.Columns().CreatedAt, in.BeginTime)
	}
	if in.EndTime != "" {
		d = d.WhereLTE(dao.Posts.Columns().CreatedAt, in.EndTime)
	}

	out.Page = in.Page
	out.Size = in.Size
	out.Total, err = d.Count()
	if err != nil {
		return out, response.NewError(err.Error())
	}
	d = d.Page(in.Page, in.Size)
	if in.OrderField != "" && in.OrderDirection != "" {
		d = d.Order(in.OrderField, in.OrderDirection)
	}

	err = d.Scan(&out.List)

	return
}

func (p *post) ToggleTop(ctx context.Context, in *define.PostToggleTopInput) (err error) {

	err = p.ExistById(ctx, in.Id)
	if err != nil {
		return err
	}
	if !in.EndTime.IsZero() && in.EndTime.Before(gtime.Now()) {
		return response.NewError("置顶截止时间不能小于当前时间")
	}
	_, err = dao.Posts.Ctx(ctx).Where(dao.Posts.Columns().Id, in.Id).Update(g.Map{
		dao.Posts.Columns().TopEndTime: in.EndTime,
	})

	return
}

func (p *post) ToggleDestroy(ctx context.Context, in *define.PostToggleDestroyInput) (err error) {

	err = p.ExistById(ctx, in.Id)
	if err != nil {
		return err
	}

	updateValue := gtime.Now()

	if !in.IsDestroy {
		updateValue = nil
	}
	_, err = dao.Posts.Ctx(ctx).Where(dao.Posts.Columns().Id, in.Id).Update(g.Map{
		dao.Posts.Columns().DeletedAt: updateValue,
	})
	return
}

func (p *post) ExistById(ctx context.Context, Id uint) (err error) {

	exist, err := dao.Posts.Ctx(ctx).WherePri(Id).Count()
	if err != nil {
		return err
	}
	if exist == 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "主题不存在")
	}
	return nil
}
