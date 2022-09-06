package service

import (
	"context"
	"fmt"
	"gf-admin/app/dao"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/utility/custom_error"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"
)

var User = user{}

type user struct {
}

func (u *user) List(ctx context.Context, in *define.UserListInput) (out *define.UserListOutput, err error) {
	out = &define.UserListOutput{}

	d := dao.Users.Ctx(ctx)
	if in.Username != "" {
		d = d.WhereLike(dao.Users.Columns().Username, fmt.Sprintf("%%%s%%", in.Username))
	}
	if in.Email != "" {
		d = d.WhereLike(dao.Users.Columns().Email, fmt.Sprintf("%%%s%%", in.Email))
	}
	if in.Status != 0 {
		d = d.Where(dao.Users.Columns().Status, in.Status)
	}
	if in.IsDestroy {
		d = d.WhereNot(dao.Users.Columns().DeletedAt, nil)
	} else {
		d = d.Where(dao.Users.Columns().DeletedAt, nil)
	}

	if in.BeginTime != "" {
		d = d.WhereGTE(dao.Users.Columns().CreatedAt, in.BeginTime)
	}

	if in.EndTime != "" {
		d = d.WhereLTE(dao.Users.Columns().CreatedAt, in.EndTime)
	}

	out.Page = in.Page
	out.Size = in.Size
	out.Total, err = d.Count()
	if err != nil {
		return out, gerror.New(err.Error())
	}
	d = d.Page(in.Page, in.Size)
	if in.OrderField != "" && in.OrderDirection != "" {
		d = d.Order(in.OrderField, in.OrderDirection)
	}
	err = d.Scan(&out.List)
	return
}

func (u *user) ToggleDestroy(ctx context.Context, in *define.UserToggleDestroyInput) (err error) {
	userCount, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, in.Id).Count()
	if err != nil {
		return err
	}
	if userCount == 0 {
		return custom_error.New("用户不存在", g.Map{})
	}
	updateValue := gtime.Now()

	if !in.IsDestroy {
		updateValue = nil
	}
	_, err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, in.Id).Update(g.Map{
		dao.Users.Columns().DeletedAt: updateValue,
	})
	return
}
