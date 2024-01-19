package models

var (
	EthRpcResponseErrorIsFalse = 0
)

type EthRpcResponseData struct {
	Error   EthRpcResponseError `json:"error"`
	Jsonrpc string              `json:"jsonrpc"`
	ID      uint                `json:"id"`
	Result  interface{}         `json:"result"`
}

type EthRpcResponseError struct {
	Code    int    `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

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
	Nonce                string   `json:"nonce"`
	R                    string   `json:"r"`
	S                    string   `json:"s"`
	To                   string   `json:"to"`
	TransactionIndex     string   `json:"transactionIndex"`
	Type                 string   `json:"type"`
	V                    string   `json:"v"`
	Value                string   `json:"value"`
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
	TransactionIndex string   `json:"transactionIndex"`
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
	Status            string                           `json:"status"`
	To                string                           `json:"to"`
	TransactionIndex  string                           `json:"transactionIndex"`
	Type              string                           `json:"type"`
}

type GetUserOperationByHashResult struct {
	UserOperation   UserOperation `json:"userOperation"`
	EntryPoint      string        `json:"entryPoint"`
	BlockNumber     int           `json:"blockNumber"`
	BlockHash       string        `json:"blockHash"`
	TransactionHash string        `json:"transactionHash"`
}
