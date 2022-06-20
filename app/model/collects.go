package model

const (
	COLLECT_STATUS_WAIT_FEE            = "wait_fee"
	COLLCT_STATUS_PROCESS_FEE          = "process_fee"
	COLLECT_STATUS_WAIT_COLLECT        = "wait_collect"
	COLLECT_STATUS_PROCESS_COLLECT     = "process_collect"
	COLLECT_STATUS_WAIT_NOTIFY         = "wait_notify"
	COLLECT_STATUS_PROCESS_NOTIFY      = "process_notify"
	COLLECT_STATUS_FINISH_NOTIFY       = "finish"
	COLLECT_STATUS_FAIL                = "fail"
	COLLECT_STATUS_FAIl_TOO_LOW_AMOUNT = "fail_to_low_amount" //充值金额过小
	COLLECT_STATUS_MISS                = "miss"
)
