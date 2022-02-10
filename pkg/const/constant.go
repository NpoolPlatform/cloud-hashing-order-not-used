package constant

import (
	"time"
)

const (
	PaymentStateWait     = "wait"
	PaymentStateDone     = "done"
	PaymentStateCanceled = "canceled"
	PaymentStateTimeout  = "timeout"

	Timeout = 6 * time.Hour
)
