package utility

import (
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
)

func EncryptPassword(username, password string) string {
	return gmd5.MustEncrypt(username + password)
}

func GetServerPath() string {
	return gfile.Join(gfile.Pwd(), g.Cfg().MustGet(gctx.New(), "server.ServerRoot").String())
}
