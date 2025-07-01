package handler

import (
	"my-gin-app/internal/model"
	"my-gin-app/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary      Health check
// @Description  Returns pong if the server is running
// @Tags         health
// @Produce      json
// @Success      200  {object}  map[string]string
// @Security     BearerAuth
// @Router       /ping [get]
// Ping is a simple handler to check if the server is running.
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// GetPermissions godoc
// @Summary      List permissions
// @Description  get all permissions
// @Tags         permissions
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.Permission
// @Router       /permissions [get]
func GetPermissions(c *gin.Context) {
	perms, err := service.GetAllPermissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, perms)
}

// CreatePermission godoc
// @Summary      Create permission
// @Description  create a new permission
// @Tags         permissions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        permission  body      model.Permission  true  "Permission to create"
// @Success      201   {object}  model.Permission
// @Failure      400   {object}  map[string]string
// @Router       /permissions [post]
func CreatePermission(c *gin.Context) {
	var perm model.Permission
	if err := c.ShouldBindJSON(&perm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.AddPermission(&perm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, perm)
}

// GetResources godoc
// @Summary      List resources
// @Description  get all resources
// @Tags         resources
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.Resource
// @Router       /resources [get]
func GetResources(c *gin.Context) {
	resources, err := service.GetAllResources()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resources)
}

// CreateResource godoc
// @Summary      Create resource
// @Description  create a new resource
// @Tags         resources
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        resource  body      model.Resource  true  "Resource to create"
// @Success      201   {object}  model.Resource
// @Failure      400   {object}  map[string]string
// @Router       /resources [post]
func CreateResource(c *gin.Context) {
	var resource model.Resource
	if err := c.ShouldBindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.AddResource(&resource); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resource)
}

// GetFeatureFlags godoc
// @Summary      List feature flags
// @Description  get all feature flags (distinct from permissions)
// @Tags         feature-flags
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  string
// @Router       /feature-flags [get]
func GetFeatureFlags(c *gin.Context) {
	flags, err := service.GetAllFeatureFlags()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, flags)
}

// CreateFeatureFlag godoc
// @Summary      Create feature flag
// @Description  create a new feature flag (as a permission with feature_flag field)
// @Tags         feature-flags
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        flag  body      map[string]string  true  "Feature flag info"
// @Success      201   {object}  model.Permission
// @Failure      400   {object}  map[string]string
// @Router       /feature-flags [post]
func CreateFeatureFlag(c *gin.Context) {
	var req map[string]string
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	flag := req["feature_flag"]
	if flag == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "feature_flag is required"})
		return
	}
	perm := model.Permission{FeatureFlag: flag}
	if err := service.AddPermission(&perm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, perm)
}
