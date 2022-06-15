package deposit_withdraw

import (
	"gf-admin/app/model"
)

// ChainClient chain concrete client
var ChainClient *BscChain

func init() {
	ChainClient = &BscChain{
		debug:       false,
		delayNumber: 6,
		netUrl:      model.CHAIN_BSC_TEST_NET,
		chainID:     model.CHAIN_BSC_TEST_NET_ID,
	}
	ChainClient.SetRpcClient(NewRpcClient(ChainClient.netUrl, ChainClient.debug))

}
