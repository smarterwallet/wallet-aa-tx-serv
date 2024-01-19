package dao

import (
	"wallet-aa-tx-serv/global"
	"wallet-aa-tx-serv/models"
)

func SaveChain(chain *models.Transaction) error {
	return global.DB.Save(chain).Error
}

func DeleteChain(chain *models.Transaction) error {
	return global.DB.Where(chain).Delete(chain).Error
}

func FindChain(chain *models.Chain) ([]models.Chain, error) {
	var chains []models.Chain
	err := global.DB.Where(chain).Find(&chains).Error
	return chains, err
}

func UpdateChain(chain *models.Chain) (*models.Chain, error) {
	err := global.DB.Where("id = ?", chain.ID).Updates(chain).Error
	return chain, err
}

func GetChainByNetworkId(chain *models.Chain, networkId uint) error {
	err := global.DB.Where("network_id = ?", networkId).First(&chain).Error
	return err
}
