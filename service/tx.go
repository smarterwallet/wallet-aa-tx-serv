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
	if strategyInfo.UserOperationHash == "" || strategyInfo.Type == 0 {
		return global.OtherError("Message is empty")
	}

	type Param struct {
		Jsonrpc string   `json:"jsonrpc"`
		Id      uint     `json:"id"`
		Method  string   `json:"method"`
		Params  []string `json:"params"`
	}

	url := "https://smarter-api-mumbai.web3idea.xyz/bundler/"
	param := Param{
		Jsonrpc: "2.0",
		Id:      strategyInfo.Type,
		Method:  "eth_getUserOperationByHash",
		Params:  []string{strategyInfo.UserOperationHash},
	}
	jsonParam, err := json.Marshal(param)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonParam))
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
		Jsonrpc         string `json:"jsonrpc"`
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
		Type:              response.Id,
		Status:            0,
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
