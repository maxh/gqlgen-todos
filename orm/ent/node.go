// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
	"github.com/maxh/gqlgen-todos/orm/ent/organization"
	"github.com/maxh/gqlgen-todos/orm/ent/tenant"
	"github.com/maxh/gqlgen-todos/orm/ent/todo"
	"github.com/maxh/gqlgen-todos/orm/ent/user"
	"github.com/maxh/gqlgen-todos/qid"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     qid.ID   `json:"id,omitempty"`     // node id.
	Type   string   `json:"type,omitempty"`   // node type.
	Fields []*Field `json:"fields,omitempty"` // node fields.
	Edges  []*Edge  `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string   `json:"type,omitempty"` // edge type.
	Name string   `json:"name,omitempty"` // edge name.
	IDs  []qid.ID `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (o *Organization) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     o.ID,
		Type:   "Organization",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(o.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(o.CreatedBy); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "qid.ID",
		Name:  "created_by",
		Value: string(buf),
	}
	if buf, err = json.Marshal(o.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(o.UpdatedBy); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "qid.ID",
		Name:  "updated_by",
		Value: string(buf),
	}
	if buf, err = json.Marshal(o.Name); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Tenant",
		Name: "tenant",
	}
	err = o.QueryTenant().
		Select(tenant.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "User",
		Name: "users",
	}
	err = o.QueryUsers().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (t *Tenant) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     t.ID,
		Type:   "Tenant",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(t.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.CreatedBy); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "qid.ID",
		Name:  "created_by",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.UpdatedBy); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "qid.ID",
		Name:  "updated_by",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.Name); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	return node, nil
}

func (t *Todo) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     t.ID,
		Type:   "Todo",
		Fields: make([]*Field, 6),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(t.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.CreatedBy); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "qid.ID",
		Name:  "created_by",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.UpdatedBy); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "qid.ID",
		Name:  "updated_by",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.Text); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "text",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.Done); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "bool",
		Name:  "done",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Tenant",
		Name: "tenant",
	}
	err = t.QueryTenant().
		Select(tenant.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "User",
		Name: "user",
	}
	err = t.QueryUser().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (u *User) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     u.ID,
		Type:   "User",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(u.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.CreatedBy); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "qid.ID",
		Name:  "created_by",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.UpdatedBy); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "qid.ID",
		Name:  "updated_by",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Name); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Tenant",
		Name: "tenant",
	}
	err = u.QueryTenant().
		Select(tenant.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Todo",
		Name: "todos",
	}
	err = u.QueryTodos().
		Select(todo.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Organization",
		Name: "organization",
	}
	err = u.QueryOrganization().
		Select(organization.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) Node(ctx context.Context, id qid.ID) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, qid.ID) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, qid.ID) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, qid.ID) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id qid.ID) (string, error) {
			return "", fmt.Errorf("cannot resolve noder (%v) without its type", id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//		c.Noder(ctx, id)
//		c.Noder(ctx, id, ent.WithNodeType(pet.Table))
//
func (c *Client) Noder(ctx context.Context, id qid.ID, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id qid.ID) (Noder, error) {
	switch table {
	case organization.Table:
		n, err := c.Organization.Query().
			Where(organization.ID(id)).
			CollectFields(ctx, "Organization").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case tenant.Table:
		n, err := c.Tenant.Query().
			Where(tenant.ID(id)).
			CollectFields(ctx, "Tenant").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case todo.Table:
		n, err := c.Todo.Query().
			Where(todo.ID(id)).
			CollectFields(ctx, "Todo").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case user.Table:
		n, err := c.User.Query().
			Where(user.ID(id)).
			CollectFields(ctx, "User").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []qid.ID, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]qid.ID)
	id2idx := make(map[qid.ID][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []qid.ID) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[qid.ID][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case organization.Table:
		nodes, err := c.Organization.Query().
			Where(organization.IDIn(ids...)).
			CollectFields(ctx, "Organization").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case tenant.Table:
		nodes, err := c.Tenant.Query().
			Where(tenant.IDIn(ids...)).
			CollectFields(ctx, "Tenant").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case todo.Table:
		nodes, err := c.Todo.Query().
			Where(todo.IDIn(ids...)).
			CollectFields(ctx, "Todo").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case user.Table:
		nodes, err := c.User.Query().
			Where(user.IDIn(ids...)).
			CollectFields(ctx, "User").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}
