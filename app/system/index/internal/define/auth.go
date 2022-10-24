package define

import (
	"gf-admin/app/model"

	"github.com/gogf/gf/v2/frame/g"
)

type AuthRegisterHtmlReq struct {
	g.Meta `path:"/register-html" method:"get" tags:"注册" summary:"注册"`
}

type AuthRegisterHtmlRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type AuthRegisterReq struct {
	g.Meta `path:"/auth-register" method:"post" summary:"注册" tags:"用户管理"`
	*AuthRegisterInput
}

type AuthRegisterRes struct {
}

type AuthRegisterInput struct {
	Username  string `v:"required|passport#请输入用户名|用户名只能字母开头，只能包含字母、数字和下划线，长度在6~18之间" dc:"用户名" d:"username" json:"username"`
	Email     string `v:"required|email#请输入邮箱|邮箱格式错误" dc:"邮箱" d:"email" json:"email"`
	Password  string `v:"required|password#请输入密码|密码长度需要在长度在6~18之间" dc:"密码" d:"password" json:"password"`
	Password2 string `v:"required|same:password#请输入确认密码|两次密码不一致" dc:"确认密码" d:"password2" json:"password2"`
	Code      string `json:"code"  v:"" dc:"验证码"`
	CaptchaId string `json:"captcha_id" v:"" dc:"后端返回的captcha标识符"`
}

type AuthLoginHtmlReq struct {
	g.Meta `path:"/login-html" method:"get" tags:"登录" summary:"登录"`
}

type AuthLoginHtmlRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type AuthLoginCaptchaReq struct {
	g.Meta `path:"/auth-captcha-get" method:"get" summary:"获取登录验证码" tags:"a全局"`
}

type AuthLoginCaptchaRes struct {
	model.CommonGenerateCaptchaOutput
}

type AuthLoginReq struct {
	g.Meta `path:"/auth-login" method:"post" summary:"登录" tags:"用户管理"`
	*AuthLoginInput
}

type AuthLoginRes struct {
	*AuthLoginOutput
}

type AuthLoginInput struct {
	Username string `v:"required#请输入用户名" dc:"用户名" d:"username" json:"username"`
	Password string `v:"required#请输入密码" dc:"密码" d:"password" json:"password"`

	Code      string `json:"code"  v:"" dc:"验证码"`
	CaptchaId string `json:"captcha_id" v:"" dc:"后端返回的captcha标识符"`
}

type AuthInfoReq struct {
	g.Meta `path:"/auth-info" method:"get" summary:"获取用户信息" tags:"用户管理"`
}

type AuthInfoRes struct {
	*AuthInfoOutput
}

type AuthInfoOutput struct {
	model.UserInfoWithoutPass
}

type AuthLoginOutput struct {
	Token string `json:"token"`
}

type AuthLogoutReq struct {
	g.Meta `path:"/auth-logout" method:"get" summary:"退出登录" tags:"用户管理"`
}

type AuthLogoutRes struct {
}
