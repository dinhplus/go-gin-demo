package handler

import (
	"my-gin-app/internal/model"
	"my-gin-app/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetStacks godoc
// @Summary      List stacks
// @Description  get all stacks
// @Tags         stacks
// @Produce      json
// @Success      200  {array}  model.Stack
// @Security     BearerAuth
// @Router       /stacks [get]
func GetStacks(c *gin.Context) {
	stacks, err := service.GetAllStacks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stacks)
}

// CreateStack godoc
// @Summary      Create stack
// @Description  create a new stack
// @Tags         stacks
// @Accept       json
// @Produce      json
// @Param        stack  body      model.Stack  true  "Stack to create"
// @Success      201   {object}  model.Stack
// @Failure      400   {object}  map[string]string
// @Security     BearerAuth
// @Router       /stacks [post]
func CreateStack(c *gin.Context) {
	var stack model.Stack
	if err := c.ShouldBindJSON(&stack); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.AddStack(&stack); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, stack)
}
