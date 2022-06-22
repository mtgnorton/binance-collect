package deposit_withdraw

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"gf-admin/app/shared"
	"gf-admin/utility/custom_error"
	"math"
	"math/big"

	"github.com/gogf/gf/v2/frame/g"

	ethereumCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/text/gstr"
)

type BscChain struct {
	rpcClient   *RpcClient
	debug       bool
	delayNumber int //区块有可能回滚，一般超过6个之后，可以确定该块已经被确认
	netUrl      string
	chainID     int64
}

func (b *BscChain) SetNetByConfig(ctx context.Context) error {

	netType, err := shared.Config.GetString(ctx, model.CONFIG_MODULE_BINNABCE, model.CONFIG_KEY_NET_TYPE)
	if err != nil {
		return custom_error.New(err.Error())
	}
	if netType == model.CONFIG_KEY_NET_TYPE_VALUE_MAIN {
		b.netUrl = model.CHAIN_BSC_MAIN_NET
		b.chainID = model.CHAIN_BSC_MAIN_NET_ID
	} else {
		b.netUrl = model.CHAIN_BSC_TEST_NET
		b.chainID = model.CHAIN_BSC_TEST_NET_ID
	}
	LogInfofDw(ctx, "SetNetByConfig netType:%s", netType)

	b.SetRpcClient(NewRpcClient(ChainClient.netUrl, ChainClient.debug))
	gcache.Remove(ctx, model.CACHE_KEY_CHAIN_NEWEST_NUMBER)

	return nil
}

// 获取用户地址map,以用户地址为key
func (b *BscChain) GetUserAddresses(ctx context.Context) (map[string]entity.UserAddresses, error) {
	//获取用户地址
	userAddressesVar, err := gcache.GetOrSetFunc(ctx, model.CACHE_KEY_USER_ADDRESSES, func(ctx context.Context) (interface{}, error) {
		var rs []entity.UserAddresses
		err := dao.UserAddresses.Ctx(ctx).Scan(&rs)
		return rs, err
	}, 0)
	if err != nil {
		return nil, err
	}

	addressSlice, ok := userAddressesVar.Val().([]entity.UserAddresses)
	if !ok {
		return nil, custom_error.New("userAddressesVar.Interface().([]entity.UserAddresses) failed")
	}
	userAddresses := make(map[string]entity.UserAddresses)
	for _, userAddress := range addressSlice {
		userAddresses[gstr.ToLower(userAddress.Address)] = userAddress
	}
	return userAddresses, nil
}

// 合约map，以合约地址为key
func (b *BscChain) GetContracts(ctx context.Context) (map[string]entity.Contracts, error) {
	//获取合约信息
	contractsVar, err := gcache.GetOrSetFunc(ctx, model.CACHE_KEY_CONTRACTS, func(ctx context.Context) (interface{}, error) {
		var rs []entity.Contracts
		err := dao.Contracts.Ctx(ctx).Scan(&rs)
		return rs, err
	}, 0)
	if err != nil {
		return nil, err
	}

	contractsSlice, ok := contractsVar.Interface().([]entity.Contracts)

	if !ok {
		return nil, custom_error.New("contractsVar.Interface().([]entity.Contracts) failed")
	}

	contracts := make(map[string]entity.Contracts)
	for _, contract := range contractsSlice {
		contracts[gstr.ToLower(contract.Address)] = contract
	}
	return contracts, nil
}

// 获取手续费提现地址
func (b *BscChain) GetFeeWithdrawAddress(ctx context.Context) (string, error) {
	feeWithdrawAddress, err := shared.Config.GetString(ctx, model.CONFIG_MODULE_BINNABCE, model.CONFIG_KEY_FEE_WITHDRAW_ADDRESS)
	if err != nil {
		return "", custom_error.New(err.Error())
	}
	return feeWithdrawAddress, nil
}

