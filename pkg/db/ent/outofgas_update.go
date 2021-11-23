// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/outofgas"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// OutOfGasUpdate is the builder for updating OutOfGas entities.
type OutOfGasUpdate struct {
	config
	hooks    []Hook
	mutation *OutOfGasMutation
}

// Where appends a list predicates to the OutOfGasUpdate builder.
func (oogu *OutOfGasUpdate) Where(ps ...predicate.OutOfGas) *OutOfGasUpdate {
	oogu.mutation.Where(ps...)
	return oogu
}

// SetOrderID sets the "order_id" field.
func (oogu *OutOfGasUpdate) SetOrderID(u uuid.UUID) *OutOfGasUpdate {
	oogu.mutation.SetOrderID(u)
	return oogu
}

// SetStart sets the "start" field.
func (oogu *OutOfGasUpdate) SetStart(u uint32) *OutOfGasUpdate {
	oogu.mutation.ResetStart()
	oogu.mutation.SetStart(u)
	return oogu
}

// AddStart adds u to the "start" field.
func (oogu *OutOfGasUpdate) AddStart(u uint32) *OutOfGasUpdate {
	oogu.mutation.AddStart(u)
	return oogu
}

// SetEnd sets the "end" field.
func (oogu *OutOfGasUpdate) SetEnd(u uint32) *OutOfGasUpdate {
	oogu.mutation.ResetEnd()
	oogu.mutation.SetEnd(u)
	return oogu
}

// AddEnd adds u to the "end" field.
func (oogu *OutOfGasUpdate) AddEnd(u uint32) *OutOfGasUpdate {
	oogu.mutation.AddEnd(u)
	return oogu
}

// SetCreateAt sets the "create_at" field.
func (oogu *OutOfGasUpdate) SetCreateAt(u uint32) *OutOfGasUpdate {
	oogu.mutation.ResetCreateAt()
	oogu.mutation.SetCreateAt(u)
	return oogu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (oogu *OutOfGasUpdate) SetNillableCreateAt(u *uint32) *OutOfGasUpdate {
	if u != nil {
		oogu.SetCreateAt(*u)
	}
	return oogu
}

// AddCreateAt adds u to the "create_at" field.
func (oogu *OutOfGasUpdate) AddCreateAt(u uint32) *OutOfGasUpdate {
	oogu.mutation.AddCreateAt(u)
	return oogu
}

// SetUpdateAt sets the "update_at" field.
func (oogu *OutOfGasUpdate) SetUpdateAt(u uint32) *OutOfGasUpdate {
	oogu.mutation.ResetUpdateAt()
	oogu.mutation.SetUpdateAt(u)
	return oogu
}

// AddUpdateAt adds u to the "update_at" field.
func (oogu *OutOfGasUpdate) AddUpdateAt(u uint32) *OutOfGasUpdate {
	oogu.mutation.AddUpdateAt(u)
	return oogu
}

// SetDeleteAt sets the "delete_at" field.
func (oogu *OutOfGasUpdate) SetDeleteAt(u uint32) *OutOfGasUpdate {
	oogu.mutation.ResetDeleteAt()
	oogu.mutation.SetDeleteAt(u)
	return oogu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (oogu *OutOfGasUpdate) SetNillableDeleteAt(u *uint32) *OutOfGasUpdate {
	if u != nil {
		oogu.SetDeleteAt(*u)
	}
	return oogu
}

// AddDeleteAt adds u to the "delete_at" field.
func (oogu *OutOfGasUpdate) AddDeleteAt(u uint32) *OutOfGasUpdate {
	oogu.mutation.AddDeleteAt(u)
	return oogu
}

// Mutation returns the OutOfGasMutation object of the builder.
func (oogu *OutOfGasUpdate) Mutation() *OutOfGasMutation {
	return oogu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oogu *OutOfGasUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	oogu.defaults()
	if len(oogu.hooks) == 0 {
		affected, err = oogu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutOfGasMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oogu.mutation = mutation
			affected, err = oogu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oogu.hooks) - 1; i >= 0; i-- {
			if oogu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oogu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oogu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (oogu *OutOfGasUpdate) SaveX(ctx context.Context) int {
	affected, err := oogu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oogu *OutOfGasUpdate) Exec(ctx context.Context) error {
	_, err := oogu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oogu *OutOfGasUpdate) ExecX(ctx context.Context) {
	if err := oogu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oogu *OutOfGasUpdate) defaults() {
	if _, ok := oogu.mutation.UpdateAt(); !ok {
		v := outofgas.UpdateDefaultUpdateAt()
		oogu.mutation.SetUpdateAt(v)
	}
}

func (oogu *OutOfGasUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outofgas.Table,
			Columns: outofgas.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: outofgas.FieldID,
			},
		},
	}
	if ps := oogu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oogu.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: outofgas.FieldOrderID,
		})
	}
	if value, ok := oogu.mutation.Start(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldStart,
		})
	}
	if value, ok := oogu.mutation.AddedStart(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldStart,
		})
	}
	if value, ok := oogu.mutation.End(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldEnd,
		})
	}
	if value, ok := oogu.mutation.AddedEnd(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldEnd,
		})
	}
	if value, ok := oogu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldCreateAt,
		})
	}
	if value, ok := oogu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldCreateAt,
		})
	}
	if value, ok := oogu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldUpdateAt,
		})
	}
	if value, ok := oogu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldUpdateAt,
		})
	}
	if value, ok := oogu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldDeleteAt,
		})
	}
	if value, ok := oogu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, oogu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outofgas.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OutOfGasUpdateOne is the builder for updating a single OutOfGas entity.
type OutOfGasUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OutOfGasMutation
}

// SetOrderID sets the "order_id" field.
func (ooguo *OutOfGasUpdateOne) SetOrderID(u uuid.UUID) *OutOfGasUpdateOne {
	ooguo.mutation.SetOrderID(u)
	return ooguo
}

// SetStart sets the "start" field.
func (ooguo *OutOfGasUpdateOne) SetStart(u uint32) *OutOfGasUpdateOne {
	ooguo.mutation.ResetStart()
	ooguo.mutation.SetStart(u)
	return ooguo
}

// AddStart adds u to the "start" field.
func (ooguo *OutOfGasUpdateOne) AddStart(u uint32) *OutOfGasUpdateOne {
	ooguo.mutation.AddStart(u)
	return ooguo
}

// SetEnd sets the "end" field.
func (ooguo *OutOfGasUpdateOne) SetEnd(u uint32) *OutOfGasUpdateOne {
	ooguo.mutation.ResetEnd()
	ooguo.mutation.SetEnd(u)
	return ooguo
}

// AddEnd adds u to the "end" field.
func (ooguo *OutOfGasUpdateOne) AddEnd(u uint32) *OutOfGasUpdateOne {
	ooguo.mutation.AddEnd(u)
	return ooguo
}

// SetCreateAt sets the "create_at" field.
func (ooguo *OutOfGasUpdateOne) SetCreateAt(u uint32) *OutOfGasUpdateOne {
	ooguo.mutation.ResetCreateAt()
	ooguo.mutation.SetCreateAt(u)
	return ooguo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (ooguo *OutOfGasUpdateOne) SetNillableCreateAt(u *uint32) *OutOfGasUpdateOne {
	if u != nil {
		ooguo.SetCreateAt(*u)
	}
	return ooguo
}

// AddCreateAt adds u to the "create_at" field.
func (ooguo *OutOfGasUpdateOne) AddCreateAt(u uint32) *OutOfGasUpdateOne {
	ooguo.mutation.AddCreateAt(u)
	return ooguo
}

// SetUpdateAt sets the "update_at" field.
func (ooguo *OutOfGasUpdateOne) SetUpdateAt(u uint32) *OutOfGasUpdateOne {
	ooguo.mutation.ResetUpdateAt()
	ooguo.mutation.SetUpdateAt(u)
	return ooguo
}

// AddUpdateAt adds u to the "update_at" field.
func (ooguo *OutOfGasUpdateOne) AddUpdateAt(u uint32) *OutOfGasUpdateOne {
	ooguo.mutation.AddUpdateAt(u)
	return ooguo
}

// SetDeleteAt sets the "delete_at" field.
func (ooguo *OutOfGasUpdateOne) SetDeleteAt(u uint32) *OutOfGasUpdateOne {
	ooguo.mutation.ResetDeleteAt()
	ooguo.mutation.SetDeleteAt(u)
	return ooguo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (ooguo *OutOfGasUpdateOne) SetNillableDeleteAt(u *uint32) *OutOfGasUpdateOne {
	if u != nil {
		ooguo.SetDeleteAt(*u)
	}
	return ooguo
}

// AddDeleteAt adds u to the "delete_at" field.
func (ooguo *OutOfGasUpdateOne) AddDeleteAt(u uint32) *OutOfGasUpdateOne {
	ooguo.mutation.AddDeleteAt(u)
	return ooguo
}

// Mutation returns the OutOfGasMutation object of the builder.
func (ooguo *OutOfGasUpdateOne) Mutation() *OutOfGasMutation {
	return ooguo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ooguo *OutOfGasUpdateOne) Select(field string, fields ...string) *OutOfGasUpdateOne {
	ooguo.fields = append([]string{field}, fields...)
	return ooguo
}

// Save executes the query and returns the updated OutOfGas entity.
func (ooguo *OutOfGasUpdateOne) Save(ctx context.Context) (*OutOfGas, error) {
	var (
		err  error
		node *OutOfGas
	)
	ooguo.defaults()
	if len(ooguo.hooks) == 0 {
		node, err = ooguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutOfGasMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ooguo.mutation = mutation
			node, err = ooguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ooguo.hooks) - 1; i >= 0; i-- {
			if ooguo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ooguo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ooguo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ooguo *OutOfGasUpdateOne) SaveX(ctx context.Context) *OutOfGas {
	node, err := ooguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ooguo *OutOfGasUpdateOne) Exec(ctx context.Context) error {
	_, err := ooguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ooguo *OutOfGasUpdateOne) ExecX(ctx context.Context) {
	if err := ooguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ooguo *OutOfGasUpdateOne) defaults() {
	if _, ok := ooguo.mutation.UpdateAt(); !ok {
		v := outofgas.UpdateDefaultUpdateAt()
		ooguo.mutation.SetUpdateAt(v)
	}
}

func (ooguo *OutOfGasUpdateOne) sqlSave(ctx context.Context) (_node *OutOfGas, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outofgas.Table,
			Columns: outofgas.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: outofgas.FieldID,
			},
		},
	}
	id, ok := ooguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing OutOfGas.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ooguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outofgas.FieldID)
		for _, f := range fields {
			if !outofgas.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != outofgas.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ooguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ooguo.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: outofgas.FieldOrderID,
		})
	}
	if value, ok := ooguo.mutation.Start(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldStart,
		})
	}
	if value, ok := ooguo.mutation.AddedStart(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldStart,
		})
	}
	if value, ok := ooguo.mutation.End(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldEnd,
		})
	}
	if value, ok := ooguo.mutation.AddedEnd(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldEnd,
		})
	}
	if value, ok := ooguo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldCreateAt,
		})
	}
	if value, ok := ooguo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldCreateAt,
		})
	}
	if value, ok := ooguo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldUpdateAt,
		})
	}
	if value, ok := ooguo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldUpdateAt,
		})
	}
	if value, ok := ooguo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldDeleteAt,
		})
	}
	if value, ok := ooguo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldDeleteAt,
		})
	}
	_node = &OutOfGas{config: ooguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ooguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outofgas.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}