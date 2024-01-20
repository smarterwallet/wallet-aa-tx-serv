package service

import (
	"wallet-aa-tx-serv/dao"
	"wallet-aa-tx-serv/global"
	"wallet-aa-tx-serv/models"
)

func SaveTransaction(transaction *models.SavedTransaction) error {
	if transaction.UserOperationHash == "" {
		return global.OtherError("UserOperationHash is empty")
	}
	var chain models.Chain
	err := dao.GetChainByNetworkId(&chain, transaction.NetworkId)
	if err != nil {
		return err
	}
	tx := &models.Transaction{
		UserOperationHash: transaction.UserOperationHash,
		Type:              transaction.Type,
		Status:            models.TransactionStatusInit,
		ChainId:           chain.ID,
	}
	return dao.SaveTransaction(tx)
}

func FindTransaction(transaction *models.Transaction) ([]models.Transaction, error) {
	return dao.FindTransaction(transaction)
}

func FindInitTransaction(transaction *models.Transaction) ([]models.Transaction, error) {
	return dao.FindInitTransaction(transaction)
}

func DeleteTransaction(transaction *models.Transaction) error {
	return dao.DeleteTransaction(transaction)
}

func UpdateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	return dao.UpdateTransaction(transaction)
}
