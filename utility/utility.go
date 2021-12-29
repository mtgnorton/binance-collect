package utility

import "github.com/gogf/gf/v2/crypto/gmd5"

func EncryptPassword(username, password string) string {
	return gmd5.MustEncrypt(username + password)
}
