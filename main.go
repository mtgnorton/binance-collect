package main

import (
	cmd "gf-admin/app/system/admin"
	_ "gf-admin/boot"
	"github.com/gogf/gf/os/gctx"
)

//初始化admin用户

func main() {

	cmd.Run(gctx.New())
}
