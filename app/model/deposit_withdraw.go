package model

const (

	// 链相关
	CHAIN_BSC = "bsc"
	CHAIN_ETH = "eth"
	CHAIN_TRC = "trc"

	CHAIN_BSC_MAIN_NET = "https://bsc-dataseed.binance.org/"
	CHAIN_BSC_TEST_NET = "https://data-seed-prebsc-2-s1.binance.org:8545/"
	//https://data-seed-prebsc-2-s1.binance.org:8545/

	CHAIN_BSC_MAIN_NET_ID = 56
	CHAIN_BSC_TEST_NET_ID = 97

	//合约相关
	CONTRACT_IS_COLLECT_OPEN  = 1
	CONTRACT_IS_COLLECT_CLOSE = 0

	CONTRACT_DEFAULT_SYMBOL        = "BNB"
	CONTRACT_DEAULT_SYMBOL_DECIMAL = 18
	CONTRACT_DEAULT_SYMBOL_ADDRESS = "0x"

	CONTRACT_TRANSFER_FUCTION_HEX = "0xa9059cbb"

	//处理器相关
	PROCESSOR_SIMPLE = "simple_processor"

	//交易相关
	TRANSACTION_TYPE_REACHRGE = "recharge"
	TRANSACTION_TYPE_COLLECT  = "collect"
	TRANSACTION_TYPE_FEE      = "fee"
	TRANSACTION_TYPE_WITHDRAW = "withdraw"

	//丢失区块状态字段
	LOSE_BLOCKS_STATE_FINISH    = 1
	LOSE_BLOCKS_STATE_NO_FINISH = 0
)
