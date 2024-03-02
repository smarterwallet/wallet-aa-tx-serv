package assetclient

import (
	"fmt"
	"github.com/spf13/viper"
	"wallet-aa-tx-serv/client/clientdto"
	"wallet-aa-tx-serv/utils/httplib"
)

var (
	aaTxServiceHost = viper.GetString("service.asset.api")
)

func GetPackage() (*clientdto.GetPackageResponse, error) {
	url := aaTxServiceHost + fmt.Sprintf("/api/v1/package")
	res := clientdto.ServiceResponse{
		Result: &clientdto.GetPackageResponse{},
	}
	_, err := httplib.GetInto(url, nil, nil, &res)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, fmt.Errorf("get package err: %s", res.Message)
	}

	return res.Result.(*clientdto.GetPackageResponse), nil
}
