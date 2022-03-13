package boot

import (
	"gf-admin/app/dao"
	"gf-admin/app/model/entity"
	_ "gf-admin/packed"
	"gf-admin/utility"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func init() {

	ctx := gctx.New()
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
