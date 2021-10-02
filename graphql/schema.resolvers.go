package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/maxh/gqlgen-todos/graphql/gql"
	"github.com/maxh/gqlgen-todos/orm/ent"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input gql.CreateTodoInput) (*ent.Todo, error) {
	client := ent.FromContext(ctx)
	user, err := client.User.Get(ctx, input.UserID)
	if err != nil {
		log.Fatal("unable to get user", err)
	}
	tenant, err := user.Tenant(ctx)
	if err != nil {
		log.Fatal("unable to get tenant", err)
	}
	created, err := client.Todo.
		Create().
		SetUserID(input.UserID).
		SetText(input.Text).
		SetTenant(tenant).
		Save(ctx)
	if err != nil {
		log.Fatal("unable to create todo", err)
	}
	print(created.String())
	return created, err
}

func (r *queryResolver) Todos(ctx context.Context) ([]*ent.Todo, error) {
	return r.client.Todo.Query().All(ctx)
}

func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().All(ctx)
}

// Mutation returns gql.MutationResolver implementation.
func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
