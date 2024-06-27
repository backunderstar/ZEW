package middleware

import (
	"github.com/backunderstar/zew/models/ctype"
	"github.com/backunderstar/zew/utils/jwt"
	"github.com/backunderstar/zew/utils/res"
	"github.com/gin-gonic/gin"
)

func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			res.FailWithMsg("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			res.FailWithMsg("token错误", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			res.FailWithMsg("未携带token", c)
			c.Abort()
			return
		}
		
		claims, err := jwt.ParseToken(token)
		if err != nil {
			res.FailWithMsg("token错误", c)
			c.Abort()
			return
		}

		if ctype.Role(claims.Role) != ctype.Admin {
			res.FailWithMsg("权限不足", c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
	}
}
