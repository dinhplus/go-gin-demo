package handler

import (
	"my-gin-app/internal/model"
	"my-gin-app/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPositions godoc
// @Summary      List positions
// @Description  get all positions
// @Tags         positions
// @Produce      json
// @Success      200  {array}  model.Position
// @Security     BearerAuth
// @Router       /positions [get]
func GetPositions(c *gin.Context) {
	positions, err := service.GetAllPositions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, positions)
}

// CreatePosition godoc
// @Summary      Create position
// @Description  create a new position
// @Tags         positions
// @Accept       json
// @Produce      json
// @Param        position  body      model.Position  true  "Position to create"
// @Success      201   {object}  model.Position
// @Failure      400   {object}  map[string]string
// @Security     BearerAuth
// @Router       /positions [post]
func CreatePosition(c *gin.Context) {
	var pos model.Position
	if err := c.ShouldBindJSON(&pos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.AddPosition(&pos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, pos)
}
