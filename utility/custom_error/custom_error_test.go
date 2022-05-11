package custom_error

import (
	"bytes"
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gutil"
	"testing"
)

var ctx = gctx.New()

func Test_New(t *testing.T) {

	//用户自定义错误
	gtest.C(t, func(t *gtest.T) {
		e := New("自定义错误", map[string]interface{}{"token": "123456"}, []string{"token"})
		catchError(e)
	})

	gtest.C(t, func(t *gtest.T) {
		e := New("自定义错误", map[string]interface{}{"token": "123456"}, []string{"token"})
		ew := Wrap(e, "自定义错误外层", map[string]interface{}{"token": "fffff"}, []string{"fff"})
		catchError(ew)
	})

	gtest.C(t, func(t *gtest.T) {
		e := errors.New("原始错误")

		token := "sdfafdafasdf"
		ew := Wrap(e, "自定义错误", map[string]interface{}{"token": token})
		catchError(ew)
	})

	gtest.C(t, func(t *gtest.T) {

		err := gerror.New("goframe 错误")
		ew := Wrap(err, "")
		catchError(ew)

	})
}

func catchError(err error) {
	code := gerror.Code(err)
	//if code == gcode.CodeNil && err != nil {
	//	code = gcode.CodeInternalError
	//}
	//堆栈日志
	g.Log().Infof(ctx, "error occur:%+v", gerror.Stack(err))

	buffer := &bytes.Buffer{}

	g.DumpTo(buffer, code.Detail(), gutil.DumpOption{})

	//上下文信息
	g.Log().Printf(ctx, "context info:%+v", buffer.String())

	//用户信息
	g.Dump(gerror.Current(err))

}
