// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/payment"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Payment is the model entity for the Payment schema.
type Payment struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID uuid.UUID `json:"order_id,omitempty"`
	// AccountID holds the value of the "account_id" field.
	AccountID uuid.UUID `json:"account_id,omitempty"`
	// StartAmount holds the value of the "start_amount" field.
	StartAmount uint64 `json:"start_amount,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount uint64 `json:"amount,omitempty"`
	// PayWithBalanceAmount holds the value of the "pay_with_balance_amount" field.
	PayWithBalanceAmount decimal.Decimal `json:"pay_with_balance_amount,omitempty"`
	// FinishAmount holds the value of the "finish_amount" field.
	FinishAmount uint64 `json:"finish_amount,omitempty"`
	// CoinUsdCurrency holds the value of the "coin_usd_currency" field.
	CoinUsdCurrency uint64 `json:"coin_usd_currency,omitempty"`
	// LocalCoinUsdCurrency holds the value of the "local_coin_usd_currency" field.
	LocalCoinUsdCurrency uint64 `json:"local_coin_usd_currency,omitempty"`
	// LiveCoinUsdCurrency holds the value of the "live_coin_usd_currency" field.
	LiveCoinUsdCurrency uint64 `json:"live_coin_usd_currency,omitempty"`
	// CoinInfoID holds the value of the "coin_info_id" field.
	CoinInfoID uuid.UUID `json:"coin_info_id,omitempty"`
	// State holds the value of the "state" field.
	State payment.State `json:"state,omitempty"`
	// ChainTransactionID holds the value of the "chain_transaction_id" field.
	ChainTransactionID string `json:"chain_transaction_id,omitempty"`
	// PlatformTransactionID holds the value of the "platform_transaction_id" field.
	PlatformTransactionID uuid.UUID `json:"platform_transaction_id,omitempty"`
	// UserSetPaid holds the value of the "user_set_paid" field.
	UserSetPaid bool `json:"user_set_paid,omitempty"`
	// UserSetCanceled holds the value of the "user_set_canceled" field.
	UserSetCanceled bool `json:"user_set_canceled,omitempty"`
	// UserPaymentTxid holds the value of the "user_payment_txid" field.
	UserPaymentTxid string `json:"user_payment_txid,omitempty"`
	// FakePayment holds the value of the "fake_payment" field.
	FakePayment bool `json:"fake_payment,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Payment) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case payment.FieldPayWithBalanceAmount:
			values[i] = new(decimal.Decimal)
		case payment.FieldUserSetPaid, payment.FieldUserSetCanceled, payment.FieldFakePayment:
			values[i] = new(sql.NullBool)
		case payment.FieldStartAmount, payment.FieldAmount, payment.FieldFinishAmount, payment.FieldCoinUsdCurrency, payment.FieldLocalCoinUsdCurrency, payment.FieldLiveCoinUsdCurrency, payment.FieldCreateAt, payment.FieldUpdateAt, payment.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case payment.FieldState, payment.FieldChainTransactionID, payment.FieldUserPaymentTxid:
			values[i] = new(sql.NullString)
		case payment.FieldID, payment.FieldAppID, payment.FieldUserID, payment.FieldGoodID, payment.FieldOrderID, payment.FieldAccountID, payment.FieldCoinInfoID, payment.FieldPlatformTransactionID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Payment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Payment fields.
func (pa *Payment) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case payment.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pa.ID = *value
			}
		case payment.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				pa.AppID = *value
			}
		case payment.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				pa.UserID = *value
			}
		case payment.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				pa.GoodID = *value
			}
		case payment.FieldOrderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value != nil {
				pa.OrderID = *value
			}
		case payment.FieldAccountID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value != nil {
				pa.AccountID = *value
			}
		case payment.FieldStartAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start_amount", values[i])
			} else if value.Valid {
				pa.StartAmount = uint64(value.Int64)
			}
		case payment.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				pa.Amount = uint64(value.Int64)
			}
		case payment.FieldPayWithBalanceAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field pay_with_balance_amount", values[i])
			} else if value != nil {
				pa.PayWithBalanceAmount = *value
			}
		case payment.FieldFinishAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field finish_amount", values[i])
			} else if value.Valid {
				pa.FinishAmount = uint64(value.Int64)
			}
		case payment.FieldCoinUsdCurrency:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field coin_usd_currency", values[i])
			} else if value.Valid {
				pa.CoinUsdCurrency = uint64(value.Int64)
			}
		case payment.FieldLocalCoinUsdCurrency:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field local_coin_usd_currency", values[i])
			} else if value.Valid {
				pa.LocalCoinUsdCurrency = uint64(value.Int64)
			}
		case payment.FieldLiveCoinUsdCurrency:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field live_coin_usd_currency", values[i])
			} else if value.Valid {
				pa.LiveCoinUsdCurrency = uint64(value.Int64)
			}
		case payment.FieldCoinInfoID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_info_id", values[i])
			} else if value != nil {
				pa.CoinInfoID = *value
			}
		case payment.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				pa.State = payment.State(value.String)
			}
		case payment.FieldChainTransactionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chain_transaction_id", values[i])
			} else if value.Valid {
				pa.ChainTransactionID = value.String
			}
		case payment.FieldPlatformTransactionID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field platform_transaction_id", values[i])
			} else if value != nil {
				pa.PlatformTransactionID = *value
			}
		case payment.FieldUserSetPaid:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field user_set_paid", values[i])
			} else if value.Valid {
				pa.UserSetPaid = value.Bool
			}
		case payment.FieldUserSetCanceled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field user_set_canceled", values[i])
			} else if value.Valid {
				pa.UserSetCanceled = value.Bool
			}
		case payment.FieldUserPaymentTxid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_payment_txid", values[i])
			} else if value.Valid {
				pa.UserPaymentTxid = value.String
			}
		case payment.FieldFakePayment:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field fake_payment", values[i])
			} else if value.Valid {
				pa.FakePayment = value.Bool
			}
		case payment.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				pa.CreateAt = uint32(value.Int64)
			}
		case payment.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				pa.UpdateAt = uint32(value.Int64)
			}
		case payment.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				pa.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Payment.
