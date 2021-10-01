package graphql

//go:generate go run github.com/99designs/gqlgen

import "github.com/maxh/gqlgen-todos/graphql/model"

type Resolver struct{
	todos []*model.Todo
}
