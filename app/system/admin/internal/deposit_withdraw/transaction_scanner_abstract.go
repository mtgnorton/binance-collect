package deposit_withdraw

import (
	"context"
)

type TransactionScannerAbstract interface {
	Scan(context.Context, chan *TransferTask, chan *NotifyTask)
}
