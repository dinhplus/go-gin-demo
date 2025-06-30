package handler

import (
	"my-gin-app/internal/model"
	"my-gin-app/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary      List users
// @Description  get all users
// @Tags         users
// @Produce      json
// @Success      200  {array}  model.User
// @Security     BearerAuth
// @Router       /users [get]
func GetUsers(c *gin.Context) {
	users, err := service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary      Create user
// @Description  create a new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      model.User  true  "User to create"
// @Success      201   {object}  model.User
// @Failure      400   {object}  map[string]string
// @Security     BearerAuth
// @Router       /users [post]
func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.AddUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}
