package controller

import (
	"context"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
)

var NoAuth = noAuth{}

type noAuth struct {
}

func (l *noAuth) Login(ctx context.Context, req *define.PersonalLoginPostReq) (res *define.PersonalLoginPostRes, err error) {
	res = &define.PersonalLoginPostRes{}

	res.PersonalLoginPostOutput, err = service.Personal.Login(ctx, req.PersonalLoginPostInput)
	return
}

// 图形验证码
func (l *noAuth) LoginCaptcha(ctx context.Context, req *define.NoAuthLoginCaptchaReq) (res *define.NoAuthLoginCaptchaRes, err error) {

	res = &define.NoAuthLoginCaptchaRes{}
	res.CommonGenerateCaptchaOutput, err = service.Common.GenerateCaptcha(ctx)
	return
}