// 获取手续费提现私钥
func (b *BscChain) GetFeeWithdrawPrivateKey(ctx context.Context) (string, error) {
	feeWithdrawPrivateKey, err := shared.Config.GetString(ctx, model.CONFIG_MODULE_BINNABCE, model.CONFIG_KEY_FEE_WITHDRAW_PRIVATE_KEY)
	if err != nil {
		return "", custom_error.New(err.Error())
	}
	return feeWithdrawPrivateKey, nil
}

// 获取归集地址
func (b *BscChain) GetCollectAddress(ctx context.Context) (string, error) {

	// 获取平台归集地址
	collectAddress, err := shared.Config.GetString(ctx, model.CONFIG_MODULE_BINNABCE, model.CONFIG_KEY_COLLECT_ADDRESS)
	if err != nil {
		return "", custom_error.New(err.Error())
	}
	return collectAddress, nil
}

//  获取最小充值金额

func (b *BscChain) GetMinCollectValue(ctx context.Context) (*big.Int, error) {
	minCollectValue, err := shared.Config.GetString(ctx, model.CONFIG_MODULE_BINNABCE, model.CONFIG_KEY_MIN_COLLECT_AMOUNT)
	if err != nil {
		return big.NewInt(0), custom_error.New(err.Error())
	}
	valueString, err := b.EtherToWei(ctx, minCollectValue, model.CONTRACT_DEFAULT_SYMBOL)
	if err != nil {
		return big.NewInt(0), custom_error.New(err.Error())
	}
	value, ok := big.NewInt(0).SetString(valueString, 10)
	if !ok {
		return big.NewInt(0), custom_error.New("big.NewInt(0).SetString(valueString, 10) failed")
	}
	return value, nil

}

func (b *BscChain) WeiToEther(ctx context.Context, value string, symbol string) (string, error) {
	if value == "" {
		return "0", nil
	}

	var valueBigFloat *big.Float
	var err error
	// 科学记数法转为10进制
	if gstr.Contains(value, "e") {
		valueBigFloat, _, err = big.ParseFloat(value, 10, 0, big.ToNearestEven)
		if err != nil {
			return "", custom_error.New(err.Error())
		}
	} else {
		var ok bool
		valueBigFloat, ok = big.NewFloat(0).SetString(value)
		if !ok {
			return "", custom_error.New("value string to big.float error", g.Map{
				"value":  value,
				"symbol": symbol,
			})
		}
	}

	contracts, err := b.GetContracts(ctx)
	if err != nil {
		return "", err
	}
	decimal := 0
	for _, contract := range contracts {
		if contract.Symbol == symbol {
			decimal = contract.Decimals
			break
		}
	}
	if decimal == 0 {
		return "", custom_error.New("wei to ether ,symbol corresponding contract not exist ", g.Map{
			"value":     value,
			"symbol":    symbol,
			"contracts": contracts,
		})
	}

	decimalBigFloat := big.NewFloat(math.Pow(10, float64(decimal)))

	valueBigFloat = valueBigFloat.Quo(valueBigFloat, decimalBigFloat)
	return valueBigFloat.String(), nil
}

func (b *BscChain) EtherToWei(ctx context.Context, value string, symbol string) (string, error) {
	if value == "" {
		return "0", nil
	}
	valueBigFloat, ok := big.NewFloat(0).SetString(value)
	if !ok {
		return "", custom_error.New("value string to big.float error", g.Map{
			"value":  value,
			"symbol": symbol,
		})
	}
	contracts, err := b.GetContracts(ctx)

	if err != nil {
		return "", err
	}
	decimal := 0
	for _, contract := range contracts {
		if contract.Symbol == symbol {
			decimal = contract.Decimals
			break
		}
	}
	if decimal == 0 {
		return "", custom_error.New("wei to ether ,symbol corresponding contract not exist ", g.Map{
			"value":     value,
			"symbol":    symbol,
			"contracts": contracts,
		})
	}
	valueBigFloat = valueBigFloat.Mul(valueBigFloat, big.NewFloat(math.Pow(10, 18)))

	valueInt, _ := valueBigFloat.Int(big.NewInt(0))

	return valueInt.String(), nil
}

