package base

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/mixin"
	"github.com/maxh/gqlgen-todos/orm/ent/privacy"
	"github.com/maxh/gqlgen-todos/orm/rule"
)

// Mixin for all schemas in the graph.
type Mixin struct {
	mixin.Schema
}

// Policy defines the privacy policy of the BaseMixin.
func (Mixin) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.DenyIfNoViewer(),
		},
		Query: privacy.QueryPolicy{
			rule.DenyIfNoViewer(),
		},
	}
}
