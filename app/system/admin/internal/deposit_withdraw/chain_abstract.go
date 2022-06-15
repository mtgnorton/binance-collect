package deposit_withdraw

import (
	"context"
	"math/big"
)

// delete
type ChainAbstractBak interface {
	SetDelayNumber(int)
	GetDelayNumber() int
	SetIsDebug(bool)
	getDebug() bool
	SetRpcClient(client *RpcClient)
	GetDetectNumber(context.Context) (int, error) // 获取等待检测的区块
	GetBlockNumber(context.Context) (int, error)  //获取当前最新区块
	GetBlockInfoByNumber(context.Context, int) (*OriginBlock, error)
	GetBalance(context.Context, string) (big.Int, error)
	GetTokenBalance(context.Context, string, string) (balance big.Int, err error)
	GetProbablyOnceGasPrice(context.Context, string) (big.Int, error)
	sendRawTransaction(context.Context, string) (string, error)
	GetTransactionReceipt(context.Context, string) (any, error)
}
