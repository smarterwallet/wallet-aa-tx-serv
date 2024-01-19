package routers

import (
	"github.com/gin-gonic/gin"
	"wallet-aa-tx-serv/controller"
)

func InitStrategyInfoRouter(router *gin.RouterGroup) {
	userInfo := router.Group("/aa-tx")
	userInfo.POST("/", controller.SaveTransaction)
	userInfo.GET("/", controller.GetTransaction)
	userInfo.DELETE("/:id", controller.DeleteTransaction)
}
