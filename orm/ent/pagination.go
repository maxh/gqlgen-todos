// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/maxh/gqlgen-todos/orm/ent/noderevision"
	"github.com/maxh/gqlgen-todos/orm/ent/organization"
	"github.com/maxh/gqlgen-todos/orm/ent/tenant"
	"github.com/maxh/gqlgen-todos/orm/ent/todo"
	"github.com/maxh/gqlgen-todos/orm/ent/user"
	"github.com/maxh/gqlgen-todos/qid"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vmihailenco/msgpack/v5"
)

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

// MarshalGQL implements graphql.Marshaler interface.
func (o OrderDirection) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(o.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (o *OrderDirection) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("order direction %T must be a string", val)
	}
	*o = OrderDirection(str)
	return o.Validate()
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

func (o OrderDirection) orderFunc(field string) OrderFunc {
	if o == OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

func cursorsToPredicates(direction OrderDirection, after, before *Cursor, field, idField string) []func(s *sql.Selector) {
	var predicates []func(s *sql.Selector)
	if after != nil {
		if after.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeGT
			} else {
				predicate = sql.CompositeLT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					after.Value, after.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.GT
			} else {
				predicate = sql.LT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					after.ID,
				))
			})
		}
	}
	if before != nil {
		if before.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeLT
			} else {
				predicate = sql.CompositeGT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					before.Value, before.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.LT
			} else {
				predicate = sql.GT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					before.ID,
				))
			})
		}
	}
	return predicates
}

// PageInfo of a connection type.
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *Cursor `json:"startCursor"`
	EndCursor       *Cursor `json:"endCursor"`
}

// Cursor of an edge type.
type Cursor struct {
	ID    qid.ID `msgpack:"i"`
	Value Value  `msgpack:"v,omitempty"`
}

// MarshalGQL implements graphql.Marshaler interface.
func (c Cursor) MarshalGQL(w io.Writer) {
	quote := []byte{'"'}
	w.Write(quote)
	defer w.Write(quote)
	wc := base64.NewEncoder(base64.RawStdEncoding, w)
	defer wc.Close()
	_ = msgpack.NewEncoder(wc).Encode(c)
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (c *Cursor) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a string", v)
	}
	if err := msgpack.NewDecoder(
		base64.NewDecoder(
			base64.RawStdEncoding,
			strings.NewReader(s),
		),
	).Decode(c); err != nil {
		return fmt.Errorf("cannot decode cursor: %w", err)
	}
	return nil
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func getCollectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	field := fc.Field

walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Name == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return getCollectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

// NodeRevisionEdge is the edge representation of NodeRevision.
type NodeRevisionEdge struct {
	Node   *NodeRevision `json:"node"`
	Cursor Cursor        `json:"cursor"`
}

// NodeRevisionConnection is the connection containing edges to NodeRevision.
type NodeRevisionConnection struct {
	Edges      []*NodeRevisionEdge `json:"edges"`
	PageInfo   PageInfo            `json:"pageInfo"`
	TotalCount int                 `json:"totalCount"`
}

// NodeRevisionPaginateOption enables pagination customization.
type NodeRevisionPaginateOption func(*nodeRevisionPager) error

