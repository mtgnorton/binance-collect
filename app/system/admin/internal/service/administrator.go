package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/utility"
	"gf-admin/utility/error_code"
	"github.com/gogf/gf/v2/frame/g"
)

var Administrator = administratorService{}

type administratorService struct {
}

func (a *administratorService) Login(ctx context.Context, in define.AdministratorLoginInput) error {

	entity, err := a.GetUserByPassportAndPassword(
		ctx,
		in.Username,
		utility.EncryptPassword(in.Username, in.Password),
	)
	if err != nil {
		return err
	}
	if entity == nil {
		return error_code.NewParams("用户名或密码错误")
	}
	if err := Session.SetAdministrator(ctx, entity); err != nil {
		return err
	}
	// 自动更新上线
	Context.SetUser(ctx, &define.ContextAdministrator{
		Id:       entity.Id,
		Username: entity.Username,
		Nickname: entity.Nickname,
		Avatar:   entity.Avatar,
	})
	return nil
}

// 根据账号和密码查询用户信息，一般用于账号密码登录。
// 注意password参数传入的是按照相同加密算法加密过后的密码字符串。
func (a *administratorService) GetUserByPassportAndPassword(ctx context.Context, username, password string) (administrator *model.Administrator, err error) {
	err = dao.Administrator.Ctx(ctx).Where(g.Map{
		dao.Administrator.Columns.Username: username,
		dao.Administrator.Columns.Password: password,
	}).Scan(&administrator)
	return
}
