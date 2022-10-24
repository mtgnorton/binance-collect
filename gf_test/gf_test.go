package test_gf

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"

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

func TestStruct(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Name string
			Age  int
		}
		var u User
		var u1 *User
		g.DumpWithType(u)
		g.Dump(u.Age)
		g.DumpWithType(u1)
	})
}

func TestStrCount(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		c := "我是1s"
		g.Dump(gstr.LenRune(c))
	})
}

func TestMatch(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// 解析回复内容中涉及到的所有其它用户
		matches, err := gregex.MatchAllString(`@(\w+)[^\w]?`, "fasdffasdf@fas1df @grdfx adsffsda @fff")
		g.Dump(matches, err)
	})
}

func TestMatchOfficial(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		buf := "fasdffasdf@fas1df @grdfx adsffsda @fff"
		//解析正则表达式，如果成功返回解释器
		reg1 := regexp.MustCompile(`@(\w+)[^\w]*`)
		if reg1 == nil { //解释失败，返回nil
			fmt.Println("regexp err")
			return
		}
		//根据规则提取关键信息
		result1 := reg1.FindAllStringSubmatch(buf, -1)
		fmt.Println("result1 = ", result1)
	})
}
