package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	//PublicGroup := Router.Group("")
	//{
	//	router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	//}
	Router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "first example.")
	})
	return Router
}
