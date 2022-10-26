package service

import (
	"context"
	"fmt"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/utility"
	"gf-admin/utility/response"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
)

var Posts = posts{}

type posts struct {
}

// 获取首页的主题列表
func (p *posts) IndexList(ctx context.Context) (res *model.PostPageList, err error) {
	return p.List(ctx, &model.PostPageInput{
		Period:  model.Day7,
		IsIndex: true,
		PageSizeInput: model.PageSizeInput{
			Page: 1,
			Size: 20,
		},
	})
}

// 获取主题列表
func (p *posts) List(ctx context.Context, in *model.PostPageInput) (res *model.PostPageList, err error) {
	res = &model.PostPageList{}
	d := dao.Posts.Ctx(ctx)

	if in.Period > 0 {

		d = d.WhereGTE(dao.Posts.Columns().CreatedAt, gtime.Now().Add(-time.Duration(in.Period)*time.Hour*24))
	}

	if in.NodeId > 0 {
		d = d.Where(dao.Posts.Columns().NodeId, in.NodeId)
	}
	if in.UserId > 0 {
		d = d.Where(dao.Posts.Columns().UserId, in.UserId)
	}

	if in.FilterKeyword != "" {
		d.Where(d.Builder().
			WhereLike(dao.Posts.Columns().Title, "%"+in.FilterKeyword+"%").
			WhereOrLike(dao.Posts.Columns().Content, "%"+in.FilterKeyword+"%"))
	}

	res.Page = in.Page
	res.Size = in.Size
	res.Total, err = d.Count()
	if err != nil {
		return
	}
	if in.IsIndex {
		d = d.Page(in.Page, in.Size).
			Order(gdb.Raw("if(top_end_time>now(),1,0) desc")).
			OrderDesc(dao.Posts.Columns().Weight).
			OrderDesc(dao.Posts.Columns().CreatedAt)
	} else {
		d = d.Page(in.Page, in.Size).
			OrderDesc(dao.Posts.Columns().CreatedAt)
	}

	// nodes.id=posts.node_id
	d = d.LeftJoin(dao.Nodes.Table(), fmt.Sprintf("%s.%s=%s.%s", dao.Posts.Table(), dao.Posts.Columns().NodeId, dao.Nodes.Table(), dao.Nodes.Columns().Id))

	err = d.FieldsPrefix(dao.Posts.Table(), "*").Fields("forum_nodes.name as node_name").Scan(&res.List)

	for _, post := range res.List {

		recentTime := post.CreatedAt
		if post.ReplyLastUserId > 0 {
			recentTime = post.ReplyLastTime
		}
		post.LastChangeTime, _ = utility.TimeFormatDivide24Hour(recentTime)

	}
	return

}

