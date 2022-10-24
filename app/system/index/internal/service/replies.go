package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/utility/response"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/text/gregex"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/text/gstr"
)

var Reply = reply{}

type reply struct {
}

func (c *reply) Store(ctx context.Context, req *define.ReplyStoreReq) (res *define.ReplyStoreRes, err error) {
	res = &define.ReplyStoreRes{}

	// 判断回复字符长度
	contentLength := gstr.LenRune(req.Content)
	settingReplyCharacterMax, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_REPLY_CHARACTER_MAX)
	if err != nil {
		return res, response.WrapError(err, "回复失败")
	}
	if gconv.Int(settingReplyCharacterMax) > 0 && contentLength > gconv.Int(settingReplyCharacterMax) {
		return res, response.NewError("内容过长")
	}

	// 判断主题是否存在
	var post entity.Posts
	err = dao.Posts.Ctx(ctx).Where(dao.Posts.Columns().Id, req.PostId).Scan(&post)
	if err != nil {
		return res, response.WrapError(err, "主题不存在")
	}

	// 获取需要扣除的余额
	settingNeedToken, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_TOKEN_ESTABLISH_REPLY_DEDUCT)

	if err != nil {
		return res, response.WrapError(err, "创建主题失败")
	}

	// 判断是否达到最大发帖数量
	user, err := FrontTokenInstance.GetUser(ctx)
	if err != nil {
		return res, response.WrapError(err, "回复失败")
	}
	settingDayReplyMax, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_REPLY_EVERY_DAY_MAX)
	if err != nil {
		return res, response.WrapError(err, "回复失败")
	}

	// 获取用户今日发帖的数量
	replyCount, err := dao.Replies.Ctx(ctx).
		Where(dao.Replies.Columns().UserId, user.Id).
		WhereGTE(dao.Replies.Columns().CreatedAt, gtime.Now().Format("Y-m-d")+" 00:00:00").
		Count()
	if err != nil {
		return res, response.WrapError(err, "回复失败")
	}

	if gconv.Int(settingDayReplyMax) > 0 && replyCount > gconv.Int(settingDayReplyMax) {
		return res, response.NewError("超过今日发帖数量")
	}

	//判断是否需要审核
	settingReplyNeedAudit, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_REPLY_IS_NEED_AUDIT)
	if err != nil {
		return res, response.WrapError(err, "回复失败")
	}
	status := model.REPLY_STATUS_NORMAL
	if settingReplyNeedAudit.Int() == 1 {
		status = model.REPLY_STATUS_NO_AUDIT
	}

	// 解析回复内容中涉及到的所有其它用户
	matches, err := gregex.MatchAllString(`@(\w+)[^\w]? `, req.Content)
	if err != nil {
		return res, response.WrapError(err, "回复失败")
	}
	relationUserIds := ""

	for _, match := range matches {
		// 获取用户
		var user entity.Users
		err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Username, match[1]).Scan(&user)
		if err != nil {
			continue
		}

		relationUserIds += gconv.String(user.Id) + ","
	}

	// 扣除余额,插入回复并修改主题最后回复相关信息
	err = dao.Replies.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {

		err = User.changeBalance(ctx, user.Id, -gconv.Int(settingNeedToken), model.BALANCE_CHANGE_TYPE_ESTABLISH_REPLY, gconv.Uint(req.PostId), "回复主题")

		if err != nil {
			return err
		}
		_, err = dao.Replies.Ctx(ctx).Insert(&entity.Replies{
			PostsId:         req.PostId,
			UserId:          user.Id,
			Username:        user.Username,
			Content:         req.Content,
			CharacterAmount: gconv.Uint(contentLength),
			RelationUserIds: relationUserIds,
			Status:          status,
		})
		if err != nil {
			return err
		}
		_, err = dao.Posts.Ctx(ctx).WherePri(req.PostId).Update(g.Map{
			dao.Posts.Columns().ReplyLastTime:     gtime.Now().Format("Y-m-d H:i:s"),
			dao.Posts.Columns().ReplyLastUserId:   user.Id,
			dao.Posts.Columns().ReplyLastUsername: user.Username,
			dao.Posts.Columns().ReplyAmount:       gdb.Raw(dao.Posts.Columns().ReplyAmount + " + 1"),
		})
		return err
	})
	return
}
