package constant

const (
	PaymentStateWait     = "wait"
	PaymentStateDone     = "done"
	PaymentStateCanceled = "canceled"
	PaymentStateTimeout  = "timeout"

	TimeoutSeconds = 6 * 60 * 60

	OrderPaymentLockKeyPrefix = "order-payment-lock"
)
