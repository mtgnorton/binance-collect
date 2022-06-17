package cmd

import (
	"context"
	"gf-admin/app/system/admin/internal/deposit_withdraw"
)

func DwRun(ctx context.Context) {
	manager := deposit_withdraw.NewManager(ctx)
	manager.Run(ctx)
}
