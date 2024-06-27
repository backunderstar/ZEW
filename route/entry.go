package route

import (
	"github.com/backunderstar/zew/global"
	"github.com/gin-gonic/gin"
)

func WebGroups(r *gin.Engine) WebRouter {
	// 设置自定义的多模板渲染器
	r.HTMLRender = createMyRender()
	return WebRouter{r}
}

func ApiGroups() []CommonRouter {
	return commonGroups()
}

func Routers() *gin.Engine {
	// 设置gin的模式
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	// 初始化api路由
	apiRouterGroup := router.Group("/api")
	var ApiGroupApp = ApiGroups()
	for _, router := range ApiGroupApp {
		router.InitRouter(apiRouterGroup)
	}

	// 初始化web路由
	webRouterGroup := router.Group("/")
	WebGroupApp := WebGroups(router)
	WebGroupApp.InitRouter(webRouterGroup)

	return router
}
