package service

import (
	"bytes"
	"encoding/json"
	"go-gin-gorm-starter/dao"
	"go-gin-gorm-starter/global"
	"go-gin-gorm-starter/models"
	"net/http"
)

func SaveTransaction(strategyInfo *models.SavedTr) error {
	if strategyInfo.UserOperationHash == "" {
		return global.OtherError("UserOperationHash is empty")
	}

	//TODO：查询chain表获取chainId
	var chain models.Chain
	err := dao.GetChainByNetworkId(&chain, strategyInfo.NetworkId)

	type Param struct {
		JsonRpc string   `json:"jsonrpc"`
		Id      uint     `json:"id"`
		Method  string   `json:"method"`
		Params  []string `json:"params"`
	}

	param := Param{
		JsonRpc: "2.0",
		Id:      strategyInfo.NetworkId,
		Method:  "eth_getUserOperationByHash",
		Params:  []string{strategyInfo.UserOperationHash},
	}
	jsonParam, err := json.Marshal(param)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", chain.BundlerApi, bytes.NewBuffer(jsonParam))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	type Result struct {
		UserOperation models.UserOperation `json:"userOperation"`
	}
	type Response struct {
		Id              uint   `json:"id"`
		JsonRpc         string `json:"jsonrpc"`
		Result          Result `json:"result"`
		Entrypoint      string `json:"entrypoint"`
		BlockNumber     uint   `json:"blockNumber"`
		BlockHash       string `json:"blockHash"`
		TransactionHash string `json:"transactionHash"`
	}
	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}
	u, err := json.Marshal(response.Result.UserOperation)
	if err != nil {
		panic(err)
	}

	tx := &models.Transaction{
		BlockHash:         response.BlockHash,
		BlockNumber:       response.BlockNumber,
		TxHash:            response.TransactionHash,
		Sender:            response.Result.UserOperation.Sender,
		EntryPointAddress: response.Entrypoint,
		UserOperation:     json.RawMessage(u),
		Type:              strategyInfo.Type,
		Status:            0,
		ChainId:           chain.ID,
	}
	return dao.SaveTransaction(tx)
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
