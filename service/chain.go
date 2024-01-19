package service

import (
	"wallet-aa-tx-serv/dao"
	"wallet-aa-tx-serv/models"
)

func FindChain(chain *models.Chain) ([]models.Chain, error) {
	return dao.FindChain(chain)
}
