package schema

import "entgo.io/ent/schema"

// QidAnnotation captures the id resourceType for a type.
type QidAnnotation struct {
	ResourceType string
}

// Name implements the ent Annotation interface.
func (a QidAnnotation) Name() string {
	return "QID"
}

// Annotations returns the annotations for a Mixin instance.
func (m QidMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		QidAnnotation{ResourceType: m.resourceType},
	}
}
