package define

import "github.com/gogf/gf/v2/frame/g"

type NoAuthLoginCaptchaReq struct {
	g.Meta `path:"/captcha-get" method:"get" summary:"获取登录验证码" tags:"a全局"`
}

type NoAuthLoginCaptchaRes struct {
	CommonGenerateCaptchaOutput
}
type CommonGenerateCaptchaOutput struct {
	CaptchaId     string `json:"captcha_id"`
	CaptchaBase64 string `json:"captcha_base64""`
}
