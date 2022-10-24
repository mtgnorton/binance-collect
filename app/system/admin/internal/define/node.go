package define

import (
	"gf-admin/app/model"
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type NodeListReq struct {
	g.Meta `path:"/node-list" method:"get" summary:"节点列表" tags:"节点管理"`
	NodeListInput
}

type NodeListRes struct {
	*NodeListOutput
}

type NodeListOutput struct {
	List []*entity.Nodes `json:"list"`
	model.PageSizeOutput
}

type NodeListInput struct {
	Name    string `json:"name"`
	IsIndex string `v:"in:1,0" json:"is_index"`
	model.PageSizeInput
	model.OrderFieldDirectionInput
}

type NodeInfoReq struct {
	g.Meta `path:"/node-info" method:"get" summary:"节点详情" tags:"节点管理"`
	Id     uint `json:"id" dc:"id" v:"min:1#请选择节点id"`
}

type NodeInfoRes struct {
	*NodeInfoOutput
}

type NodeInfoOutput struct {
	entity.Nodes
}

type NodeStoreInput struct {
	Name        string `json:"name" v:"required#节点名称不能为空"`
	Description string `json:"description"`
	IsIndex     int    `json:"is_index"`
}

type NodeStoreReq struct {
	g.Meta `path:"/node-store" method:"post" summary:"添加节点" tags:"节点管理"`
	*NodeStoreInput
}

type NodeStoreRes struct {
}

type NodeUpdateReq struct {
	g.Meta `path:"/node-update" method:"put" summary:"更新节点" tags:"节点管理"`
	*NodeUpdateInput
}

type NodeUpdateRes struct {
}
type NodeUpdateInput struct {
	Id          uint   `json:"id" v:"required#节点id不能为空"`
	Name        string `json:"name" v:"required#节点名称不能为空"`
	Description string `json:"description"`
	IsIndex     int    `json:"is_index"`
}

type NodeDestroyReq struct {
	g.Meta `path:"/node-destroy" method:"delete" summary:"删除节点" tags:"节点管理"`
	Id     uint `json:"id" v:"required#节点id不能为空"`
}

type NodeDestroyRes struct {
}
