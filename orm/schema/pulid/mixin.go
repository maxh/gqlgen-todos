package pulid

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	"github.com/maxh/gqlgen-todos/qrn"
)

// MixinWithPrefix creates a Mixin that encodes the provided resourceType.
func MixinWithPrefix(prefix string) *Mixin {
	return &Mixin{resourceType: prefix}
}

// Mixin defines an ent Mixin that captures the PULID resourceType for a type.
type Mixin struct {
	mixin.Schema
	resourceType string
}

// Fields provides the id field.
func (m Mixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(qrn.ID("")).
			DefaultFunc(func() qrn.ID { return qrn.MustNew(m.resourceType) }),
	}
}

// Annotation captures the id resourceType for a type.
type Annotation struct {
	Prefix string
}

// Name implements the ent Annotation interface.
func (a Annotation) Name() string {
	return "PULID"
}

// Annotations returns the annotations for a Mixin instance.
func (m Mixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		Annotation{Prefix: m.resourceType},
	}
}
