package graphql

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/maxh/gqlgen-todos/graphql/gql"
	"github.com/maxh/gqlgen-todos/orm/ent"
)

type Resolver struct{
	client *ent.Client
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return gql.NewExecutableSchema(gql.Config{
		Resolvers: &Resolver{client},
	})
}