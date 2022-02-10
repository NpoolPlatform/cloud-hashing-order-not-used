package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	constant "github.com/NpoolPlatform/cloud-hashing-order/pkg/const"
	"github.com/google/uuid"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("order_id", uuid.UUID{}).
			Unique(),
		field.UUID("account_id", uuid.UUID{}),
		field.Uint64("start_amount"),
		field.Uint64("amount"),
		field.Uint64("coin_usd_currency"),
		field.UUID("coin_info_id", uuid.UUID{}),
		field.Enum("state").
			Values(
				constant.PaymentStateWait,
				constant.PaymentStateDone,
				constant.PaymentStateCanceled,
				constant.PaymentStateTimeout,
			),
		field.String("chain_transaction_id"),
		field.UUID("platform_transaction_id", uuid.UUID{}),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("delete_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return nil
}
