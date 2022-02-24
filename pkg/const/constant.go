package constant

import (
	inspireconst "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/const"
)

const (
	PaymentStateWait     = "wait"
	PaymentStateDone     = "done"
	PaymentStateCanceled = "canceled"
	PaymentStateTimeout  = "timeout"

	TimeoutSeconds = 6 * 60 * 60

	OrderPaymentLockKeyPrefix = "order-payment-lock"

	FixAmountCoupon            = inspireconst.CouponTypeCoupon
	DiscountCoupon             = inspireconst.CouponTypeDiscount
	UserSpecialReductionCoupon = "user-special-reduction"
)