/**
 创建主题
 创建主题的前提条件
	- 标题字符长度小于255
	- 内容长度小于后台设定的长度
	- 用户没有被禁止发帖
	- 用户今天发帖数量没有超过后台设定的数量
	- 用户积分大于后台设定的创建主题需要的积分
*/
func (p *posts) Store(ctx context.Context, req *define.PostsStoreReq) (res *define.PostsStoreRes, err error) {

	res = &define.PostsStoreRes{}

	nodeCount, err := dao.Nodes.Ctx(ctx).Where(dao.Nodes.Columns().Id, req.NodeId).Count()
	if err != nil {
		return res, response.WrapError(err, "创建主题失败")
	}

	if nodeCount == 0 {
		return res, response.NewError("节点不存在")
	}
	if gstr.LenRune(req.Title) > 255 {
		return res, response.NewError("标题过长")
	}

	contentLength := gstr.LenRune(req.Content)
	settingPostsCharacterMax, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_POSTS_CHARACTER_MAX)
	if err != nil {
		return res, response.WrapError(err, "创建主题失败")
	}
	if gconv.Int(settingPostsCharacterMax) > 0 && contentLength > gconv.Int(settingPostsCharacterMax) {
		return res, response.NewError("内容过长")
	}

	user, err := FrontTokenInstance.GetUser(ctx)
	if err != nil {
		return res, response.WrapError(err, "创建主题失败")
	}
	if ok, _ := User.canPost(ctx, user.Id); !ok {
		return res, response.NewError("禁止发帖")
	}

	settingPostsDayMax, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_POSTS_EVERY_DAY_MAX)

	hasPostsAmount, err := dao.Posts.Ctx(ctx).
		Where(dao.Posts.Columns().UserId, user.Id).
		WhereBetween(dao.Posts.Columns().CreatedAt, gconv.String(gtime.Now().Format("Y-m-d 00:00:00")), gconv.String(gtime.Now().Format("Y-m-d 23:59:59"))).Count()
	if settingPostsDayMax.Int() > 0 && hasPostsAmount > settingPostsDayMax.Int() {
		return res, response.NewError("今日创建主题数量已达上限")
	}
	if err != nil {
		return res, response.WrapError(err, "创建主题失败")
	}
	settingNeedToken, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_TOKEN_ESTABLISH_POSTS_DEDUCT)

	if err != nil {
		return res, response.WrapError(err, "创建主题失败")
	}

	entity := &entity.Posts{}

	err = gconv.Scan(req, entity)

	entity.UserId = user.Id
	entity.Username = user.Username
	entity.CharacterAmount = uint(contentLength)
	if err != nil {
		return res, response.WrapError(err, "创建主题失败")
	}

	err = dao.Posts.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		id, err := dao.Posts.Ctx(ctx).InsertAndGetId(entity)
		if err != nil {
			return err
		}
		return User.changeBalance(ctx, user.Id, -settingNeedToken.Int(), model.BALANCE_CHANGE_TYPE_ESTABLISH_POST, uint(id), "")

	})
	if err != nil {
		return res, err
	}
	return res, response.NewSuccess("创建成功")
}

/**
更新主题
更新主题的前提条件
	- 主题创建之后的 CONFIG_POSTS_CAN_UPDATE_TIME 分钟内，如果主题获得的回复数量小于 CONFIG_POSTS_CAN_UPDATE_REPLY_AMOUNT
	- 每次编辑消耗后台设定CONFIG_TOKEN_UPDATE_POSTS_DEDUCT
*/
func (p *posts) Update(ctx context.Context, req *define.PostsUpdateReq) (res *define.PostsUpdateRes, err error) {
	res = &define.PostsUpdateRes{}

	posts := entity.Posts{}

	err = dao.Posts.Ctx(ctx).WherePri(req.Id).Scan(&posts)
	if err != nil {
		return res, response.WrapError(err, "更新主题失败")
	}
	if posts.Id == 0 {
		return res, response.NewError("主题不存在")
	}
	settingCanUpdateTime, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_POSTS_CAN_UPDATE_TIME)
	if err != nil {
		return res, response.WrapError(err, "更新主题失败")
	}

	if gtime.Now().Sub(posts.CreatedAt) > time.Duration(settingCanUpdateTime.Int())*time.Minute {
		return res, response.NewError(fmt.Sprintf("主题创建之后的 %d 分钟内才能编辑", settingCanUpdateTime.Int()))
	}
	settingCanUpdateReplyAmount, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_POSTS_CAN_UPDATE_REPLY_AMOUNT)
	if err != nil {
		return res, response.WrapError(err, "更新主题失败")
	}
	replyAmount, err := dao.Replies.Ctx(ctx).Where(dao.Replies.Columns().PostsId, posts.Id).Count()
	if replyAmount > settingCanUpdateReplyAmount.Int() {
		return res, response.NewError(fmt.Sprintf("主题获得的回复数量大于 %d 之后不能编辑", settingCanUpdateReplyAmount.Int()))
	}
	return
}

func (p *posts) Detail(ctx context.Context, Id uint) (res *model.PostWithComments, err error) {
	res = &model.PostWithComments{}
	err = dao.Posts.Ctx(ctx).WherePri(Id).WithAll().Scan(res)
	return
}

// 访问主题后，进行的相关操作
func (p *posts) Visit(ctx context.Context, postId uint) (err error) {
	_, err = dao.Posts.Ctx(ctx).WherePri(postId).Increment(dao.Posts.Columns().VisitsAmount, 1)
	return
}

