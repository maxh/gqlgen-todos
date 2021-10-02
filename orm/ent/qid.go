// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/maxh/gqlgen-todos/orm/ent/organization"
	"github.com/maxh/gqlgen-todos/orm/ent/tenant"
	"github.com/maxh/gqlgen-todos/orm/ent/todo"
	"github.com/maxh/gqlgen-todos/orm/ent/user"
	"github.com/maxh/gqlgen-todos/qid"
)

// resourceTypeMap maps qid resource types to table names.
var resourceTypeMap = map[qid.ID]string{
	"organization": organization.Table,
	"tenant":       tenant.Table,
	"todo":         todo.Table,
	"user":         user.Table,
}