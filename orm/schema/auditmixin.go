package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/maxh/gqlgen-todos/qid"
	"time"
)

// AuditMixin implements the ent.Mixin for sharing
// audit-log capabilities with package schemas.
type AuditMixin struct {
	mixin.Schema
}

// Fields of the AuditMixin.
func (AuditMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.String("created_by").
			GoType(qid.ID("")).
			Optional(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("updated_by").
			GoType(qid.ID("")).
			Optional(),
	}
}

// Hooks of the AuditMixin.
func (AuditMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		AuditHook,
	}
}
