package service

import (
	"fmt"
	"wallet-aa-tx-serv/global"
	"wallet-aa-tx-serv/models"
	"wallet-aa-tx-serv/utils/httplib"
)

type getOPByHashResult struct {
	Entrypoint      string                `json:"entrypoint"`
	BlockNumber     uint                  `json:"blockNumber"`
	BlockHash       string                `json:"blockHash"`
	TransactionHash string                `json:"transactionHash"`
	UserOperation   *models.UserOperation `json:"userOperation"`
}

type GetOPByHashResponse struct {
	Id      uint               `json:"id"`
	JsonRpc string             `json:"jsonrpc"`
	Result  *getOPByHashResult `json:"result"`
}

type bundlerRequestParam struct {
	JsonRpc string   `json:"jsonrpc"`
	Id      uint     `json:"id"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
}

func GetUserOperationByHash(api string, hash string) (*GetOPByHashResponse, error) {
	param := bundlerRequestParam{
		JsonRpc: "2.0",
		Id:      1,
		Method:  "eth_getUserOperationByHash",
		Params:  []string{hash},
	}

	response := &GetOPByHashResponse{}
	res, err := httplib.PostInto(api, param, map[string]string{}, &response)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, global.OtherError(fmt.Sprintf("bundler api response status is %v", res.StatusCode))
	}

	if response == nil || response.Result == nil || response.Result.UserOperation == nil {
		return nil, global.OtherError("UserOperation is empty or expired")
	}

	return response, nil
}
