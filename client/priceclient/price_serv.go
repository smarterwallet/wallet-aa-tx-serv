package priceclient

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
	"wallet-aa-tx-serv/client/clientdto"
	"wallet-aa-tx-serv/utils/httplib"
	"wallet-aa-tx-serv/utils/time2"
)

var (
	tokenServiceHost = viper.GetString("service.token.price.api")
)

func GetUSDByTokenName(tokenName string, timeQuery *time.Time) (*decimal.Decimal, error) {
	defer time2.TimeConsume(time.Now())

	url := tokenServiceHost + "/api/v1/coin-price?coinName=" + strings.ToUpper(tokenName)

	if timeQuery != nil {
		url += "&timeQuery=" + strconv.Itoa(int(timeQuery.Unix()))
	}

	res := clientdto.ServiceResponse{}
	_, err := httplib.GetInto(url, nil, nil, &res)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, fmt.Errorf("get token price err: %s", res.Message)
	}

	price, err := decimal.NewFromString(res.Result.(string))
	if err != nil {
		return nil, err
	}
	return &price, nil
}
