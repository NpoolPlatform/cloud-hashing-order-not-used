package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("good_id", uuid.UUID{}),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.Uint32("units"),
		field.Uint32("discount").
			Default(0),
		field.Uint64("special_reduction_amount").
			Default(0),
		field.Enum("state").
			Values("created", "paying", "paid", "timeout", "canceled"),
		field.UUID("good_pay_id", uuid.UUID{}),
		field.Uint32("start"),
		field.Uint32("end"),
		field.Uint32("compensate_minutes").
			Default(0),
		field.Uint32("compensate_elapsed_minutes").
			Default(0),
		field.Uint32("gas_start"),
		field.Uint32("gas_end"),
		field.JSON("gas_pay_ids", []uuid.UUID{}),
		field.UUID("coupon_id", uuid.UUID{}),
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

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return nil
}
