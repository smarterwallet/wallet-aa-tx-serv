package service

import (
	"errors"
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

	res := &models.EthRpcResponseData{
		Result: &models.GetTransactionByHashResult{},
	}
	resPost, err := httplib.PostInto(
		rpcUrl,
		data,
		header,
		res,
	)
	defer resPost.Body.Close()

	if res.Error.Code != models.EthRpcResponseErrorIsFalse {
		return res, errors.New("fail to GetTransactionByHashResponse")
	}

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

	res := &models.EthRpcResponseData{
		Result: &models.GetTransactionReceiptResult{},
	}
	_, err := httplib.PostInto(
		rpcUrl,
		data,
		header,
		res,
	)

	if res.Error.Code != models.EthRpcResponseErrorIsFalse {
		return res, errors.New("fail to GetTransactionReceiptResponse. Tx hash: " + txHash)
	}

	return res, err
}

func GetUserOperationByHashResponse(bundlerUrl string, userOperationHash string) (*models.EthRpcResponseData, error) {
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

	res := &models.EthRpcResponseData{
		Result: &models.GetUserOperationByHashResult{},
	}
	_, err := httplib.PostInto(
		bundlerUrl,
		data,
		header,
		res,
	)

	if res.Error.Code != models.EthRpcResponseErrorIsFalse {
		return res, errors.New("fail to GetUserOperationByHashResponse")
	}

	return res, err
}
