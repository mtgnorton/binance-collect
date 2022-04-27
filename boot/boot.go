package boot

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model/entity"
	_ "gf-admin/packed"
	"gf-admin/utility"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
)

var EnvName = "prod"

func init() {

	ctx := context.TODO()
	loadConfigFile(ctx)
	newDefaultAdministrator(ctx)
}

// 根据环境变量或命令行参数家在对应的配置文件，命令行参数优先级高于环境变量,默认使用prod环境
func loadConfigFile(ctx context.Context) {
	// 获取环境变量
	EnvName = gcmd.GetOptWithEnv("gf_admin_env_file").String()
	g.Log().Infof(ctx, "从命令或环境变量获取环境：%s", EnvName)

	if EnvName == "" {
		EnvName = "prod"
		g.Log().Infof(ctx, "从命令 或环境变量获取环境失败，使用默认环境：%s", EnvName)

	}
	switch EnvName {
	case "dev":
		g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config-dev.toml")
	case "prod":
		g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config-prod.toml")
	case "local":
		g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config-local.toml")
	}
	g.Log().Infof(ctx, "当前配置环境：%s", EnvName)

}

//系统启动时，判断是否存在管理员，如果不存在，则创建一个管理员
func newDefaultAdministrator(ctx context.Context) {
	administrator, err := dao.Administrator.Ctx(ctx).Where("username", "admin").One()
	if err != nil {
		g.Log().Fatalf(ctx, "query administrator error ,error msg following: %s", err)
	}
	if administrator.IsEmpty() {
		g.Log().Debug(ctx, "开始创建admin管理员")

		_, err = dao.Administrator.Ctx(ctx).Insert(entity.Administrator{
			Username: "admin",
			Password: utility.EncryptPassword("admin", "admin"),
			Nickname: "admin",
			Avatar:   "",
			Status:   "",
			Remark:   "",
		})
		if err != nil {
			g.Log().Fatalf(ctx, "init admin error,%s", err)
		}
	} else {
		g.Log().Debug(ctx, "admin管理员已经存在")
	}
}
