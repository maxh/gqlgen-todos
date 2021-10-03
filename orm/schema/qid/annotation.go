package qid

// Annotation captures the id resourceType for a type.
type Annotation struct {
	ResourceType string
}

// Name implements the ent Annotation interface.
func (a Annotation) Name() string {
	return "QID"
}
