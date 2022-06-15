package deposit_withdraw

import (
	"encoding/json"
	"fmt"
	"gf-admin/app/model"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/test/gtest"
	"strings"
	"testing"
)

const ()

var client *RpcClient

func TestRpcClient_GetBlockByNumber(t *testing.T) {

	//gtest.C(t, func(t *gtest.T) {
	//	client = NewRpcClient(model.CHAIN_BSC_MAIN_NET, false)
	//	block, err := client.GetBlockInfoByNumber(15913847, true)
	//	t.Assert(err, nil)
	//	g.DumpWithType(block)
	//})
	//os.Exit(1)
	//with tx
	gtest.C(t, func(t *gtest.T) {
		client = NewRpcClient(model.CHAIN_BSC_MAIN_NET, false)
		block, err := client.GetBlockInfoByNumber(18133145, true)
		t.Assert(err, nil)
		jsonBytes, err := json.Marshal(block)
		t.Assert(err, nil)

		jsonString := strings.Trim(string(jsonBytes), " ")

		filePath := gfile.Pwd() + "/rpc_test_response/GetBlockByNumberWithTx.json"
		c := strings.Trim(gfile.GetContents(filePath), " \n")

		t.Assert(jsonString, c)
	})
	//without tx
	gtest.C(t, func(t *gtest.T) {
		client = NewRpcClient(model.CHAIN_BSC_MAIN_NET, false)
		block, err := client.GetBlockInfoByNumber(18133145, true)
		t.Assert(err, nil)
		jsonBytes, err := json.Marshal(block)
		t.Assert(err, nil)

		jsonString := strings.Trim(string(jsonBytes), " ")

		filePath := gfile.Pwd() + "/rpc_test_response/GetBlockByNumberWithoutTx.json"
		c := strings.Trim(gfile.GetContents(filePath), " \n")
		t.Assert(jsonString, c)
	})

}

func TestRpcClient_BlockNumber(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		client = NewRpcClient(model.CHAIN_BSC_MAIN_NET, false)
		blockNumber, err := client.GetBlockNumber()
		t.Assert(err, nil)
		fmt.Printf("current block number : %v\n", blockNumber)
	})
}

func TestRpcClient_GetBalance(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		client = NewRpcClient(model.CHAIN_BSC_MAIN_NET, true)
		balance, err := client.GetBalance("0x9b8f2b8a9b8f2b8a9b8f2b8a9b8f2b8a9b8f2b8a")
		t.Assert(err, nil)
		fmt.Printf("balance : %v\n", balance.String())
	})
}

func TestRpcClient_GetGasPrice(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		client = NewRpcClient(model.CHAIN_BSC_MAIN_NET, true)
		gasPrice, err := client.GetGasPrice()
		t.Assert(err, nil)
		fmt.Printf("gas price : %v\n", gasPrice.String())
	})
}

func TestRpcClient_GetTransactionReceipt(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		client = NewRpcClient(model.CHAIN_BSC_MAIN_NET, true)
		txReceipt, err := client.GetTransactionReceipt("0xdee9aaffa4f20533139ea6bb9fa05158799c987d647c8df2e5f7828d7d4f90a3")
		t.Assert(err, nil)
		fmt.Println(txReceipt)
	})
}
