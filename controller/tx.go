package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"wallet-aa-tx-serv/client/clientdto"
	"wallet-aa-tx-serv/models"
	"wallet-aa-tx-serv/service"
	"wallet-aa-tx-serv/utils/gin2"
)

func GetTransaction(ctx *gin.Context) {
	chainName := ctx.Query("chainName")
	txHash := ctx.Query("txHash")
	opHash := ctx.Query("opHash")
	address := ctx.Query("address")
	status := ctx.Query("status")

	var chain *clientdto.Chain
	var err error
	if chainName != "" {
		chain, err = service.GetChainByName(chainName)
		if err != nil {
			gin2.HttpResponse(ctx, "", err)
			return
		}
	}
	statusUint, err := strconv.ParseUint(status, 10, 64)
	data, err := service.FindTransaction(&models.Transaction{
		UserOperationHash: opHash,
		TxHash:            txHash,
		Sender:            address,
		ChainId:           chain.ID,
		Status:            uint(statusUint),
	})
	if err != nil {
		gin2.HttpResponse(ctx, "", err)
		return
	}

	for i := range data {
		data[i].UserOperationJson = nil
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

func GetEstimateFee(ctx *gin.Context) {
	chainId := ctx.Query("chainId")

	if chainId == "" {
		gin2.HttpResponse(ctx, "", fmt.Errorf("chainId is empty"))
		return
	}

	networkIdUint, err := strconv.ParseInt(chainId, 10, 64)
	if err != nil {
		gin2.HttpResponse(ctx, "", err)
		return
	}
	fee, err := service.GetEstimateFee(int(networkIdUint))
	if err != nil {
		return
	}

	gin2.HttpResponse(ctx, fee, err)
}
