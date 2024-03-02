package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
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

	var chain models.Chain
	if chainName != "" {
		chains, err := service.FindChain(&models.Chain{
			Name: chainName,
		})
		if err != nil {
			gin2.HttpResponse(ctx, "", err)
			return
		}
		if len(chains) == 0 {
			gin2.HttpResponse(ctx, "", fmt.Errorf("chain not found"))
			return
		}
		chain = chains[0]
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
	networkId := ctx.Query("networkId")

	if networkId == "" {
		gin2.HttpResponse(ctx, "", fmt.Errorf("networkId is empty"))
		return
	}

	networkIdUint, err := strconv.ParseUint(networkId, 10, 64)
	if err != nil {
		gin2.HttpResponse(ctx, "", err)
		return
	}
	fee, err := service.GetEstimateFee(networkIdUint)
	if err != nil {
		return
	}

	gin2.HttpResponse(ctx, fee, err)
}
