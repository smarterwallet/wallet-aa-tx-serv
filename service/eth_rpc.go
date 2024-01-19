package service

import (
	"wallet-aa-tx-serv/models"
	"wallet-aa-tx-serv/utils/httplib"
)

func GetTransactionByHashResponse(rpcUrl string, txHash string) (*models.EthRpcResponseData, error) {
	header := map[string]string{
		"Content-Type": "application/json",
	}
	hash := []string{txHash}
	data := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "eth_getTransactionByHash",
		"params":  hash,
	}

	res := &models.EthRpcResponseData{}
	resPost, err := httplib.PostInto(
		rpcUrl,
		data,
		header,
		res,
	)
	defer resPost.Body.Close()

	return res, err
}

func GetTransactionReceiptResponse(rpcUrl string, txHash string) (*models.EthRpcResponseData, error) {
	header := map[string]string{
		"Content-Type": "application/json",
	}
	hash := []string{txHash}
	data := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "eth_getTransactionReceipt",
		"params":  hash,
	}

	res := &models.EthRpcResponseData{}
	resPost, err := httplib.PostInto(
		rpcUrl,
		data,
		header,
		res,
	)
	defer resPost.Body.Close()

	return res, err
}

func GetUserOperationByHashResponse(rpcUrl string, userOperationHash string) (*models.EthRpcResponseData, error) {
	header := map[string]string{
		"Content-Type": "application/json",
	}
	hash := []string{userOperationHash}
	data := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "eth_getUserOperationByHash",
		"params":  hash,
	}

	res := &models.EthRpcResponseData{}
	resPost, err := httplib.PostInto(
		rpcUrl,
		data,
		header,
		res,
	)
	defer resPost.Body.Close()

	return res, err
}
