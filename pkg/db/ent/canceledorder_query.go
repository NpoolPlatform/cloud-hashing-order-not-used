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
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/canceledorder"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CanceledOrderQuery is the builder for querying CanceledOrder entities.
type CanceledOrderQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CanceledOrder
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CanceledOrderQuery builder.
func (coq *CanceledOrderQuery) Where(ps ...predicate.CanceledOrder) *CanceledOrderQuery {
	coq.predicates = append(coq.predicates, ps...)
	return coq
}

// Limit adds a limit step to the query.
func (coq *CanceledOrderQuery) Limit(limit int) *CanceledOrderQuery {
	coq.limit = &limit
	return coq
}

// Offset adds an offset step to the query.
func (coq *CanceledOrderQuery) Offset(offset int) *CanceledOrderQuery {
	coq.offset = &offset
	return coq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (coq *CanceledOrderQuery) Unique(unique bool) *CanceledOrderQuery {
	coq.unique = &unique
	return coq
}

// Order adds an order step to the query.
func (coq *CanceledOrderQuery) Order(o ...OrderFunc) *CanceledOrderQuery {
	coq.order = append(coq.order, o...)
	return coq
}

// First returns the first CanceledOrder entity from the query.
// Returns a *NotFoundError when no CanceledOrder was found.
func (coq *CanceledOrderQuery) First(ctx context.Context) (*CanceledOrder, error) {
	nodes, err := coq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{canceledorder.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (coq *CanceledOrderQuery) FirstX(ctx context.Context) *CanceledOrder {
	node, err := coq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CanceledOrder ID from the query.
// Returns a *NotFoundError when no CanceledOrder ID was found.
func (coq *CanceledOrderQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = coq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{canceledorder.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (coq *CanceledOrderQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := coq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CanceledOrder entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one CanceledOrder entity is not found.
// Returns a *NotFoundError when no CanceledOrder entities are found.
func (coq *CanceledOrderQuery) Only(ctx context.Context) (*CanceledOrder, error) {
	nodes, err := coq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{canceledorder.Label}
	default:
		return nil, &NotSingularError{canceledorder.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (coq *CanceledOrderQuery) OnlyX(ctx context.Context) *CanceledOrder {
	node, err := coq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CanceledOrder ID in the query.
// Returns a *NotSingularError when exactly one CanceledOrder ID is not found.
// Returns a *NotFoundError when no entities are found.
func (coq *CanceledOrderQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = coq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{canceledorder.Label}
	default:
		err = &NotSingularError{canceledorder.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (coq *CanceledOrderQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := coq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CanceledOrders.
func (coq *CanceledOrderQuery) All(ctx context.Context) ([]*CanceledOrder, error) {
	if err := coq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return coq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (coq *CanceledOrderQuery) AllX(ctx context.Context) []*CanceledOrder {
	nodes, err := coq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CanceledOrder IDs.
func (coq *CanceledOrderQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := coq.Select(canceledorder.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (coq *CanceledOrderQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := coq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (coq *CanceledOrderQuery) Count(ctx context.Context) (int, error) {
	if err := coq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return coq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (coq *CanceledOrderQuery) CountX(ctx context.Context) int {
	count, err := coq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (coq *CanceledOrderQuery) Exist(ctx context.Context) (bool, error) {
	if err := coq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return coq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (coq *CanceledOrderQuery) ExistX(ctx context.Context) bool {
	exist, err := coq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CanceledOrderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (coq *CanceledOrderQuery) Clone() *CanceledOrderQuery {
	if coq == nil {
		return nil
	}
	return &CanceledOrderQuery{
		config:     coq.config,
		limit:      coq.limit,
		offset:     coq.offset,
		order:      append([]OrderFunc{}, coq.order...),
		predicates: append([]predicate.CanceledOrder{}, coq.predicates...),
		// clone intermediate query.
		sql:  coq.sql.Clone(),
		path: coq.path,
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
//	client.CanceledOrder.Query().
//		GroupBy(canceledorder.FieldOrderID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (coq *CanceledOrderQuery) GroupBy(field string, fields ...string) *CanceledOrderGroupBy {
	group := &CanceledOrderGroupBy{config: coq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := coq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return coq.sqlQuery(ctx), nil
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
//	client.CanceledOrder.Query().
//		Select(canceledorder.FieldOrderID).
//		Scan(ctx, &v)
//
func (coq *CanceledOrderQuery) Select(fields ...string) *CanceledOrderSelect {
	coq.fields = append(coq.fields, fields...)
	return &CanceledOrderSelect{CanceledOrderQuery: coq}
}

func (coq *CanceledOrderQuery) prepareQuery(ctx context.Context) error {
	for _, f := range coq.fields {
		if !canceledorder.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if coq.path != nil {
		prev, err := coq.path(ctx)
		if err != nil {
			return err
		}
		coq.sql = prev
	}
	return nil
}

func (coq *CanceledOrderQuery) sqlAll(ctx context.Context) ([]*CanceledOrder, error) {
	var (
		nodes = []*CanceledOrder{}
		_spec = coq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &CanceledOrder{config: coq.config}
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
	if err := sqlgraph.QueryNodes(ctx, coq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (coq *CanceledOrderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := coq.querySpec()
	return sqlgraph.CountNodes(ctx, coq.driver, _spec)
}

func (coq *CanceledOrderQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := coq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (coq *CanceledOrderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   canceledorder.Table,
			Columns: canceledorder.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: canceledorder.FieldID,
			},
		},
		From:   coq.sql,
		Unique: true,
	}
	if unique := coq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := coq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, canceledorder.FieldID)
		for i := range fields {
			if fields[i] != canceledorder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := coq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := coq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := coq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := coq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (coq *CanceledOrderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(coq.driver.Dialect())
	t1 := builder.Table(canceledorder.Table)
	columns := coq.fields
	if len(columns) == 0 {
		columns = canceledorder.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if coq.sql != nil {
		selector = coq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range coq.predicates {
		p(selector)
	}
	for _, p := range coq.order {
		p(selector)
	}
	if offset := coq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := coq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CanceledOrderGroupBy is the group-by builder for CanceledOrder entities.
type CanceledOrderGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cogb *CanceledOrderGroupBy) Aggregate(fns ...AggregateFunc) *CanceledOrderGroupBy {
	cogb.fns = append(cogb.fns, fns...)
	return cogb
}

// Scan applies the group-by query and scans the result into the given value.
func (cogb *CanceledOrderGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cogb.path(ctx)
	if err != nil {
		return err
	}
	cogb.sql = query
	return cogb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cogb *CanceledOrderGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cogb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (cogb *CanceledOrderGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cogb.fields) > 1 {
		return nil, errors.New("ent: CanceledOrderGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cogb *CanceledOrderGroupBy) StringsX(ctx context.Context) []string {
	v, err := cogb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cogb *CanceledOrderGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cogb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{canceledorder.Label}
	default:
		err = fmt.Errorf("ent: CanceledOrderGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cogb *CanceledOrderGroupBy) StringX(ctx context.Context) string {
	v, err := cogb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (cogb *CanceledOrderGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cogb.fields) > 1 {
		return nil, errors.New("ent: CanceledOrderGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cogb *CanceledOrderGroupBy) IntsX(ctx context.Context) []int {
	v, err := cogb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cogb *CanceledOrderGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cogb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{canceledorder.Label}
	default:
		err = fmt.Errorf("ent: CanceledOrderGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cogb *CanceledOrderGroupBy) IntX(ctx context.Context) int {
	v, err := cogb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (cogb *CanceledOrderGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cogb.fields) > 1 {
		return nil, errors.New("ent: CanceledOrderGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cogb *CanceledOrderGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cogb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cogb *CanceledOrderGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cogb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{canceledorder.Label}
	default:
		err = fmt.Errorf("ent: CanceledOrderGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cogb *CanceledOrderGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cogb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (cogb *CanceledOrderGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cogb.fields) > 1 {
		return nil, errors.New("ent: CanceledOrderGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cogb *CanceledOrderGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cogb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cogb *CanceledOrderGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cogb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{canceledorder.Label}
	default:
		err = fmt.Errorf("ent: CanceledOrderGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cogb *CanceledOrderGroupBy) BoolX(ctx context.Context) bool {
	v, err := cogb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cogb *CanceledOrderGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cogb.fields {
		if !canceledorder.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cogb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cogb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cogb *CanceledOrderGroupBy) sqlQuery() *sql.Selector {
	selector := cogb.sql.Select()
	aggregation := make([]string, 0, len(cogb.fns))
	for _, fn := range cogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cogb.fields)+len(cogb.fns))
		for _, f := range cogb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cogb.fields...)...)
}

// CanceledOrderSelect is the builder for selecting fields of CanceledOrder entities.
type CanceledOrderSelect struct {
	*CanceledOrderQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cos *CanceledOrderSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cos.prepareQuery(ctx); err != nil {
		return err
	}
	cos.sql = cos.CanceledOrderQuery.sqlQuery(ctx)
	return cos.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cos *CanceledOrderSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cos.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (cos *CanceledOrderSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cos.fields) > 1 {
		return nil, errors.New("ent: CanceledOrderSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cos *CanceledOrderSelect) StringsX(ctx context.Context) []string {
	v, err := cos.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (cos *CanceledOrderSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cos.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{canceledorder.Label}
	default:
		err = fmt.Errorf("ent: CanceledOrderSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cos *CanceledOrderSelect) StringX(ctx context.Context) string {
	v, err := cos.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (cos *CanceledOrderSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cos.fields) > 1 {
		return nil, errors.New("ent: CanceledOrderSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cos *CanceledOrderSelect) IntsX(ctx context.Context) []int {
	v, err := cos.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (cos *CanceledOrderSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cos.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{canceledorder.Label}
	default:
		err = fmt.Errorf("ent: CanceledOrderSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cos *CanceledOrderSelect) IntX(ctx context.Context) int {
	v, err := cos.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (cos *CanceledOrderSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cos.fields) > 1 {
		return nil, errors.New("ent: CanceledOrderSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cos *CanceledOrderSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cos.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (cos *CanceledOrderSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cos.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{canceledorder.Label}
	default:
		err = fmt.Errorf("ent: CanceledOrderSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cos *CanceledOrderSelect) Float64X(ctx context.Context) float64 {
	v, err := cos.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (cos *CanceledOrderSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cos.fields) > 1 {
		return nil, errors.New("ent: CanceledOrderSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cos *CanceledOrderSelect) BoolsX(ctx context.Context) []bool {
	v, err := cos.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (cos *CanceledOrderSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cos.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{canceledorder.Label}
	default:
		err = fmt.Errorf("ent: CanceledOrderSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cos *CanceledOrderSelect) BoolX(ctx context.Context) bool {
	v, err := cos.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cos *CanceledOrderSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cos.sql.Query()
	if err := cos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
