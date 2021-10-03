package util

import "github.com/maxh/gqlgen-todos/qid"

type FieldMap map[string]string
type EdgeMap map[string][]qid.ID
type EntityValue struct {
	Fields FieldMap
	Edges  EdgeMap
}
