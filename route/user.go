package route

import (
	"github.com/backunderstar/zew/api"
	"github.com/backunderstar/zew/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (r *UserRouter) InitRouter(Router *gin.RouterGroup) {
	userApi := api.ApiGroupApp.UsersApi
	userRouter := Router.Group("users")
	userRouter.Use(middleware.AdminAuth())
	{
		userRouter.GET("", userApi.UsersList)
	}
}