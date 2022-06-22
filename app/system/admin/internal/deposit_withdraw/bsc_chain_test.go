package deposit_withdraw

import (
	"fmt"
	"gf-admin/app/model"
	"math/big"
	"testing"
	"time"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/gogf/gf/v2/test/gtest"
)

var TestClient *BscChain

func init() {
	TestClient = &BscChain{
		debug:       true,
		delayNumber: 6,
		netUrl:      model.CHAIN_BSC_TEST_NET,
		chainID:     model.CHAIN_BSC_TEST_NET_ID,
	}
	TestClient.SetRpcClient(NewRpcClient(TestClient.netUrl, TestClient.debug))

}

func TestBscChain_GetOnceGasPrice(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		client := ChainClient

		total, err := client.GetProbablyOnceGasPrice(ctx, "BNB")
		t.AssertNil(err)
		fmt.Println(total.String())
		t.AssertGT(total.Int64(), 0)

		total, err = client.GetProbablyOnceGasPrice(ctx, "bnb")
		t.AssertNil(err)
		t.AssertGT(total.Int64(), 0)

		total, err = client.GetProbablyOnceGasPrice(ctx, "BSC-USD")
		t.AssertNil(err)
		t.AssertGT(total.Int64(), 0)

		total, err = client.GetProbablyOnceGasPrice(ctx, "bsc-usd")
		t.AssertNil(err)
		t.AssertGT(total.Int64(), 0)
	})
}

func TestBscChain_SendTransaction(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		from := "0x991195b40a5bdf4725afbd4f10f579bca25308f5"
		fromPrivateKey := "0x81bef96cefdc28f6e77256dff93a6c56f9953796a251ae796aa27ec196b26c03"
		to := "0xdf4767b601e4ec936f2ff1b5630e13270805dffc"

		contractAddress := "0x2151F2B84134C6df6690E8E3E11AEf1AC3594145"

		nonce, err := TestClient.GetLastNonce(ctx, from)
		t.AssertNil(err)
		gasLimit := TestClient.GetGasLimit(ctx, model.CONTRACT_DEFAULT_SYMBOL)
		gasPrice, err := TestClient.GetGasPrice(ctx)
		t.AssertNil(err)

		toAddress := common.HexToAddress(to)
		legacyTx := &types.LegacyTx{
			Nonce:    nonce,
			GasPrice: gasPrice,
			Gas:      gasLimit,
			To:       &toAddress,
			Value:    big.NewInt(1000000),
			Data:     nil,
		}

		// eth 交易
		hash, err := TestClient.SendTransaction(ctx, legacyTx, fromPrivateKey, "")
		t.AssertNil(err)

		fmt.Printf("eth hash:%s \n", hash)
		time.Sleep(time.Second)

		legacyTx = &types.LegacyTx{
			Nonce:    nonce,
			GasPrice: gasPrice,
			Gas:      gasLimit,
			To:       &toAddress,
			Value:    big.NewInt(1000000),
			Data:     nil,
		}
		// erc-20 交易
		legacyTx.Gas = TestClient.GetGasLimit(ctx, "MTG-USD")

		legacyTx.Nonce = nonce + 1

		hash, err = TestClient.SendTransaction(ctx, legacyTx, fromPrivateKey, contractAddress)
		t.AssertNil(err)
		fmt.Printf("erc20 hash:%s \n", hash)
	})

}

func TestBscChain_SendTransaction2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {

		from := "0x84b2d9c9b870ca47719e17e8cf790d66686743c8"
		fromPrivateKey := "0x3e525c9e6a3687f342021efee8258a21f66eee5bdfe6975e981a7ac438b95425"

		to := "0x8520e2ea780e400ab87322d04c158267f36f733a"
		toAddress := common.HexToAddress(to)

		contractAddress := "0x2151f2b84134c6df6690e8e3e11aef1ac3594145"

		nonce, err := TestClient.GetLastNonce(ctx, from)
		t.AssertNil(err)

		gasPrice, err := TestClient.GetGasPrice(ctx)
		t.AssertNil(err)

		legacyTx := &types.LegacyTx{
			Nonce:    nonce,
			GasPrice: gasPrice,
			Gas:      60000,
			To:       &toAddress,
			Value:    big.NewInt(100000000000000000), // 0.1
			Data:     nil,
		}

		hash, err := TestClient.SendTransaction(ctx, legacyTx, fromPrivateKey, contractAddress)
		t.AssertNil(err)
		fmt.Printf("erc20 hash:%s \n", hash)
	})
}

func TestBscChain_GetLastNonce(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		address := "0x991195b40a5bdf4725afbd4f10f579bca25308f5"
		nonce, err := TestClient.GetLastNonce(ctx, address)
		t.AssertNil(err)
		fmt.Println(nonce)
	})

}

func TestBscChain_WeiToEther(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		wei := "100000000000000000"
		ether, err := TestClient.WeiToEther(ctx, wei, model.CONTRACT_DEFAULT_SYMBOL)
		t.AssertNil(err)
		fmt.Println(ether)
	})
	gtest.C(t, func(t *gtest.T) {
		wei := "1e+16"
		ether, err := TestClient.WeiToEther(ctx, wei, model.CONTRACT_DEFAULT_SYMBOL)
		t.AssertNil(err)
		fmt.Println(ether)
	})
}

func TestBscChain_EtherToWei(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ether := "0.1"
		wei, err := TestClient.EtherToWei(ctx, ether, model.CONTRACT_DEFAULT_SYMBOL)
		t.AssertNil(err)
		fmt.Println(wei)
	})
}

func TestBscChain_GetMinCollectValue(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		value, err := TestClient.GetMinCollectValue(ctx)
		t.AssertNil(err)
		g.Dump(value.String())
	})
}
