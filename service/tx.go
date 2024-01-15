package service

import (
	"fmt"
	"go-gin-gorm-starter/dao"
	"go-gin-gorm-starter/global"
	"go-gin-gorm-starter/models"
	"go-gin-gorm-starter/utils/httplib"
)

type GetOPByHashResult struct {
	Entrypoint      string                `json:"entrypoint"`
	BlockNumber     uint                  `json:"blockNumber"`
	BlockHash       string                `json:"blockHash"`
	TransactionHash string                `json:"transactionHash"`
	UserOperation   *models.UserOperation `json:"userOperation"`
}

type GetOPByHashResponse struct {
	Id      uint               `json:"id"`
	JsonRpc string             `json:"jsonrpc"`
	Result  *GetOPByHashResult `json:"result"`
}

type BundlerRequestParam struct {
	JsonRpc string   `json:"jsonrpc"`
	Id      uint     `json:"id"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
}

func SaveTransaction(info *models.SavedTr) error {
	if info.UserOperationHash == "" {
		return global.OtherError("UserOperationHash is empty")
	}

	var chain models.Chain
	err := dao.GetChainByNetworkId(&chain, info.NetworkId)

	resInDb, err := dao.FindTransaction(&models.Transaction{
		ChainId:           chain.ID,
		UserOperationHash: info.UserOperationHash,
	})
	if err != nil {
		return err
	}
	if len(resInDb) > 0 {
		return global.OtherError("UserOperationHash already exists")
	}

	response, err := getUserOperationByHash(chain.BundlerApi, info.UserOperationHash)

	tx := &models.Transaction{
		BlockHash:         response.Result.BlockHash,
		BlockNumber:       response.Result.BlockNumber,
		TxHash:            response.Result.TransactionHash,
		Sender:            response.Result.UserOperation.Sender,
		EntryPointAddress: response.Result.Entrypoint,
		UserOperation:     response.Result.UserOperation,
		UserOperationHash: info.UserOperationHash,
		Type:              info.Type,
		Status:            models.TransactionStatusInit,
		ChainId:           chain.ID,
	}
	return dao.SaveTransaction(tx)
}

func getUserOperationByHash(api string, hash string) (*GetOPByHashResponse, error) {
	param := BundlerRequestParam{
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

func FindTransaction(strategyInfo *models.Transaction) ([]models.Transaction, error) {
	return dao.FindTransaction(strategyInfo)
}

func DeleteTransaction(strategyInfo *models.Transaction) error {
	return dao.DeleteTransaction(strategyInfo)
}

func UpdateTransaction(strategyInfo *models.Transaction) (*models.Transaction, error) {
	return dao.UpdateTransaction(strategyInfo)
}
