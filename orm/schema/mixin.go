package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/maxh/gqlgen-todos/orm/ent/privacy"
	"github.com/maxh/gqlgen-todos/qid"
)

/////////////
// BASE MIXIN
/////////////

// BaseMixin for all schemas in the graph.
type BaseMixin struct {
	mixin.Schema
}

// Policy defines the privacy policy of the BaseMixin.
func (BaseMixin) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.AlwaysAllowRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}

///////////////
// TENANT MIXIN
///////////////

// TenantMixin for embedding the tenant info in different schemas.
type TenantMixin struct {
	mixin.Schema
}

// Edges for all schemas that embed TenantMixin.
func (TenantMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tenant", Tenant.Type).
			Unique().
			Required(),
	}
}

////////////
// QID MIXIN
////////////

// QidMixinWithPrefix creates a Mixin that encodes the provided resourceType.
func QidMixinWithPrefix(resourceType string) *QidMixin {
	return &QidMixin{resourceType: resourceType}
}

// QidMixin defines an ent Mixin that captures the QID resourceType for a type.
type QidMixin struct {
	mixin.Schema
	resourceType string
}

// Fields provides the id field.
func (m QidMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(qid.ID("")).
			DefaultFunc(func() qid.ID { return qid.MustNew(m.resourceType) }),
	}
}

// Annotations returns the annotations for a Mixin instance.
func (m QidMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		QidAnnotation{ResourceType: m.resourceType},
	}
}
