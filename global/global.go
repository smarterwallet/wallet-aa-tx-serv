package global

import (
	"gorm.io/gorm"
	"wallet-aa-tx-serv/client/clientdto"
)

var (
	DB *gorm.DB

	// CacheConfigChainIdAndChain 缓存的链 key: chainId value: chain
	CacheConfigChainIdAndChain = make(map[int]*clientdto.Chain)
)
