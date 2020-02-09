package gql

import (
	"go_api_boilerplate/gql/gen"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
)

// GraphqlHandler defines the GQLGen GraphQL server handler
func GraphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	conf := gen.Config{Resolvers: &Resolver{}}
	exec := gen.NewExecutableSchema(conf)
	h := handler.GraphQL(exec)
	return func(c *gin.Context) { h.ServeHTTP(c.Writer, c.Request) }
}

// PlaygroundHandler Defines the Playground handler to expose our playground
func PlaygroundHandler(path string) gin.HandlerFunc {
	h := handler.Playground("GraphQL Playground", path)
	return func(c *gin.Context) { h.ServeHTTP(c.Writer, c.Request) }
}
