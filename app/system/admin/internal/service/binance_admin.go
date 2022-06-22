package service

import (
	"context"
	"fmt"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/deposit_withdraw"
	"gf-admin/utility/custom_error"

	"github.com/gogf/gf/v2/container/garray"

	"github.com/gogf/gf/v2/os/gcache"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/database/gdb"
)

var BinanceAdmin = binanceAdmin{}

type binanceAdmin struct {
}

func (ba *binanceAdmin) UserAddressList(ctx context.Context, in *define.UserAddressListInput) (out *define.UserAddressListOutput, err error) {
	out = &define.UserAddressListOutput{}

	d := dao.UserAddresses.Ctx(ctx)

	if in.ExternalUserId != 0 {
		d = d.Where(dao.UserAddresses.Columns().ExternalUserId, in.ExternalUserId)
	}

	if in.Address != "" {
		d = d.WhereLike(dao.UserAddresses.Columns().Address, fmt.Sprintf("%%%s%%", in.Address))
	}

	out.Page = in.Page
	out.Size = in.Size
	out.Total, err = d.Count()
	if err != nil {
		return out, custom_error.New(err.Error())
	}

	d = d.Order(dao.UserAddresses.Columns().Id, "desc").Page(in.Page, in.Size)
	err = d.Scan(&out.List)

	return

}

func (ba *binanceAdmin) CollectList(ctx context.Context, in *define.CollectListInput) (out *define.CollectListOutput, err error) {
	out = &define.CollectListOutput{}

	d := dao.Collects.Ctx(ctx)

	if in.RechargeHash != "" {
		d = d.Where(dao.Collects.Columns().RechargeHash, in.RechargeHash)
	}

	if in.HandfeeHash != "" {
		d = d.Where(dao.Collects.Columns().HandfeeHash, in.HandfeeHash)
	}
	if in.CollectHash != "" {
		d = d.Where(dao.Collects.Columns().CollectHash, in.CollectHash)
	}

	if in.UserAddress != "" {
		d = d.Where(dao.Collects.Columns().UserAddress, in.UserAddress)
	}
	if in.Status != "" {
		d = d.Where(dao.Collects.Columns().Status, in.Status)
	}
	out.Page = in.Page
	out.Size = in.Size
	out.Total, err = d.Count()
	if err != nil {
		return out, custom_error.New(err.Error())
	}
	d = d.Order(dao.Collects.Columns().Id, "desc").Page(in.Page, in.Size)
	err = d.Scan(&out.List)
	for _, item := range out.List {
		item.Value, err = deposit_withdraw.ChainClient.WeiToEther(ctx, item.Value, item.Symbol)
		if err != nil {
			return out, custom_error.New(err.Error(), g.Map{
				"item": item,
			})
		}
		item.RechargeValue, err = deposit_withdraw.ChainClient.WeiToEther(ctx, item.RechargeValue, item.Symbol)
		if err != nil {
			return out, custom_error.New(err.Error(), g.Map{
				"item": item,
			})
		}
	}
	return
}
func (ba *binanceAdmin) CollectUpdate(ctx context.Context, in *define.CollectUpdateInput) (err error) {
	d := dao.Collects.Ctx(ctx)
	if in.Id != 0 {
		d = d.Where(dao.Collects.Columns().Id, in.Id)
	}
	idVar, err := d.Value(dao.Collects.Columns().Id)

	if idVar.Int() == 0 {
		return custom_error.New("id is not exist")
	}
	_, err = d.Update(g.Map{
		dao.Collects.Columns().Status: in.Status,
	})

	return

}

func (ba *binanceAdmin) CollectDestroy(ctx context.Context, in *define.CollectDestroyInput) (err error) {
	d := dao.Collects.Ctx(ctx)
	if in.Id != 0 {
		d = d.Where(dao.Collects.Columns().Id, in.Id)
	}
	idVar, err := d.Value(dao.Collects.Columns().Id)

	if idVar.Int() == 0 {
		return custom_error.New("id is not exist")
	}

	_, err = d.Delete()

	return err
}

