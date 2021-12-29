package service

import (
	"context"
	"gf-admin/app/model"
)

// Session管理服务
var Session = serviceSession{}

type serviceSession struct{}

const (
	sessionKeyAdministrator = "SessionKeyAdministrator" // 用户信息存放在Session中的Key
)




// 设置用户Session.
func (s *serviceSession) SetAdministrator(ctx context.Context, user *model.Administrator) error {
	return Context.Get(ctx).Session.Set(sessionKeyAdministrator, user)
}

// 获取当前登录的用户信息对象，如果用户未登录返回nil。
func (s *serviceSession) GetAdministrator(ctx context.Context) *model.Administrator {
	customCtx := Context.Get(ctx)
	if customCtx != nil {
		v, _ := customCtx.Session.Get(sessionKeyAdministrator)
		if !v.IsNil() {
			var user *model.Administrator
			_ = v.Struct(&user)
			return user
		}
	}
	return &model.Administrator{}
}

// 删除用户Session。
func (s *serviceSession) RemoveAdministrator(ctx context.Context) error {
	customCtx := Context.Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(sessionKeyAdministrator)
	}
	return nil
}
