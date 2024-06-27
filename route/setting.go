package route

import (
	"github.com/backunderstar/zew/api"
	"github.com/gin-gonic/gin"
)

type SettingRouter struct{}

func (s *SettingRouter) InitRouter(Router *gin.RouterGroup) {
	settingGroup := Router.Group("setting")
	settingApi := api.ApiGroupApp.SettingsApi
	{
		settingGroup.GET("/:name", settingApi.SettingsInfo)
		settingGroup.PUT("/:name", settingApi.SettingsUpdate)
	}
}
