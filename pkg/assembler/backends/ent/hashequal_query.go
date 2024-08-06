// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/artifact"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/hashequal"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
)

// HashEqualQuery is the builder for querying HashEqual entities.
type HashEqualQuery struct {
	config
	ctx           *QueryContext
	order         []hashequal.OrderOption
	inters        []Interceptor
	predicates    []predicate.HashEqual
	withArtifactA *ArtifactQuery
	withArtifactB *ArtifactQuery
	modifiers     []func(*sql.Selector)
	loadTotal     []func(context.Context, []*HashEqual) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HashEqualQuery builder.
func (heq *HashEqualQuery) Where(ps ...predicate.HashEqual) *HashEqualQuery {
	heq.predicates = append(heq.predicates, ps...)
	return heq
}

// Limit the number of records to be returned by this query.
func (heq *HashEqualQuery) Limit(limit int) *HashEqualQuery {
	heq.ctx.Limit = &limit
	return heq
}

// Offset to start from.
func (heq *HashEqualQuery) Offset(offset int) *HashEqualQuery {
	heq.ctx.Offset = &offset
	return heq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (heq *HashEqualQuery) Unique(unique bool) *HashEqualQuery {
	heq.ctx.Unique = &unique
	return heq
}

// Order specifies how the records should be ordered.
func (heq *HashEqualQuery) Order(o ...hashequal.OrderOption) *HashEqualQuery {
	heq.order = append(heq.order, o...)
	return heq
}

// QueryArtifactA chains the current query on the "artifact_a" edge.
func (heq *HashEqualQuery) QueryArtifactA() *ArtifactQuery {
	query := (&ArtifactClient{config: heq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := heq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := heq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(hashequal.Table, hashequal.FieldID, selector),
			sqlgraph.To(artifact.Table, artifact.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, hashequal.ArtifactATable, hashequal.ArtifactAColumn),
		)
		fromU = sqlgraph.SetNeighbors(heq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryArtifactB chains the current query on the "artifact_b" edge.
func (heq *HashEqualQuery) QueryArtifactB() *ArtifactQuery {
	query := (&ArtifactClient{config: heq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := heq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := heq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(hashequal.Table, hashequal.FieldID, selector),
			sqlgraph.To(artifact.Table, artifact.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, hashequal.ArtifactBTable, hashequal.ArtifactBColumn),
		)
		fromU = sqlgraph.SetNeighbors(heq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first HashEqual entity from the query.
// Returns a *NotFoundError when no HashEqual was found.
func (heq *HashEqualQuery) First(ctx context.Context) (*HashEqual, error) {
	nodes, err := heq.Limit(1).All(setContextOp(ctx, heq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{hashequal.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (heq *HashEqualQuery) FirstX(ctx context.Context) *HashEqual {
	node, err := heq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first HashEqual ID from the query.
// Returns a *NotFoundError when no HashEqual ID was found.
func (heq *HashEqualQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = heq.Limit(1).IDs(setContextOp(ctx, heq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{hashequal.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (heq *HashEqualQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := heq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single HashEqual entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one HashEqual entity is found.
// Returns a *NotFoundError when no HashEqual entities are found.
func (heq *HashEqualQuery) Only(ctx context.Context) (*HashEqual, error) {
	nodes, err := heq.Limit(2).All(setContextOp(ctx, heq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{hashequal.Label}
	default:
		return nil, &NotSingularError{hashequal.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (heq *HashEqualQuery) OnlyX(ctx context.Context) *HashEqual {
	node, err := heq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only HashEqual ID in the query.
// Returns a *NotSingularError when more than one HashEqual ID is found.
// Returns a *NotFoundError when no entities are found.
func (heq *HashEqualQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = heq.Limit(2).IDs(setContextOp(ctx, heq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{hashequal.Label}
	default:
		err = &NotSingularError{hashequal.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (heq *HashEqualQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := heq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of HashEquals.
func (heq *HashEqualQuery) All(ctx context.Context) ([]*HashEqual, error) {
	ctx = setContextOp(ctx, heq.ctx, ent.OpQueryAll)
	if err := heq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*HashEqual, *HashEqualQuery]()
	return withInterceptors[[]*HashEqual](ctx, heq, qr, heq.inters)
}

// AllX is like All, but panics if an error occurs.
func (heq *HashEqualQuery) AllX(ctx context.Context) []*HashEqual {
	nodes, err := heq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of HashEqual IDs.
func (heq *HashEqualQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if heq.ctx.Unique == nil && heq.path != nil {
		heq.Unique(true)
	}
	ctx = setContextOp(ctx, heq.ctx, ent.OpQueryIDs)
	if err = heq.Select(hashequal.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (heq *HashEqualQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := heq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (heq *HashEqualQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, heq.ctx, ent.OpQueryCount)
	if err := heq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, heq, querierCount[*HashEqualQuery](), heq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (heq *HashEqualQuery) CountX(ctx context.Context) int {
	count, err := heq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (heq *HashEqualQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, heq.ctx, ent.OpQueryExist)
	switch _, err := heq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (heq *HashEqualQuery) ExistX(ctx context.Context) bool {
	exist, err := heq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HashEqualQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (heq *HashEqualQuery) Clone() *HashEqualQuery {
	if heq == nil {
		return nil
	}
	return &HashEqualQuery{
		config:        heq.config,
		ctx:           heq.ctx.Clone(),
		order:         append([]hashequal.OrderOption{}, heq.order...),
		inters:        append([]Interceptor{}, heq.inters...),
		predicates:    append([]predicate.HashEqual{}, heq.predicates...),
		withArtifactA: heq.withArtifactA.Clone(),
		withArtifactB: heq.withArtifactB.Clone(),
		// clone intermediate query.
		sql:  heq.sql.Clone(),
		path: heq.path,
	}
}

// WithArtifactA tells the query-builder to eager-load the nodes that are connected to
// the "artifact_a" edge. The optional arguments are used to configure the query builder of the edge.
func (heq *HashEqualQuery) WithArtifactA(opts ...func(*ArtifactQuery)) *HashEqualQuery {
	query := (&ArtifactClient{config: heq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	heq.withArtifactA = query
	return heq
}

// WithArtifactB tells the query-builder to eager-load the nodes that are connected to
// the "artifact_b" edge. The optional arguments are used to configure the query builder of the edge.
func (heq *HashEqualQuery) WithArtifactB(opts ...func(*ArtifactQuery)) *HashEqualQuery {
	query := (&ArtifactClient{config: heq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	heq.withArtifactB = query
	return heq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ArtID uuid.UUID `json:"art_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.HashEqual.Query().
//		GroupBy(hashequal.FieldArtID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (heq *HashEqualQuery) GroupBy(field string, fields ...string) *HashEqualGroupBy {
	heq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &HashEqualGroupBy{build: heq}
	grbuild.flds = &heq.ctx.Fields
	grbuild.label = hashequal.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ArtID uuid.UUID `json:"art_id,omitempty"`
//	}
//
//	client.HashEqual.Query().
//		Select(hashequal.FieldArtID).
//		Scan(ctx, &v)
func (heq *HashEqualQuery) Select(fields ...string) *HashEqualSelect {
	heq.ctx.Fields = append(heq.ctx.Fields, fields...)
	sbuild := &HashEqualSelect{HashEqualQuery: heq}
	sbuild.label = hashequal.Label
	sbuild.flds, sbuild.scan = &heq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a HashEqualSelect configured with the given aggregations.
func (heq *HashEqualQuery) Aggregate(fns ...AggregateFunc) *HashEqualSelect {
	return heq.Select().Aggregate(fns...)
}

func (heq *HashEqualQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range heq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, heq); err != nil {
				return err
			}
		}
	}
	for _, f := range heq.ctx.Fields {
		if !hashequal.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if heq.path != nil {
		prev, err := heq.path(ctx)
		if err != nil {
			return err
		}
		heq.sql = prev
	}
	return nil
}

func (heq *HashEqualQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*HashEqual, error) {
	var (
		nodes       = []*HashEqual{}
		_spec       = heq.querySpec()
		loadedTypes = [2]bool{
			heq.withArtifactA != nil,
			heq.withArtifactB != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*HashEqual).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &HashEqual{config: heq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(heq.modifiers) > 0 {
		_spec.Modifiers = heq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, heq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := heq.withArtifactA; query != nil {
		if err := heq.loadArtifactA(ctx, query, nodes, nil,
			func(n *HashEqual, e *Artifact) { n.Edges.ArtifactA = e }); err != nil {
			return nil, err
		}
	}
	if query := heq.withArtifactB; query != nil {
		if err := heq.loadArtifactB(ctx, query, nodes, nil,
			func(n *HashEqual, e *Artifact) { n.Edges.ArtifactB = e }); err != nil {
			return nil, err
		}
	}
	for i := range heq.loadTotal {
		if err := heq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (heq *HashEqualQuery) loadArtifactA(ctx context.Context, query *ArtifactQuery, nodes []*HashEqual, init func(*HashEqual), assign func(*HashEqual, *Artifact)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*HashEqual)
	for i := range nodes {
		fk := nodes[i].ArtID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(artifact.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "art_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (heq *HashEqualQuery) loadArtifactB(ctx context.Context, query *ArtifactQuery, nodes []*HashEqual, init func(*HashEqual), assign func(*HashEqual, *Artifact)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*HashEqual)
	for i := range nodes {
		fk := nodes[i].EqualArtID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(artifact.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "equal_art_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (heq *HashEqualQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := heq.querySpec()
	if len(heq.modifiers) > 0 {
		_spec.Modifiers = heq.modifiers
	}
	_spec.Node.Columns = heq.ctx.Fields
	if len(heq.ctx.Fields) > 0 {
		_spec.Unique = heq.ctx.Unique != nil && *heq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, heq.driver, _spec)
}

func (heq *HashEqualQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(hashequal.Table, hashequal.Columns, sqlgraph.NewFieldSpec(hashequal.FieldID, field.TypeUUID))
	_spec.From = heq.sql
	if unique := heq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if heq.path != nil {
		_spec.Unique = true
	}
	if fields := heq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hashequal.FieldID)
		for i := range fields {
			if fields[i] != hashequal.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if heq.withArtifactA != nil {
			_spec.Node.AddColumnOnce(hashequal.FieldArtID)
		}
		if heq.withArtifactB != nil {
			_spec.Node.AddColumnOnce(hashequal.FieldEqualArtID)
		}
	}
	if ps := heq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := heq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := heq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := heq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (heq *HashEqualQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(heq.driver.Dialect())
	t1 := builder.Table(hashequal.Table)
	columns := heq.ctx.Fields
	if len(columns) == 0 {
		columns = hashequal.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if heq.sql != nil {
		selector = heq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if heq.ctx.Unique != nil && *heq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range heq.predicates {
		p(selector)
	}
	for _, p := range heq.order {
		p(selector)
	}
	if offset := heq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := heq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// HashEqualGroupBy is the group-by builder for HashEqual entities.
type HashEqualGroupBy struct {
	selector
	build *HashEqualQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hegb *HashEqualGroupBy) Aggregate(fns ...AggregateFunc) *HashEqualGroupBy {
	hegb.fns = append(hegb.fns, fns...)
	return hegb
}

// Scan applies the selector query and scans the result into the given value.
func (hegb *HashEqualGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hegb.build.ctx, ent.OpQueryGroupBy)
	if err := hegb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HashEqualQuery, *HashEqualGroupBy](ctx, hegb.build, hegb, hegb.build.inters, v)
}

func (hegb *HashEqualGroupBy) sqlScan(ctx context.Context, root *HashEqualQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(hegb.fns))
	for _, fn := range hegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*hegb.flds)+len(hegb.fns))
		for _, f := range *hegb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*hegb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hegb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// HashEqualSelect is the builder for selecting fields of HashEqual entities.
type HashEqualSelect struct {
	*HashEqualQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (hes *HashEqualSelect) Aggregate(fns ...AggregateFunc) *HashEqualSelect {
	hes.fns = append(hes.fns, fns...)
	return hes
}

// Scan applies the selector query and scans the result into the given value.
func (hes *HashEqualSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hes.ctx, ent.OpQuerySelect)
	if err := hes.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HashEqualQuery, *HashEqualSelect](ctx, hes.HashEqualQuery, hes, hes.inters, v)
}

func (hes *HashEqualSelect) sqlScan(ctx context.Context, root *HashEqualQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(hes.fns))
	for _, fn := range hes.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*hes.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
