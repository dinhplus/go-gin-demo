package handler

import (
	"my-gin-app/internal/model"
	"my-gin-app/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetDepartments godoc
// @Summary      List departments
// @Description  get all departments
// @Tags         departments
// @Produce      json
// @Success      200  {array}  model.Department
// @Security     BearerAuth
// @Router       /departments [get]
func GetDepartments(c *gin.Context) {
	depts, err := service.GetAllDepartments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, depts)
}

// CreateDepartment godoc
// @Summary      Create department
// @Description  create a new department
// @Tags         departments
// @Accept       json
// @Produce      json
// @Param        department  body      model.Department  true  "Department to create"
// @Success      201   {object}  model.Department
// @Failure      400   {object}  map[string]string
// @Security     BearerAuth
// @Router       /departments [post]
func CreateDepartment(c *gin.Context) {
	var dept model.Department
	if err := c.ShouldBindJSON(&dept); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.AddDepartment(&dept); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dept)
}
