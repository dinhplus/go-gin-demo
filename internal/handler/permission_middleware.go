package handler

import (
	"net/http"
	"strconv"

	"my-gin-app/internal/service"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// RequirePermission middleware checks if the user has the required permission for a specific resource and action.
func RequirePermission(resource, action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		roleID, _ := strconv.Atoi(claims["role_id"].(string))
		role, err := service.GetRoleByID(roleID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "role not found"})
			return
		}
		for _, perm := range role.Permissions {

			if perm.Action == "manage" && (perm.Resource.Name == resource || perm.Resource.Name == "all") {
				c.Next()
				return
			}

			if (perm.Resource.Name == "all") && (perm.Action == action || perm.Action == "manage") {
				c.Next()
				return
			}

			if perm.Resource.Name == resource && perm.Action == action {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "permission denied"})
	}
}
