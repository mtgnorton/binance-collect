package api

import (
	"context"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
)

var Administrator = administrator{}

type administrator struct {
}

func (a *administrator) List(ctx context.Context,req *define.) {

}

func (a *administrator) Login(ctx context.Context, req *define.LoginPostReq) (res *define.LoginPostRes, err error) {
	res = &define.LoginPostRes{
		"登录成功",
	}

	err = service.Administrator.Login(ctx, define.AdministratorLoginInput{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		res.Message = err.Error()
	}
	return
}
