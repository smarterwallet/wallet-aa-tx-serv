package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin-gorm-starter/models"
	"go-gin-gorm-starter/service"
	"go-gin-gorm-starter/utils/gin2"
	"strconv"
)

func GetTransaction(ctx *gin.Context) {
	txHash := ctx.Query("hash")
	address := ctx.Query("address")

	data, err := service.FindTransaction(&models.Transaction{
		TxHash: txHash,
		Sender: address,
	})
	if err != nil {
		gin2.HttpResponse(ctx, "", err)
		return
	}

	gin2.HttpResponse(ctx, data, err)
}

func SaveTransaction(ctx *gin.Context) {
	var (
		reqUser models.SavedTr
	)

	ctx.Bind(&reqUser)

	err := service.SaveTransaction(&reqUser)
	if err != nil {
		gin2.HttpResponse(ctx, "", err)
		return
	}

	gin2.HttpResponse(ctx, err == nil, err)
}

func DeleteTransaction(ctx *gin.Context) {
	var (
		params models.Transaction
	)

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		gin2.HttpResponse(ctx, "", err)
		return
	}

	params.ID = uint(id)

	err = service.DeleteTransaction(&params)
	if err != nil {
		gin2.HttpResponse(ctx, "", err)
		return
	}

	gin2.HttpResponse(ctx, err == nil, err)
}
