package deposit_withdraw

import (
	"bytes"
	"context"
	"gf-admin/utility/custom_error"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gutil"
)

func LogErrorfDw(ctx context.Context, err error) {

	code := gerror.Code(err)

	buffer := &bytes.Buffer{}

	g.DumpTo(buffer, code.Detail(), gutil.DumpOption{})

	g.Log("deposit_withdraw").Warningf(ctx, "context variables :%+v.\n error stack :%+v", buffer.String(), custom_error.Stack(err))

}

func LogInfofDw(ctx context.Context, format string, v ...any) {
	g.Log("deposit_withdraw").Infof(ctx, format, v...)
}
