// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/maxh/gqlgen-todos/orm/ent/entityrevision"
	"github.com/maxh/gqlgen-todos/qid"
	"github.com/maxh/gqlgen-todos/nodevalue"
)

// EntityRevisionCreate is the builder for creating a EntityRevision entity.
type EntityRevisionCreate struct {
	config
	mutation *EntityRevisionMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (erc *EntityRevisionCreate) SetCreatedAt(t time.Time) *EntityRevisionCreate {
	erc.mutation.SetCreatedAt(t)
	return erc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (erc *EntityRevisionCreate) SetNillableCreatedAt(t *time.Time) *EntityRevisionCreate {
	if t != nil {
		erc.SetCreatedAt(*t)
	}
	return erc
}

// SetCreatedBy sets the "created_by" field.
func (erc *EntityRevisionCreate) SetCreatedBy(q qid.ID) *EntityRevisionCreate {
	erc.mutation.SetCreatedBy(q)
	return erc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (erc *EntityRevisionCreate) SetNillableCreatedBy(q *qid.ID) *EntityRevisionCreate {
	if q != nil {
		erc.SetCreatedBy(*q)
	}
	return erc
}

// SetUpdatedAt sets the "updated_at" field.
func (erc *EntityRevisionCreate) SetUpdatedAt(t time.Time) *EntityRevisionCreate {
	erc.mutation.SetUpdatedAt(t)
	return erc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (erc *EntityRevisionCreate) SetNillableUpdatedAt(t *time.Time) *EntityRevisionCreate {
	if t != nil {
		erc.SetUpdatedAt(*t)
	}
	return erc
}

// SetUpdatedBy sets the "updated_by" field.
func (erc *EntityRevisionCreate) SetUpdatedBy(q qid.ID) *EntityRevisionCreate {
	erc.mutation.SetUpdatedBy(q)
	return erc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (erc *EntityRevisionCreate) SetNillableUpdatedBy(q *qid.ID) *EntityRevisionCreate {
	if q != nil {
		erc.SetUpdatedBy(*q)
	}
	return erc
}

// SetEntityID sets the "entity_id" field.
func (erc *EntityRevisionCreate) SetEntityID(s string) *EntityRevisionCreate {
	erc.mutation.SetEntityID(s)
	return erc
}

// SetEntityRevision sets the "entity_revision" field.
func (erc *EntityRevisionCreate) SetEntityRevision(s string) *EntityRevisionCreate {
	erc.mutation.SetEntityRevision(s)
	return erc
}

// SetEntityValue sets the "entity_value" field.
func (erc *EntityRevisionCreate) SetEntityValue(uv *nodevalue.NodeValue) *EntityRevisionCreate {
	erc.mutation.SetEntityValue(uv)
	return erc
}

// SetID sets the "id" field.
func (erc *EntityRevisionCreate) SetID(q qid.ID) *EntityRevisionCreate {
	erc.mutation.SetID(q)
	return erc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (erc *EntityRevisionCreate) SetNillableID(q *qid.ID) *EntityRevisionCreate {
	if q != nil {
		erc.SetID(*q)
	}
	return erc
}

// Mutation returns the EntityRevisionMutation object of the builder.
func (erc *EntityRevisionCreate) Mutation() *EntityRevisionMutation {
	return erc.mutation
}

// Save creates the EntityRevision in the database.
func (erc *EntityRevisionCreate) Save(ctx context.Context) (*EntityRevision, error) {
	var (
		err  error
		node *EntityRevision
	)
	if err := erc.defaults(); err != nil {
		return nil, err
	}
	if len(erc.hooks) == 0 {
		if err = erc.check(); err != nil {
			return nil, err
		}
		node, err = erc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntityRevisionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = erc.check(); err != nil {
				return nil, err
			}
			erc.mutation = mutation
			if node, err = erc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(erc.hooks) - 1; i >= 0; i-- {
			if erc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = erc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, erc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (erc *EntityRevisionCreate) SaveX(ctx context.Context) *EntityRevision {
	v, err := erc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (erc *EntityRevisionCreate) Exec(ctx context.Context) error {
	_, err := erc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (erc *EntityRevisionCreate) ExecX(ctx context.Context) {
	if err := erc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (erc *EntityRevisionCreate) defaults() error {
	if _, ok := erc.mutation.CreatedAt(); !ok {
		if entityrevision.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized entityrevision.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := entityrevision.DefaultCreatedAt()
		erc.mutation.SetCreatedAt(v)
	}
	if _, ok := erc.mutation.UpdatedAt(); !ok {
		if entityrevision.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized entityrevision.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := entityrevision.DefaultUpdatedAt()
		erc.mutation.SetUpdatedAt(v)
	}
	if _, ok := erc.mutation.ID(); !ok {
		if entityrevision.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized entityrevision.DefaultID (forgotten import ent/runtime?)")
		}
		v := entityrevision.DefaultID()
		erc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (erc *EntityRevisionCreate) check() error {
	if _, ok := erc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := erc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := erc.mutation.EntityID(); !ok {
		return &ValidationError{Name: "entity_id", err: errors.New(`ent: missing required field "entity_id"`)}
	}
	if v, ok := erc.mutation.EntityID(); ok {
		if err := entityrevision.EntityIDValidator(v); err != nil {
			return &ValidationError{Name: "entity_id", err: fmt.Errorf(`ent: validator failed for field "entity_id": %w`, err)}
		}
	}
	if _, ok := erc.mutation.EntityRevision(); !ok {
		return &ValidationError{Name: "entity_revision", err: errors.New(`ent: missing required field "entity_revision"`)}
	}
	if v, ok := erc.mutation.EntityRevision(); ok {
		if err := entityrevision.EntityRevisionValidator(v); err != nil {
			return &ValidationError{Name: "entity_revision", err: fmt.Errorf(`ent: validator failed for field "entity_revision": %w`, err)}
		}
	}
	if _, ok := erc.mutation.EntityValue(); !ok {
		return &ValidationError{Name: "entity_value", err: errors.New(`ent: missing required field "entity_value"`)}
	}
	return nil
}

func (erc *EntityRevisionCreate) sqlSave(ctx context.Context) (*EntityRevision, error) {
	_node, _spec := erc.createSpec()
	if err := sqlgraph.CreateNode(ctx, erc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(qid.ID)
	}
	return _node, nil
}

func (erc *EntityRevisionCreate) createSpec() (*EntityRevision, *sqlgraph.CreateSpec) {
	var (
		_node = &EntityRevision{config: erc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: entityrevision.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: entityrevision.FieldID,
			},
		}
	)
	if id, ok := erc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := erc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entityrevision.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := erc.mutation.CreatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entityrevision.FieldCreatedBy,
		})
		_node.CreatedBy = value
	}
	if value, ok := erc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entityrevision.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := erc.mutation.UpdatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entityrevision.FieldUpdatedBy,
		})
		_node.UpdatedBy = value
	}
	if value, ok := erc.mutation.EntityID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entityrevision.FieldEntityID,
		})
		_node.EntityID = value
	}
	if value, ok := erc.mutation.EntityRevision(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entityrevision.FieldEntityRevision,
		})
		_node.EntityRevision = value
	}
	if value, ok := erc.mutation.EntityValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: entityrevision.FieldEntityValue,
		})
		_node.EntityValue = value
	}
	return _node, _spec
}

// EntityRevisionCreateBulk is the builder for creating many EntityRevision entities in bulk.
type EntityRevisionCreateBulk struct {
	config
	builders []*EntityRevisionCreate
}

// Save creates the EntityRevision entities in the database.
func (ercb *EntityRevisionCreateBulk) Save(ctx context.Context) ([]*EntityRevision, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ercb.builders))
	nodes := make([]*EntityRevision, len(ercb.builders))
	mutators := make([]Mutator, len(ercb.builders))
	for i := range ercb.builders {
		func(i int, root context.Context) {
			builder := ercb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EntityRevisionMutation)
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
					_, err = mutators[i+1].Mutate(root, ercb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ercb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ercb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ercb *EntityRevisionCreateBulk) SaveX(ctx context.Context) []*EntityRevision {
	v, err := ercb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ercb *EntityRevisionCreateBulk) Exec(ctx context.Context) error {
	_, err := ercb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ercb *EntityRevisionCreateBulk) ExecX(ctx context.Context) {
	if err := ercb.Exec(ctx); err != nil {
		panic(err)
	}
}
