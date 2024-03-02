package routers

import (
	"github.com/gin-gonic/gin"
	"wallet-aa-tx-serv/controller"
)

func InitConfigCacheRouter(router *gin.RouterGroup) {
	userInfo := router.Group("/cache")
	userInfo.GET("/config", controller.FlushCacheConfig)
}
