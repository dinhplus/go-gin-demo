package handler

import (
	"my-gin-app/internal/graphql/generated"
	"my-gin-app/internal/graphql/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

// GraphQLHandler creates a new GraphQL handler with authentication
func GraphQLHandler() gin.HandlerFunc {
	// Create a new GraphQL server with the generated schema and resolvers
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolvers.Resolver{},
	}))

	return gin.HandlerFunc(func(c *gin.Context) {
		// Apply GraphQL auth middleware
		GraphQLAuthMiddleware()(c)
		if c.IsAborted() {
			return
		}

		// Wrap the GraphQL handler
		gin.WrapH(srv)(c)
	})
}

// PlaygroundHandler creates a GraphQL Playground handler for development
func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL Playground", "/graphql")
	return gin.WrapH(h)
}
