package deposit_withdraw

import (
	"crypto/ecdsa"
	"fmt"
	"gf-admin/app/model"
	"gf-admin/utility/custom_error"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogf/gf/v2/text/gstr"
	"math/big"
	"strconv"
	"strings"
)

// IntToHex convert int to hexadecimal representation
func IntToHex[T model.Integer](i T) string {
	return fmt.Sprintf("0x%x", i)
}

// ParseInt parse hex string value to int
func ParseInt(value string) (int, error) {
	i, err := strconv.ParseInt(strings.TrimPrefix(value, "0x"), 16, 64)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

// ParseBigInt parse hex string value to big.Int
func ParseBigInt(value string) (big.Int, error) {
	i := big.Int{}
	_, ok := i.SetString(strings.TrimPrefix(value, "0x"), 16)
	if !ok {
		return i, custom_error.New(fmt.Sprintf("invalid hex string %s", value))
	}

	return i, nil
}

// BigToHex covert big.Int to hexadecimal representation
func BigToHex(bigInt big.Int) string {
	if bigInt.BitLen() == 0 {
		return "0x0"
	}

	return "0x" + strings.TrimPrefix(fmt.Sprintf("%x", bigInt.Bytes()), "0")
}

// 创建eth地址,返回eth地址和私钥
func CreateAddress() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", custom_error.New(err.Error())
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	privateKeyString := hexutil.Encode(privateKeyBytes)

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		return "", "", custom_error.Wrap(err, "cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	address = gstr.ToLower(address)

	return address, privateKeyString, nil
}
