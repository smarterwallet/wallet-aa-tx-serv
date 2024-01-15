package dao

import (
	"wallet-aa-tx-serv/global"
	"wallet-aa-tx-serv/models"
)

func GetChainByNetworkId(chain *models.Chain, networkId uint) error {
	err := global.DB.Where("network_id = ?", networkId).First(&chain).Error
	return err
}
