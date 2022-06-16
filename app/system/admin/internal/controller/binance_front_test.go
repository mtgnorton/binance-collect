package controller

import (
	"gf-admin/app/system/admin/internal/define"
	"testing"

	"github.com/gogf/gf/v2/util/grand"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

func TestBinanceApi_CreateAddress(t *testing.T) {

	type RES struct {
		Code    int                     `json:"code"`
		Message string                  `json:"message"`
		Data    define.CreateAddressRes `json:"data"`
	}
	var res1, res2, res3 RES

	client := NewTestClient(NewTestServer(func(group *ghttp.RouterGroup) {
		group.Bind(
			new(binanceApi),
		)
	}))

	gtest.C(t, func(t *gtest.T) {

		content1 := client.GetContent(ctx, "dw-create-address", "user_id=ff")

		err := gconv.Scan(content1, &res1)

		t.Assert(err, nil)

		t.Assert(res1.Message, "用户ID必须是整数")

		content2 := client.GetContent(ctx, "dw-create-address", "user_id=6")
		err = gconv.Scan(content2, &res2)
		t.Assert(err, nil)
		resCopy := res2
		t.Assert(res2.Code, 0)
		if len(res2.Data.Address) == 0 {
			t.Error("address is empty")
		}

		content3 := client.GetContent(ctx, "dw-create-address", "user_id=6")

		err = gconv.Scan(content3, &res3)
		t.Assert(err, nil)
		t.Assert(resCopy.Data.Address, res3.Data.Address)
	})

}

func TestBinanceApi_ApplyWithdraw(t *testing.T) {
	type RES struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	var res1 RES

	client := NewTestClient(NewTestServer(func(group *ghttp.RouterGroup) {
		group.Bind(
			new(binanceApi),
		)
	}))

	gtest.C(t, func(t *gtest.T) {
		//"order_id=1&user_id=1&value=0.1&to_address=0x991195b40a5bDF4725AfbD4f10F579BCa25308F5&user_address=0x84b2d9c9b870ca47719e17e8cf790d66686743c8&Symbol=BNB"
		c1 := client.PostContent(ctx, "/dw-apply-withdraw", g.Map{
			"ExternalOrderId": grand.N(1, 100000),
			"ExternalUserId":  1,
			"Value":           0.1,
			"To":              "0x991195b40a5bDF4725AfbD4f10F579BCa25308F5",
			"BinanceAdmin":    "0x84b2d9c9b870ca47719e17e8cf790d66686743c8",
			"Symbol":          "BNB",
		})

		err := gconv.Scan(c1, &res1)

		g.Dump(res1)
		t.Assert(err, nil)

		t.Assert(res1.Message, "")

	})

}
