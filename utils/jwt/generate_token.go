package jwt

import (
	"time"

	"github.com/backunderstar/zew/global"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(user JwtPayLoad) (string, error) {
	var Secret = []byte(global.Config.Jwt.Secret)
	claim := CustomClaims{
		JwtPayLoad: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires)).Unix(),
			Issuer:    global.Config.Jwt.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(Secret)
}
