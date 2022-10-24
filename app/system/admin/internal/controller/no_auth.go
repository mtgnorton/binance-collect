package controller

import (
	"context"
	"gf-admin/app/shared"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
)

var NoAuth = noAuth{}

type noAuth struct {
}

func (l *noAuth) LoginInfo(ctx context.Context, req *define.PersonalLoginInfoReq) (res *define.PersonalLoginInfoRes, err error) {

	res = &define.PersonalLoginInfoRes{}
	res.PersonalLoginInfoOutput, err = service.Personal.LoginInfo(ctx)
	return
}

func (l *noAuth) Login(ctx context.Context, req *define.PersonalLoginReq) (res *define.PersonalLoginRes, err error) {
	res = &define.PersonalLoginRes{}

	res.PersonalLoginOutput, err = service.Personal.Login(ctx, req.PersonalLoginInput)
	return

}

// 图形验证码
func (l *noAuth) LoginCaptcha(ctx context.Context, req *define.NoAuthLoginCaptchaReq) (res *define.NoAuthLoginCaptchaRes, err error) {

	res = &define.NoAuthLoginCaptchaRes{}
	res.CommonGenerateCaptchaOutput, err = shared.Captcha.Generate(ctx)
	return
}
