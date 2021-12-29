package boot

import (
	"gf-admin/app/dao"
	"gf-admin/app/model"
	_ "gf-admin/packed"
	"gf-admin/utility"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func init() {

	ctx := gctx.New()
	entity, err := dao.Administrator.Ctx(ctx).Where("username", "admin").One()
	if err != nil {
		g.Log().Fatalf(ctx, "query administrator error ,error msg following: %s", err)
	}
	if entity.IsEmpty() {
		g.Log().Debug(ctx, "开始创建admin管理员")

		dao.Administrator.Ctx(ctx).Insert(model.Administrator{
			Username:      "admin",
			Password:      utility.EncryptPassword("admin","admin"),
			Nickname:      "admin",
			Avatar:        "",
			Status:        "",
			Remark:        "",

		})
	}else{
		g.Log().Debug(ctx,"admin管理员已经存在")
	}
}
