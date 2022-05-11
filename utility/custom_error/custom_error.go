package custom_error

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

//自定义错误状态码
const CustomCodeNumber = 1000

func New(message string, contextArgs ...interface{}) error {
	if message == "" {
		message = "未知错误"
	}
	code := gcode.New(CustomCodeNumber, message, contextArgs)
	return gerror.NewCodeSkip(code, 1)
}

func Wrap(err error, message string, contextArgs ...interface{}) error {
	if message == "" {
		message = "未知错误"
	}
	code := gcode.New(CustomCodeNumber, message, contextArgs)

	return gerror.WrapCodeSkip(code, 1, err)
}

// 获取嵌套最深的gerror的堆栈信息
func Stack(err error) string {
	message := "\n0. " + err.Error() //先获取完整的错误信息
	layer := 0
	for {
		e, ok := err.(*gerror.Error)
		if !ok {
			if layer > 0 {
				return message + "\n" + gerror.Stack(err)
			}
			return gerror.Stack(err)
		}
		if _, ok := e.Next().(*gerror.Error); !ok {
			if layer > 0 {
				return message + "\n" + gerror.Stack(err)
			}
			return gerror.Stack(err)
		}
		layer++
		err = e.Next()
	}

}
