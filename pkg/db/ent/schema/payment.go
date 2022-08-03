package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"

	constant "github.com/NpoolPlatform/cloud-hashing-order/pkg/const"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("good_id", uuid.UUID{}),
		field.UUID("order_id", uuid.UUID{}).
			Unique(),
		field.UUID("account_id", uuid.UUID{}),
		field.Uint64("start_amount"),
		field.Uint64("amount"),
		field.
			Float("pay_with_balance_amount").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional().
			Nillable(),
		field.Uint64("finish_amount"),
		field.Uint64("coin_usd_currency"),
		field.Uint64("local_coin_usd_currency"),
		field.Uint64("live_coin_usd_currency"),
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
		field.Bool("user_set_paid").Default(false),
		field.Bool("user_set_canceled").Default(false),
		field.String("user_payment_txid").Default(""),
		field.Bool("fake_payment").Default(false),
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
