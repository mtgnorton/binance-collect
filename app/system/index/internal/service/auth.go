package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/utility"
	"gf-admin/utility/response"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/util/gconv"
)

var Auth = auth{}

type auth struct {
}

func (a *auth) Register(ctx context.Context, in *define.AuthRegisterInput) (err error) {
	if !shared.Captcha.Verify(ctx, in.Code, in.CaptchaId) {
		return response.NewError("验证码错误")

	}
	d := dao.Users.Ctx(ctx)
	count, err := d.Where(dao.Users.Columns().Username, in.Username).WhereOr(dao.Users.Columns().Email, in.Email).Count()
	if err != nil {
		return response.WrapError(err, "注册失败")
	}
	if count > 0 {
		return response.NewError("用户名或邮箱已存在")
	}
	var user *entity.Users

	if err = gconv.Struct(in, &user); err != nil {
		return err
	}
	user.Password = utility.EncryptPassword(user.Username, user.Password)

	err = dao.Users.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		userId, err := dao.Users.Ctx(ctx).OmitEmptyData().InsertAndGetId(user)
		if err != nil {
			return err
		}

		settingRegisterGiveAmount, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_TOKEN_REGISTER_GIVE)
		if err != nil {
			return err
		}

		err = User.changeBalance(ctx, uint(userId), settingRegisterGiveAmount.Int(), model.BALANCE_CHANGE_TYPE_REGSITER, 0, "")

		return err
	})
	if err != nil {
		return response.WrapError(err, "注册失败")
	}

	return response.NewSuccess("注册成功")
}

func (a *auth) Login(ctx context.Context, in *define.AuthLoginInput) (out *define.AuthLoginOutput, err error) {
	out = &define.AuthLoginOutput{}
	if !shared.Captcha.Verify(ctx, in.Code, in.CaptchaId) {
		return out, response.NewError("验证码错误")
	}

	var user entity.Users
	err = dao.Users.Ctx(ctx).Where(g.Map{
		dao.Users.Columns().Username: in.Username,
		dao.Users.Columns().Password: utility.EncryptPassword(in.Username, in.Password),
	}).Scan(&user)

	if err != nil {
		return out, response.WrapError(err, "用户名或密码错误")
	}

	if user.Id == 0 {
		return out, response.NewError("用户名或密码错误")
	}

	// 每日登录赠送积分
	count, err := dao.BalanceChangeLog.Ctx(ctx).
		Where(dao.BalanceChangeLog.Columns().UserId, user.Id).
		WhereGTE(dao.BalanceChangeLog.Columns().CreatedAt, gtime.Date()).
		Where(dao.BalanceChangeLog.Columns().Type, model.BALANCE_CHANGE_TYPE_LOGIN).Count()
	if err != nil {
		return out, response.WrapError(err, "登录失败")
	}
	if count == 0 {
		settingLoginGiveAmount, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_TOKEN_LOGIN_GIVE)
		if err != nil {
			return out, response.WrapError(err, "登录失败")
		}
		err = User.changeBalance(ctx, user.Id, settingLoginGiveAmount.Int(), model.BALANCE_CHANGE_TYPE_LOGIN, 0, "")
		if err != nil {
			return out, response.WrapError(err, "登录失败")
		}
	}

	// todo 登录状态判断
	token, err := FrontTokenInstance.LoadConfig().TokenHandler.GenerateAndSaveData(ctx, in.Username, user)
	if err != nil {
		return
	}
	out.Token = token
	shared.Context.SetUser(ctx, user)
	return out, response.NewSuccess("登录成功")

}

func (a *auth) Info(ctx context.Context) (out *define.AuthInfoOutput, err error) {
	out = &define.AuthInfoOutput{}

	userId, err := FrontTokenInstance.GetUserId(ctx)
	if err != nil {
		return
	}

	err = dao.Users.Ctx(ctx).WherePri(userId).Scan(&out)

	if err != nil {
		return out, response.WrapError(err, "获取失败")
	}
	return out, response.NewSuccess("获取成功")
}

func (a *auth) Logout(ctx context.Context) (err error) {
	token, err := FrontTokenInstance.TokenHandler.GetTokenFromRequest(ctx, g.RequestFromCtx(ctx))
	if err != nil {
		return
	}
	err = FrontTokenInstance.Remove(ctx, token)
	if err != nil {
		return err
	}
	return response.NewSuccess("退出成功")

}
