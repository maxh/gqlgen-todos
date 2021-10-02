// Code generated by entc, DO NOT EDIT.

package todo

import (
	"github.com/maxh/gqlgen-todos/qrn"
)

const (
	// Label holds the string label denoting the todo type in the database.
	Label = "todo"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldDone holds the string denoting the done field in the database.
	FieldDone = "done"
	// EdgeTenant holds the string denoting the tenant edge name in mutations.
	EdgeTenant = "tenant"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the todo in the database.
	Table = "todos"
	// TenantTable is the table that holds the tenant relation/edge.
	TenantTable = "todos"
	// TenantInverseTable is the table name for the Tenant entity.
	// It exists in this package in order to avoid circular dependency with the "tenant" package.
	TenantInverseTable = "tenants"
	// TenantColumn is the table column denoting the tenant relation/edge.
	TenantColumn = "todo_tenant"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "todos"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_todos"
)

// Columns holds all SQL columns for todo fields.
var Columns = []string{
	FieldID,
	FieldText,
	FieldDone,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "todos"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"todo_tenant",
	"user_todos",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultText holds the default value on creation for the "text" field.
	DefaultText string
	// DefaultDone holds the default value on creation for the "done" field.
	DefaultDone bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() qrn.ID
)
