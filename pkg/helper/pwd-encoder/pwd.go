package pwd_encoder

import (
	"crypto/md5"
	"github.com/anaskhan96/go-password-encoder"
)

const (
	SaltLen    = 16
	Iterations = 100
	KeyLen     = 32
)

//加密
func PwdEncode(pass string) (salt, encodePwd string) {
	//加密
	options := password.Options{
		SaltLen:      SaltLen,
		Iterations:   Iterations,
		KeyLen:       KeyLen,
		HashFunction: md5.New,
	}
	salt, encodePwd = password.Encode(pass, &options)
	return
}

//检查密码
func PwdDecode(rawPwd, encodedPwd, salt string) bool {
	options := password.Options{
		SaltLen:      SaltLen,
		Iterations:   Iterations,
		KeyLen:       KeyLen,
		HashFunction: md5.New,
	}
	return password.Verify(rawPwd, salt, encodedPwd, &options)
}
