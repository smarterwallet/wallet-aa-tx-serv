package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wallet-aa-tx-serv/models"
	"wallet-aa-tx-serv/service"
	"wallet-aa-tx-serv/utils/gin2"
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
		reqUser models.SavedTransaction
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

func GetStatusOfUserSendingTransaction(ctx *gin.Context) {
	txHash := ctx.Param("txHash")

	res, err := service.GetTransactionReceiptResponse(txHash)
	if err != nil {
		gin2.HttpResponse(ctx, "", err)
		return
	}

	gin2.HttpResponse(ctx, res, err)
}
