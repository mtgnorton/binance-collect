package logging

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
)

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
