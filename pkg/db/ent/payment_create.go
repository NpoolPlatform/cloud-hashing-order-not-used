// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/payment"
	"github.com/google/uuid"
)

// PaymentCreate is the builder for creating a Payment entity.
type PaymentCreate struct {
	config
	mutation *PaymentMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetOrderID sets the "order_id" field.
func (pc *PaymentCreate) SetOrderID(u uuid.UUID) *PaymentCreate {
	pc.mutation.SetOrderID(u)
	return pc
}

// SetAccountID sets the "account_id" field.
func (pc *PaymentCreate) SetAccountID(u uuid.UUID) *PaymentCreate {
	pc.mutation.SetAccountID(u)
	return pc
}

// SetStartAmount sets the "start_amount" field.
func (pc *PaymentCreate) SetStartAmount(u uint64) *PaymentCreate {
	pc.mutation.SetStartAmount(u)
	return pc
}

// SetAmount sets the "amount" field.
func (pc *PaymentCreate) SetAmount(u uint64) *PaymentCreate {
	pc.mutation.SetAmount(u)
	return pc
}

// SetCoinInfoID sets the "coin_info_id" field.
func (pc *PaymentCreate) SetCoinInfoID(u uuid.UUID) *PaymentCreate {
	pc.mutation.SetCoinInfoID(u)
	return pc
}

// SetState sets the "state" field.
func (pc *PaymentCreate) SetState(pa payment.State) *PaymentCreate {
	pc.mutation.SetState(pa)
	return pc
}

// SetChainTransactionID sets the "chain_transaction_id" field.
func (pc *PaymentCreate) SetChainTransactionID(s string) *PaymentCreate {
	pc.mutation.SetChainTransactionID(s)
	return pc
}

// SetPlatformTransactionID sets the "platform_transaction_id" field.
func (pc *PaymentCreate) SetPlatformTransactionID(u uuid.UUID) *PaymentCreate {
	pc.mutation.SetPlatformTransactionID(u)
	return pc
}

// SetCreateAt sets the "create_at" field.
func (pc *PaymentCreate) SetCreateAt(u uint32) *PaymentCreate {
	pc.mutation.SetCreateAt(u)
	return pc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (pc *PaymentCreate) SetNillableCreateAt(u *uint32) *PaymentCreate {
	if u != nil {
		pc.SetCreateAt(*u)
	}
	return pc
}

// SetUpdateAt sets the "update_at" field.
func (pc *PaymentCreate) SetUpdateAt(u uint32) *PaymentCreate {
	pc.mutation.SetUpdateAt(u)
	return pc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (pc *PaymentCreate) SetNillableUpdateAt(u *uint32) *PaymentCreate {
	if u != nil {
		pc.SetUpdateAt(*u)
	}
	return pc
}

// SetDeleteAt sets the "delete_at" field.
func (pc *PaymentCreate) SetDeleteAt(u uint32) *PaymentCreate {
	pc.mutation.SetDeleteAt(u)
	return pc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (pc *PaymentCreate) SetNillableDeleteAt(u *uint32) *PaymentCreate {
	if u != nil {
		pc.SetDeleteAt(*u)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PaymentCreate) SetID(u uuid.UUID) *PaymentCreate {
	pc.mutation.SetID(u)
	return pc
}

// Mutation returns the PaymentMutation object of the builder.
func (pc *PaymentCreate) Mutation() *PaymentMutation {
	return pc.mutation
}

// Save creates the Payment in the database.
func (pc *PaymentCreate) Save(ctx context.Context) (*Payment, error) {
	var (
		err  error
		node *Payment
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PaymentCreate) SaveX(ctx context.Context) *Payment {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PaymentCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PaymentCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PaymentCreate) defaults() {
	if _, ok := pc.mutation.CreateAt(); !ok {
		v := payment.DefaultCreateAt()
		pc.mutation.SetCreateAt(v)
	}
	if _, ok := pc.mutation.UpdateAt(); !ok {
		v := payment.DefaultUpdateAt()
		pc.mutation.SetUpdateAt(v)
	}
	if _, ok := pc.mutation.DeleteAt(); !ok {
		v := payment.DefaultDeleteAt()
		pc.mutation.SetDeleteAt(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := payment.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PaymentCreate) check() error {
	if _, ok := pc.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`ent: missing required field "order_id"`)}
	}
	if _, ok := pc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account_id", err: errors.New(`ent: missing required field "account_id"`)}
	}
	if _, ok := pc.mutation.StartAmount(); !ok {
		return &ValidationError{Name: "start_amount", err: errors.New(`ent: missing required field "start_amount"`)}
	}
	if _, ok := pc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "amount"`)}
	}
	if _, ok := pc.mutation.CoinInfoID(); !ok {
		return &ValidationError{Name: "coin_info_id", err: errors.New(`ent: missing required field "coin_info_id"`)}
	}
	if _, ok := pc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "state"`)}
	}
	if v, ok := pc.mutation.State(); ok {
		if err := payment.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "state": %w`, err)}
		}
	}
	if _, ok := pc.mutation.ChainTransactionID(); !ok {
		return &ValidationError{Name: "chain_transaction_id", err: errors.New(`ent: missing required field "chain_transaction_id"`)}
	}
	if _, ok := pc.mutation.PlatformTransactionID(); !ok {
		return &ValidationError{Name: "platform_transaction_id", err: errors.New(`ent: missing required field "platform_transaction_id"`)}
	}
	if _, ok := pc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := pc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := pc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (pc *PaymentCreate) sqlSave(ctx context.Context) (*Payment, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (pc *PaymentCreate) createSpec() (*Payment, *sqlgraph.CreateSpec) {
	var (
		_node = &Payment{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: payment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: payment.FieldID,
			},
		}
	)
	_spec.OnConflict = pc.conflict
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.OrderID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: payment.FieldOrderID,
		})
		_node.OrderID = value
	}
	if value, ok := pc.mutation.AccountID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: payment.FieldAccountID,
		})
		_node.AccountID = value
	}
	if value, ok := pc.mutation.StartAmount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: payment.FieldStartAmount,
		})
		_node.StartAmount = value
	}
	if value, ok := pc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: payment.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := pc.mutation.CoinInfoID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: payment.FieldCoinInfoID,
		})
		_node.CoinInfoID = value
	}
	if value, ok := pc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: payment.FieldState,
		})
		_node.State = value
	}
	if value, ok := pc.mutation.ChainTransactionID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payment.FieldChainTransactionID,
		})
		_node.ChainTransactionID = value
	}
	if value, ok := pc.mutation.PlatformTransactionID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: payment.FieldPlatformTransactionID,
		})
		_node.PlatformTransactionID = value
	}
	if value, ok := pc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: payment.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := pc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: payment.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := pc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: payment.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Payment.Create().
//		SetOrderID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PaymentUpsert) {
//			SetOrderID(v+v).
//		}).
//		Exec(ctx)
//
func (pc *PaymentCreate) OnConflict(opts ...sql.ConflictOption) *PaymentUpsertOne {
	pc.conflict = opts
	return &PaymentUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Payment.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (pc *PaymentCreate) OnConflictColumns(columns ...string) *PaymentUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PaymentUpsertOne{
		create: pc,
	}
}

type (
	// PaymentUpsertOne is the builder for "upsert"-ing
	//  one Payment node.
	PaymentUpsertOne struct {
		create *PaymentCreate
	}

	// PaymentUpsert is the "OnConflict" setter.
	PaymentUpsert struct {
		*sql.UpdateSet
	}
)

// SetOrderID sets the "order_id" field.
func (u *PaymentUpsert) SetOrderID(v uuid.UUID) *PaymentUpsert {
	u.Set(payment.FieldOrderID, v)
	return u
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *PaymentUpsert) UpdateOrderID() *PaymentUpsert {
	u.SetExcluded(payment.FieldOrderID)
	return u
}

// SetAccountID sets the "account_id" field.
func (u *PaymentUpsert) SetAccountID(v uuid.UUID) *PaymentUpsert {
	u.Set(payment.FieldAccountID, v)
	return u
}

// UpdateAccountID sets the "account_id" field to the value that was provided on create.
func (u *PaymentUpsert) UpdateAccountID() *PaymentUpsert {
	u.SetExcluded(payment.FieldAccountID)
	return u
}

// SetStartAmount sets the "start_amount" field.
func (u *PaymentUpsert) SetStartAmount(v uint64) *PaymentUpsert {
	u.Set(payment.FieldStartAmount, v)
	return u
}

// UpdateStartAmount sets the "start_amount" field to the value that was provided on create.
func (u *PaymentUpsert) UpdateStartAmount() *PaymentUpsert {
	u.SetExcluded(payment.FieldStartAmount)
	return u
}

// SetAmount sets the "amount" field.
func (u *PaymentUpsert) SetAmount(v uint64) *PaymentUpsert {
	u.Set(payment.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *PaymentUpsert) UpdateAmount() *PaymentUpsert {
	u.SetExcluded(payment.FieldAmount)
	return u
}

// SetCoinInfoID sets the "coin_info_id" field.
func (u *PaymentUpsert) SetCoinInfoID(v uuid.UUID) *PaymentUpsert {
	u.Set(payment.FieldCoinInfoID, v)
	return u
}

// UpdateCoinInfoID sets the "coin_info_id" field to the value that was provided on create.
func (u *PaymentUpsert) UpdateCoinInfoID() *PaymentUpsert {
	u.SetExcluded(payment.FieldCoinInfoID)
	return u
}

// SetState sets the "state" field.
func (u *PaymentUpsert) SetState(v payment.State) *PaymentUpsert {
	u.Set(payment.FieldState, v)
	return u
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *PaymentUpsert) UpdateState() *PaymentUpsert {
	u.SetExcluded(payment.FieldState)
	return u
}

// SetChainTransactionID sets the "chain_transaction_id" field.
func (u *PaymentUpsert) SetChainTransactionID(v string) *PaymentUpsert {
	u.Set(payment.FieldChainTransactionID, v)
	return u
}

// UpdateChainTransactionID sets the "chain_transaction_id" field to the value that was provided on create.
func (u *PaymentUpsert) UpdateChainTransactionID() *PaymentUpsert {
	u.SetExcluded(payment.FieldChainTransactionID)
	return u
}

// SetPlatformTransactionID sets the "platform_transaction_id" field.
func (u *PaymentUpsert) SetPlatformTransactionID(v uuid.UUID) *PaymentUpsert {
	u.Set(payment.FieldPlatformTransactionID, v)
	return u
}

// UpdatePlatformTransactionID sets the "platform_transaction_id" field to the value that was provided on create.
func (u *PaymentUpsert) UpdatePlatformTransactionID() *PaymentUpsert {
	u.SetExcluded(payment.FieldPlatformTransactionID)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *PaymentUpsert) SetCreateAt(v uint32) *PaymentUpsert {
	u.Set(payment.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *PaymentUpsert) UpdateCreateAt() *PaymentUpsert {
	u.SetExcluded(payment.FieldCreateAt)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *PaymentUpsert) SetUpdateAt(v uint32) *PaymentUpsert {
	u.Set(payment.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *PaymentUpsert) UpdateUpdateAt() *PaymentUpsert {
	u.SetExcluded(payment.FieldUpdateAt)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *PaymentUpsert) SetDeleteAt(v uint32) *PaymentUpsert {
	u.Set(payment.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *PaymentUpsert) UpdateDeleteAt() *PaymentUpsert {
	u.SetExcluded(payment.FieldDeleteAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Payment.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(payment.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *PaymentUpsertOne) UpdateNewValues() *PaymentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(payment.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Payment.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *PaymentUpsertOne) Ignore() *PaymentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PaymentUpsertOne) DoNothing() *PaymentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PaymentCreate.OnConflict
// documentation for more info.
func (u *PaymentUpsertOne) Update(set func(*PaymentUpsert)) *PaymentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PaymentUpsert{UpdateSet: update})
	}))
	return u
}

// SetOrderID sets the "order_id" field.
func (u *PaymentUpsertOne) SetOrderID(v uuid.UUID) *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *PaymentUpsertOne) UpdateOrderID() *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdateOrderID()
	})
}

// SetAccountID sets the "account_id" field.
func (u *PaymentUpsertOne) SetAccountID(v uuid.UUID) *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.SetAccountID(v)
	})
}

// UpdateAccountID sets the "account_id" field to the value that was provided on create.
func (u *PaymentUpsertOne) UpdateAccountID() *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdateAccountID()
	})
}

// SetStartAmount sets the "start_amount" field.
func (u *PaymentUpsertOne) SetStartAmount(v uint64) *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.SetStartAmount(v)
	})
}

// UpdateStartAmount sets the "start_amount" field to the value that was provided on create.
func (u *PaymentUpsertOne) UpdateStartAmount() *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdateStartAmount()
	})
}

// SetAmount sets the "amount" field.
func (u *PaymentUpsertOne) SetAmount(v uint64) *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.SetAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *PaymentUpsertOne) UpdateAmount() *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdateAmount()
	})
}

// SetCoinInfoID sets the "coin_info_id" field.
func (u *PaymentUpsertOne) SetCoinInfoID(v uuid.UUID) *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.SetCoinInfoID(v)
	})
}

// UpdateCoinInfoID sets the "coin_info_id" field to the value that was provided on create.
func (u *PaymentUpsertOne) UpdateCoinInfoID() *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdateCoinInfoID()
	})
}

// SetState sets the "state" field.
func (u *PaymentUpsertOne) SetState(v payment.State) *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.SetState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *PaymentUpsertOne) UpdateState() *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdateState()
	})
}

// SetChainTransactionID sets the "chain_transaction_id" field.
func (u *PaymentUpsertOne) SetChainTransactionID(v string) *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.SetChainTransactionID(v)
	})
}

// UpdateChainTransactionID sets the "chain_transaction_id" field to the value that was provided on create.
func (u *PaymentUpsertOne) UpdateChainTransactionID() *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdateChainTransactionID()
	})
}

// SetPlatformTransactionID sets the "platform_transaction_id" field.
func (u *PaymentUpsertOne) SetPlatformTransactionID(v uuid.UUID) *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.SetPlatformTransactionID(v)
	})
}

// UpdatePlatformTransactionID sets the "platform_transaction_id" field to the value that was provided on create.
func (u *PaymentUpsertOne) UpdatePlatformTransactionID() *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdatePlatformTransactionID()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *PaymentUpsertOne) SetCreateAt(v uint32) *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *PaymentUpsertOne) UpdateCreateAt() *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *PaymentUpsertOne) SetUpdateAt(v uint32) *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *PaymentUpsertOne) UpdateUpdateAt() *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *PaymentUpsertOne) SetDeleteAt(v uint32) *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *PaymentUpsertOne) UpdateDeleteAt() *PaymentUpsertOne {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *PaymentUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PaymentCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PaymentUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PaymentUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: PaymentUpsertOne.ID is not supported by MySQL driver. Use PaymentUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PaymentUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PaymentCreateBulk is the builder for creating many Payment entities in bulk.
type PaymentCreateBulk struct {
	config
	builders []*PaymentCreate
	conflict []sql.ConflictOption
}

// Save creates the Payment entities in the database.
func (pcb *PaymentCreateBulk) Save(ctx context.Context) ([]*Payment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Payment, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PaymentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PaymentCreateBulk) SaveX(ctx context.Context) []*Payment {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PaymentCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PaymentCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Payment.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PaymentUpsert) {
//			SetOrderID(v+v).
//		}).
//		Exec(ctx)
//
func (pcb *PaymentCreateBulk) OnConflict(opts ...sql.ConflictOption) *PaymentUpsertBulk {
	pcb.conflict = opts
	return &PaymentUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Payment.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (pcb *PaymentCreateBulk) OnConflictColumns(columns ...string) *PaymentUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PaymentUpsertBulk{
		create: pcb,
	}
}

// PaymentUpsertBulk is the builder for "upsert"-ing
// a bulk of Payment nodes.
type PaymentUpsertBulk struct {
	create *PaymentCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Payment.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(payment.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *PaymentUpsertBulk) UpdateNewValues() *PaymentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(payment.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Payment.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *PaymentUpsertBulk) Ignore() *PaymentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PaymentUpsertBulk) DoNothing() *PaymentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PaymentCreateBulk.OnConflict
// documentation for more info.
func (u *PaymentUpsertBulk) Update(set func(*PaymentUpsert)) *PaymentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PaymentUpsert{UpdateSet: update})
	}))
	return u
}

// SetOrderID sets the "order_id" field.
func (u *PaymentUpsertBulk) SetOrderID(v uuid.UUID) *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *PaymentUpsertBulk) UpdateOrderID() *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdateOrderID()
	})
}

// SetAccountID sets the "account_id" field.
func (u *PaymentUpsertBulk) SetAccountID(v uuid.UUID) *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.SetAccountID(v)
	})
}

// UpdateAccountID sets the "account_id" field to the value that was provided on create.
func (u *PaymentUpsertBulk) UpdateAccountID() *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdateAccountID()
	})
}

// SetStartAmount sets the "start_amount" field.
func (u *PaymentUpsertBulk) SetStartAmount(v uint64) *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.SetStartAmount(v)
	})
}

// UpdateStartAmount sets the "start_amount" field to the value that was provided on create.
func (u *PaymentUpsertBulk) UpdateStartAmount() *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdateStartAmount()
	})
}

// SetAmount sets the "amount" field.
func (u *PaymentUpsertBulk) SetAmount(v uint64) *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.SetAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *PaymentUpsertBulk) UpdateAmount() *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdateAmount()
	})
}

// SetCoinInfoID sets the "coin_info_id" field.
func (u *PaymentUpsertBulk) SetCoinInfoID(v uuid.UUID) *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.SetCoinInfoID(v)
	})
}

// UpdateCoinInfoID sets the "coin_info_id" field to the value that was provided on create.
func (u *PaymentUpsertBulk) UpdateCoinInfoID() *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdateCoinInfoID()
	})
}

// SetState sets the "state" field.
func (u *PaymentUpsertBulk) SetState(v payment.State) *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.SetState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *PaymentUpsertBulk) UpdateState() *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdateState()
	})
}

// SetChainTransactionID sets the "chain_transaction_id" field.
func (u *PaymentUpsertBulk) SetChainTransactionID(v string) *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.SetChainTransactionID(v)
	})
}

// UpdateChainTransactionID sets the "chain_transaction_id" field to the value that was provided on create.
func (u *PaymentUpsertBulk) UpdateChainTransactionID() *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdateChainTransactionID()
	})
}

// SetPlatformTransactionID sets the "platform_transaction_id" field.
func (u *PaymentUpsertBulk) SetPlatformTransactionID(v uuid.UUID) *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.SetPlatformTransactionID(v)
	})
}

// UpdatePlatformTransactionID sets the "platform_transaction_id" field to the value that was provided on create.
func (u *PaymentUpsertBulk) UpdatePlatformTransactionID() *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdatePlatformTransactionID()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *PaymentUpsertBulk) SetCreateAt(v uint32) *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *PaymentUpsertBulk) UpdateCreateAt() *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *PaymentUpsertBulk) SetUpdateAt(v uint32) *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *PaymentUpsertBulk) UpdateUpdateAt() *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *PaymentUpsertBulk) SetDeleteAt(v uint32) *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *PaymentUpsertBulk) UpdateDeleteAt() *PaymentUpsertBulk {
	return u.Update(func(s *PaymentUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *PaymentUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PaymentCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PaymentCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PaymentUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
