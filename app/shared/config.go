package shared

import (
	"context"
	"gf-admin/app/dao"
	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

var Config = config{}

type config struct {
}

//根据传递的module和key获取对应的配置值
//module 可以为空

func (c *config) Get(ctx context.Context, module, key string) (value *gvar.Var, err error) {

	condition := g.Map{
		dao.Config.Columns.Key: key,
	}
	if module != "" {
		condition[dao.Config.Columns.Module] = module
	}
	record, err := dao.Config.Ctx(ctx).Where(condition).One()
	if err != nil {
		return gvar.New(record[dao.Config.Columns.Value]), err
	}
	return gvar.New(record[dao.Config.Columns.Value]), nil

}

//根据传递的module和keys批量获取配置值
//module 可以为空,keys可以为空
func (c *config) Gets(ctx context.Context, module string, keys ...string) (values map[string]*gvar.Var, err error) {
	values = make(map[string]*gvar.Var)

	condition := g.Map{}
	if len(keys) > 0 {
		condition[dao.Config.Columns.Key] = keys

	}

	if module != "" {
		condition[dao.Config.Columns.Module] = module
	}
	records, err := dao.Config.Ctx(ctx).Where(condition).All()

	if err != nil {
		return values, err
	}
	for _, record := range records {
		values[record[dao.Config.Columns.Key].String()] = gvar.New(record[dao.Config.Columns.Value])
	}

	return values, err

}

// 根据传递的module和key设置对应的配置值,存在则更新，不存在则插入
func (c *config) Set(ctx context.Context, module, key string, value interface{}) (err error) {

	data := g.Map{
		dao.Config.Columns.Module: module,
		dao.Config.Columns.Key:    key,
		dao.Config.Columns.Value:  value,
	}
	_, err = dao.Config.Ctx(ctx).Save(data)
	return
}

// 根据传递的module和mapping批量设置对应的配置值,存在则更新，不存在则插入
func (c *config) Sets(ctx context.Context, module string, mapping map[string]interface{}) (err error) {
	if len(mapping) == 0 {
		return
	}

	return dao.Config.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		for key, value := range mapping {
			data := g.Map{
				dao.Config.Columns.Module: module,
				dao.Config.Columns.Key:    key,
				dao.Config.Columns.Value:  value,
			}
			_, err = tx.Ctx(ctx).Save(dao.Config.Table, data)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
