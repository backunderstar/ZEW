package route

import (
	"github.com/backunderstar/zew/api"
	"github.com/gin-gonic/gin"
)

type LoginRouter struct{}

func (r *LoginRouter) InitRouter(Router *gin.RouterGroup) {
	loginRouuter := Router.Group("login")
	loginApi := api.ApiGroupApp.LoginApi
	{
		loginRouuter.POST("/email", loginApi.EmailLogin)
	}
}