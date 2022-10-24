package controller

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/app/system/index/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

var NoAuth = noAuth{}

type noAuth struct {
}

func (n *noAuth) RegisterHtml(ctx context.Context, req *define.AuthRegisterHtmlReq) (res *define.AuthRegisterHtmlRes, err error) {

	shared.View().Render(ctx, model.View{
		Title:   "注册",
		MainTpl: "/register.html",
		Data: g.Map{
			"captchaInfo": model.CommonGenerateCaptchaOutput{
				CaptchaId:     "",
				CaptchaBase64: "",
			},
		},
	})
	return
}

func (n *noAuth) LoginHtml(ctx context.Context, req *define.AuthLoginHtmlReq) (res *define.AuthLoginHtmlRes, err error) {

	shared.View().Render(ctx, model.View{
		Title:   "登录",
		MainTpl: "/login.html",
		Data: g.Map{
			"captchaInfo": model.CommonGenerateCaptchaOutput{
				CaptchaId:     "",
				CaptchaBase64: "",
			},
		},
	})
	return
}

// 图形验证码
func (l *noAuth) LoginCaptcha(ctx context.Context, req *define.AuthLoginCaptchaReq) (res *define.AuthLoginCaptchaRes, err error) {
	res = &define.AuthLoginCaptchaRes{}
	res.CommonGenerateCaptchaOutput, err = shared.Captcha.Generate(ctx)
	return
}

func (n *noAuth) Register(ctx context.Context, req *define.AuthRegisterReq) (res *define.AuthRegisterRes, err error) {
	res = &define.AuthRegisterRes{}
	err = service.Auth.Register(ctx, req.AuthRegisterInput)
	return
}

func (n *noAuth) Login(ctx context.Context, req *define.AuthLoginReq) (res *define.AuthLoginRes, err error) {
	res = &define.AuthLoginRes{}
	res.AuthLoginOutput, err = service.Auth.Login(ctx, req.AuthLoginInput)
	return
}

func (p *noAuth) PostsDetail(ctx context.Context, req *define.PostsDetailReq) (res *define.PostsDetailRes, err error) {
	user, err := service.FrontTokenInstance.GetUser(ctx)
	if err != nil {
		return
	}
	shared.View().Render(ctx, model.View{
		Title:   "详情",
		MainTpl: "posts-detail.html",
		User:    user,
		Data: g.Map{
			"comments": []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
	})
	return
}
