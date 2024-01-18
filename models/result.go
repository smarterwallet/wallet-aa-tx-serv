package models

// 可能存在一部分类型未改为 uint
type GetTransactionByHashResult struct {
	BlockHash            string   `json:"blockHash"`
	BlockNumber          string   `json:"blockNumber"`
	Hash                 string   `json:"hash"`
	AccessList           []string `json:"accessList"`
	ChainID              uint     `json:"chainId"`
	From                 string   `json:"from"`
	Gas                  string   `json:"gas"`
	GasPrice             string   `json:"gasPrice"`
	Input                string   `json:"input"`
	MaxFeePerGas         string   `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string   `json:"maxPriorityFeePerGas"`
	Nonce                uint     `json:"nonce"`
	R                    string   `json:"r"`
	S                    string   `json:"s"`
	To                   string   `json:"to"`
	TransactionIndex     uint     `json:"transactionIndex"`
	Type                 uint     `json:"type"`
	V                    uint     `json:"v"`
	Value                uint     `json:"value"`
}

type GetTransactionReceiptResultLog struct {
	TransactionHash  string   `json:"transactionHash"`
	Address          string   `json:"address"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      string   `json:"blockNumber"`
	Data             string   `json:"data"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
	Topics           []string `json:"topics"`
	TransactionIndex uint     `json:"transactionIndex"`
}

type GetTransactionReceiptResult struct {
	TransactionHash   string                           `json:"transactionHash"`
	BlockHash         string                           `json:"blockHash"`
	BlockNumber       string                           `json:"blockNumber"`
	Logs              []GetTransactionReceiptResultLog `json:"logs"`
	ContractAddress   string                           `json:"contractAddress"`
	EffectiveGasPrice string                           `json:"effectiveGasPrice"`
	CumulativeGasUsed string                           `json:"cumulativeGasUsed"`
	From              string                           `json:"from"`
	GasUsed           string                           `json:"gasUsed"`
	LogsBloom         string                           `json:"logsBloom"`
	Status            uint                             `json:"status"`
	To                string                           `json:"to"`
	TransactionIndex  uint                             `json:"transactionIndex"`
	Type              uint                             `json:"type"`
}

type GetTransactionReceiptResponseData struct {
	Result GetTransactionReceiptResult `json:"result"`
}
