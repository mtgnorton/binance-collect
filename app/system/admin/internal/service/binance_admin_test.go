package service

import (
	"gf-admin/app/model"
	"gf-admin/app/system/admin/internal/define"
	"testing"

	"github.com/gogf/gf/v2/os/gcfg"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestBinanceAdmin_UserAddressList(t *testing.T) {

	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config-unit.toml") //单元测试时，日志输出会作为响应返回，所以暂时建立专门的测试配置文件关闭所有log输出
	gtest.C(t, func(t *gtest.T) {
		var out, err = BinanceAdmin.UserAddressList(ctx, &define.UserAddressListInput{
			Address:        "",
			ExternalUserId: 0,
			PageSizeInput: model.PageSizeInput{
				Page: 1,
				Size: 10,
			},
		})

		t.AssertNil(err)
		g.Dump(out)
	})
}

func TestBinanceAdmin_CollectList(t *testing.T) {

	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config-unit.toml") //单元测试时，日志输出会作为响应返回，所以暂时建立专门的测试配置文件关闭所有log输出
	gtest.C(t, func(t *gtest.T) {
		var out, err = BinanceAdmin.CollectList(ctx, &define.CollectListInput{
			PageSizeInput: model.PageSizeInput{
				Page: 1,
				Size: 10,
			},
		})

		t.AssertNil(err)
		g.Dump(out)
	})
}

func TestBinanceAdmin_WithdrawList(t *testing.T) {
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config-unit.toml") //单元测试时，日志输出会作为响应返回，所以暂时建立专门的测试配置文件关闭所有log输出

	gtest.C(t, func(t *gtest.T) {
		var out, err = BinanceAdmin.WithdrawList(ctx, &define.WithdrawListInput{
			PageSizeInput: model.PageSizeInput{
				Page: 1,
				Size: 10,
			},
		})
		t.AssertNil(err)
		g.Dump(out)
	})
}
