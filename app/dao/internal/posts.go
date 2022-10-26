// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PostsDao is the data access object for table forum_posts.
type PostsDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns PostsColumns // columns contains all the column names of Table for convenient usage.
}

// PostsColumns defines and stores column names for table forum_posts.
type PostsColumns struct {
	Id                string //
	NodeId            string // 节点id
	UserId            string // 用户id
	Username          string // 用户名
	Title             string // 标题
	Content           string // 内容
	TopEndTime        string // 置顶截止时间,为空说明没有置顶
	CharacterAmount   string // 字符长度
	VisitsAmount      string // 访问次数
	CollectionAmount  string // 收藏次数
	ReplyAmount       string // 回复次数
	ThanksAmount      string // 感谢次数
	ShieldedAmount    string // 被屏蔽次数
	Status            string // 状态：no_audit, normal, shielded
	Weight            string // 权重
	ReplyLastUserId   string // 最后回复用户id
	ReplyLastUsername string // 最后回复用户名
	ReplyLastTime     string // 最后回复时间
	CreatedAt         string // 主题创建时间
	UpdatedAt         string // 主题更新时间
	DeletedAt         string // 删除时间
}

//  PostsColumns holds the columns for table forum_posts.
var postsColumns = PostsColumns{
	Id:                "id",
	NodeId:            "node_id",
	UserId:            "user_id",
	Username:          "username",
	Title:             "title",
	Content:           "content",
	TopEndTime:        "top_end_time",
	CharacterAmount:   "character_amount",
	VisitsAmount:      "visits_amount",
	CollectionAmount:  "collection_amount",
	ReplyAmount:       "reply_amount",
	ThanksAmount:      "thanks_amount",
	ShieldedAmount:    "shielded_amount",
	Status:            "status",
	Weight:            "weight",
	ReplyLastUserId:   "reply_last_user_id",
	ReplyLastUsername: "reply_last_username",
	ReplyLastTime:     "reply_last_time",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
	DeletedAt:         "deleted_at",
}

// NewPostsDao creates and returns a new DAO object for table data access.
func NewPostsDao() *PostsDao {
	return &PostsDao{
		group:   "default",
		table:   "forum_posts",
		columns: postsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PostsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PostsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PostsDao) Columns() PostsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PostsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PostsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PostsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}