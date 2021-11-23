// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/payment"
	"github.com/google/uuid"
)

// Payment is the model entity for the Payment schema.
type Payment struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID uuid.UUID `json:"order_id,omitempty"`
	// AccountID holds the value of the "account_id" field.
	AccountID uuid.UUID `json:"account_id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount uint64 `json:"amount,omitempty"`
	// CoinInfoID holds the value of the "coin_info_id" field.
	CoinInfoID uuid.UUID `json:"coin_info_id,omitempty"`
	// State holds the value of the "state" field.
	State payment.State `json:"state,omitempty"`
	// ChainTransactionID holds the value of the "chain_transaction_id" field.
	ChainTransactionID string `json:"chain_transaction_id,omitempty"`
	// PlatformTransactionID holds the value of the "platform_transaction_id" field.
	PlatformTransactionID uuid.UUID `json:"platform_transaction_id,omitempty"`
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
		case payment.FieldAmount, payment.FieldCreateAt, payment.FieldUpdateAt, payment.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case payment.FieldState, payment.FieldChainTransactionID:
			values[i] = new(sql.NullString)
		case payment.FieldID, payment.FieldOrderID, payment.FieldAccountID, payment.FieldCoinInfoID, payment.FieldPlatformTransactionID:
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
		case payment.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				pa.Amount = uint64(value.Int64)
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
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Payment is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Payment) String() string {
	var builder strings.Builder
	builder.WriteString("Payment(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", order_id=")
	builder.WriteString(fmt.Sprintf("%v", pa.OrderID))
	builder.WriteString(", account_id=")
	builder.WriteString(fmt.Sprintf("%v", pa.AccountID))
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", pa.Amount))
	builder.WriteString(", coin_info_id=")
	builder.WriteString(fmt.Sprintf("%v", pa.CoinInfoID))
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", pa.State))
	builder.WriteString(", chain_transaction_id=")
	builder.WriteString(pa.ChainTransactionID)
	builder.WriteString(", platform_transaction_id=")
	builder.WriteString(fmt.Sprintf("%v", pa.PlatformTransactionID))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", pa.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", pa.UpdateAt))
	builder.WriteString(", delete_at=")
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