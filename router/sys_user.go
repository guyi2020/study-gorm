package router

import (
	"github.com/gin-gonic/gin"
	v1 "study-gorm/api/v1"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("getUserList", v1.GetUserList)           // 分页获取用户列表
	}
}
