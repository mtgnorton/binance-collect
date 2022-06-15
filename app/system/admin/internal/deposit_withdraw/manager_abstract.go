package deposit_withdraw

import "context"

type ManagerAbstract interface {
	Run(context.Context)                          // 运行
	Detect(context.Context) (*OriginBlock, error) //检测交易
}
