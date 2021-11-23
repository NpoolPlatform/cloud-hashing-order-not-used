// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/payment"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// PaymentUpdate is the builder for updating Payment entities.
type PaymentUpdate struct {
	config
	hooks    []Hook
	mutation *PaymentMutation
}

// Where appends a list predicates to the PaymentUpdate builder.
func (pu *PaymentUpdate) Where(ps ...predicate.Payment) *PaymentUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetOrderID sets the "order_id" field.
func (pu *PaymentUpdate) SetOrderID(u uuid.UUID) *PaymentUpdate {
	pu.mutation.SetOrderID(u)
	return pu
}

// SetAccountID sets the "account_id" field.
func (pu *PaymentUpdate) SetAccountID(u uuid.UUID) *PaymentUpdate {
	pu.mutation.SetAccountID(u)
	return pu
}

// SetAmount sets the "amount" field.
func (pu *PaymentUpdate) SetAmount(u uint64) *PaymentUpdate {
	pu.mutation.ResetAmount()
	pu.mutation.SetAmount(u)
	return pu
}

// AddAmount adds u to the "amount" field.
func (pu *PaymentUpdate) AddAmount(u uint64) *PaymentUpdate {
	pu.mutation.AddAmount(u)
	return pu
}

// SetCoinInfoID sets the "coin_info_id" field.
func (pu *PaymentUpdate) SetCoinInfoID(u uuid.UUID) *PaymentUpdate {
	pu.mutation.SetCoinInfoID(u)
	return pu
}

// SetState sets the "state" field.
func (pu *PaymentUpdate) SetState(pa payment.State) *PaymentUpdate {
	pu.mutation.SetState(pa)
	return pu
}

// SetChainTransactionID sets the "chain_transaction_id" field.
func (pu *PaymentUpdate) SetChainTransactionID(s string) *PaymentUpdate {
	pu.mutation.SetChainTransactionID(s)
	return pu
}

// SetPlatformTransactionID sets the "platform_transaction_id" field.
func (pu *PaymentUpdate) SetPlatformTransactionID(u uuid.UUID) *PaymentUpdate {
	pu.mutation.SetPlatformTransactionID(u)
	return pu
}

// SetCreateAt sets the "create_at" field.
func (pu *PaymentUpdate) SetCreateAt(u uint32) *PaymentUpdate {
	pu.mutation.ResetCreateAt()
	pu.mutation.SetCreateAt(u)
	return pu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableCreateAt(u *uint32) *PaymentUpdate {
	if u != nil {
		pu.SetCreateAt(*u)
	}
	return pu
}

// AddCreateAt adds u to the "create_at" field.
func (pu *PaymentUpdate) AddCreateAt(u uint32) *PaymentUpdate {
	pu.mutation.AddCreateAt(u)
	return pu
}

// SetUpdateAt sets the "update_at" field.
func (pu *PaymentUpdate) SetUpdateAt(u uint32) *PaymentUpdate {
	pu.mutation.ResetUpdateAt()
	pu.mutation.SetUpdateAt(u)
	return pu
}

// AddUpdateAt adds u to the "update_at" field.
func (pu *PaymentUpdate) AddUpdateAt(u uint32) *PaymentUpdate {
	pu.mutation.AddUpdateAt(u)
	return pu
}

// SetDeleteAt sets the "delete_at" field.
func (pu *PaymentUpdate) SetDeleteAt(u uint32) *PaymentUpdate {
	pu.mutation.ResetDeleteAt()
	pu.mutation.SetDeleteAt(u)
	return pu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableDeleteAt(u *uint32) *PaymentUpdate {
	if u != nil {
		pu.SetDeleteAt(*u)
	}
	return pu
}

// AddDeleteAt adds u to the "delete_at" field.
func (pu *PaymentUpdate) AddDeleteAt(u uint32) *PaymentUpdate {
	pu.mutation.AddDeleteAt(u)
	return pu
}

// Mutation returns the PaymentMutation object of the builder.
func (pu *PaymentUpdate) Mutation() *PaymentMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PaymentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pu.defaults()
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PaymentUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PaymentUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PaymentUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PaymentUpdate) defaults() {
	if _, ok := pu.mutation.UpdateAt(); !ok {
		v := payment.UpdateDefaultUpdateAt()
		pu.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PaymentUpdate) check() error {
	if v, ok := pu.mutation.State(); ok {
		if err := payment.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("ent: validator failed for field \"state\": %w", err)}
		}
	}
	return nil
}

