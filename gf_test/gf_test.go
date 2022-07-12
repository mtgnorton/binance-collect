package test_gf

import (
	"testing"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gogf/gf/v2/test/gtest"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
)

// 动态设置配置值
func TestSetConfig(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		ctx := gctx.New()

		g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config-unit.toml")

		// 动态设置配置值
		err := g.Cfg().GetAdapter().(*gcfg.AdapterFile).Set("database.test", "1111")

		g.Dump(err)

		// 获取所有配置的内容
		g.Dump(g.Cfg().Data(ctx))

	})

}
