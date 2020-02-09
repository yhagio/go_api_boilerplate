package gql

import (
	"context"
	"go_api_boilerplate/gql/gen"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() gen.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gen.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input gen.NewTodo) (*gen.Todo, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*gen.Todo, error) {
	panic("not implemented")
}

func (r *queryResolver) Users(ctx context.Context) ([]*gen.User, error) {
	records := []*gen.User{
		&gen.User{
			ID:    "abc123",
			Email: "alice@email.com",
			Name:  "alice",
			Age:   23,
		},
	}
	return records, nil
}