func (ba *binanceAdmin) WithdrawList(ctx context.Context, in *define.WithdrawListInput) (out *define.WithdrawListOutput, err error) {
	out = &define.WithdrawListOutput{}

	d := dao.Withdraws.Ctx(ctx)

	if in.Hash != "" {
		d = d.Where(dao.Withdraws.Columns().Hash, in.Hash)
	}

	if in.To != "" {
		d = d.Where(dao.Withdraws.Columns().To, in.To)
	}

	if in.UserAddress != "" {
		d = d.Where(dao.Withdraws.Columns().UserAddress, in.UserAddress)
	}
	if in.ExternalOrderId != "" {
		d = d.Where(dao.Withdraws.Columns().ExternalOrderId, in.ExternalOrderId)
	}

	if in.ExternalUserId != 0 {
		d = d.Where(dao.Withdraws.Columns().ExternalUserId, in.ExternalUserId)
	}
	if in.Status != "" {
		d = d.Where(dao.Withdraws.Columns().Status, in.Status)
	}
	out.Page = in.Page
	out.Size = in.Size
	out.Total, err = d.Count()
	if err != nil {
		return out, custom_error.New(err.Error())
	}
	d = d.Order(dao.Withdraws.Columns().Id, "desc").Page(in.Page, in.Size)
	err = d.Scan(&out.List)
	for _, item := range out.List {
		item.Value, err = deposit_withdraw.ChainClient.WeiToEther(ctx, item.Value, item.Symbol)
		if err != nil {
			return out, custom_error.New(err.Error())
		}
	}
	return
}

func (ba *binanceAdmin) WithdrawUpdate(ctx context.Context, in *define.WithdrawUpdateInput) (err error) {
	d := dao.Withdraws.Ctx(ctx)
	if in.Id != 0 {
		d = d.Where(dao.Withdraws.Columns().Id, in.Id)
	}
	idVar, err := d.Value(dao.Withdraws.Columns().Id)

	if idVar.Int() == 0 {
		return custom_error.New("id is not exist")
	}
	_, err = d.Update(g.Map{
		dao.Withdraws.Columns().Status: in.Status,
	})
	return
}
func (ba *binanceAdmin) WithdrawDestroy(ctx context.Context, in *define.WithdrawDestroyInput) (err error) {
	d := dao.Withdraws.Ctx(ctx)
	if in.Id != 0 {
		d = d.Where(dao.Withdraws.Columns().Id, in.Id)
	}
	idVar, err := d.Value(dao.Withdraws.Columns().Id)

	if idVar.Int() == 0 {
		return custom_error.New("id is not exist")
	}
	_, err = d.Delete()
	return
}

func (ba *binanceAdmin) QueueTaskList(ctx context.Context, in *define.QueueTaskListInput) (out *define.QueueTaskListOutput, err error) {
	out = &define.QueueTaskListOutput{}
	d := dao.QueueTask.Ctx(ctx)

	if in.Hash != "" {
		d = d.Where(dao.QueueTask.Columns().Hash, in.Hash)
	}

	if in.From != "" {
		d = d.Where(dao.QueueTask.Columns().From, in.From)
	}
	if in.To != "" {
		d = d.Where(dao.QueueTask.Columns().To, in.To)
	}
	if in.Status != "" {
		d = d.Where(dao.QueueTask.Columns().Status, in.Status)
	}

	out.Page = in.Page
	out.Size = in.Size
	out.Total, err = d.Count()
	if err != nil {
		return out, custom_error.New(err.Error())
	}
	d = d.Order(dao.Withdraws.Columns().Id, "desc").Page(in.Page, in.Size)
	err = d.ScanList(&out.List, "Task")

	if err != nil {
		return
	}
	err = dao.QueueTaskLog.Ctx(ctx).
		Where(dao.QueueTaskLog.Columns().QueueTaskId, gdb.ListItemValuesUnique(out.List, "Task", "Id")).
		ScanList(&out.List, "Logs", "Task", "queue_task_id:Id")

	for _, item := range out.List {
		item.Task.Value, err = deposit_withdraw.ChainClient.WeiToEther(ctx, item.Task.Value, item.Task.Symbol)
		if err != nil {
			return out, custom_error.New(err.Error())
		}
	}
	return
}