// 用户收藏/取消主题
func (p *posts) ToggleCollect(ctx context.Context, postId uint, userId uint, username string) (err error) {

	d := dao.Posts.Ctx(ctx)
	var post entity.Posts
	err = d.WherePri(postId).Scan(&post)
	if err != nil {
		return
	}
	if post.Id == 0 {
		return response.NewError("主题不存在")
	}

	relationId, err := User.WhetherCollectPost(ctx, userId, postId)

	if err != nil {
		return
	}

	if relationId == 0 {
		err = d.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
			_, err = d.WherePri(postId).Increment(dao.Posts.Columns().CollectionAmount, 1)
			if err != nil {
				return
			}
			_, err = dao.ThanksOrShieldOrCollectContentRelation.Ctx(ctx).Insert(&entity.ThanksOrShieldOrCollectContentRelation{
				UserId:         userId,
				Username:       username,
				TargetId:       postId,
				TargetUserId:   post.UserId,
				TargetUsername: post.Username,
				Type:           model.ContentRelationTypeCollectPosts,
			})
			return
		})
	} else {
		err = d.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
			_, err = d.WherePri(postId).Decrement(dao.Posts.Columns().CollectionAmount, 1)
			if err != nil {
				return
			}
			_, err = dao.ThanksOrShieldOrCollectContentRelation.Ctx(ctx).WherePri(relationId).Delete()
			return
		})
	}

	return
}

// 用户忽略/取消主题
func (p *posts) ToggleShield(ctx context.Context, postId uint, userId uint, username string) (err error) {

	d := dao.Posts.Ctx(ctx)
	var post entity.Posts
	err = d.WherePri(postId).Scan(&post)
	if err != nil {
		return
	}
	if post.Id == 0 {
		return response.NewError("主题不存在")
	}

	relationId, err := User.WhetherShieldPost(ctx, userId, postId)

	if err != nil {
		return
	}

	if relationId == 0 {
		err = d.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
			_, err = d.WherePri(postId).Increment(dao.Posts.Columns().ShieldedAmount, 1)
			if err != nil {
				return
			}
			_, err = dao.ThanksOrShieldOrCollectContentRelation.Ctx(ctx).Insert(&entity.ThanksOrShieldOrCollectContentRelation{
				UserId:         userId,
				Username:       username,
				TargetId:       postId,
				TargetUserId:   post.UserId,
				TargetUsername: post.Username,
				Type:           model.ContentRelationTypeShieldPosts,
			})
			return
		})
	} else {
		err = d.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
			_, err = d.WherePri(postId).Decrement(dao.Posts.Columns().ShieldedAmount, 1)
			if err != nil {
				return
			}
			_, err = dao.ThanksOrShieldOrCollectContentRelation.Ctx(ctx).WherePri(relationId).Delete()
			return
		})
	}

	return
}

// 用户感谢主题
func (p *posts) Thanks(ctx context.Context, postId uint, userId uint, username string) (err error) {

	d := dao.Posts.Ctx(ctx)
	var post entity.Posts
	err = d.WherePri(postId).Scan(&post)
	if err != nil {
		return
	}
	if post.Id == 0 {
		return response.NewError("主题不存在")
	}

	relationId, err := User.WhetherThanksPost(ctx, userId, postId)

	if err != nil {
		return
	}

	if relationId > 0 {
		return response.NewError("您已经感谢过该主题")
	}

	err = d.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
		_, err = d.WherePri(postId).Increment(dao.Posts.Columns().ThanksAmount, 1)
		if err != nil {
			return
		}
		_, err = dao.ThanksOrShieldOrCollectContentRelation.Ctx(ctx).Insert(&entity.ThanksOrShieldOrCollectContentRelation{
			UserId:         userId,
			Username:       username,
			TargetId:       postId,
			TargetUserId:   post.UserId,
			TargetUsername: post.Username,
			Type:           model.ContentRelationTypeThanksPosts,
		})
		return
	})

	return
}
