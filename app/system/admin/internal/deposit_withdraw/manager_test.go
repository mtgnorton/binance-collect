package deposit_withdraw

import (
	"fmt"
	"gf-admin/app/model"
	"math"
	"math/big"
	"testing"

	"github.com/gogf/gf/v2/container/garray"

	"github.com/gogf/gf/v2/util/grand"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/test/gtest"
)

var manager *Manager

func InitManager() {

	manager = NewManager(ctx)

	//_, _ = g.DB().Exec(ctx, "truncate table collects;")
	//_, _ = g.DB().Exec(ctx, "truncate table queue_task;")
	//_, _ = g.DB().Exec(ctx, "truncate table queue_task_log;")
	//_, _ = g.DB().Exec(ctx, "truncate table withdraws;")
	//_, _ = g.DB().Exec(ctx, "truncate table notify;")
	//_, _ = g.DB().Exec(ctx, "truncate table notify_log;")
}

func Test_Manager_Detect(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		InitManager()
		manager.Run(ctx)
	})
}

// 测试某一个区块
func Test_Manager_Single_Block(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		InitManager()

		m := manager
		detectNumber := 20148458
		blockInfo, err := ChainClient.GetBlockInfoByNumber(ctx, detectNumber)
		t.AssertNil(err)
		transactions, err := m.transactionProcessor.DistinguishAndParse(ctx, blockInfo)

		t.AssertNil(err)

		m.Dispatch(ctx, transactions)

	})
}

