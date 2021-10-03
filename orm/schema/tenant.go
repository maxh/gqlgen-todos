package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/maxh/gqlgen-todos/orm/schema/qid"

	"github.com/maxh/gqlgen-todos/orm/schema/base"
)

// Tenant holds the schema definition for the Tenant entity.
type Tenant struct {
	ent.Schema
}

// Mixin of the Tenant schema.
func (Tenant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		base.Mixin{},
		qid.MixinWithPrefix("tenant"),
		AuditMixin{},
	}
}

// Fields of the Tenant.
func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
	}
}
