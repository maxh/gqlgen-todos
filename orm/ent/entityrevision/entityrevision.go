// Code generated by entc, DO NOT EDIT.

package entityrevision

import (
	"time"

	"entgo.io/ent"
	"github.com/maxh/gqlgen-todos/qid"
)

const (
	// Label holds the string label denoting the entityrevision type in the database.
	Label = "entity_revision"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldEntityID holds the string denoting the entity_id field in the database.
	FieldEntityID = "entity_id"
	// FieldEntityRevision holds the string denoting the entity_revision field in the database.
	FieldEntityRevision = "entity_revision"
	// FieldEntityValue holds the string denoting the entity_value field in the database.
	FieldEntityValue = "entity_value"
	// Table holds the table name of the entityrevision in the database.
	Table = "entity_revisions"
)

// Columns holds all SQL columns for entityrevision fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldEntityID,
	FieldEntityRevision,
	FieldEntityValue,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/maxh/gqlgen-todos/orm/ent/runtime"
//
var (
	Hooks  [2]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// EntityIDValidator is a validator for the "entity_id" field. It is called by the builders before save.
	EntityIDValidator func(string) error
	// EntityRevisionValidator is a validator for the "entity_revision" field. It is called by the builders before save.
	EntityRevisionValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() qid.ID
)