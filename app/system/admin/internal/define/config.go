package define

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
)

type ConfigListInput struct {
	Modules []string `json:"module"`
}
type ConfigListOutput struct {
	Data map[string]map[string]*gvar.Var `json:"data"`
}
type ConfigListReq struct {
	g.Meta `path:"/config-list" method:"get" summary:"配置管理" tags:"配置管理"`
	*ConfigListInput
}

type ConfigListRes struct {
	*ConfigListOutput
}

type ConfigUpdateInput struct {
	Module      string                 `json:"module"`
	KeyValueMap map[string]interface{} `json:"key_value_map"`
}

type ConfigUpdateReq struct {
	g.Meta `path:"/config-update" method:"put" summary:"保存配置" tags:"配置管理"`
	*ConfigUpdateInput
}

type ConfigUpdateRes struct {
}
