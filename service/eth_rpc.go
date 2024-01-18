package service

import (
	"encoding/json"
	"io"
	"wallet-aa-tx-serv/models"
	"wallet-aa-tx-serv/utils/httplib"
)

func GetTransactionReceiptResponse(txHash string) (*models.GetTransactionReceiptResponseData, error) {
	// FIXME: 需要从Chain表中获取Api
	url := "https://mumbai-rpc.web3idea.xyz"
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

	res := models.GetTransactionReceiptResponseData{}
	// FIXME: 这里使用httplib.PostInto就不用自己Unmarshal了
	resPost, err := httplib.Post(
		url,
		data,
		header,
	)
	defer resPost.Body.Close()
	body, err := io.ReadAll(resPost.Body)
	json.Unmarshal(body, &res)

	return &res, err
}