func (ba *binanceAdmin) QueueTaskUpdate(ctx context.Context, in *define.QueueTaskUpdateInput) (err error) {
	d := dao.QueueTask.Ctx(ctx)
	if in.Id != 0 {
		d = d.Where(dao.QueueTask.Columns().Id, in.Id)
	}
	idVar, err := d.Value(dao.QueueTask.Columns().Id)
	if idVar.Int() == 0 {
		return custom_error.New("id is not exist")
	}
	statusArray := garray.NewStrArrayFrom([]string{model.QUEUE_TASK_STATUS_WAIT, model.QUEUE_TASK_STATUS_PROCESS, model.QUEUE_TASK_STATUS_SUCCESS, model.QUEUE_TASK_STATUS_FAIL})
	if !statusArray.Contains(in.Status) {
		return custom_error.New("status is not exist")
	}

	_, err = d.Update(g.Map{
		dao.QueueTask.Columns().Status:     in.Status,
		dao.QueueTask.Columns().FailAmount: in.FailAmount,
	})

	return err

}

func (ba *binanceAdmin) QueueTaskDestroy(ctx context.Context, in *define.QueueTaskDestroyInput) (err error) {
	d := dao.QueueTask.Ctx(ctx)
	if in.Id != 0 {
		d = d.Where(dao.QueueTask.Columns().Id, in.Id)
	}
	idVar, err := d.Value(dao.QueueTask.Columns().Id)
	if idVar.Int() == 0 {
		return custom_error.New("id is not exist")
	}
	_, err = d.Delete()
	return err
}

func (ba *binanceAdmin) NotifyList(ctx context.Context, in *define.NotifyListInput) (out *define.NotifyListOutput, err error) {
	out = &define.NotifyListOutput{}
	d := dao.Notify.Ctx(ctx)

	if in.UniqueId != "" {
		d = d.Where(dao.Notify.Columns().UniqueId, in.UniqueId)
	}
	out.Page = in.Page
	out.Size = in.Size
	out.Total, err = d.Count()
	if err != nil {
		return out, custom_error.New(err.Error())
	}
	d = d.Order(dao.Notify.Columns().Id, "desc").Page(in.Page, in.Size)
	err = d.ScanList(&out.List, "Notify")

	if err != nil {
		return
	}
	err = dao.NotifyLog.Ctx(ctx).
		Where(dao.NotifyLog.Columns().NotifyId, gdb.ListItemValuesUnique(out.List, "Notify", "Id")).
		ScanList(&out.List, "Logs", "Notify", "notify_id:Id")

	return
}
func (ba *binanceAdmin) NotifyUpdate(ctx context.Context, in *define.NotifyUpdateInput) (err error) {
	d := dao.Notify.Ctx(ctx)
	if in.Id != 0 {
		d = d.Where(dao.Notify.Columns().Id, in.Id)
	}
	idVar, err := d.Value(dao.Notify.Columns().Id)
	if idVar.Int() == 0 {
		return custom_error.New("id is not exist")
	}
	_, err = d.Update(g.Map{
		dao.Notify.Columns().Status:             in.Status,
		dao.Notify.Columns().IsImmediatelyRetry: in.IsImmediatelyRetry,
	})
	return
}

func (ba *binanceAdmin) NotifyDestroy(ctx context.Context, in *define.NotifyDestroyInput) (err error) {
	d := dao.Notify.Ctx(ctx)
	if in.Id != 0 {
		d = d.Where(dao.Notify.Columns().Id, in.Id)
	}
	idVar, err := d.Value(dao.Notify.Columns().Id)
	if idVar.Int() == 0 {
		return custom_error.New("id is not exist")
	}
	_, err = d.Delete()
	return
}

func (ba *binanceAdmin) ContractList(ctx context.Context, in *define.ContractListInput) (out *define.ContractListOutput, err error) {
	out = &define.ContractListOutput{}
	d := dao.Contracts.Ctx(ctx)

	out.Page = in.Page
	out.Size = in.Size
	out.Total, err = d.Count()
	if err != nil {
		return out, custom_error.New(err.Error())
	}
	d = d.Order(dao.Contracts.Columns().Id, "desc").Page(in.Page, in.Size)

	err = d.Scan(&out.List)

	return
}

