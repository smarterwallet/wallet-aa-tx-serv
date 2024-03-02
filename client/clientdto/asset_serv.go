package clientdto

import (
	"time"
	"wallet-aa-tx-serv/models"
)

type GetPackageResponse struct {
	Common struct {
		ID        int         `json:"ID"`
		CreatedAt time.Time   `json:"CreatedAt"`
		UpdatedAt time.Time   `json:"UpdatedAt"`
		DeletedAt interface{} `json:"DeletedAt"`
		Name      string      `json:"name"`
		Version   string      `json:"version"`
		Config    struct {
			Url struct {
				Mpc struct {
					Api  string `json:"api"`
					Wasm string `json:"wasm"`
				} `json:"mpc"`
				AutoTrading struct {
					Mumbai string `json:"mumbai"`
				} `json:"autoTrading"`
				Asset   string `json:"asset"`
				Storage string `json:"storage"`
			} `json:"url"`
			ContractAddress struct {
				AutoTrading string `json:"autoTrading"`
			} `json:"contractAddress"`
		} `json:"config"`
	} `json:"common"`
	Chain []struct {
		ID                     int            `json:"ID"`
		CreatedAt              time.Time      `json:"CreatedAt"`
		UpdatedAt              time.Time      `json:"UpdatedAt"`
		DeletedAt              interface{}    `json:"DeletedAt"`
		NetWorkId              uint64         `json:"netWorkId"`
		Name                   string         `json:"name"`
		Icon                   string         `json:"icon"`
		Tokens                 []models.Token `json:"tokens"`
		Erc4337ContractAddress struct {
			SimpleAccountFactory string `json:"simpleAccountFactory"`
			TokenPaymaster       struct {
				Swt string `json:"swt"`
			} `json:"tokenPaymaster"`
			Entrypoint string `json:"entrypoint"`
		} `json:"erc4337ContractAddress"`
		RpcApi          string `json:"rpcApi"`
		BundlerApi      string `json:"bundlerApi"`
		BlockScanUrl    string `json:"blockScanUrl"`
		CreateWalletApi string `json:"createWalletApi"`
		ApiType         int    `json:"apiType"`
		ProduceBlock24H int    `json:"produceBlock24h"`
	} `json:"chain"`
}
