package route

import (
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

type WebRouter struct {
	*gin.Engine
}

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/dist/index.html")
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}

func (w *WebRouter) InitRouter(r *gin.RouterGroup) {
	// 设置自定义的多模板渲染器
	w.HTMLRender = createMyRender()

	// 配置静态文件服务
	r.Static("/assets", "./web/admin/dist/assets")
	r.Static("/statics", "./web/front/dist/statics")

	// 路由处理：将所有以"/home"开头的路由映射到admin模板
	r.GET("/home/*any", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin", gin.H{})
	})

	// 未匹配到任何路由时，返回front模板
	w.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "front", gin.H{})
	})
}