// WithNodeRevisionOrder configures pagination ordering.
func WithNodeRevisionOrder(order *NodeRevisionOrder) NodeRevisionPaginateOption {
	if order == nil {
		order = DefaultNodeRevisionOrder
	}
	o := *order
	return func(pager *nodeRevisionPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultNodeRevisionOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithNodeRevisionFilter configures pagination filter.
func WithNodeRevisionFilter(filter func(*NodeRevisionQuery) (*NodeRevisionQuery, error)) NodeRevisionPaginateOption {
	return func(pager *nodeRevisionPager) error {
		if filter == nil {
			return errors.New("NodeRevisionQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type nodeRevisionPager struct {
	order  *NodeRevisionOrder
	filter func(*NodeRevisionQuery) (*NodeRevisionQuery, error)
}

func newNodeRevisionPager(opts []NodeRevisionPaginateOption) (*nodeRevisionPager, error) {
	pager := &nodeRevisionPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultNodeRevisionOrder
	}
	return pager, nil
}

func (p *nodeRevisionPager) applyFilter(query *NodeRevisionQuery) (*NodeRevisionQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *nodeRevisionPager) toCursor(nr *NodeRevision) Cursor {
	return p.order.Field.toCursor(nr)
}

func (p *nodeRevisionPager) applyCursors(query *NodeRevisionQuery, after, before *Cursor) *NodeRevisionQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultNodeRevisionOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *nodeRevisionPager) applyOrder(query *NodeRevisionQuery, reverse bool) *NodeRevisionQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultNodeRevisionOrder.Field {
		query = query.Order(direction.orderFunc(DefaultNodeRevisionOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to NodeRevision.
func (nr *NodeRevisionQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...NodeRevisionPaginateOption,
) (*NodeRevisionConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newNodeRevisionPager(opts)
	if err != nil {
		return nil, err
	}

	if nr, err = pager.applyFilter(nr); err != nil {
		return nil, err
	}

	conn := &NodeRevisionConnection{Edges: []*NodeRevisionEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := nr.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := nr.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	nr = pager.applyCursors(nr, after, before)
	nr = pager.applyOrder(nr, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		nr = nr.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		nr = nr.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := nr.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *NodeRevision
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *NodeRevision {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *NodeRevision {
			return nodes[i]
		}
	}

	conn.Edges = make([]*NodeRevisionEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &NodeRevisionEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

// NodeRevisionOrderField defines the ordering field of NodeRevision.
type NodeRevisionOrderField struct {
	field    string
	toCursor func(*NodeRevision) Cursor
}

// NodeRevisionOrder defines the ordering of NodeRevision.
type NodeRevisionOrder struct {
	Direction OrderDirection          `json:"direction"`
	Field     *NodeRevisionOrderField `json:"field"`
}

// DefaultNodeRevisionOrder is the default ordering of NodeRevision.
var DefaultNodeRevisionOrder = &NodeRevisionOrder{
	Direction: OrderDirectionAsc,
	Field: &NodeRevisionOrderField{
		field: noderevision.FieldID,
		toCursor: func(nr *NodeRevision) Cursor {
			return Cursor{ID: nr.ID}
		},
	},
}

// ToEdge converts NodeRevision into NodeRevisionEdge.
func (nr *NodeRevision) ToEdge(order *NodeRevisionOrder) *NodeRevisionEdge {
	if order == nil {
		order = DefaultNodeRevisionOrder
	}
	return &NodeRevisionEdge{
		Node:   nr,
		Cursor: order.Field.toCursor(nr),
	}
}

// OrganizationEdge is the edge representation of Organization.
type OrganizationEdge struct {
	Node   *Organization `json:"node"`
	Cursor Cursor        `json:"cursor"`
}

// OrganizationConnection is the connection containing edges to Organization.
type OrganizationConnection struct {
	Edges      []*OrganizationEdge `json:"edges"`
	PageInfo   PageInfo            `json:"pageInfo"`
	TotalCount int                 `json:"totalCount"`
}

// OrganizationPaginateOption enables pagination customization.
type OrganizationPaginateOption func(*organizationPager) error

// WithOrganizationOrder configures pagination ordering.
func WithOrganizationOrder(order *OrganizationOrder) OrganizationPaginateOption {
	if order == nil {
		order = DefaultOrganizationOrder
	}
	o := *order
	return func(pager *organizationPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultOrganizationOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithOrganizationFilter configures pagination filter.
func WithOrganizationFilter(filter func(*OrganizationQuery) (*OrganizationQuery, error)) OrganizationPaginateOption {
	return func(pager *organizationPager) error {
		if filter == nil {
			return errors.New("OrganizationQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type organizationPager struct {
	order  *OrganizationOrder
	filter func(*OrganizationQuery) (*OrganizationQuery, error)
}

func newOrganizationPager(opts []OrganizationPaginateOption) (*organizationPager, error) {
	pager := &organizationPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultOrganizationOrder
	}
	return pager, nil
}

func (p *organizationPager) applyFilter(query *OrganizationQuery) (*OrganizationQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *organizationPager) toCursor(o *Organization) Cursor {
	return p.order.Field.toCursor(o)
}

func (p *organizationPager) applyCursors(query *OrganizationQuery, after, before *Cursor) *OrganizationQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultOrganizationOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *organizationPager) applyOrder(query *OrganizationQuery, reverse bool) *OrganizationQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultOrganizationOrder.Field {
		query = query.Order(direction.orderFunc(DefaultOrganizationOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to Organization.
func (o *OrganizationQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...OrganizationPaginateOption,
) (*OrganizationConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newOrganizationPager(opts)
	if err != nil {
		return nil, err
	}

	if o, err = pager.applyFilter(o); err != nil {
		return nil, err
	}

	conn := &OrganizationConnection{Edges: []*OrganizationEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := o.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := o.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	o = pager.applyCursors(o, after, before)
	o = pager.applyOrder(o, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		o = o.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		o = o.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := o.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *Organization
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Organization {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Organization {
			return nodes[i]
		}
	}

	conn.Edges = make([]*OrganizationEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &OrganizationEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

// OrganizationOrderField defines the ordering field of Organization.
type OrganizationOrderField struct {
	field    string
	toCursor func(*Organization) Cursor
}

// OrganizationOrder defines the ordering of Organization.
type OrganizationOrder struct {
	Direction OrderDirection          `json:"direction"`
	Field     *OrganizationOrderField `json:"field"`
}

// DefaultOrganizationOrder is the default ordering of Organization.
var DefaultOrganizationOrder = &OrganizationOrder{
	Direction: OrderDirectionAsc,
	Field: &OrganizationOrderField{
		field: organization.FieldID,
		toCursor: func(o *Organization) Cursor {
			return Cursor{ID: o.ID}
		},
	},
}

// ToEdge converts Organization into OrganizationEdge.
func (o *Organization) ToEdge(order *OrganizationOrder) *OrganizationEdge {
	if order == nil {
		order = DefaultOrganizationOrder
	}
	return &OrganizationEdge{
		Node:   o,
		Cursor: order.Field.toCursor(o),
	}
}

// TenantEdge is the edge representation of Tenant.
type TenantEdge struct {
	Node   *Tenant `json:"node"`
	Cursor Cursor  `json:"cursor"`
}

// TenantConnection is the connection containing edges to Tenant.
type TenantConnection struct {
	Edges      []*TenantEdge `json:"edges"`
	PageInfo   PageInfo      `json:"pageInfo"`
	TotalCount int           `json:"totalCount"`
}

// TenantPaginateOption enables pagination customization.
type TenantPaginateOption func(*tenantPager) error

// WithTenantOrder configures pagination ordering.
func WithTenantOrder(order *TenantOrder) TenantPaginateOption {
	if order == nil {
		order = DefaultTenantOrder
	}
	o := *order
	return func(pager *tenantPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultTenantOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithTenantFilter configures pagination filter.
func WithTenantFilter(filter func(*TenantQuery) (*TenantQuery, error)) TenantPaginateOption {
	return func(pager *tenantPager) error {
		if filter == nil {
			return errors.New("TenantQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type tenantPager struct {
	order  *TenantOrder
	filter func(*TenantQuery) (*TenantQuery, error)
}

func newTenantPager(opts []TenantPaginateOption) (*tenantPager, error) {
	pager := &tenantPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultTenantOrder
	}
	return pager, nil
}

func (p *tenantPager) applyFilter(query *TenantQuery) (*TenantQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *tenantPager) toCursor(t *Tenant) Cursor {
	return p.order.Field.toCursor(t)
}

func (p *tenantPager) applyCursors(query *TenantQuery, after, before *Cursor) *TenantQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultTenantOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *tenantPager) applyOrder(query *TenantQuery, reverse bool) *TenantQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultTenantOrder.Field {
		query = query.Order(direction.orderFunc(DefaultTenantOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to Tenant.
func (t *TenantQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...TenantPaginateOption,
) (*TenantConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newTenantPager(opts)
	if err != nil {
		return nil, err
	}

	if t, err = pager.applyFilter(t); err != nil {
		return nil, err
	}

	conn := &TenantConnection{Edges: []*TenantEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := t.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := t.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	t = pager.applyCursors(t, after, before)
	t = pager.applyOrder(t, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		t = t.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		t = t.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := t.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *Tenant
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Tenant {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Tenant {
			return nodes[i]
		}
	}

	conn.Edges = make([]*TenantEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &TenantEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

// TenantOrderField defines the ordering field of Tenant.
type TenantOrderField struct {
	field    string
	toCursor func(*Tenant) Cursor
}

// TenantOrder defines the ordering of Tenant.
type TenantOrder struct {
	Direction OrderDirection    `json:"direction"`
	Field     *TenantOrderField `json:"field"`
}

// DefaultTenantOrder is the default ordering of Tenant.
var DefaultTenantOrder = &TenantOrder{
	Direction: OrderDirectionAsc,
	Field: &TenantOrderField{
		field: tenant.FieldID,
		toCursor: func(t *Tenant) Cursor {
			return Cursor{ID: t.ID}
		},
	},
}

// ToEdge converts Tenant into TenantEdge.
func (t *Tenant) ToEdge(order *TenantOrder) *TenantEdge {
	if order == nil {
		order = DefaultTenantOrder
	}
	return &TenantEdge{
		Node:   t,
		Cursor: order.Field.toCursor(t),
	}
}

// TodoEdge is the edge representation of Todo.
type TodoEdge struct {
	Node   *Todo  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// TodoConnection is the connection containing edges to Todo.
type TodoConnection struct {
	Edges      []*TodoEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

// TodoPaginateOption enables pagination customization.
type TodoPaginateOption func(*todoPager) error

// WithTodoOrder configures pagination ordering.
func WithTodoOrder(order *TodoOrder) TodoPaginateOption {
	if order == nil {
		order = DefaultTodoOrder
	}
	o := *order
	return func(pager *todoPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultTodoOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithTodoFilter configures pagination filter.
func WithTodoFilter(filter func(*TodoQuery) (*TodoQuery, error)) TodoPaginateOption {
	return func(pager *todoPager) error {
		if filter == nil {
			return errors.New("TodoQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type todoPager struct {
	order  *TodoOrder
	filter func(*TodoQuery) (*TodoQuery, error)
}

func newTodoPager(opts []TodoPaginateOption) (*todoPager, error) {
	pager := &todoPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultTodoOrder
	}
	return pager, nil
}

func (p *todoPager) applyFilter(query *TodoQuery) (*TodoQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *todoPager) toCursor(t *Todo) Cursor {
	return p.order.Field.toCursor(t)
}

func (p *todoPager) applyCursors(query *TodoQuery, after, before *Cursor) *TodoQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultTodoOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *todoPager) applyOrder(query *TodoQuery, reverse bool) *TodoQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultTodoOrder.Field {
		query = query.Order(direction.orderFunc(DefaultTodoOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to Todo.
func (t *TodoQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...TodoPaginateOption,
) (*TodoConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newTodoPager(opts)
	if err != nil {
		return nil, err
	}

	if t, err = pager.applyFilter(t); err != nil {
		return nil, err
	}

	conn := &TodoConnection{Edges: []*TodoEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := t.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := t.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	t = pager.applyCursors(t, after, before)
	t = pager.applyOrder(t, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		t = t.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		t = t.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := t.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *Todo
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Todo {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Todo {
			return nodes[i]
		}
	}

	conn.Edges = make([]*TodoEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &TodoEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

// TodoOrderField defines the ordering field of Todo.
type TodoOrderField struct {
	field    string
	toCursor func(*Todo) Cursor
}

// TodoOrder defines the ordering of Todo.
type TodoOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *TodoOrderField `json:"field"`
}

// DefaultTodoOrder is the default ordering of Todo.
var DefaultTodoOrder = &TodoOrder{
	Direction: OrderDirectionAsc,
	Field: &TodoOrderField{
		field: todo.FieldID,
		toCursor: func(t *Todo) Cursor {
			return Cursor{ID: t.ID}
		},
	},
}

// ToEdge converts Todo into TodoEdge.
func (t *Todo) ToEdge(order *TodoOrder) *TodoEdge {
	if order == nil {
		order = DefaultTodoOrder
	}
	return &TodoEdge{
		Node:   t,
		Cursor: order.Field.toCursor(t),
	}
}

// UserEdge is the edge representation of User.
type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// UserConnection is the connection containing edges to User.
type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*userPager) error

// WithUserOrder configures pagination ordering.
func WithUserOrder(order *UserOrder) UserPaginateOption {
	if order == nil {
		order = DefaultUserOrder
	}
	o := *order
	return func(pager *userPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultUserOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithUserFilter configures pagination filter.
func WithUserFilter(filter func(*UserQuery) (*UserQuery, error)) UserPaginateOption {
	return func(pager *userPager) error {
		if filter == nil {
			return errors.New("UserQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type userPager struct {
	order  *UserOrder
	filter func(*UserQuery) (*UserQuery, error)
}

func newUserPager(opts []UserPaginateOption) (*userPager, error) {
	pager := &userPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserOrder
	}
	return pager, nil
}

func (p *userPager) applyFilter(query *UserQuery) (*UserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *userPager) toCursor(u *User) Cursor {
	return p.order.Field.toCursor(u)
}

func (p *userPager) applyCursors(query *UserQuery, after, before *Cursor) *UserQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultUserOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *userPager) applyOrder(query *UserQuery, reverse bool) *UserQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultUserOrder.Field {
		query = query.Order(direction.orderFunc(DefaultUserOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to User.
func (u *UserQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserPaginateOption,
) (*UserConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserPager(opts)
	if err != nil {
		return nil, err
	}

	if u, err = pager.applyFilter(u); err != nil {
		return nil, err
	}

	conn := &UserConnection{Edges: []*UserEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := u.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := u.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	u = pager.applyCursors(u, after, before)
	u = pager.applyOrder(u, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		u = u.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		u = u.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := u.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *User
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *User {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *User {
			return nodes[i]
		}
	}

	conn.Edges = make([]*UserEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &UserEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

// UserOrderField defines the ordering field of User.
type UserOrderField struct {
	field    string
	toCursor func(*User) Cursor
}

// UserOrder defines the ordering of User.
type UserOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *UserOrderField `json:"field"`
}

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = &UserOrder{
	Direction: OrderDirectionAsc,
	Field: &UserOrderField{
		field: user.FieldID,
		toCursor: func(u *User) Cursor {
			return Cursor{ID: u.ID}
		},
	},
}

// ToEdge converts User into UserEdge.
func (u *User) ToEdge(order *UserOrder) *UserEdge {
	if order == nil {
		order = DefaultUserOrder
	}
	return &UserEdge{
		Node:   u,
		Cursor: order.Field.toCursor(u),
	}
}
