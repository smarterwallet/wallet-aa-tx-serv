package schedule

import (
	log "github.com/cihub/seelog"
	"strings"
	"wallet-aa-tx-serv/models"
	"wallet-aa-tx-serv/service"
	"wallet-aa-tx-serv/utils/common"
)

func PeriodicalUpdateStatusOfUserSendingTransaction() {
	// Find transaction status is TransactionStatusInit
	infos, err := service.FindInitTransaction(&models.Transaction{Status: models.TransactionStatusInit})
	if err != nil {
		log.Error(err)
		return
	}
	for _, info := range infos {
		// maybe url could be cached
		chain, err := service.GetChainByChainId(info.ChainId)
		if err != nil {
			log.Error(err)
			return
		}

		// get and update parts of userOperation details
		details, err := service.GetUserOperationByHashResponse(chain.BundlerApi, info.UserOperationHash)
		if err != nil {
			log.Error(err)
			return
		}
		opResult, ok := details.Result.(*models.GetUserOperationByHashResult)
		if !ok {
			log.Error("fail to type assertion. (details.Result.(models.GetUserOperationByHashResult))")
			return
		}
		info.BlockHash = strings.ToLower(opResult.BlockHash)
		info.BlockNumber = uint(opResult.BlockNumber)
		info.TxHash = strings.ToLower(opResult.TransactionHash)
		info.Sender = strings.ToLower(opResult.UserOperation.Sender)
		info.EntryPointAddress = strings.ToLower(opResult.EntryPoint)
		info.UserOperation = &opResult.UserOperation

		// get and update transaction status
		receipt, err := service.GetTransactionReceiptResponse(chain.RpcApi, info.TxHash)
		if err != nil {
			log.Error(err)
			return
		}
		// type assertion
		receiptResult, ok := receipt.Result.(*models.GetTransactionReceiptResult)
		if !ok {
			log.Error("fail to type assertion. (receipt.Result.(models.GetTransactionReceiptResult))")
			return
		}

		resultStatus := common.ParseUint(receiptResult.Status)
		switch resultStatus {
		case 0:
			resultStatus = models.TransactionStatusFail
			break
		case 1:
			resultStatus = models.TransactionStatusSuccess
			break
		}
		info.Status = resultStatus

		// update transaction
		_, err = service.UpdateTransaction(&info)
		if err != nil {
			log.Error(err)
			return
		}
		log.Info("Updated status of tx:", info.TxHash)
	}
}
