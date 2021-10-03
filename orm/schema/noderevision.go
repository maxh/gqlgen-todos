package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/maxh/gqlgen-todos/nodevalue"
	"github.com/maxh/gqlgen-todos/orm/schema/base"
	"github.com/maxh/gqlgen-todos/orm/schema/qid"
)

type NodeRevision struct {
	ent.Schema
}

func (NodeRevision) Fields() []ent.Field {
	return []ent.Field{
		field.String("node_id").
			NotEmpty().
			Immutable(),
		field.String("node_revision").
			NotEmpty().
			Immutable(),
		field.JSON("node_value", &nodevalue.NodeValue{}).
			Immutable(),
	}
}

func (NodeRevision) Mixin() []ent.Mixin {
	return []ent.Mixin{
		base.Mixin{},
		qid.MixinWithPrefix("node_revision"),
		AuditMixin{},
	}
}