func (b *BscChain) GetDelayNumber() int {

	return b.delayNumber
}

//func (b *BscChain) getDebug() bool {
//	return b.debug
//}

func (b *BscChain) SetDelayNumber(number int) {
	b.delayNumber = number
}

func (b *BscChain) SetIsDebug(isDebug bool) {
	b.debug = isDebug
}

func (b *BscChain) SetRpcClient(client *RpcClient) {
	b.rpcClient = client
}

// 获取待检测的区块，如果数据库中有需要检测的区块，则优先返回数据库中的区块，否则返回最新的区块
// 如果获取最新的区块和上次相同，为了避免重复检测，返回0
func (b *BscChain) GetDetectNumber(ctx context.Context) (newestNumber, detectNumber int, err error) {

	newestNumber, err = b.GetBlockNumber(ctx)
	if err != nil {
		return 0, 0, custom_error.New(err.Error())
	}
	newestNumber -= b.delayNumber

	v, err := dao.LoseBlocks.Ctx(ctx).Where(dao.LoseBlocks.Columns().Status, model.LOSE_BLOCK_STATUS_WAIT).Value(dao.LoseBlocks.Columns().Number)
	if err != nil {
		return newestNumber, 0, custom_error.New(err.Error())
	}
	if v.Int() > 0 {
		// 为了避免检测后每次都查询 loseBlocks 表，在此时直接将区块改成已经检测
		_, err = dao.LoseBlocks.Ctx(ctx).Where(dao.LoseBlocks.Columns().Status, model.LOSE_BLOCK_STATUS_WAIT).Update(g.Map{
			dao.LoseBlocks.Columns().Status: model.LOSE_BLOCK_STATUS_FINISH,
		})
		if err != nil {
			return newestNumber, 0, err
		}
		return newestNumber, v.Int(), nil
	}

	cacheNumber, err := gcache.Get(ctx, model.CACHE_KEY_CHAIN_NEWEST_NUMBER)
	if err != nil {
		return newestNumber, 0, custom_error.New(err.Error())
	}
	detectNumber = newestNumber

	// 如果已经扫描过，不再扫描
	if newestNumber == cacheNumber.Int() {
		return newestNumber, 0, nil
	}
	// 区块是连续产生的，为了避免跳块，此时直接使用缓存中记录的区块加一
	if cacheNumber.Int() != 0 {
		detectNumber = cacheNumber.Int() + 1
	}

	err = gcache.Set(ctx, model.CACHE_KEY_CHAIN_NEWEST_NUMBER, detectNumber, 0)

	return newestNumber, detectNumber, err
}

func (b *BscChain) GetTokenBalance(ctx context.Context, contractAddress, address string) (balance big.Int, err error) {
	//TODO implement me
	panic("implement me")
}

func (b *BscChain) GetBlockNumber(ctx context.Context) (int, error) {
	return b.rpcClient.GetBlockNumber()
}

func (b *BscChain) GetBlockInfoByNumber(ctx context.Context, i int) (*OriginBlock, error) {
	return b.rpcClient.GetBlockInfoByNumber(i, true)
}

func (b *BscChain) GetBalance(ctx context.Context, address string) (big.Int, error) {
	return b.rpcClient.GetBalance(address)
}

func (b *BscChain) GetProbablyOnceGasPrice(ctx context.Context, symbol string) (big.Int, error) {

	gasNumberUint := b.GetGasLimit(ctx, symbol)
	gasNumber := big.NewInt(0).SetUint64(gasNumberUint)
	price, err := b.rpcClient.GetGasPrice()

	if err != nil {
		return big.Int{}, err
	}
	total := gasNumber.Mul(gasNumber, &price)
	return *total, nil
}

func (b *BscChain) GetGasPrice(ctx context.Context) (*big.Int, error) {
	gasPrice, err := b.rpcClient.GetGasPrice()
	return &gasPrice, err
}

