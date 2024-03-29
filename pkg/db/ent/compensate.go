// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/compensate"
	"github.com/google/uuid"
)

// Compensate is the model entity for the Compensate schema.
type Compensate struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID uuid.UUID `json:"order_id,omitempty"`
	// Start holds the value of the "start" field.
	Start uint32 `json:"start,omitempty"`
	// End holds the value of the "end" field.
	End uint32 `json:"end,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Compensate) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case compensate.FieldStart, compensate.FieldEnd, compensate.FieldCreateAt, compensate.FieldUpdateAt, compensate.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case compensate.FieldMessage:
			values[i] = new(sql.NullString)
		case compensate.FieldID, compensate.FieldOrderID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Compensate", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Compensate fields.
func (c *Compensate) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case compensate.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case compensate.FieldOrderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value != nil {
				c.OrderID = *value
			}
		case compensate.FieldStart:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start", values[i])
			} else if value.Valid {
				c.Start = uint32(value.Int64)
			}
		case compensate.FieldEnd:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field end", values[i])
			} else if value.Valid {
				c.End = uint32(value.Int64)
			}
		case compensate.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				c.Message = value.String
			}
		case compensate.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				c.CreateAt = uint32(value.Int64)
			}
		case compensate.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				c.UpdateAt = uint32(value.Int64)
			}
		case compensate.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				c.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Compensate.
// Note that you need to call Compensate.Unwrap() before calling this method if this Compensate
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Compensate) Update() *CompensateUpdateOne {
	return (&CompensateClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Compensate entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Compensate) Unwrap() *Compensate {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Compensate is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Compensate) String() string {
	var builder strings.Builder
	builder.WriteString("Compensate(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", c.OrderID))
	builder.WriteString(", ")
	builder.WriteString("start=")
	builder.WriteString(fmt.Sprintf("%v", c.Start))
	builder.WriteString(", ")
	builder.WriteString("end=")
	builder.WriteString(fmt.Sprintf("%v", c.End))
	builder.WriteString(", ")
	builder.WriteString("message=")
	builder.WriteString(c.Message)
	builder.WriteString(", ")
	builder.WriteString("create_at=")
	builder.WriteString(fmt.Sprintf("%v", c.CreateAt))
	builder.WriteString(", ")
	builder.WriteString("update_at=")
	builder.WriteString(fmt.Sprintf("%v", c.UpdateAt))
	builder.WriteString(", ")
	builder.WriteString("delete_at=")
	builder.WriteString(fmt.Sprintf("%v", c.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// Compensates is a parsable slice of Compensate.
type Compensates []*Compensate

func (c Compensates) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
