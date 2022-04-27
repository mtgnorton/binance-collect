package service

import (
	"context"
	"gf-admin/app/system/admin/internal/define"
	"github.com/mojocn/base64Captcha"
)

var Common = common{}

type common struct {
}

func (c *common) GenerateCaptcha(ctx context.Context) (out define.CommonGenerateCaptchaOutput, err error) {
	var store = base64Captcha.DefaultMemStore

	driver := base64Captcha.NewDriverDigit(50, 120, 5, 0, 50)

	instance := base64Captcha.NewCaptcha(driver, store)

	id, b64s, err := instance.Generate()

	return define.CommonGenerateCaptchaOutput{
		CaptchaId:     id,
		CaptchaBase64: b64s,
	}, err

}

func (c *common) VerifyCaptcha(ctx context.Context, code string, id string) bool {
	var store = base64Captcha.DefaultMemStore
	return store.Verify(id, code, true)

}
