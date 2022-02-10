// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/order"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks    []Hook
	mutation *OrderMutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetGoodID sets the "good_id" field.
func (ou *OrderUpdate) SetGoodID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetGoodID(u)
	return ou
}

// SetAppID sets the "app_id" field.
func (ou *OrderUpdate) SetAppID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetAppID(u)
	return ou
}

// SetUserID sets the "user_id" field.
func (ou *OrderUpdate) SetUserID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetUserID(u)
	return ou
}

// SetUnits sets the "units" field.
func (ou *OrderUpdate) SetUnits(u uint32) *OrderUpdate {
	ou.mutation.ResetUnits()
	ou.mutation.SetUnits(u)
	return ou
}

// AddUnits adds u to the "units" field.
func (ou *OrderUpdate) AddUnits(u int32) *OrderUpdate {
	ou.mutation.AddUnits(u)
	return ou
}

// SetDiscountCouponID sets the "discount_coupon_id" field.
func (ou *OrderUpdate) SetDiscountCouponID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetDiscountCouponID(u)
	return ou
}

// SetUserSpecialReductionID sets the "user_special_reduction_id" field.
func (ou *OrderUpdate) SetUserSpecialReductionID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetUserSpecialReductionID(u)
	return ou
}

// SetStart sets the "start" field.
func (ou *OrderUpdate) SetStart(u uint32) *OrderUpdate {
	ou.mutation.ResetStart()
	ou.mutation.SetStart(u)
	return ou
}

// AddStart adds u to the "start" field.
func (ou *OrderUpdate) AddStart(u int32) *OrderUpdate {
	ou.mutation.AddStart(u)
	return ou
}

// SetEnd sets the "end" field.
func (ou *OrderUpdate) SetEnd(u uint32) *OrderUpdate {
	ou.mutation.ResetEnd()
	ou.mutation.SetEnd(u)
	return ou
}

// AddEnd adds u to the "end" field.
func (ou *OrderUpdate) AddEnd(u int32) *OrderUpdate {
	ou.mutation.AddEnd(u)
	return ou
}

// SetCouponID sets the "coupon_id" field.
func (ou *OrderUpdate) SetCouponID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetCouponID(u)
	return ou
}

// SetCreateAt sets the "create_at" field.
func (ou *OrderUpdate) SetCreateAt(u uint32) *OrderUpdate {
	ou.mutation.ResetCreateAt()
	ou.mutation.SetCreateAt(u)
	return ou
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableCreateAt(u *uint32) *OrderUpdate {
	if u != nil {
		ou.SetCreateAt(*u)
	}
	return ou
}

// AddCreateAt adds u to the "create_at" field.
func (ou *OrderUpdate) AddCreateAt(u int32) *OrderUpdate {
	ou.mutation.AddCreateAt(u)
	return ou
}

// SetUpdateAt sets the "update_at" field.
func (ou *OrderUpdate) SetUpdateAt(u uint32) *OrderUpdate {
	ou.mutation.ResetUpdateAt()
	ou.mutation.SetUpdateAt(u)
	return ou
}

// AddUpdateAt adds u to the "update_at" field.
func (ou *OrderUpdate) AddUpdateAt(u int32) *OrderUpdate {
	ou.mutation.AddUpdateAt(u)
	return ou
}

// SetDeleteAt sets the "delete_at" field.
func (ou *OrderUpdate) SetDeleteAt(u uint32) *OrderUpdate {
	ou.mutation.ResetDeleteAt()
	ou.mutation.SetDeleteAt(u)
	return ou
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableDeleteAt(u *uint32) *OrderUpdate {
	if u != nil {
		ou.SetDeleteAt(*u)
	}
	return ou
}

// AddDeleteAt adds u to the "delete_at" field.
func (ou *OrderUpdate) AddDeleteAt(u int32) *OrderUpdate {
	ou.mutation.AddDeleteAt(u)
	return ou
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ou.defaults()
	if len(ou.hooks) == 0 {
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			if ou.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *OrderUpdate) defaults() {
	if _, ok := ou.mutation.UpdateAt(); !ok {
		v := order.UpdateDefaultUpdateAt()
		ou.mutation.SetUpdateAt(v)
	}
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: order.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldGoodID,
		})
	}
	if value, ok := ou.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldAppID,
		})
	}
	if value, ok := ou.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldUserID,
		})
	}
	if value, ok := ou.mutation.Units(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUnits,
		})
	}
	if value, ok := ou.mutation.AddedUnits(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUnits,
		})
	}
	if value, ok := ou.mutation.DiscountCouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldDiscountCouponID,
		})
	}
	if value, ok := ou.mutation.UserSpecialReductionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldUserSpecialReductionID,
		})
	}
	if value, ok := ou.mutation.Start(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldStart,
		})
	}
	if value, ok := ou.mutation.AddedStart(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldStart,
		})
	}
	if value, ok := ou.mutation.End(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldEnd,
		})
	}
	if value, ok := ou.mutation.AddedEnd(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldEnd,
		})
	}
	if value, ok := ou.mutation.CouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldCouponID,
		})
	}
	if value, ok := ou.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldCreateAt,
		})
	}
	if value, ok := ou.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldCreateAt,
		})
	}
	if value, ok := ou.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUpdateAt,
		})
	}
	if value, ok := ou.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUpdateAt,
		})
	}
	if value, ok := ou.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldDeleteAt,
		})
	}
	if value, ok := ou.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderMutation
}

// SetGoodID sets the "good_id" field.
func (ouo *OrderUpdateOne) SetGoodID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetGoodID(u)
	return ouo
}

