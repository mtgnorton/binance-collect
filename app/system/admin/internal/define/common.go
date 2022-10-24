package define

import (
	"gf-admin/app/model"

	"github.com/gogf/gf/v2/frame/g"
)

type NoAuthLoginCaptchaReq struct {
	g.Meta `path:"/captcha-get" method:"get" summary:"获取登录验证码" tags:"a全局"`
}

type NoAuthLoginCaptchaRes struct {
	model.CommonGenerateCaptchaOutput
}
