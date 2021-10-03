package qid

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/maxh/gqlgen-todos/qid"
)

func MixinWithPrefix(resourceType string) *Mixin {
	return &Mixin{resourceType: resourceType}
}

// Mixin defines an ent TenantMixin that captures the QID resourceType for a type.
type Mixin struct {
	mixin.Schema
	resourceType string
}

// Fields provides the id field.
func (m Mixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(qid.ID("")).
			DefaultFunc(func() qid.ID { return qid.MustNew(m.resourceType) }),
	}
}

// Annotations returns the annotations for a TenantMixin instance.
func (m Mixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		Annotation{ResourceType: m.resourceType},
	}
}
