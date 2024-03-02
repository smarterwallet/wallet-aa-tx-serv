package controller

import (
	"github.com/gin-gonic/gin"
	"wallet-aa-tx-serv/service"
	"wallet-aa-tx-serv/utils/gin2"
)

func FlushCacheConfig(ctx *gin.Context) {
	err := service.FlushCacheConfig()
	if err != nil {
		gin2.HttpResponse(ctx, "", err)
		return
	}

	gin2.HttpResponse(ctx, err == nil, err)
}
