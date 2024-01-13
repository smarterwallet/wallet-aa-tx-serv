package service

import (
	"go-gin-gorm-starter/dao"
	"go-gin-gorm-starter/models"
)

func SaveTransaction(strategyInfo *models.Transaction) error {
	return dao.SaveTransaction(strategyInfo)
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