func (b *BscChain) GetGasLimit(ctx context.Context, symbol string) uint64 {
	gasNumber := 21000
	if gstr.ToUpper(symbol) != model.CONTRACT_DEFAULT_SYMBOL {
		gasNumber = 80000
	}
	return uint64(gasNumber)
}

func (b *BscChain) sendRawTransaction(ctx context.Context, s string) (string, error) {
	return b.rpcClient.SendRawTransaction(s)
}

func (b *BscChain) GetTransactionReceipt(ctx context.Context, hash string) (any, error) {
	return b.rpcClient.GetTransactionReceipt(hash)
}

// 发送交易
func (b *BscChain) SendTransaction(ctx context.Context, legacyTx *types.LegacyTx, privateKey string, contractAddress string) (string, error) {

	if contractAddress == model.CONTRACT_DEAULT_SYMBOL_ADDRESS {
		contractAddress = ""
	}

	// contractAddress不为空,说明是erc20 交易
	if contractAddress != "" {

		dataBytes, err := b.makeERC20TransferData(legacyTx.To.String(), legacyTx.Value)
		if err != nil {
			return "", custom_error.Wrap(err, "", g.Map{
				"legacyTx": *legacyTx,
			})
		}

		toAddress := ethereumCommon.HexToAddress(contractAddress)

		legacyTx.To = &toAddress
		legacyTx.Value = big.NewInt(0)

		legacyTx.Data = dataBytes
	}
	tx := types.NewTx(legacyTx)
	g.Dump(b.chainID, tx, privateKey)
	data, err := b.signTransaction(b.chainID, tx, privateKey)
	if err != nil {
		return "", err
	}
	return b.sendRawTransaction(ctx, "0x"+data)
}

func (b *BscChain) GetLastNonce(ctx context.Context, address string) (uint64, error) {
	nonceDatabase, err := dao.QueueTask.Ctx(ctx).
		Where(dao.QueueTask.Columns().From, address).
		Where(dao.QueueTask.Columns().Status, g.Slice{model.QUEUE_TASK_STATUS_SUCCESS, model.QUEUE_TASK_STATUS_PROCESS}).
		Max(dao.QueueTask.Columns().Nonce)
	if err != nil {
		return 0, err
	}

	nonceNet, err := b.rpcClient.GetTransactionCount(address, "pending")

	if err != nil {
		return 0, err
	}
	if uint64(nonceDatabase) > uint64(nonceNet) {
		return uint64(nonceDatabase), nil
	}
	return uint64(nonceNet), nil
}

func (b *BscChain) signTransaction(chainID int64, tx *types.Transaction, privateKeyStr string) (string, error) {

	privateKey, err := b.stringToPrivateKey(privateKeyStr)
	if err != nil {
		return "", err
	}
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(chainID)), privateKey)
	if err != nil {
		return "", nil
	}

	bytes, err := rlp.EncodeToBytes(signTx)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (b *BscChain) stringToPrivateKey(privateKeyStr string) (*ecdsa.PrivateKey, error) {
	privateKeyByte, err := hexutil.Decode(privateKeyStr)
	if err != nil {
		return nil, err
	}
	privateKey, err := crypto.ToECDSA(privateKeyByte)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// MakeERC20TransferData 交易数据
func (b *BscChain) makeERC20TransferData(toAddress string, amount *big.Int) ([]byte, error) {
	methodID := crypto.Keccak256Hash([]byte("transfer(address,uint256)"))
	var data []byte
	data = append(data, methodID[:4]...)
	paddedAddress := ethereumCommon.LeftPadBytes(ethereumCommon.HexToAddress(toAddress).Bytes(), 32)
	data = append(data, paddedAddress...)
	paddedAmount := ethereumCommon.LeftPadBytes(amount.Bytes(), 32)
	data = append(data, paddedAmount...)
	return data, nil
}
