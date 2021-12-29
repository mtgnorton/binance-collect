package error_code

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

//用户输入错误码
const ParamErrorCode = 1001

func NewParams(params ...string) error {
	var msg string
	var detail interface{}
	if len(params) > 0 {
		msg = params[0]
	}
	if len(params) > 1 {
		detail = params[1]
	}
	return gerror.NewCode(gcode.New(ParamErrorCode, msg, detail))
}
