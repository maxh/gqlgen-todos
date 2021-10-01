package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/maxh/gqlgen-todos/graphql/gql"
	"github.com/maxh/gqlgen-todos/graphql/model"
	"github.com/maxh/gqlgen-todos/orm/ent"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*ent.Todo, error) {
	return nil, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*ent.Todo, error) {
	return []*ent.Todo{}, nil
}

func (r *todoResolver) ID(ctx context.Context, obj *ent.Todo) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) User(ctx context.Context, obj *ent.Todo) (*ent.User, error) {
	return nil, nil
}

func (r *userResolver) ID(ctx context.Context, obj *ent.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns gql.MutationResolver implementation.
func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

// Todo returns gql.TodoResolver implementation.
func (r *Resolver) Todo() gql.TodoResolver { return &todoResolver{r} }

// User returns gql.UserResolver implementation.
func (r *Resolver) User() gql.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
