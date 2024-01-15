package dao

import (
	"go-gin-gorm-starter/global"
	"go-gin-gorm-starter/models"
)

func GetChainByNetworkId(chain *models.Chain, networkId uint) error {
	err := global.DB.Where("network_id = ?", networkId).First(&chain).Error
	return err
}
