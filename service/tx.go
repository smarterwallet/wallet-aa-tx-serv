package service

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
	"wallet-aa-tx-serv/client/clientdto"
	"wallet-aa-tx-serv/client/priceclient"
	"wallet-aa-tx-serv/dao"
	"wallet-aa-tx-serv/global"
	"wallet-aa-tx-serv/models"
)

func SaveTransaction(transaction *models.SavedTransaction) error {
	if transaction.UserOperationHash == "" {
		return global.OtherError("UserOperationHash is empty")
	}
	chain, err := GetChainByChainId(transaction.ChainId)
	if err != nil {
		return err
	}
	tx := &models.Transaction{
		UserOperationHash: transaction.UserOperationHash,
		TxSource:          transaction.TxSource,
		Status:            models.TransactionStatusInit,
		ChainId:           chain.ID,
		ExtraData:         transaction.ExtraData,
	}
	return dao.SaveTransaction(tx)
}

func FindTransaction(transaction *models.Transaction) ([]models.Transaction, error) {
	return dao.FindTransaction(transaction)
}

func FindInitTransaction(transaction *models.Transaction) ([]models.Transaction, error) {
	return dao.FindInitTransaction(transaction)
}

func DeleteTransaction(transaction *models.Transaction) error {
	return dao.DeleteTransaction(transaction)
}

func UpdateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	return dao.UpdateTransaction(transaction)
}

func GetEstimateFee(chainId int) (*models.EstimateFeeResponse, error) {
	chain, ok := global.CacheConfigChainIdAndChain[chainId]
	if !ok {
		return nil, fmt.Errorf("chainId(%d) not found rpc", chainId)
	}
	nativeToken := &clientdto.Token{}
	for _, t := range chain.Tokens {
		if t.Type == 0 {
			nativeToken = &t
			break
		}
	}

	gasLimit := big.NewInt(500000)

	// 预估需要native token数量
	// nativeGasfee = gasPrice * gasLimit
	response, err := GetGasPriceResponse(chain.RpcApi)
	if err != nil {
		return nil, err
	}
	gasPrice, _ := new(big.Int).SetString(response.Result.(string)[2:], 16)
	gasFeeBigInt := new(big.Int).Mul(gasPrice, gasLimit)

	result := new(big.Int)
	nativeTokenDeciBigInt := result.Exp(big.NewInt(10), big.NewInt(int64(nativeToken.Decimal)), nil)
	nativeTokenDeciDecimal, _ := decimal.NewFromString(nativeTokenDeciBigInt.String())

	gasFeeDecimal, _ := decimal.NewFromString(gasFeeBigInt.String())
	gasFeeDecimal = gasFeeDecimal.Div(nativeTokenDeciDecimal)

	nativeTokenPrice, err := priceclient.GetUSDByTokenName(nativeToken.Name, nil)
	if err != nil {
		return nil, err
	}
	gasFeeDecimalUSD := gasFeeDecimal.Mul(*nativeTokenPrice)

	// 最后输出的结果
	feeResult := &models.EstimateFeeResponse{
		ChainId:        chain.ID,
		GasPrice:       gasPrice,
		PayFeeUSDValue: gasFeeDecimalUSD,
		PayFeeByToken:  []models.TokenFee{},
	}
	feeResult.PayFeeByToken = append(feeResult.PayFeeByToken, models.TokenFee{
		Token:      nativeToken,
		NeedAmount: gasFeeDecimal,
	})

	// 计算其他token
	for i, token := range chain.Tokens {
		// 要求ERC20token并且可以作为手续费
		if token.Type == 1 && token.Fee == 1 {
			// ERC20Token = nativeGasfee * native token price / ERC20Token price
			erc20TokenPrice, err := priceclient.GetUSDByTokenName(token.Name, nil)
			if err != nil {
				return nil, err
			}
			if erc20TokenPrice == nil || erc20TokenPrice.String() == "0" {
				continue
			}
			feeResult.PayFeeByToken = append(feeResult.PayFeeByToken, models.TokenFee{
				Token:      &chain.Tokens[i],
				NeedAmount: gasFeeDecimalUSD.Div(*erc20TokenPrice),
			})
		}
	}

	return feeResult, nil

}
