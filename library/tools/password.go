package tools

import "github.com/gogf/gf/crypto/gmd5"

func GenPassword(pwd string) string {
	p := gmd5.MustEncryptString(pwd)
	return gmd5.MustEncryptString(p + "MctwGHs6BOdTfOiS")
}
