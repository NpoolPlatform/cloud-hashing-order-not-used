// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/canceledorder"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/predicate"
)

// CanceledOrderDelete is the builder for deleting a CanceledOrder entity.
type CanceledOrderDelete struct {
	config
	hooks    []Hook
	mutation *CanceledOrderMutation
}

// Where appends a list predicates to the CanceledOrderDelete builder.
func (cod *CanceledOrderDelete) Where(ps ...predicate.CanceledOrder) *CanceledOrderDelete {
	cod.mutation.Where(ps...)
	return cod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cod *CanceledOrderDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cod.hooks) == 0 {
		affected, err = cod.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CanceledOrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cod.mutation = mutation
			affected, err = cod.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cod.hooks) - 1; i >= 0; i-- {
			if cod.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cod.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cod.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cod *CanceledOrderDelete) ExecX(ctx context.Context) int {
	n, err := cod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cod *CanceledOrderDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: canceledorder.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: canceledorder.FieldID,
			},
		},
	}
	if ps := cod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cod.driver, _spec)
}

// CanceledOrderDeleteOne is the builder for deleting a single CanceledOrder entity.
type CanceledOrderDeleteOne struct {
	cod *CanceledOrderDelete
}

// Exec executes the deletion query.
func (codo *CanceledOrderDeleteOne) Exec(ctx context.Context) error {
	n, err := codo.cod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{canceledorder.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (codo *CanceledOrderDeleteOne) ExecX(ctx context.Context) {
	codo.cod.ExecX(ctx)
}
