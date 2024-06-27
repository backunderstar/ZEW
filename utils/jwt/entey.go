package jwt

import (
	"github.com/golang-jwt/jwt"
)

type JwtPayLoad struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Role     int    `json:"role"`
	UserID   uint   `json:"user_id"`
}

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}
