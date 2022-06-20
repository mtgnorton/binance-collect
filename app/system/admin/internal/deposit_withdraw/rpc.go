package deposit_withdraw

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/errors/gcode"
)

// EthError - ethereum error
type EthError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err EthError) Error() string {
	return fmt.Sprintf("Error %d (%s)", err.Code, err.Message)
}

type RpcRes struct {
	ID      int             `json:"id"`
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   *EthError       `json:"error"`
}

type RpcReq struct {
	ID      int           `json:"id"`
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type RpcClient struct {
	url    string
	client *http.Client
	Debug  bool
}

// 获取一个rpc客户端的实例
func NewRpcClient(url string, debug bool) *RpcClient {
	rpc := &RpcClient{
		url:    url,
		client: http.DefaultClient,
		Debug:  debug,
	}
	return rpc
}

// 获取当前最新的块编号
func (rpc *RpcClient) GetBlockNumber() (number int, err error) {
	var response string
	if err = rpc.callRetryAndAssignRes("eth_blockNumber", &response); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

// 根据块编号获取块信息
func (rpc *RpcClient) GetBlockInfoByNumber(number int, withTransactions bool) (*OriginBlock, error) {
	return rpc.getBlockInfo("eth_getBlockByNumber", withTransactions, IntToHex(number), withTransactions)
}

// 获取当前最新的余额
func (rpc *RpcClient) GetBalance(address string) (balance big.Int, err error) {
	var response string
	if err = rpc.callRetryAndAssignRes("eth_getBalance", &response, address, "latest"); err != nil {
		return big.Int{}, err
	}

	return ParseBigInt(response)
}

// 获取gas价格
func (rpc *RpcClient) GetGasPrice() (price big.Int, err error) {
	var response string
	if err = rpc.callRetryAndAssignRes("eth_gasPrice", &response); err != nil {
		return big.Int{}, err
	}

	return ParseBigInt(response)
}

func (rpc *RpcClient) getBlockInfo(method string, withTransactions bool, params ...interface{}) (*OriginBlock, error) {
	result, err := rpc.callRetry(method, params...)
	if err != nil {
		return nil, err
	}
	// err = gfile.PutContents(gfile.Pwd()+"/test_data/block_15913847.json", string(result))
	if err != nil {
		return nil, err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil, nil
	}

	var response proxyBlock

	if withTransactions {
		response = new(proxyBlockWithTransactions)
	} else {
		response = new(proxyBlockWithoutTransactions)
	}

	err = json.Unmarshal(result, response)
	if err != nil {
		return nil, err
	}
	block := response.toBlock()
	return &block, nil
}

//发送交易
func (rpc *RpcClient) SendRawTransaction(data string) (string, error) {
	var hash string

	err := rpc.callAndAssignRes("eth_sendRawTransaction", &hash, data)
	return hash, err
}

// 获取交易是否成功的相关信息
func (rpc *RpcClient) GetTransactionReceipt(hash string) (*TransactionReceipt, error) {
	transactionReceipt := new(TransactionReceipt)

	err := rpc.callRetryAndAssignRes("eth_getTransactionReceipt", transactionReceipt, hash)
	if err != nil {
		return nil, err
	}

	return transactionReceipt, nil
}

func (rpc *RpcClient) GetTransactionCount(address string, block string) (int, error) {
	var response string

	if err := rpc.callRetryAndAssignRes("eth_getTransactionCount", &response, address, block); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

func (rpc *RpcClient) callAndAssignRes(method string, res interface{}, params ...interface{}) (err error) {
	result, err := rpc.call(method, params...)
	if err != nil {
		return err
	}

	if res == nil {
		return nil
	}

	return json.Unmarshal(result, res)
}

func (rpc *RpcClient) callRetryAndAssignRes(method string, res interface{}, params ...interface{}) (err error) {
	result, err := rpc.callRetry(method, params...)
	if err != nil {
		return err
	}

	if res == nil {
		return nil
	}

	return json.Unmarshal(result, res)
}

// 重试调用rpc,如果call调用返回错误，每隔500ms重试一次，最多重试三次
func (rpc *RpcClient) callRetry(method string, params ...interface{}) (json.RawMessage, error) {
	sleepTime := time.Millisecond * time.Duration(500)
	var (
		res json.RawMessage
		err error
	)
	attempts := 3
	for i := 0; ; i++ {
		res, err = rpc.call(method, params...)
		if err == nil {
			return res, err
		}
		if i >= (attempts - 1) {
			break
		}
		time.Sleep(sleepTime * time.Duration(i+1))
	}

	LogInfofDw(context.TODO(), "after %d attempts, last error: %v", attempts, err)
	return res, err

}

// 底层请求方法
func (rpc *RpcClient) call(method string, params ...interface{}) (json.RawMessage, error) {
	ctx := context.TODO()
	request := RpcReq{
		ID:      1,
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	response, err := rpc.client.Post(rpc.url, "application/json", bytes.NewBuffer(body))
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if rpc.Debug {
		LogInfofDw(ctx, "%s\nRequest: %s\nResponse: %s\n", method, body, data)
	}

	resp := new(RpcRes)
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, gerror.NewCode(gcode.New(resp.Error.Code, resp.Error.Error(), nil))
	}
	return resp.Result, nil

}
