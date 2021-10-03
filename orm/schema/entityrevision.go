package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/maxh/gqlgen-todos/orm/schema/base"
	"github.com/maxh/gqlgen-todos/orm/schema/qid"
	"github.com/maxh/gqlgen-todos/util"
)

type EntityRevision struct {
	ent.Schema
}

func (EntityRevision) Fields() []ent.Field {
	return []ent.Field{
		field.String("entity_id").
			NotEmpty().
			Immutable(),
		field.String("entity_revision").
			NotEmpty().
			Immutable(),
		field.JSON("entity_value", &util.Any).
			Immutable(),
	}
}

func (EntityRevision) Mixin() []ent.Mixin {
	return []ent.Mixin{
		base.Mixin{},
		qid.MixinWithPrefix("entity_revision"),
		AuditMixin{},
	}
}