// SetAppID sets the "app_id" field.
func (ouo *OrderUpdateOne) SetAppID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetAppID(u)
	return ouo
}

// SetUserID sets the "user_id" field.
func (ouo *OrderUpdateOne) SetUserID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetUserID(u)
	return ouo
}

// SetUnits sets the "units" field.
func (ouo *OrderUpdateOne) SetUnits(u uint32) *OrderUpdateOne {
	ouo.mutation.ResetUnits()
	ouo.mutation.SetUnits(u)
	return ouo
}

// AddUnits adds u to the "units" field.
func (ouo *OrderUpdateOne) AddUnits(u int32) *OrderUpdateOne {
	ouo.mutation.AddUnits(u)
	return ouo
}

// SetDiscountCouponID sets the "discount_coupon_id" field.
func (ouo *OrderUpdateOne) SetDiscountCouponID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetDiscountCouponID(u)
	return ouo
}

// SetUserSpecialReductionID sets the "user_special_reduction_id" field.
func (ouo *OrderUpdateOne) SetUserSpecialReductionID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetUserSpecialReductionID(u)
	return ouo
}

// SetStart sets the "start" field.
func (ouo *OrderUpdateOne) SetStart(u uint32) *OrderUpdateOne {
	ouo.mutation.ResetStart()
	ouo.mutation.SetStart(u)
	return ouo
}

// AddStart adds u to the "start" field.
func (ouo *OrderUpdateOne) AddStart(u int32) *OrderUpdateOne {
	ouo.mutation.AddStart(u)
	return ouo
}

// SetEnd sets the "end" field.
func (ouo *OrderUpdateOne) SetEnd(u uint32) *OrderUpdateOne {
	ouo.mutation.ResetEnd()
	ouo.mutation.SetEnd(u)
	return ouo
}

// AddEnd adds u to the "end" field.
func (ouo *OrderUpdateOne) AddEnd(u int32) *OrderUpdateOne {
	ouo.mutation.AddEnd(u)
	return ouo
}

// SetCouponID sets the "coupon_id" field.
func (ouo *OrderUpdateOne) SetCouponID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetCouponID(u)
	return ouo
}

// SetCreateAt sets the "create_at" field.
func (ouo *OrderUpdateOne) SetCreateAt(u uint32) *OrderUpdateOne {
	ouo.mutation.ResetCreateAt()
	ouo.mutation.SetCreateAt(u)
	return ouo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableCreateAt(u *uint32) *OrderUpdateOne {
	if u != nil {
		ouo.SetCreateAt(*u)
	}
	return ouo
}

// AddCreateAt adds u to the "create_at" field.
func (ouo *OrderUpdateOne) AddCreateAt(u int32) *OrderUpdateOne {
	ouo.mutation.AddCreateAt(u)
	return ouo
}

// SetUpdateAt sets the "update_at" field.
func (ouo *OrderUpdateOne) SetUpdateAt(u uint32) *OrderUpdateOne {
	ouo.mutation.ResetUpdateAt()
	ouo.mutation.SetUpdateAt(u)
	return ouo
}

// AddUpdateAt adds u to the "update_at" field.
func (ouo *OrderUpdateOne) AddUpdateAt(u int32) *OrderUpdateOne {
	ouo.mutation.AddUpdateAt(u)
	return ouo
}

// SetDeleteAt sets the "delete_at" field.
func (ouo *OrderUpdateOne) SetDeleteAt(u uint32) *OrderUpdateOne {
	ouo.mutation.ResetDeleteAt()
	ouo.mutation.SetDeleteAt(u)
	return ouo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableDeleteAt(u *uint32) *OrderUpdateOne {
	if u != nil {
		ouo.SetDeleteAt(*u)
	}
	return ouo
}

// AddDeleteAt adds u to the "delete_at" field.
func (ouo *OrderUpdateOne) AddDeleteAt(u int32) *OrderUpdateOne {
	ouo.mutation.AddDeleteAt(u)
	return ouo
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	var (
		err  error
		node *Order
	)
	ouo.defaults()
	if len(ouo.hooks) == 0 {
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			if ouo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *OrderUpdateOne) defaults() {
	if _, ok := ouo.mutation.UpdateAt(); !ok {
		v := order.UpdateDefaultUpdateAt()
		ouo.mutation.SetUpdateAt(v)
	}
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: order.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Order.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldGoodID,
		})
	}
	if value, ok := ouo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldAppID,
		})
	}
	if value, ok := ouo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldUserID,
		})
	}
	if value, ok := ouo.mutation.Units(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUnits,
		})
	}
	if value, ok := ouo.mutation.AddedUnits(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUnits,
		})
	}
	if value, ok := ouo.mutation.DiscountCouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldDiscountCouponID,
		})
	}
	if value, ok := ouo.mutation.UserSpecialReductionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldUserSpecialReductionID,
		})
	}
	if value, ok := ouo.mutation.Start(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldStart,
		})
	}
	if value, ok := ouo.mutation.AddedStart(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldStart,
		})
	}
	if value, ok := ouo.mutation.End(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldEnd,
		})
	}
	if value, ok := ouo.mutation.AddedEnd(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldEnd,
		})
	}
	if value, ok := ouo.mutation.CouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldCouponID,
		})
	}
	if value, ok := ouo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldCreateAt,
		})
	}
	if value, ok := ouo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldCreateAt,
		})
	}
	if value, ok := ouo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUpdateAt,
		})
	}
	if value, ok := ouo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUpdateAt,
		})
	}
	if value, ok := ouo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldDeleteAt,
		})
	}
	if value, ok := ouo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldDeleteAt,
		})
	}
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
