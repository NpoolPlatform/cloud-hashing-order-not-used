// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/goodpaying"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// GoodPayingQuery is the builder for querying GoodPaying entities.
type GoodPayingQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GoodPaying
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GoodPayingQuery builder.
func (gpq *GoodPayingQuery) Where(ps ...predicate.GoodPaying) *GoodPayingQuery {
	gpq.predicates = append(gpq.predicates, ps...)
	return gpq
}

// Limit adds a limit step to the query.
func (gpq *GoodPayingQuery) Limit(limit int) *GoodPayingQuery {
	gpq.limit = &limit
	return gpq
}

// Offset adds an offset step to the query.
func (gpq *GoodPayingQuery) Offset(offset int) *GoodPayingQuery {
	gpq.offset = &offset
	return gpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gpq *GoodPayingQuery) Unique(unique bool) *GoodPayingQuery {
	gpq.unique = &unique
	return gpq
}

// Order adds an order step to the query.
func (gpq *GoodPayingQuery) Order(o ...OrderFunc) *GoodPayingQuery {
	gpq.order = append(gpq.order, o...)
	return gpq
}

// First returns the first GoodPaying entity from the query.
// Returns a *NotFoundError when no GoodPaying was found.
func (gpq *GoodPayingQuery) First(ctx context.Context) (*GoodPaying, error) {
	nodes, err := gpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{goodpaying.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gpq *GoodPayingQuery) FirstX(ctx context.Context) *GoodPaying {
	node, err := gpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GoodPaying ID from the query.
// Returns a *NotFoundError when no GoodPaying ID was found.
func (gpq *GoodPayingQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{goodpaying.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gpq *GoodPayingQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := gpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GoodPaying entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GoodPaying entity is found.
// Returns a *NotFoundError when no GoodPaying entities are found.
func (gpq *GoodPayingQuery) Only(ctx context.Context) (*GoodPaying, error) {
	nodes, err := gpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{goodpaying.Label}
	default:
		return nil, &NotSingularError{goodpaying.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gpq *GoodPayingQuery) OnlyX(ctx context.Context) *GoodPaying {
	node, err := gpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GoodPaying ID in the query.
// Returns a *NotSingularError when more than one GoodPaying ID is found.
// Returns a *NotFoundError when no entities are found.
func (gpq *GoodPayingQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{goodpaying.Label}
	default:
		err = &NotSingularError{goodpaying.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gpq *GoodPayingQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := gpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoodPayings.
func (gpq *GoodPayingQuery) All(ctx context.Context) ([]*GoodPaying, error) {
	if err := gpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gpq *GoodPayingQuery) AllX(ctx context.Context) []*GoodPaying {
	nodes, err := gpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GoodPaying IDs.
func (gpq *GoodPayingQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := gpq.Select(goodpaying.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gpq *GoodPayingQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := gpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gpq *GoodPayingQuery) Count(ctx context.Context) (int, error) {
	if err := gpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gpq *GoodPayingQuery) CountX(ctx context.Context) int {
	count, err := gpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gpq *GoodPayingQuery) Exist(ctx context.Context) (bool, error) {
	if err := gpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gpq *GoodPayingQuery) ExistX(ctx context.Context) bool {
	exist, err := gpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GoodPayingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gpq *GoodPayingQuery) Clone() *GoodPayingQuery {
	if gpq == nil {
		return nil
	}
	return &GoodPayingQuery{
		config:     gpq.config,
		limit:      gpq.limit,
		offset:     gpq.offset,
		order:      append([]OrderFunc{}, gpq.order...),
		predicates: append([]predicate.GoodPaying{}, gpq.predicates...),
		// clone intermediate query.
		sql:    gpq.sql.Clone(),
		path:   gpq.path,
		unique: gpq.unique,
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
//	client.GoodPaying.Query().
//		GroupBy(goodpaying.FieldOrderID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gpq *GoodPayingQuery) GroupBy(field string, fields ...string) *GoodPayingGroupBy {
	grbuild := &GoodPayingGroupBy{config: gpq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gpq.sqlQuery(ctx), nil
	}
	grbuild.label = goodpaying.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
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
//	client.GoodPaying.Query().
//		Select(goodpaying.FieldOrderID).
//		Scan(ctx, &v)
//
func (gpq *GoodPayingQuery) Select(fields ...string) *GoodPayingSelect {
	gpq.fields = append(gpq.fields, fields...)
	selbuild := &GoodPayingSelect{GoodPayingQuery: gpq}
	selbuild.label = goodpaying.Label
	selbuild.flds, selbuild.scan = &gpq.fields, selbuild.Scan
	return selbuild
}

func (gpq *GoodPayingQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gpq.fields {
		if !goodpaying.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gpq.path != nil {
		prev, err := gpq.path(ctx)
		if err != nil {
			return err
		}
		gpq.sql = prev
	}
	return nil
}

func (gpq *GoodPayingQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GoodPaying, error) {
	var (
		nodes = []*GoodPaying{}
		_spec = gpq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*GoodPaying).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &GoodPaying{config: gpq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(gpq.modifiers) > 0 {
		_spec.Modifiers = gpq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gpq *GoodPayingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gpq.querySpec()
	if len(gpq.modifiers) > 0 {
		_spec.Modifiers = gpq.modifiers
	}
	_spec.Node.Columns = gpq.fields
	if len(gpq.fields) > 0 {
		_spec.Unique = gpq.unique != nil && *gpq.unique
	}
	return sqlgraph.CountNodes(ctx, gpq.driver, _spec)
}

func (gpq *GoodPayingQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (gpq *GoodPayingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodpaying.Table,
			Columns: goodpaying.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodpaying.FieldID,
			},
		},
		From:   gpq.sql,
		Unique: true,
	}
	if unique := gpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodpaying.FieldID)
		for i := range fields {
			if fields[i] != goodpaying.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gpq *GoodPayingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gpq.driver.Dialect())
	t1 := builder.Table(goodpaying.Table)
	columns := gpq.fields
	if len(columns) == 0 {
		columns = goodpaying.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gpq.sql != nil {
		selector = gpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gpq.unique != nil && *gpq.unique {
		selector.Distinct()
	}
	for _, m := range gpq.modifiers {
		m(selector)
	}
	for _, p := range gpq.predicates {
		p(selector)
	}
	for _, p := range gpq.order {
		p(selector)
	}
	if offset := gpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (gpq *GoodPayingQuery) ForUpdate(opts ...sql.LockOption) *GoodPayingQuery {
	if gpq.driver.Dialect() == dialect.Postgres {
		gpq.Unique(false)
	}
	gpq.modifiers = append(gpq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return gpq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (gpq *GoodPayingQuery) ForShare(opts ...sql.LockOption) *GoodPayingQuery {
	if gpq.driver.Dialect() == dialect.Postgres {
		gpq.Unique(false)
	}
	gpq.modifiers = append(gpq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return gpq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gpq *GoodPayingQuery) Modify(modifiers ...func(s *sql.Selector)) *GoodPayingSelect {
	gpq.modifiers = append(gpq.modifiers, modifiers...)
	return gpq.Select()
}

// GoodPayingGroupBy is the group-by builder for GoodPaying entities.
type GoodPayingGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gpgb *GoodPayingGroupBy) Aggregate(fns ...AggregateFunc) *GoodPayingGroupBy {
	gpgb.fns = append(gpgb.fns, fns...)
	return gpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (gpgb *GoodPayingGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gpgb.path(ctx)
	if err != nil {
		return err
	}
	gpgb.sql = query
	return gpgb.sqlScan(ctx, v)
}

func (gpgb *GoodPayingGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gpgb.fields {
		if !goodpaying.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gpgb *GoodPayingGroupBy) sqlQuery() *sql.Selector {
	selector := gpgb.sql.Select()
	aggregation := make([]string, 0, len(gpgb.fns))
	for _, fn := range gpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(gpgb.fields)+len(gpgb.fns))
		for _, f := range gpgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(gpgb.fields...)...)
}

// GoodPayingSelect is the builder for selecting fields of GoodPaying entities.
type GoodPayingSelect struct {
	*GoodPayingQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gps *GoodPayingSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gps.prepareQuery(ctx); err != nil {
		return err
	}
	gps.sql = gps.GoodPayingQuery.sqlQuery(ctx)
	return gps.sqlScan(ctx, v)
}

func (gps *GoodPayingSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gps.sql.Query()
	if err := gps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gps *GoodPayingSelect) Modify(modifiers ...func(s *sql.Selector)) *GoodPayingSelect {
	gps.modifiers = append(gps.modifiers, modifiers...)
	return gps
}
