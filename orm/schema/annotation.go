package schema

// QidAnnotation captures the id resourceType for a type.
type QidAnnotation struct {
	ResourceType string
}

// Name implements the ent Annotation interface.
func (a QidAnnotation) Name() string {
	return "QID"
}
