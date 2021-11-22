// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/outofgas"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// OutOfGasQuery is the builder for querying OutOfGas entities.
type OutOfGasQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OutOfGas
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OutOfGasQuery builder.
func (oogq *OutOfGasQuery) Where(ps ...predicate.OutOfGas) *OutOfGasQuery {
	oogq.predicates = append(oogq.predicates, ps...)
	return oogq
}

// Limit adds a limit step to the query.
func (oogq *OutOfGasQuery) Limit(limit int) *OutOfGasQuery {
	oogq.limit = &limit
	return oogq
}

// Offset adds an offset step to the query.
func (oogq *OutOfGasQuery) Offset(offset int) *OutOfGasQuery {
	oogq.offset = &offset
	return oogq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oogq *OutOfGasQuery) Unique(unique bool) *OutOfGasQuery {
	oogq.unique = &unique
	return oogq
}

// Order adds an order step to the query.
func (oogq *OutOfGasQuery) Order(o ...OrderFunc) *OutOfGasQuery {
	oogq.order = append(oogq.order, o...)
	return oogq
}

// First returns the first OutOfGas entity from the query.
// Returns a *NotFoundError when no OutOfGas was found.
func (oogq *OutOfGasQuery) First(ctx context.Context) (*OutOfGas, error) {
	nodes, err := oogq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{outofgas.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oogq *OutOfGasQuery) FirstX(ctx context.Context) *OutOfGas {
	node, err := oogq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OutOfGas ID from the query.
// Returns a *NotFoundError when no OutOfGas ID was found.
func (oogq *OutOfGasQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = oogq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{outofgas.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oogq *OutOfGasQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := oogq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OutOfGas entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one OutOfGas entity is not found.
// Returns a *NotFoundError when no OutOfGas entities are found.
func (oogq *OutOfGasQuery) Only(ctx context.Context) (*OutOfGas, error) {
	nodes, err := oogq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{outofgas.Label}
	default:
		return nil, &NotSingularError{outofgas.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oogq *OutOfGasQuery) OnlyX(ctx context.Context) *OutOfGas {
	node, err := oogq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OutOfGas ID in the query.
// Returns a *NotSingularError when exactly one OutOfGas ID is not found.
// Returns a *NotFoundError when no entities are found.
func (oogq *OutOfGasQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = oogq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{outofgas.Label}
	default:
		err = &NotSingularError{outofgas.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oogq *OutOfGasQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := oogq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OutOfGasSlice.
func (oogq *OutOfGasQuery) All(ctx context.Context) ([]*OutOfGas, error) {
	if err := oogq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oogq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oogq *OutOfGasQuery) AllX(ctx context.Context) []*OutOfGas {
	nodes, err := oogq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OutOfGas IDs.
func (oogq *OutOfGasQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := oogq.Select(outofgas.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oogq *OutOfGasQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := oogq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oogq *OutOfGasQuery) Count(ctx context.Context) (int, error) {
	if err := oogq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oogq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oogq *OutOfGasQuery) CountX(ctx context.Context) int {
	count, err := oogq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oogq *OutOfGasQuery) Exist(ctx context.Context) (bool, error) {
	if err := oogq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oogq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oogq *OutOfGasQuery) ExistX(ctx context.Context) bool {
	exist, err := oogq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OutOfGasQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oogq *OutOfGasQuery) Clone() *OutOfGasQuery {
	if oogq == nil {
		return nil
	}
	return &OutOfGasQuery{
		config:     oogq.config,
		limit:      oogq.limit,
		offset:     oogq.offset,
		order:      append([]OrderFunc{}, oogq.order...),
		predicates: append([]predicate.OutOfGas{}, oogq.predicates...),
		// clone intermediate query.
		sql:  oogq.sql.Clone(),
		path: oogq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OrderID uuid.UUID `json:"order_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OutOfGas.Query().
//		GroupBy(outofgas.FieldOrderID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (oogq *OutOfGasQuery) GroupBy(field string, fields ...string) *OutOfGasGroupBy {
	group := &OutOfGasGroupBy{config: oogq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oogq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oogq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OrderID uuid.UUID `json:"order_id,omitempty"`
//	}
//
//	client.OutOfGas.Query().
//		Select(outofgas.FieldOrderID).
//		Scan(ctx, &v)
//
func (oogq *OutOfGasQuery) Select(fields ...string) *OutOfGasSelect {
	oogq.fields = append(oogq.fields, fields...)
	return &OutOfGasSelect{OutOfGasQuery: oogq}
}

func (oogq *OutOfGasQuery) prepareQuery(ctx context.Context) error {
	for _, f := range oogq.fields {
		if !outofgas.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oogq.path != nil {
		prev, err := oogq.path(ctx)
		if err != nil {
			return err
		}
		oogq.sql = prev
	}
	return nil
}

func (oogq *OutOfGasQuery) sqlAll(ctx context.Context) ([]*OutOfGas, error) {
	var (
		nodes = []*OutOfGas{}
		_spec = oogq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &OutOfGas{config: oogq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, oogq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (oogq *OutOfGasQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oogq.querySpec()
	return sqlgraph.CountNodes(ctx, oogq.driver, _spec)
}

func (oogq *OutOfGasQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := oogq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (oogq *OutOfGasQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outofgas.Table,
			Columns: outofgas.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: outofgas.FieldID,
			},
		},
		From:   oogq.sql,
		Unique: true,
	}
	if unique := oogq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oogq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outofgas.FieldID)
		for i := range fields {
			if fields[i] != outofgas.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oogq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oogq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oogq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oogq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oogq *OutOfGasQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oogq.driver.Dialect())
	t1 := builder.Table(outofgas.Table)
	columns := oogq.fields
	if len(columns) == 0 {
		columns = outofgas.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oogq.sql != nil {
		selector = oogq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range oogq.predicates {
		p(selector)
	}
	for _, p := range oogq.order {
		p(selector)
	}
	if offset := oogq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oogq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OutOfGasGroupBy is the group-by builder for OutOfGas entities.
type OutOfGasGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ooggb *OutOfGasGroupBy) Aggregate(fns ...AggregateFunc) *OutOfGasGroupBy {
	ooggb.fns = append(ooggb.fns, fns...)
	return ooggb
}

// Scan applies the group-by query and scans the result into the given value.
func (ooggb *OutOfGasGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ooggb.path(ctx)
	if err != nil {
		return err
	}
	ooggb.sql = query
	return ooggb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ooggb *OutOfGasGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ooggb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ooggb *OutOfGasGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ooggb.fields) > 1 {
		return nil, errors.New("ent: OutOfGasGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ooggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ooggb *OutOfGasGroupBy) StringsX(ctx context.Context) []string {
	v, err := ooggb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ooggb *OutOfGasGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ooggb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outofgas.Label}
	default:
		err = fmt.Errorf("ent: OutOfGasGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ooggb *OutOfGasGroupBy) StringX(ctx context.Context) string {
	v, err := ooggb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ooggb *OutOfGasGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ooggb.fields) > 1 {
		return nil, errors.New("ent: OutOfGasGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ooggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ooggb *OutOfGasGroupBy) IntsX(ctx context.Context) []int {
	v, err := ooggb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ooggb *OutOfGasGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ooggb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outofgas.Label}
	default:
		err = fmt.Errorf("ent: OutOfGasGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ooggb *OutOfGasGroupBy) IntX(ctx context.Context) int {
	v, err := ooggb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ooggb *OutOfGasGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ooggb.fields) > 1 {
		return nil, errors.New("ent: OutOfGasGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ooggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ooggb *OutOfGasGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ooggb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ooggb *OutOfGasGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ooggb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outofgas.Label}
	default:
		err = fmt.Errorf("ent: OutOfGasGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ooggb *OutOfGasGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ooggb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ooggb *OutOfGasGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ooggb.fields) > 1 {
		return nil, errors.New("ent: OutOfGasGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ooggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ooggb *OutOfGasGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ooggb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ooggb *OutOfGasGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ooggb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outofgas.Label}
	default:
		err = fmt.Errorf("ent: OutOfGasGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ooggb *OutOfGasGroupBy) BoolX(ctx context.Context) bool {
	v, err := ooggb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ooggb *OutOfGasGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ooggb.fields {
		if !outofgas.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ooggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ooggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ooggb *OutOfGasGroupBy) sqlQuery() *sql.Selector {
	selector := ooggb.sql.Select()
	aggregation := make([]string, 0, len(ooggb.fns))
	for _, fn := range ooggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ooggb.fields)+len(ooggb.fns))
		for _, f := range ooggb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ooggb.fields...)...)
}

// OutOfGasSelect is the builder for selecting fields of OutOfGas entities.
type OutOfGasSelect struct {
	*OutOfGasQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (oogs *OutOfGasSelect) Scan(ctx context.Context, v interface{}) error {
	if err := oogs.prepareQuery(ctx); err != nil {
		return err
	}
	oogs.sql = oogs.OutOfGasQuery.sqlQuery(ctx)
	return oogs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (oogs *OutOfGasSelect) ScanX(ctx context.Context, v interface{}) {
	if err := oogs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (oogs *OutOfGasSelect) Strings(ctx context.Context) ([]string, error) {
	if len(oogs.fields) > 1 {
		return nil, errors.New("ent: OutOfGasSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := oogs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (oogs *OutOfGasSelect) StringsX(ctx context.Context) []string {
	v, err := oogs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (oogs *OutOfGasSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = oogs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outofgas.Label}
	default:
		err = fmt.Errorf("ent: OutOfGasSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (oogs *OutOfGasSelect) StringX(ctx context.Context) string {
	v, err := oogs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (oogs *OutOfGasSelect) Ints(ctx context.Context) ([]int, error) {
	if len(oogs.fields) > 1 {
		return nil, errors.New("ent: OutOfGasSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := oogs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (oogs *OutOfGasSelect) IntsX(ctx context.Context) []int {
	v, err := oogs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (oogs *OutOfGasSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = oogs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outofgas.Label}
	default:
		err = fmt.Errorf("ent: OutOfGasSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (oogs *OutOfGasSelect) IntX(ctx context.Context) int {
	v, err := oogs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (oogs *OutOfGasSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(oogs.fields) > 1 {
		return nil, errors.New("ent: OutOfGasSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := oogs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (oogs *OutOfGasSelect) Float64sX(ctx context.Context) []float64 {
	v, err := oogs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (oogs *OutOfGasSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = oogs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outofgas.Label}
	default:
		err = fmt.Errorf("ent: OutOfGasSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (oogs *OutOfGasSelect) Float64X(ctx context.Context) float64 {
	v, err := oogs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (oogs *OutOfGasSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(oogs.fields) > 1 {
		return nil, errors.New("ent: OutOfGasSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := oogs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (oogs *OutOfGasSelect) BoolsX(ctx context.Context) []bool {
	v, err := oogs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (oogs *OutOfGasSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = oogs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outofgas.Label}
	default:
		err = fmt.Errorf("ent: OutOfGasSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (oogs *OutOfGasSelect) BoolX(ctx context.Context) bool {
	v, err := oogs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oogs *OutOfGasSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := oogs.sql.Query()
	if err := oogs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
