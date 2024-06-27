package pwd

import (
	"github.com/backunderstar/zew/global"
	"golang.org/x/crypto/bcrypt"
)

// Hash 密码
func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		global.Log.Error(err)
	}
	return string(hash)
}

// 验证
func Verify(hashpwd string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpwd), []byte(pwd))
	if err != nil {
		global.Log.Error(err)
		return false
	}
	return true
}
