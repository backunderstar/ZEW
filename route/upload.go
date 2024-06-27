package route

import (
	"github.com/backunderstar/zew/api"
	"github.com/gin-gonic/gin"
)

type UploadRouter struct{}

func (u *UploadRouter) InitRouter(Router *gin.RouterGroup) {
	uploadRouter := Router.Group("upload")
	uploadApi := api.ApiGroupApp.UploadApi
	{
		uploadRouter.POST("/image", uploadApi.UploadImage)
	}
}