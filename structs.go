package transactions

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Holds the Parameters for an RPC call
type ParamInfo struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int           `json:"id"`
}

// Holds the eth_blockNumber response
type BlockNumberResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  string `json:"result"`
}

// Stores an JSON RPC Transaction
type TxBlockRPC struct {
	Hash             string `json:"hash"`
	Nonce            string `json:"nonce"`
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	TransactionIndex string `json:"transactionIndex"`
	From             string `json:"from"`
	To               string `json:"to"`
	Value            string `json:"value"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Input            string `json:"input"`
}

type TransactionType struct {
	Number           string       `json:"number"`
	Hash             string       `json:"hash"`
	ParentHash       string       `json:"parentHash"`
	Nonce            string       `json:"nonce"`
	Sha3Uncles       string       `json:"sha3Uncles"`
	LogsBloom        string       `json:"logsBloom"`
	TransactionsRoot string       `json:"transactionsRoot"`
	StateRoot        string       `json:"stateRoot"`
	Miner            string       `json:"miner"`
	Difficulty       string       `json:"difficulty"`
	TotalDifficulty  string       `json:"totalDifficulty"`
	ExtraData        string       `json:"extraData"`
	Size             string       `json:"size"`
	GasLimit         string       `json:"gasLimit"`
	GasUsed          string       `json:"gasUsed"`
	Timestamp        *hexutil.Big `json:"timestamp"`
	Transactions     []TxBlockRPC `json:"transactions"`
	Uncles           []string     `json:"uncles"`
}

// The return type for eth_getBlockByNumber method
type BlockByNumberResponse struct {
	Jsonrpc string          `json:"jsonrpc"`
	Id      int             `json:"id"`
	Result  TransactionType `json:"result"`
}

// A single transaction for the EtherScan.io Response
type TransactionBlockEtherScan struct {
	BlockHash         string `json:"blockHash"`
	BlockNumber       string `json:"blockNumber`
	Confirmations     string `json:"confirmations"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	From              string `json:"from"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           string `json:"gasUsed"`
	Hash              string `json:"hash"`
	Input             string `json:"input"`
	IsError           string `json:"isError"`
	Nonce             string `json:"nonce"`
	TimeStamp         string `json:"timeStamp"`
	To                string `json:"to"`
	TransactionIndex  string `json:"transactionIndex"`
	Txreceipt_status  string `json:"txreceipt_status"`
	Value             string `json:"value"`
}

// Holds the EtherScan.io API response
type Configuration struct {
	Status  string                      `json:"status"`
	Message string                      `json:"message"`
	Result  []TransactionBlockEtherScan `json:"result"`
}
