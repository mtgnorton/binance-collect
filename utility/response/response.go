package response

import (
	"gf-admin/boot"
	"gf-admin/utility/logging"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gvalid"
)

// JsonRes 数据返回通用JSON数据结构
type JsonRes struct {
	Code     int         `json:"code"`     // 错误码((0:成功, 1:失败, >1:错误码))
	Message  string      `json:"message"`  // 提示信息
	Data     interface{} `json:"data"`     // 返回数据(业务接口定义具体数据结构)
	Redirect string      `json:"redirect"` // 引导客户端跳转到指定路由
}

// Json 返回标准JSON数据。
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	} else {
		responseData = g.Map{}
	}

	r.Response.WriteJson(JsonRes{
		Code:    code,
		Message: message,
		Data:    responseData,
	})

}

//可能的错误类型
// 1. go 原生或第三方代码返回的最基础错误 -> 无状态码
// 2. gf 代码返回的错误 -> 有状态码
// 3. 显式使用gerror.New()生成的错误 ->无状态码
// 4. 显式使用gcode.Code()的错误 ->有状态码,可能为gcode.CodeInvalidParameter,gcode.CodeNotAuthorized

// 无状态码默认使用gcode.CodeInternalError
// 在非生成环境下,给前端返回所有错误的相关信息，在生产环境下只返回gcode.CodeInvalidParameter和gcode.CodeNotAuthorized和gvalid.Error对应的具体错误，其他全部显示为系统内部错误
func JsonErrorLogExit(r *ghttp.Request, err error, codeSlice ...gcode.Code) {

	message := ""
	var code gcode.Code

	if len(codeSlice) > 0 {
		code = codeSlice[0]
	}
	if code == gcode.CodeNil || code == nil {
		if code = gerror.Code(err); code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
	}

	_, isGvalid := err.(gvalid.Error)

	if boot.EnvName == boot.ENV_PROD {

		if code == gcode.CodeInvalidParameter || code == gcode.CodeNotAuthorized || isGvalid {
			message = gerror.Current(err).Error()
		} else {
			message = "系统内部错误"
		}
	} else {
		message = gerror.Current(err).Error()
	}

	go logging.ErrorLog(err)

	r.Response.ClearBuffer()
	JsonExit(r, code.Code(), message)
}

// JsonExit 返回标准JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, code int, message string, data ...interface{}) {
	Json(r, code, message, data...)
	r.Exit()
}

// JsonRedirect 返回标准JSON数据引导客户端跳转。
func JsonRedirect(r *ghttp.Request, code int, message, redirect string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(JsonRes{
		Code:     code,
		Message:  message,
		Data:     responseData,
		Redirect: redirect,
	})
}

// JsonRedirectExit 返回标准JSON数据引导客户端跳转，并退出当前HTTP执行函数。
func JsonRedirectExit(r *ghttp.Request, code int, message, redirect string, data ...interface{}) {
	JsonRedirect(r, code, message, redirect, data...)
	r.Exit()
}
