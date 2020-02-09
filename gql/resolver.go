package gql

import (
	"go_api_boilerplate/gql/gen"
	"go_api_boilerplate/services/authservice"
	"go_api_boilerplate/services/userservice"
)

// Resolver struct
type Resolver struct {
	UserService userservice.UserService
	AuthService authservice.AuthService
}

// Mutation graphql
func (r *Resolver) Mutation() gen.MutationResolver {
	return &mutationResolver{r}
}

// Query graphql
func (r *Resolver) Query() gen.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
