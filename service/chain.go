package service

import (
	"fmt"
	"strings"
	"wallet-aa-tx-serv/client/clientdto"
	"wallet-aa-tx-serv/global"
)

func GetChainByName(chainName string) (*clientdto.Chain, error) {
	for _, chain := range global.CacheConfigChainIdAndChain {
		if strings.ToLower(chain.Name) == strings.ToLower(chainName) {
			return chain, nil
		}
	}
	return nil, fmt.Errorf("can not find chain by chain name(%s)", chainName)
}

func GetChainByChainId(chainId int) (*clientdto.Chain, error) {
	chain, ok := global.CacheConfigChainIdAndChain[chainId]
	if !ok {
		return nil, fmt.Errorf("can not find chain by chain id(%d)", chainId)
	}
	return chain, nil
}
