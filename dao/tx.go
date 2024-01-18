package dao

import (
	"wallet-aa-tx-serv/global"
	"wallet-aa-tx-serv/models"
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

// FindTransactionNeededToCheckStatus FIXME: 这个sql只需要查询为0的，直接用FindTransaction就行，不用单独写一个
func FindTransactionNeededToCheckStatus() ([]models.Transaction, error) {
	var infos []models.Transaction
	err := global.DB.Where("status != 1 and status != 2").Find(&infos).Error
	return infos, err
}

func UpdateTransaction(info *models.Transaction) (*models.Transaction, error) {
	err := global.DB.Where("id = ?", info.ID).Updates(info).Error
	return info, err
}
