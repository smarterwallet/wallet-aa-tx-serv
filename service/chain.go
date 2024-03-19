package service

import (
	"fmt"
	"wallet-aa-tx-serv/client/clientdto"
	"wallet-aa-tx-serv/global"
)

func GetChainByChainId(chainId int) (*clientdto.Chain, error) {
	chain, ok := global.CacheConfigChainIdAndChain[chainId]
	if !ok {
		return nil, fmt.Errorf("can not find chain by chain id(%d)", chainId)
	}
	return chain, nil
}
