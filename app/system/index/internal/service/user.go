package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"gf-admin/utility/response"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/v2/text/gstr"
)

var User = user{}

type user struct {
}

func (u *user) canLogin(ctx context.Context, userId uint) (bool, error) {
	status, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, userId).Value(dao.Users.Columns().Status)
	if err != nil {
		return false, err
	}
	if gstr.Contains(status.String(), model.USER_STATUS_DISABLE_LOGIN) {
		return false, nil
	}

	return true, nil
}

func (u *user) canPost(ctx context.Context, userId uint) (bool, error) {
	status, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, userId).Value(dao.Users.Columns().Status)
	if err != nil {
		return false, err
	}
	if gstr.Contains(status.String(), model.USER_STATUS_DISABLE_POST) {
		return false, nil
	}

	return true, nil
}

func (u *user) canReply(ctx context.Context, userId uint) (bool, error) {
	status, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, userId).Value(dao.Users.Columns().Status)
	if err != nil {
		return false, err
	}
	if gstr.Contains(status.String(), model.USER_STATUS_DISABLE_REPLY) {
		return false, nil
	}

	return true, nil
}

func (u *user) balance(ctx context.Context, userId uint) (uint, error) {
	balanceVar, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, userId).Value(dao.Users.Columns().Balance)

	return balanceVar.Uint(), err
}

func (u *user) changeBalance(ctx context.Context, userId uint, amount int, changeType model.BalanceChangeType, relationId uint, remark string) error {

	if amount == 0 {
		return nil
	}
	return dao.Users.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {

		balanceVar, err := dao.Users.Ctx(ctx).WherePri(userId).LockUpdate().Value(dao.Users.Columns().Balance)
		balance := balanceVar.Uint()

		g.Dump(amount, balance, 3333)
		if amount < 0 && balance < uint(-amount) {
			return response.NewError("积分不足")
		}

		_, err = dao.Users.Ctx(ctx).WherePri(userId).Increment(dao.Users.Columns().Balance, amount)
		if err != nil {
			return err
		}
		username, err := dao.Users.Ctx(ctx).WherePri(userId).Value(dao.Users.Columns().Username)
		if err != nil {
			return err
		}
		var log = entity.BalanceChangeLog{
			UserId:     userId,
			Username:   username.String(),
			Type:       string(changeType),
			Amount:     amount,
			Before:     uint(balance),
			After:      uint(int(balance) + amount),
			RelationId: relationId,
			Remark:     remark,
		}
		_, err = dao.BalanceChangeLog.Ctx(ctx).Insert(log)
		return err
	})

}
