package deposit_withdraw

// ChainClient chain concrete client
var ChainClient *BscChain

func init() {

	ChainClient = &BscChain{
		debug:       false,
		delayNumber: 6,
	}

}