func (pu *PaymentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   payment.Table,
			Columns: payment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: payment.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: payment.FieldOrderID,
		})
	}
	if value, ok := pu.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: payment.FieldAccountID,
		})
	}
	if value, ok := pu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: payment.FieldAmount,
		})
	}
	if value, ok := pu.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: payment.FieldAmount,
		})
	}
	if value, ok := pu.mutation.CoinInfoID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: payment.FieldCoinInfoID,
		})
	}
	if value, ok := pu.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: payment.FieldState,
		})
	}
	if value, ok := pu.mutation.ChainTransactionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payment.FieldChainTransactionID,
		})
	}
	if value, ok := pu.mutation.PlatformTransactionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: payment.FieldPlatformTransactionID,
		})
	}
	if value, ok := pu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: payment.FieldCreateAt,
		})
	}
	if value, ok := pu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: payment.FieldCreateAt,
		})
	}
	if value, ok := pu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: payment.FieldUpdateAt,
		})
	}
	if value, ok := pu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: payment.FieldUpdateAt,
		})
	}
	if value, ok := pu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: payment.FieldDeleteAt,
		})
	}
	if value, ok := pu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: payment.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PaymentUpdateOne is the builder for updating a single Payment entity.
type PaymentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PaymentMutation
}

// SetOrderID sets the "order_id" field.
func (puo *PaymentUpdateOne) SetOrderID(u uuid.UUID) *PaymentUpdateOne {
	puo.mutation.SetOrderID(u)
	return puo
}

// SetAccountID sets the "account_id" field.
func (puo *PaymentUpdateOne) SetAccountID(u uuid.UUID) *PaymentUpdateOne {
	puo.mutation.SetAccountID(u)
	return puo
}

// SetAmount sets the "amount" field.
func (puo *PaymentUpdateOne) SetAmount(u uint64) *PaymentUpdateOne {
	puo.mutation.ResetAmount()
	puo.mutation.SetAmount(u)
	return puo
}

// AddAmount adds u to the "amount" field.
func (puo *PaymentUpdateOne) AddAmount(u uint64) *PaymentUpdateOne {
	puo.mutation.AddAmount(u)
	return puo
}

// SetCoinInfoID sets the "coin_info_id" field.
func (puo *PaymentUpdateOne) SetCoinInfoID(u uuid.UUID) *PaymentUpdateOne {
	puo.mutation.SetCoinInfoID(u)
	return puo
}

// SetState sets the "state" field.
func (puo *PaymentUpdateOne) SetState(pa payment.State) *PaymentUpdateOne {
	puo.mutation.SetState(pa)
	return puo
}

// SetChainTransactionID sets the "chain_transaction_id" field.
func (puo *PaymentUpdateOne) SetChainTransactionID(s string) *PaymentUpdateOne {
	puo.mutation.SetChainTransactionID(s)
	return puo
}

// SetPlatformTransactionID sets the "platform_transaction_id" field.
func (puo *PaymentUpdateOne) SetPlatformTransactionID(u uuid.UUID) *PaymentUpdateOne {
	puo.mutation.SetPlatformTransactionID(u)
	return puo
}

