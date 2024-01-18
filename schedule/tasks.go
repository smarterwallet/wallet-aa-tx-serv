package schedule

import (
	log "github.com/cihub/seelog"
	"wallet-aa-tx-serv/service"
)

func PeriodicalUpdateStatusOfUserSendingTransaction() {
	infos, err := service.FindTransactionNeededToCheckStatus()
	if err != nil {
		log.Error(err)
		return
	}

	var txHashs []string
	for _, info := range infos {
		txHashs = append(txHashs, info.TxHash)
	}

	for i, txHash := range txHashs {

		transaction, err := service.GetTransactionReceiptResponse(txHash)
		if err != nil {
			log.Error(err)
			return
		}

		if transaction.Result.Status == infos[i].Status {
			continue
		}

		infos[i].Status = transaction.Result.Status

		_, err = service.UpdateTransaction(&infos[i])
		if err != nil {
			log.Error(err)
			return
		}

		log.Info("Updated status of tx:", txHash)

	}
}
