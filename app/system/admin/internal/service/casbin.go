package service

import (
	"context"
	"gf-admin/app/dao"

	"github.com/gogf/gf/v2/os/gres"

	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	gdb "github.com/vance-liu/gdb-adapter"
)

var Enforcer = enforcer{
	modelConfigPath: "config/casbin_model.conf",
}

type enforcer struct {
	modelConfigPath string
	instance        *casbin.Enforcer
}

func (e *enforcer) Register(ctx context.Context) {
	gv, err := g.Cfg().Get(ctx, "casbin")

	if err != nil {
		g.Log().Fatalf(ctx, "casbin 配置读取错误,错误信息如下:%s", err)
	}
	configs := gv.MapStrStr()
	g.Log().Debug(ctx, "casbin mysql config following", configs)
	opts := &gdb.Adapter{
		DriverName:     configs["driverName"],
		DataSourceName: configs["dataSourceName"],
		TableName:      configs["tableName"],
		// or reuse an existing connection:
		// Db: yourDBConn,
	}

	adapter, err := gdb.NewAdapterFromOptions(opts)

	if err != nil {
		g.Log().Fatalf(ctx, "casbin adapter init error,error info following:%s", err)
	}
	//程序的执行路径
	var absolutePath = gfile.Join(gfile.SelfPath(), e.modelConfigPath)

	if !gfile.IsFile(absolutePath) {
		//源码的路径
		absolutePath = gfile.Join(gfile.Pwd(), e.modelConfigPath)
	}

	if !gfile.IsFile(absolutePath) {
		content := gres.GetContent("config/casbin_model.conf")
		if string(content) == "" {
			panic("casbin model config not found")
		}
		err := gfile.PutBytes(absolutePath, content)
		if err != nil {
			panic(err)
		}

	}

	e.instance, err = casbin.NewEnforcer(absolutePath, adapter)
	if err != nil {
		g.Log().Fatalf(ctx, "casbin new enforce error,error info following: %s", err)
	}
	err = e.instance.LoadPolicy()

	if err != nil {
		g.Log().Fatalf(ctx, "casbin load policy error,error info following: %s", err)
	}
	g.Log().Debug(ctx, "casbin start finish")
}

func (e *enforcer) AddPermissionForUserOrRole(ctx context.Context, userOrRole string, permissions ...string) (bool, error) {
	return e.instance.AddPermissionForUser(userOrRole, permissions...)

}

func (e *enforcer) DeletePermissionsForUser(ctx context.Context, userOrRole string) (bool, error) {
	return e.instance.DeletePermissionsForUser(userOrRole)
}

// 添加用户角色关联关系
func (e *enforcer) AddRoleForUser(username, roleName string) (bool, error) {
	return e.instance.AddRoleForUser(username, roleName)
}

// 删除用户角色关联关系
func (e *enforcer) DeleteRoleForUser(username, roleName string) (bool, error) {
	return e.instance.DeleteRoleForUser(username, roleName)
}

//当删除管理员时，删除casbin中对应的数据
func (e *enforcer) DeleteUser(user string) (bool, error) {
	return e.instance.DeleteUser(user)
}

//当更新菜单时，更新casbin中对应的数据,因为驱动器没有实现相关方法，自己实现
func (e *enforcer) UpdatePolicy(old []string, new []string) (bool, error) {
	if len(old) != 2 || len(new) != 2 {
		return false, gerror.NewCode(gcode.CodeInvalidParameter, "更新casbin策略参数错误")
	}
	r, err := dao.CasbinRule.Ctx(gctx.New()).Where("ptype='p' && v1 = ? && v2= ?", old[0], old[1]).Update(g.Map{
		"v1": new[0],
		"v2": gstr.ToLower(new[1]),
	})
	if err != nil {
		return false, err
	}
	row, err := r.RowsAffected()
	return row > 0, err
}

func (e *enforcer) Auth(username, path, method string) (bool, error) {
	return e.instance.Enforce(username, gstr.ToLower(path), gstr.ToLower(method))
}