// SetCreateAt sets the "create_at" field.
func (puo *PaymentUpdateOne) SetCreateAt(u uint32) *PaymentUpdateOne {
	puo.mutation.ResetCreateAt()
	puo.mutation.SetCreateAt(u)
	return puo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableCreateAt(u *uint32) *PaymentUpdateOne {
	if u != nil {
		puo.SetCreateAt(*u)
	}
	return puo
}

// AddCreateAt adds u to the "create_at" field.
func (puo *PaymentUpdateOne) AddCreateAt(u uint32) *PaymentUpdateOne {
	puo.mutation.AddCreateAt(u)
	return puo
}

// SetUpdateAt sets the "update_at" field.
func (puo *PaymentUpdateOne) SetUpdateAt(u uint32) *PaymentUpdateOne {
	puo.mutation.ResetUpdateAt()
	puo.mutation.SetUpdateAt(u)
	return puo
}

// AddUpdateAt adds u to the "update_at" field.
func (puo *PaymentUpdateOne) AddUpdateAt(u uint32) *PaymentUpdateOne {
	puo.mutation.AddUpdateAt(u)
	return puo
}

// SetDeleteAt sets the "delete_at" field.
func (puo *PaymentUpdateOne) SetDeleteAt(u uint32) *PaymentUpdateOne {
	puo.mutation.ResetDeleteAt()
	puo.mutation.SetDeleteAt(u)
	return puo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableDeleteAt(u *uint32) *PaymentUpdateOne {
	if u != nil {
		puo.SetDeleteAt(*u)
	}
	return puo
}

// AddDeleteAt adds u to the "delete_at" field.
func (puo *PaymentUpdateOne) AddDeleteAt(u uint32) *PaymentUpdateOne {
	puo.mutation.AddDeleteAt(u)
	return puo
}

// Mutation returns the PaymentMutation object of the builder.
func (puo *PaymentUpdateOne) Mutation() *PaymentMutation {
	return puo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PaymentUpdateOne) Select(field string, fields ...string) *PaymentUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Payment entity.
func (puo *PaymentUpdateOne) Save(ctx context.Context) (*Payment, error) {
	var (
		err  error
		node *Payment
	)
	puo.defaults()
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PaymentUpdateOne) SaveX(ctx context.Context) *Payment {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PaymentUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PaymentUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PaymentUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdateAt(); !ok {
		v := payment.UpdateDefaultUpdateAt()
		puo.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PaymentUpdateOne) check() error {
	if v, ok := puo.mutation.State(); ok {
		if err := payment.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("ent: validator failed for field \"state\": %w", err)}
		}
	}
	return nil
}

func (puo *PaymentUpdateOne) sqlSave(ctx context.Context) (_node *Payment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   payment.Table,
			Columns: payment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: payment.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Payment.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, payment.FieldID)
		for _, f := range fields {
			if !payment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != payment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: payment.FieldOrderID,
		})
	}
	if value, ok := puo.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: payment.FieldAccountID,
		})
	}
	if value, ok := puo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: payment.FieldAmount,
		})
	}
	if value, ok := puo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: payment.FieldAmount,
		})
	}
	if value, ok := puo.mutation.CoinInfoID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: payment.FieldCoinInfoID,
		})
	}
	if value, ok := puo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: payment.FieldState,
		})
	}
	if value, ok := puo.mutation.ChainTransactionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payment.FieldChainTransactionID,
		})
	}
	if value, ok := puo.mutation.PlatformTransactionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: payment.FieldPlatformTransactionID,
		})
	}
	if value, ok := puo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: payment.FieldCreateAt,
		})
	}
	if value, ok := puo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: payment.FieldCreateAt,
		})
	}
	if value, ok := puo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: payment.FieldUpdateAt,
		})
	}
	if value, ok := puo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: payment.FieldUpdateAt,
		})
	}
	if value, ok := puo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: payment.FieldDeleteAt,
		})
	}
	if value, ok := puo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: payment.FieldDeleteAt,
		})
	}
	_node = &Payment{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}