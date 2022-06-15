package deposit_withdraw

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"gf-admin/utility/custom_error"
	"math"
	"math/big"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

type Transaction struct {
	Symbol          string //
	SymbolDecimal   int
	ContractAddress string
	Hash            string    // 交易hash
	Nonce           int       // 交易nonce
	From            string    // 转出地址
	To              string    // 转入地址
	Value           big.Int   // 转账金额,单位为wei
	ValueDecimal    big.Float // 转账金额,单位为ether
	GasUsed         int       // gas 实际消耗
	GasPrice        big.Int   // gas 实际价格
	BlockNumber     int       // 块号
	BlockHash       string    // 块hash
	Type            string    // 转账类型
	UserID          uint      // 用户id
}

// NewTransaction 需要判断tx交易类型是eth还是合约
// 如果是合约，则需要判断是否是需要归集的合约
// 如果是合约，则需要判断是否是转账类型
func NewTransaction(ctx context.Context, origin *OriginTransaction) (*Transaction, error) {
	t := &Transaction{
		Hash:        origin.Hash,
		Nonce:       origin.Nonce,
		From:        origin.From,
		To:          origin.To,
		Value:       origin.Value,
		GasUsed:     origin.Gas,
		GasPrice:    origin.GasPrice,
		BlockNumber: *origin.BlockNumber,
		BlockHash:   origin.BlockHash,
	}

	if origin.Input == "0x" { //eth交易
		t.Symbol = model.CONTRACT_DEFAULT_SYMBOL
		t.SymbolDecimal = model.CONTRACT_DEAULT_SYMBOL_DECIMAL
		t.ContractAddress = model.CONTRACT_DEAULT_SYMBOL_ADDRESS
	} else { //合约交易

		contracts, err := ChainClient.GetContracts(ctx)
		if err != nil {
			return nil, err
		}
		var contract entity.Contracts
		var ok bool
		if contract, ok = contracts[gstr.ToLower(origin.To)]; !ok {
			return nil, nil
		}
		if len(origin.Input) < 11 { //transfer的长度肯定大于11，所以不是转账交易
			return nil, nil
		}

		methodName := origin.Input[:10]
		if methodName != model.CONTRACT_TRANSFER_FUCTION_HEX {
			return nil, nil
		}

		t.ContractAddress = gstr.ToLower(t.To)
		t.Symbol = contract.Symbol
		t.SymbolDecimal = contract.Decimals

		t.To = gstr.ToLower("0x" + origin.Input[34:74]) //实际to转账地址

		value := origin.Input[74:] //实际转账金额

		_, ok = t.Value.SetString(value, 16) //将金额转为10进制

		if !ok {
			return nil, custom_error.New("合约交易设置转账金额失败")
		}
	}

	t.ValueDecimal.Quo(new(big.Float).SetInt(&t.Value), new(big.Float).SetFloat64(math.Pow(10, float64(t.SymbolDecimal))))

	//判断具体的交易类型，是充值，归集，手续费，还是提现
	userAddresses, err := ChainClient.GetUserAddresses(ctx)
	if err != nil {
		return nil, err
	}
	feeWithdrawAddress, err := ChainClient.GetFeeWithdrawAddress(ctx)
	if err != nil {
		return nil, err
	}
	collectAddress, err := ChainClient.GetCollectAddress(ctx)
	if err != nil {
		return nil, err
	}

	fromUser, fromAddressIsUser := userAddresses[t.From]
	toUser, toAddressIsUser := userAddresses[t.To]

	logInfofDw(ctx, "tx symbol is %s,from address is %s, to address is %s, feewithdraw address is %s,collect address is %s \n", t.Symbol, t.From, t.To, feeWithdrawAddress, collectAddress)
	// 充值: 转入地址是平台生成的用户地址,转出地址不是手续费提现地址为充值
	if toAddressIsUser && t.From != feeWithdrawAddress {

		t.Type = model.TRANSACTION_TYPE_REACHRGE
		t.UserID = toUser.Id

		//手续费: 转出地址是手续费提现地址地址,转入地址是 用户地址,交易需要为eth交易 为转出手续费
	} else if t.From == feeWithdrawAddress && toAddressIsUser && t.Symbol == model.CONTRACT_DEFAULT_SYMBOL {

		t.Type = model.TRANSACTION_TYPE_FEE
		t.UserID = toUser.Id

		//提现： 转出地址是手续费提现地址为提现 ,并且提现表中存在提现记录
	} else if t.From == feeWithdrawAddress {

		userIdVar, err := dao.Withdraws.Ctx(ctx).Where(dao.Withdraws.Columns().Hash, t.Hash).Value(dao.Withdraws.Columns().UserId)
		if err != nil {
			return nil, custom_error.Wrap(err, "查询提现记录失败", g.Map{
				"hash": t.Hash,
			})
		}
		if userIdVar.Int() > 0 {
			t.Type = model.TRANSACTION_TYPE_WITHDRAW
			t.UserID = userIdVar.Uint()
		}
		//归集： 转出地址是用户地址,转入地址是归集地址
	} else if fromAddressIsUser && t.To == collectAddress {

		t.Type = model.TRANSACTION_TYPE_COLLECT
		t.UserID = fromUser.Id
	}
	return t, nil
}
