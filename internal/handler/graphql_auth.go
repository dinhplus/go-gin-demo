package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

// GraphQLAuthMiddleware provides JWT authentication for GraphQL endpoints
func GraphQLAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// For queries that don't require auth, we can continue
			c.Next()
			return
		}

		// Check if it's a Bearer token
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			return
		}

		// Extract the token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse and validate the token (you'll need to use your JWT secret)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Make sure the signing method is what you expect
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Use the same secret key as the JWT middleware
			return []byte("secret key"), nil
		})

		if err != nil || !token.Valid {
			fmt.Println("Invalid token:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// Store user claims in context for GraphQL resolvers
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Add claims to Gin context
			c.Set("user_claims", claims)

			// Also add to GraphQL context for resolvers
			ctx := context.WithValue(c.Request.Context(), "user_claims", claims)
			c.Request = c.Request.WithContext(ctx)
		}

		c.Next()
	}
}

// GetUserFromContext extracts user claims from GraphQL context
func GetUserFromContext(ctx context.Context) (jwt.MapClaims, error) {
	claims, ok := ctx.Value("user_claims").(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("no user claims found in context")
	}
	return claims, nil
}

// RequireAuth is a GraphQL directive function to require authentication
func RequireAuth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	_, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("authentication required")
	}
	return next(ctx)
}
