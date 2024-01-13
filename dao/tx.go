package dao

import (
	"go-gin-gorm-starter/global"
	"go-gin-gorm-starter/models"
)

func SaveTransaction(strategyInfo *models.Transaction) error {
	return global.DB.Save(strategyInfo).Error
}

func DeleteTransaction(strategyInfo *models.Transaction) error {
	return global.DB.Where(strategyInfo).Delete(strategyInfo).Error
}

func FindTransaction(info *models.Transaction) ([]models.Transaction, error) {
	var infos []models.Transaction
	err := global.DB.Where(info).Find(&infos).Error
	return infos, err
}

func UpdateTransaction(info *models.Transaction) (*models.Transaction, error) {
	err := global.DB.Where("id = ?", info.ID).Updates(info).Error
	return info, err
}
