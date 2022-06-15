package deposit_withdraw

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
)

var ctx context.Context = context.Background()

func init() {
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config-unit.toml")

}
