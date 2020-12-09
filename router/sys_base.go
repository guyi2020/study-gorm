package router

import (
	"github.com/gin-gonic/gin"
	v1 "study-gorm/api/v1"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("login", v1.Login)
	}
	return BaseRouter
}
