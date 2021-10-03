package schema

import (
	"context"
	"entgo.io/ent"
	"fmt"
	"github.com/maxh/gqlgen-todos/qid"
	"github.com/maxh/gqlgen-todos/viewer"

	"time"
)

// AuditHook for audit-log hook.
func AuditHook(next ent.Mutator) ent.Mutator {
	// AuditLogger wraps the methods that are shared between all mutations of
	// schemas that embed the AuditLog mixin. The variable "exists" is true, if
	// the field already exists in the mutation (e.g. was set by a different hook).
	type AuditLogger interface {
		SetCreatedAt(time.Time)
		CreatedAt() (value time.Time, exists bool)
		SetCreatedBy(qid.ID)
		CreatedBy() (id qid.ID, exists bool)
		SetUpdatedAt(time.Time)
		UpdatedAt() (value time.Time, exists bool)
		SetUpdatedBy(qid.ID)
		UpdatedBy() (id qid.ID, exists bool)
	}
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		ml, ok := m.(AuditLogger)
		if !ok {
			return nil, fmt.Errorf("unexpected audit-log call from mutation type %T", m)
		}
		usr, err := viewer.UserFromContext(ctx)
		if err != nil {
			// TODO: assert acting in god mode
			return next.Mutate(ctx, m)
			// return nil, err
		}
		switch op := m.Op(); {
		case op.Is(ent.OpCreate):
			ml.SetCreatedAt(time.Now())
			if _, exists := ml.CreatedBy(); !exists {
				ml.SetCreatedBy(usr.ID)
			}
		case op.Is(ent.OpUpdateOne | ent.OpUpdate):
			ml.SetUpdatedAt(time.Now())
			if _, exists := ml.UpdatedBy(); !exists {
				ml.SetUpdatedBy(usr.ID)
			}
		}
		return next.Mutate(ctx, m)
	})
}
