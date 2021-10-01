package graphql

import "github.com/maxh/gqlgen-todos/orm/ent"

type Resolver struct{
	todos []*ent.Todo
}
