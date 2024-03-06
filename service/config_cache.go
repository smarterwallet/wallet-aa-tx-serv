package service

import (
	"wallet-aa-tx-serv/client/assetclient"
	"wallet-aa-tx-serv/global"
)

func init() {
	err := FlushCacheConfig()
	if err != nil {
		panic(err)
	}
}

func FlushCacheConfig() error {
	config, err := assetclient.GetPackage()
	if err != nil {
		return err
	}

	for i, c := range config.Chain {
		global.CacheConfigChainIdAndChain[c.ID] = &config.Chain[i]
	}
	return nil
}
