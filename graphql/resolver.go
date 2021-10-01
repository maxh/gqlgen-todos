package graphql

import "github.com/maxh/gqlgen-todos/orm/ent"

//go:generate go run github.com/99designs/gqlgen

type Resolver struct{
	todos []*ent.Todo
}
