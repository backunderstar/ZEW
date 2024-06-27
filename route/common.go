package route

import "github.com/gin-gonic/gin"

type CommonRouter interface {
	InitRouter(Router *gin.RouterGroup)
}

func commonGroups() []CommonRouter {
	return []CommonRouter{
		&LoginRouter{},
		&SettingRouter{},
		&UploadRouter{},
		&UserRouter{},
	}
}