package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary      Health check
// @Description  Returns pong if the server is running
// @Tags         health
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /ping [get]
// Ping is a simple handler to check if the server is running.
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
