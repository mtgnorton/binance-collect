package define

import (
	"gf-admin/app/model"
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type UserListInput struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	IsDestroy bool   `json:"is_destroy"`
	Status    int    `json:"status"`
	BeginTime string `json:"begin_time"`
	EndTime   string `json:"end_time"`
	model.OrderFieldDirectionInput
	model.PageSizeInput
}

type UserListOutput struct {
	List  []entity.Users `json:"users"`
	Page  int            `json:"page"`
	Size  int            `json:"size"`
	Total int            `json:"total"`
}

type UserListReq struct {
	g.Meta `path:"/user-list" method:"get" summary:"用户列表" tags:"用户管理"`
	*UserListInput
}

type UserListRes struct {
	*UserListOutput
}

type UserToggleDestroyInput struct {
	Id        uint `json:"id"`
	IsDestroy bool `json:"is_destroy"`
}

type UserToggleDestroyReq struct {
	g.Meta `path:"/user-toggle-destroy" method:"delete" summary:"删除用户" tags:"用户管理"`
	*UserToggleDestroyInput
}

type UserToggleDestroyRes struct {
}
