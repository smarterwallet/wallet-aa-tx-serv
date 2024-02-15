package dao

import (
	"fmt"
	"wallet-aa-tx-serv/global"
	"wallet-aa-tx-serv/models"
)

func SaveTransaction(info *models.Transaction) error {
	return global.DB.Save(info).Error
}

func DeleteTransaction(info *models.Transaction) error {
	return global.DB.Where(info).Delete(info).Error
}

func FindTransaction(info *models.Transaction) ([]models.Transaction, error) {
	var infos []models.Transaction
	err := global.DB.Where(info).Where(fmt.Sprintf("status = %d", info.Status)).Find(&infos).Error
	return infos, err
}

func FindInitTransaction(info *models.Transaction) ([]models.Transaction, error) {
	var infos []models.Transaction
	err := global.DB.Where(info).Where("status = 0").Find(&infos).Error
	return infos, err
}

func FindTransactionNeededToCheckStatus() ([]models.Transaction, error) {
	var infos []models.Transaction
	err := global.DB.Where("status != ? and status != ?", models.TransactionStatusSuccess, models.TransactionStatusFail).Find(&infos).Error
	return infos, err
}

func UpdateTransaction(info *models.Transaction) (*models.Transaction, error) {
	err := global.DB.Where("id = ?", info.ID).Updates(info).Error
	return info, err
}
