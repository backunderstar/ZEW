package jwt

import (
	"errors"
	"fmt"

	"github.com/backunderstar/zew/global"
	"github.com/golang-jwt/jwt"
)

func ParseToken(tokenString string) (*CustomClaims, error) {
	var Secret = []byte(global.Config.Jwt.Secret)
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		global.Log.Error(fmt.Sprintf("token解析失败: %s", err.Error()))
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token验证错误")
}
