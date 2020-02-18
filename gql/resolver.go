package gql

import (
	"github.com/yhagio/go_api_boilerplate/gql/gen"
	"github.com/yhagio/go_api_boilerplate/services/authservice"
	"github.com/yhagio/go_api_boilerplate/services/emailservice"
	"github.com/yhagio/go_api_boilerplate/services/userservice"
)

// Resolver struct
type Resolver struct {
	UserService  userservice.UserService
	AuthService  authservice.AuthService
	EmailService emailservice.EmailService
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
