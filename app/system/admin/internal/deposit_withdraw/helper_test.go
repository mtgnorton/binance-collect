package deposit_withdraw

import (
	"fmt"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func Test_IntToHex(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		r := IntToHex[int](1)
		t.Assert(r, "0x1")
	})

	gtest.C(t, func(t *gtest.T) {
		r := IntToHex[uint](uint(1))
		t.Assert(r, "0x1")

	})
}

func TestCreateAddress(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		address, privateKey, err := CreateAddress()
		t.Assert(err, nil)
		fmt.Println(address)
		fmt.Println(privateKey)
	})
}