// Note that you need to call Payment.Unwrap() before calling this method if this Payment
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Payment) Update() *PaymentUpdateOne {
	return (&PaymentClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the Payment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Payment) Unwrap() *Payment {
	_tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Payment is not a transactional entity")
	}
	pa.config.driver = _tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Payment) String() string {
	var builder strings.Builder
	builder.WriteString("Payment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pa.ID))
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", pa.AppID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", pa.UserID))
	builder.WriteString(", ")
	builder.WriteString("good_id=")
	builder.WriteString(fmt.Sprintf("%v", pa.GoodID))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", pa.OrderID))
	builder.WriteString(", ")
	builder.WriteString("account_id=")
	builder.WriteString(fmt.Sprintf("%v", pa.AccountID))
	builder.WriteString(", ")
	builder.WriteString("start_amount=")
	builder.WriteString(fmt.Sprintf("%v", pa.StartAmount))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", pa.Amount))
	builder.WriteString(", ")
	builder.WriteString("pay_with_balance_amount=")
	builder.WriteString(fmt.Sprintf("%v", pa.PayWithBalanceAmount))
	builder.WriteString(", ")
	builder.WriteString("finish_amount=")
	builder.WriteString(fmt.Sprintf("%v", pa.FinishAmount))
	builder.WriteString(", ")
	builder.WriteString("coin_usd_currency=")
	builder.WriteString(fmt.Sprintf("%v", pa.CoinUsdCurrency))
	builder.WriteString(", ")
	builder.WriteString("local_coin_usd_currency=")
	builder.WriteString(fmt.Sprintf("%v", pa.LocalCoinUsdCurrency))
	builder.WriteString(", ")
	builder.WriteString("live_coin_usd_currency=")
	builder.WriteString(fmt.Sprintf("%v", pa.LiveCoinUsdCurrency))
	builder.WriteString(", ")
	builder.WriteString("coin_info_id=")
	builder.WriteString(fmt.Sprintf("%v", pa.CoinInfoID))
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", pa.State))
	builder.WriteString(", ")
	builder.WriteString("chain_transaction_id=")
	builder.WriteString(pa.ChainTransactionID)
	builder.WriteString(", ")
	builder.WriteString("platform_transaction_id=")
	builder.WriteString(fmt.Sprintf("%v", pa.PlatformTransactionID))
	builder.WriteString(", ")
	builder.WriteString("user_set_paid=")
	builder.WriteString(fmt.Sprintf("%v", pa.UserSetPaid))
	builder.WriteString(", ")
	builder.WriteString("user_set_canceled=")
	builder.WriteString(fmt.Sprintf("%v", pa.UserSetCanceled))
	builder.WriteString(", ")
	builder.WriteString("user_payment_txid=")
	builder.WriteString(pa.UserPaymentTxid)
	builder.WriteString(", ")
	builder.WriteString("fake_payment=")
	builder.WriteString(fmt.Sprintf("%v", pa.FakePayment))
	builder.WriteString(", ")
	builder.WriteString("create_at=")
	builder.WriteString(fmt.Sprintf("%v", pa.CreateAt))
	builder.WriteString(", ")
	builder.WriteString("update_at=")
	builder.WriteString(fmt.Sprintf("%v", pa.UpdateAt))
	builder.WriteString(", ")
	builder.WriteString("delete_at=")
	builder.WriteString(fmt.Sprintf("%v", pa.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// Payments is a parsable slice of Payment.
type Payments []*Payment

func (pa Payments) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
