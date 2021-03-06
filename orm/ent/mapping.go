// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/maxh/gqlgen-todos/orm/ent/noderevision"
	"github.com/maxh/gqlgen-todos/orm/ent/organization"
	"github.com/maxh/gqlgen-todos/orm/ent/tenant"
	"github.com/maxh/gqlgen-todos/orm/ent/todo"
	"github.com/maxh/gqlgen-todos/orm/ent/user"
)

// TablesByResourceType maps qid resource types to table names.
var TablesByResourceType = map[string]string{
	"node_revision": noderevision.Table,
	"organization":  organization.Table,
	"tenant":        tenant.Table,
	"todo":          todo.Table,
	"user":          user.Table,
}

// TablesByEntType maps qid resource types to table names.
var TablesByEntType = map[string]string{
	"NodeRevision": noderevision.Table,
	"Organization": organization.Table,
	"Tenant":       tenant.Table,
	"Todo":         todo.Table,
	"User":         user.Table,
}
