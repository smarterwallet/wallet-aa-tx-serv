package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

var (
	// TransactionStatusInit 交易状态初始化
	TransactionStatusInit = 0
	// TransactionStatusSuccess 交易状态成功
	TransactionStatusSuccess = 1
	// TransactionStatusFail 交易状态失败
	TransactionStatusFail = 2
	// TransactionStatusRollback 交易所在区块回滚
	TransactionStatusRollback = 3
	// TransactionStatusUnKnow 交易状态无法判断
	TransactionStatusUnKnow = 4

	// TransactionTypeFromDemandAbstraction 交易类型来自需求抽象
	TransactionTypeFromDemandAbstraction = 1
	// TransactionTypeFromAutTrading 交易类型来自自动交易
	TransactionTypeFromAutTrading = 2
)

// Transaction 交易信息
type Transaction struct {
	gorm.Model
	ChainId           uint            `gorm:"comment:chain表id" json:"chainId,omitempty"`
	BlockHash         string          `gorm:"comment:交易所在块Hash" json:"blockHash"`
	BlockNumber       string          `gorm:"comment:交易所在块高度" json:"blockNumber"`
	TxHash            string          `gorm:"comment:交易Hash" json:"txHash"`
	Sender            string          `gorm:"comment:发送方" json:"sender"`
	EntryPointAddress string          `gorm:"comment:接收方" json:"entryPointAddress"`
	UserOperation     json.RawMessage `gorm:"comment:op详情" json:"userOperation"`
	Type              uint            `gorm:"not null;comment:类型" json:"type"`
	Status            uint            `gorm:"not null;comment:状态" json:"status"`
}

func (Transaction) TableName() string {
	return "transaction"
}
