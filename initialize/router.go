package initialize

import (
	"github.com/LRboyz/gin-api/middleware"
	"github.com/LRboyz/gin-api/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	// 加入中间件
	Router.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	// 路由分组
	ApiGroup := Router.Group("/v1")
	router.UserRouter(ApiGroup)
	return Router
}
