package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// GoodPaying holds the schema definition for the GoodPaying entity.
type GoodPaying struct {
	ent.Schema
}

// Fields of the GoodPaying.
func (GoodPaying) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("order_id", uuid.UUID{}),
		field.UUID("account_id", uuid.UUID{}),
		field.Enum("state").
			Values("wait", "done", "canceled", "timeout"),
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

// Edges of the GoodPaying.
func (GoodPaying) Edges() []ent.Edge {
	return nil
}
