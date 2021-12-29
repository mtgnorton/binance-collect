package define

import "github.com/gogf/gf/v2/frame/g"


type LoginPostReq struct {
	g.Meta   `path:"/login" method:"post" summary:"执行登录请求" tags:"登录-aaaa"`
	Username string `json:"username" v:"required#请输入账号"   dc:"账号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
	// Captcha  string `json:"captcha"  v:"required#请输入验证码" dc:"验证码"`
}
type LoginPostRes struct {
	Message string `json:"message" dc:"账号或密码错误|登录成功"`
}

type LoginInput struct {
	Username string
	Password string
}
