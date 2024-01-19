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
	if err != nil {
		return err
	}
	tx := &models.Transaction{
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

func FindTransactionNeededToCheckStatus() ([]models.Transaction, error) {
	return dao.FindTransactionNeededToCheckStatus()
}

func DeleteTransaction(strategyInfo *models.Transaction) error {
	return dao.DeleteTransaction(strategyInfo)
}

func UpdateTransaction(strategyInfo *models.Transaction) (*models.Transaction, error) {
	return dao.UpdateTransaction(strategyInfo)
}
