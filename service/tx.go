package service

import (
	"wallet-aa-tx-serv/dao"
	"wallet-aa-tx-serv/global"
	"wallet-aa-tx-serv/models"
)

func SaveTransaction(info *models.SavedTransaction) error {
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

	response, err := GetUserOperationByHash(chain.BundlerApi, info.UserOperationHash)

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

func FindTransaction(strategyInfo *models.Transaction) ([]models.Transaction, error) {
	return dao.FindTransaction(strategyInfo)
}

func DeleteTransaction(strategyInfo *models.Transaction) error {
	return dao.DeleteTransaction(strategyInfo)
}

func UpdateTransaction(strategyInfo *models.Transaction) (*models.Transaction, error) {
	return dao.UpdateTransaction(strategyInfo)
}
