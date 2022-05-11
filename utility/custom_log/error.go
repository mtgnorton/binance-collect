package custom_log

import (
	"bytes"
	"context"
	"gf-admin/utility/custom_error"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gutil"
	"strings"
)

func Log(r *ghttp.Request, err error) {

	code := gerror.Code(err)

	buffer := &bytes.Buffer{}

	g.DumpTo(buffer, code.Detail(), gutil.DumpOption{})

	//g.Log().Warningf()(r.Context(), "request params :%+v", r.GetMap())

	//stackContentSlice := strings.Split(gerror.Stack(err), "\n")
	//
	//finalStackContent := ""
	//if len(stackContentSlice) > 0 {
	//	for _, s := range stackContentSlice { //去除无用gf源码堆栈
	//		if !strings.Contains(s, "/ghttp/") {
	//			finalStackContent += s + "\n"
	//		}
	//	}
	//}
	//堆栈日志和上下文信息
	g.Log().Warningf(r.Context(), "context variables :%+v.\n error stack :%+v", buffer.String(), custom_error.Stack(err))

}

func ErrorLog(err error, loggerTypeSlice ...string) {
	loggerType := ""
	if len(loggerTypeSlice) > 0 {
		loggerType = loggerTypeSlice[0]
	}
	stackContentSlice := strings.Split(gerror.Stack(err), "\n")

	finalStackContent := ""
	if len(stackContentSlice) > 0 {
		for _, s := range stackContentSlice { //去除无用gf源码堆栈
			if !strings.Contains(s, "/v2/") && !strings.Contains(s, "/v2@") {
				finalStackContent += s + "\n"
			}
		}
	}

	g.Log(loggerType).Warning(context.TODO(), finalStackContent)
}
