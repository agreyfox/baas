package storm

import (
	"math/big"
)

type EthereumAPI interface {
	Web3ClientVersion() (string, error)
	Web3Sha3(data []byte) (string, error)
	NetVersion() (string, error)
	NetListening() (bool, error)
	NetPeerCount() (int, error)
	EthProtocolVersion() (string, error)
	EthSyncing() (*Syncing, error)
	EthCoinbase() (string, error)
	EthMining() (bool, error)
	EthHashrate() (int, error)
	EthGasPrice() (big.Int, error)
	EthAccounts() ([]string, error)
	EthBlockNumber() (int, error)
	EthGetBalance(address, block string) (big.Int, error)
	EthGetStorageAt(data string, position int, tag string) (string, error)
	EthGetTransactionCount(address, block string) (int, error)
	EthGetBlockTransactionCountByHash(hash string) (int, error)
	EthGetBlockTransactionCountByNumber(number int) (int, error)
	EthGetUncleCountByBlockHash(hash string) (int, error)
	EthGetUncleCountByBlockNumber(number int) (int, error)
	EthGetCode(address, block string) (string, error)
	EthSign(address, data string) (string, error)
	EthSendTransaction(transaction T) (string, error)
	EthSendRawTransaction(data string) (string, error)
	EthCall(transaction T, tag string) (string, error)
	EthEstimateGas(transaction T) (int, error)
	EthGetBlockByHash(hash string, withTransactions bool) (*Block, error)
	EthGetBlockByNumber(number int, withTransactions bool) (*Block, error)
	EthGetTransactionByHash(hash string) (*Transaction, error)
	EthGetTransactionByBlockHashAndIndex(blockHash string, transactionIndex int) (*Transaction, error)
	EthGetTransactionByBlockNumberAndIndex(blockNumber, transactionIndex int) (*Transaction, error)
	EthGetTransactionReceipt(hash string) (*TransactionReceipt, error)
	EthGetCompilers() ([]string, error)
	EthNewFilter(params FilterParams) (string, error)
	EthNewBlockFilter() (string, error)
	EthNewPendingTransactionFilter() (string, error)
	EthUninstallFilter(filterID string) (bool, error)
	EthGetFilterChanges(filterID string) ([]Log, error)
	EthGetFilterLogs(filterID string) ([]Log, error)
	EthGetLogs(params FilterParams) ([]Log, error)
}

var StormCmd = map[string]string{
	"Web3ClientVersion":                   "web3_clientVersion",
	"Web3Sha3":                            "web3_sha3",
	"NetVersion":                          "net_version",
	"NetListening":                        "net_listening",
	"NetPeerCount":                        "net_peerCount",
	"ProtocolVersion":                     "fst_protocolVersion",
	"Syncing":                             "fst_syncing",
	"Coinbase":                            "fst_coinbase",
	"Mining":                              "fst_mining",
	"Hashrate":                            "fst_hashrate",
	"GasPrice":                            "fst_gasPrice",
	"Accounts":                            "fst_accounts",
	"BlockNumber":                         "fst_blockNumber",
	"GetBalance":                          "fst_getBalance",
	"GetStorageAt":                        "fst_getStorageAt",
	"GetTransactionCount":                 "fst_getTransactionCount",
	"GetBlockTransactionCountByHash":      "fst_getBlockTransactionCountByHash",
	"GetBlockTransactionCountByNumber":    "fst_getBlockTransactionCountByNumber",
	"GetUncleCountByBlockHash":            "fst_getUncleCountByBlockHash",
	"GetUncleCountByBlockNumber":          "fst_getUncleCountByBlockNumber",
	"GetCode":                             "fst_getCode",
	"Sign":                                "fst_sign",
	"SendTransaction":                     "fst_sendTransaction",
	"SendRawTransaction":                  "fst_sendRawTransaction",
	"Call":                                "fst_call",
	"EstimateGas":                         "fst_estimateGas",
	"GetBlockByHash":                      "fst_getBlockByHash",
	"GetBlockByNumber":                    "fst_getBlockByNumber",
	"GetTransactionByHash":                "fst_getTransactionByHash",
	"GetTransactionByBlockHashAndIndex":   "fst_getTransactionByBlockHashAndIndex",
	"GetTransactionByBlockNumberAndIndex": "fst_getTransactionByBlockNumberAndIndex",
	"GetTransactionReceipt":               "fst_getTransactionReceipt",
	"GetCompilers":                        "fst_getCompilers",
	"NewFilter":                           "fst_newFilter",
	"NewBlockFilter":                      "fst_newBlockFilter",
	"NewPendingTransactionFilter":         "fst_newPendingTransactionFilter",
	"UninstallFilter":                     "fst_uninstallFilter",
	"GetFilterChanges":                    "fst_getFilterChanges",
	"GetFilterLogs":                       "fst_getFilterLogs",
	"GetLogs":                             "fst_getLogs",
}

var _ EthereumAPI = (*EthRPC)(nil)
