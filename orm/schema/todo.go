package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/maxh/gqlgen-todos/orm/schema/pulid"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").
			Default("unknown"),
		field.Bool("done").
			Default(false),
	}
}

// Mixin of the Todo.
func (Todo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinWithPrefix("todo"),
		TenantMixin{},
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("todos").
			Unique(),
	}
}