func (ba *binanceAdmin) ContractStore(ctx context.Context, in *define.ContractStoreInput) (err error) {
	d := dao.Contracts.Ctx(ctx)

	idVar, err := d.Where(dao.Contracts.Columns().Symbol, in.Symbol).WhereOr(dao.Contracts.Columns().Address, in.Address).Value(dao.Contracts.Columns().Id)

	if err != nil {
		return custom_error.New(err.Error())
	}
	if idVar.Int() > 0 {
		return custom_error.New("货币或地址重复")
	}
	_, err = gcache.Remove(ctx, model.CACHE_KEY_CONTRACTS)
	if err != nil {
		return err
	}
	_, err = d.Insert(g.Map{
		dao.Contracts.Columns().Symbol:        in.Symbol,
		dao.Contracts.Columns().Address:       in.Address,
		dao.Contracts.Columns().Decimals:      in.Decimals,
		dao.Contracts.Columns().IsCollectOpen: in.IsCollectOpen,
	})
	return

}

func (ba *binanceAdmin) ContractUpdate(ctx context.Context, in *define.ContractUpdateInput) (err error) {
	d := dao.Contracts.Ctx(ctx)
	idVar, err := d.
		Where("id != ? and (symbol = ? or  address = ?)", in.Id, in.Symbol, in.Address).
		Value(dao.Contracts.Columns().Id)
	if err != nil {
		return custom_error.New(err.Error())
	}
	if idVar.Int() > 0 {
		return custom_error.New("货币或地址重复")
	}
	_, err = gcache.Remove(ctx, model.CACHE_KEY_CONTRACTS)
	if err != nil {
		return err
	}
	_, err = d.Where(dao.Contracts.Columns().Id, in.Id).Update(g.Map{
		dao.Contracts.Columns().Symbol:        in.Symbol,
		dao.Contracts.Columns().Address:       in.Address,
		dao.Contracts.Columns().Decimals:      in.Decimals,
		dao.Contracts.Columns().IsCollectOpen: in.IsCollectOpen,
	})
	return
}
func (ba *binanceAdmin) ContractInfo(ctx context.Context, in *define.ContractInfoInput) (out *define.ContractInfoOutput, err error) {
	out = &define.ContractInfoOutput{}
	d := dao.Contracts.Ctx(ctx)
	err = d.Where(dao.Contracts.Columns().Id, in.Id).Scan(&out.Contract)
	return
}

func (ba *binanceAdmin) ContractDestroy(ctx context.Context, in *define.ContractDestroyInput) (err error) {
	d := dao.Contracts.Ctx(ctx)
	_, err = gcache.Remove(ctx, model.CACHE_KEY_CONTRACTS)
	if err != nil {
		return err
	}
	_, err = d.Where(dao.Contracts.Columns().Id, in.Id).Delete()
	return
}

func (ba *binanceAdmin) LoseBlockList(ctx context.Context, in *define.LoseBlockListInput) (out *define.LoseBlockListOutput, err error) {
	out = &define.LoseBlockListOutput{}
	d := dao.LoseBlocks.Ctx(ctx)
	out.Page = in.Page
	out.Size = in.Size
	out.Total, err = d.Count()
	if err != nil {
		return out, custom_error.New(err.Error())
	}
	d = d.Order(dao.Collects.Columns().Id, "desc").Page(in.Page, in.Size)
	err = d.Scan(&out.List)
	return
}
func (ba *binanceAdmin) LoseBlockStore(ctx context.Context, in *define.LoseBlockStoreInput) (err error) {
	d := dao.LoseBlocks.Ctx(ctx)
	idVar, err := d.Where(dao.LoseBlocks.Columns().Number, in.Number).Value(dao.LoseBlocks.Columns().Id)
	if err != nil {
		return custom_error.New(err.Error())
	}
	if idVar.Int() > 0 {
		return custom_error.New("区块号重复")
	}
	_, err = d.Insert(g.Map{
		dao.LoseBlocks.Columns().Number: in.Number,
	})
	return
}

func (ba *binanceAdmin) LoseBlockDestroy(ctx context.Context, in *define.LoseBlockDestroyInput) (err error) {
	d := dao.LoseBlocks.Ctx(ctx)
	_, err = d.Delete(dao.LoseBlocks.Columns().Id, in.Id)
	return
}
