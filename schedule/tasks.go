package schedule

import (
	"encoding/json"
	log "github.com/cihub/seelog"
	"wallet-aa-tx-serv/models"
	"wallet-aa-tx-serv/service"
	"wallet-aa-tx-serv/utils/common"
)

func PeriodicalUpdateStatusOfUserSendingTransaction() {
	// Find transaction status is TransactionStatusInit
	infos, err := service.FindTransaction(&models.Transaction{Status: models.TransactionStatusInit})
	if err != nil {
		log.Error(err)
		return
	}
	for _, info := range infos {
		// maybe url could be cached
		chain := &models.Chain{}
		chain.ID = info.ChainId
		chains, err := service.FindChain(chain)
		if err != nil {
			log.Error(err)
			return
		}
		rpcUrl := chains[0].RpcApi

		receipt, err := service.GetTransactionReceiptResponse(rpcUrl, info.TxHash)
		if err != nil {
			log.Error(err)
			return
		}

		// type assertion
		marshal, err := json.Marshal(receipt.Result)
		if err != nil {
			log.Errorf("fail to marshal receipt.Result: %v", err)
			return
		}

		receiptResult := &models.GetTransactionReceiptResult{}
		err = json.Unmarshal(marshal, &receiptResult)
		if err != nil {
			log.Errorf("fail to unmarshal receipt.Result: %v", err)
			return
		}

		resultStatus := common.ParseUint(receiptResult.Status)
		if resultStatus == info.Status {
			continue
		}

		info.Status = resultStatus

		// get and update parts of userOperation details
		details, err := service.GetUserOperationByHashResponse(rpcUrl, info.UserOperationHash)
		if err != nil {
			log.Error(err)
			return
		}
		opResult, ok := details.Result.(models.GetUserOperationByHashResult)
		if !ok {
			log.Error("fail to type assertion. (details.Result.(models.GetUserOperationByHashResult))")
			return
		}
		info.BlockHash = opResult.BlockHash
		info.BlockNumber = uint(opResult.BlockNumber)
		info.TxHash = opResult.TransactionHash
		info.Sender = opResult.UserOperation.Sender
		info.EntryPointAddress = opResult.EntryPoint
		info.UserOperation = &opResult.UserOperation

		_, err = service.UpdateTransaction(&info)
		if err != nil {
			log.Error(err)
			return
		}
		log.Info("Updated status of tx:", info.TxHash)
	}
}
