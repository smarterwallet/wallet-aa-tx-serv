package models

import (
	"encoding/json"
	"errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

var (
	// TransactionStatusInit 交易状态初始化
	TransactionStatusInit = uint(0)
	// TransactionStatusSuccess 交易状态成功
	TransactionStatusSuccess = uint(1)
	// TransactionStatusFail 交易状态失败
	TransactionStatusFail = uint(2)
	// TransactionStatusRollback 交易所在区块回滚
	TransactionStatusRollback = uint(3)
	// TransactionStatusUnKnow 交易状态无法判断
	TransactionStatusUnKnow = uint(4)

	// TxSourceFromDemandAbstraction 交易类型来自需求抽象
	TxSourceFromDemandAbstraction = uint(1)
	// TxSourceFromAutTrading 交易类型来自自动交易
	TxSourceFromAutTrading = uint(2)
)

type SavedTransaction struct {
	NetworkId         uint            `json:"networkId"`
	TxSource          uint            `gorm:"not null;comment:交易来源" json:"txSource"`
	UserOperationHash string          `json:"userOperationHash"`
	ExtraData         json.RawMessage `json:"extraData"`
}

type UserOperation struct {
	Sender               string `json:"sender"`
	Nonce                string `json:"nonce"`
	InitCode             string `json:"initCode"`
	CallData             string `json:"callData"`
	CallGasLimit         string `json:"callGasLimit"`
	VerificationGasLimit string `json:"verificationGasLimit"`
	PreVerificationGas   string `json:"preVerificationGas"`
	MaxFeePerGas         string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
	PaymasterAndData     string `json:"paymasterAndData"`
	Signature            string `json:"signature"`
}

// Transaction 交易信息
type Transaction struct {
	gorm.Model
	ChainId           uint            `gorm:"comment:chain表id" json:"chainId,omitempty"`
	BlockHash         string          `gorm:"comment:交易所在块Hash" json:"blockHash"`
	BlockNumber       uint            `gorm:"comment:交易所在块高度" json:"blockNumber"`
	TxHash            string          `gorm:"comment:交易Hash" json:"txHash"`
	Sender            string          `gorm:"comment:发送方" json:"sender"`
	EntryPointAddress string          `gorm:"comment:接收方" json:"entryPointAddress"`
	UserOperationHash string          `gorm:"comment:op hash" json:"userOperationHash"`
	UserOperation     *UserOperation  `gorm:"-" json:"userOperation"`
	UserOperationJson json.RawMessage `gorm:"comment:op详情;type:json" json:"userOperationJson"`
	ExtraData         json.RawMessage `gorm:"comment:额外数据;type:json" json:"extraData"`
	TxSource          uint            `gorm:"not null;comment:交易来源" json:"txSource"`
	Status            uint            `gorm:"not null;comment:状态" json:"status"`
}

func (Transaction) TableName() string {
	return "transaction"
}

func (s *Transaction) BeforeSave(tx *gorm.DB) (err error) {
	if s.UserOperation != nil {
		bytes, err := json.Marshal(s.UserOperation)
		if err != nil {
			return errors.New("failed to marshal Tokens field to JSON")
		}
		s.UserOperationJson = bytes
	}
	return
}

func (s *Transaction) AfterFind(tx *gorm.DB) (err error) {
	if s.UserOperationJson != nil {
		err = json.Unmarshal(s.UserOperationJson, &s.UserOperation)
		if err != nil {
			return errors.New("failed to unmarshal Tokens field from JSON")
		}
	}
	return
}

type EstimateFeeResponse struct {
	NetworkId uint64 `json:"networkId"`
	// USDValue 估算的手续费美元金额
	USDValue decimal.Decimal `json:"usdValue"`
	// EstimateFee 估算的手续费 key: tokenId value: estimateFee
	EstimateFee map[int]decimal.Decimal `json:"estimateFee"`
}