func Test_Concurrency_Collect(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {

		_, _ = g.DB().Exec(ctx, "truncate table collects;")
		_, _ = g.DB().Exec(ctx, "truncate table queue_task;")
		_, _ = g.DB().Exec(ctx, "truncate table queue_task_log;")
		_, _ = g.DB().Exec(ctx, "truncate table withdraws;")
		_, _ = g.DB().Exec(ctx, "truncate table notify;")
		_, _ = g.DB().Exec(ctx, "truncate table notify_log;")

		TestClient = &BscChain{
			debug:       true,
			delayNumber: 6,
			netUrl:      model.CHAIN_BSC_TEST_NET,
			chainID:     model.CHAIN_BSC_TEST_NET_ID,
		}
		TestClient.SetRpcClient(NewRpcClient(TestClient.netUrl, TestClient.debug))

		// 第三方地址和私钥
		//  0xf0E49B017EbC2C52DD1210Dd88FB2370d393B784
		//	9b2d63625bc88ceffb207e0c37b08935e021d87f24c9c69b2a52d05f33fcf259
		//
		//
		//	0x050E33f56373b99F88e22CF80D6576463b88442E
		//	13f6ed079f11d515966ed00c9a4d0b1cb222d03d0bf7dddee2c0c46857e8ea16
		//
		//	0x92ca31682C42a1c3b663F2f691e62646c5A9cfd5
		//	f51070f4a3a7a16b9aa16e022b59edc608eab6044e845a328d081d81f117bf62
		//
		//	0x906600336C35308bFa5c5233075531cA14047bbE
		//	d7c13b7baf714972573d46225238f7d404768da5e4739cf2d0d06fdd4c6f7eb5
		//
		//	0x832232A77406D899E47089E521a7FE593457f253
		//	357576af27163a78fb9501d938f441801c7d5a25c79efdad150cda0a71541156
		//
		//	0x82BabaA5ffaB8774605bbF4481BFF6cC72Aa87AA
		//	08c7e923c28fab3d1377ee84ce31f27b895df4978449770e3a01b228837a3764
		//
		//	0xcD5A0D11DFf194Ad316e8bE4Ae3A0fEb63344522
		//	2c4d68f6b38c0f071a178abe065c2b3e8db4316b3f4d243879afab2184a91a3e
		//
		//	0x4108c486198acC0898BFd0a88198fD207A4501Df
		//	f2a3bb54514614a3bee08553be02fdee34d48e9d12ec05b9e6e07aca1220484d
		//
		//	0x8A2fe8C3970CB0a610843D6d3ED8D7B6821Afb71
		//	4f94476e95307e0949738b73be62d3fe9909742c3c0101db4a42f79cd10ceff7
		//
		//	0x82b7Bdb35E1F4c3443543Fa6fc13A92d8dC03443
		//	404550111b3316990c19aeac839831524e67810e8e375d14eb8e5f3bfcbb7fe5
		thirdAddresses := g.MapStrStr{
			"0xf0E49B017EbC2C52DD1210Dd88FB2370d393B784": "0x9b2d63625bc88ceffb207e0c37b08935e021d87f24c9c69b2a52d05f33fcf259",
			"0x050E33f56373b99F88e22CF80D6576463b88442E": "0x13f6ed079f11d515966ed00c9a4d0b1cb222d03d0bf7dddee2c0c46857e8ea16",
			"0x92ca31682C42a1c3b663F2f691e62646c5A9cfd5": "0xf51070f4a3a7a16b9aa16e022b59edc608eab6044e845a328d081d81f117bf62",
			"0x906600336C35308bFa5c5233075531cA14047bbE": "0xd7c13b7baf714972573d46225238f7d404768da5e4739cf2d0d06fdd4c6f7eb5",
			"0x832232A77406D899E47089E521a7FE593457f253": "0x357576af27163a78fb9501d938f441801c7d5a25c79efdad150cda0a71541156",
			"0x82BabaA5ffaB8774605bbF4481BFF6cC72Aa87AA": "0x08c7e923c28fab3d1377ee84ce31f27b895df4978449770e3a01b228837a3764",
			"0xcD5A0D11DFf194Ad316e8bE4Ae3A0fEb63344522": "0x2c4d68f6b38c0f071a178abe065c2b3e8db4316b3f4d243879afab2184a91a3e",
			"0x4108c486198acC0898BFd0a88198fD207A4501Df": "0xf2a3bb54514614a3bee08553be02fdee34d48e9d12ec05b9e6e07aca1220484d",
			"0x8A2fe8C3970CB0a610843D6d3ED8D7B6821Afb71": "0x4f94476e95307e0949738b73be62d3fe9909742c3c0101db4a42f79cd10ceff7",
			"0x82b7Bdb35E1F4c3443543Fa6fc13A92d8dC03443": "0x404550111b3316990c19aeac839831524e67810e8e375d14eb8e5f3bfcbb7fe5",
		}
		userAddresses := garray.NewStrArrayFrom([]string{
			//"0x81023633832221b512018a21f8a3c6a6fe774913",
			"0x84b2d9c9b870ca47719e17e8cf790d66686743c8",
		})

		contractAddresses := gmap.NewStrStrMapFrom(g.MapStrStr{
			"0x2151F2B84134C6df6690E8E3E11AEf1AC3594145": "MTG-USD",
			"0x": "BNB",
		})

		amount := 0
		for thirdAddress, thirdPrivateKey := range thirdAddresses {

			userAddress, _ := userAddresses.Rand()

			toAddress := common.HexToAddress(userAddress)

			nonce, err := TestClient.GetLastNonce(ctx, thirdAddress)
			t.AssertNil(err)

			contractAddress, symbol := contractAddresses.Pop()
			contractAddresses.Set(contractAddress, symbol)

			gasLimit := TestClient.GetGasLimit(ctx, symbol)

			gasPrice, err := TestClient.GetGasPrice(ctx)

			t.AssertNil(err)

			valueEther := big.NewFloat(float64(grand.N(1, 10)) / float64(100))

			valueFloat := valueEther.Mul(valueEther, big.NewFloat(math.Pow(10, 18)))

			valueInt, _ := valueFloat.Int(big.NewInt(0))

			legacyTx := &types.LegacyTx{
				Nonce:    nonce,
				GasPrice: gasPrice,
				Gas:      gasLimit,
				To:       &toAddress,
				Value:    valueInt,
				Data:     nil,
			}
			g.Dump(symbol, thirdAddress)
			// eth 交易
			hash, err := TestClient.SendTransaction(ctx, legacyTx, thirdPrivateKey, contractAddress)
			t.AssertNil(err)
			amount++
			fmt.Printf("第%d次交易hash为%s", amount, hash)

		}

	})
}
