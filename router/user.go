package router

import (
	"github.com/LRboyz/gin-api/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("login", controller.PasswordLogin)
		UserRouter.POST("list", controller.GetUserList)
	}
}
