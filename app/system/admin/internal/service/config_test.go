package service

import (
	"gf-admin/app/system/admin/internal/define"
	"testing"

	"github.com/gogf/gf/v2/os/gcfg"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestConfig_GetModules(t *testing.T) {
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config-unit.toml") //单元测试时，日志输出会作为响应返回，所以暂时建立专门的测试配置文件关闭所有log输出
	gtest.C(t, func(t *gtest.T) {
		out, err := Config.GetModules(ctx, &define.ConfigListInput{
			Modules: []string{"backend", "binance"},
		})
		t.AssertNil(err)
		g.Dump(out)
	})

}
