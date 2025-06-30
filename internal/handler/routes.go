package handler

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", Ping)

	r.GET("/users", GetUsers)
	r.POST("/users", CreateUser)

	r.GET("/departments", GetDepartments)
	r.POST("/departments", CreateDepartment)

	r.GET("/stacks", GetStacks)
	r.POST("/stacks", CreateStack)

	r.GET("/positions", GetPositions)
	r.POST("/positions", CreatePosition)

	return r
}
