package global

import (
	"gorm.io/gorm"
	"wallet-aa-tx-serv/models"
)

var (
	DB *gorm.DB

	// CacheConfigNetworkIdAndRPC 缓存的链 key: networkId value: rpcApi
	CacheConfigNetworkIdAndRPC = make(map[uint64]string)
	// CacheConfigNetworkIdAndTokens 缓存的链 key: networkId value: []token
	CacheConfigNetworkIdAndTokens = make(map[uint64][]models.Token)
)
