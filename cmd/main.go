package main

import (
	"fmt"
	_ "my-gin-app/docs" // Import generated Swagger docs
	"my-gin-app/internal/handler"
	"my-gin-app/internal/repository"
)

// @title My Gin App API
// @version 1.0
// @description This is a sample server for a Gin application.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description JWT Authorization header using the Bearer scheme. Example: "Authorization: Bearer {token}"
// @Security ApiKeyAuth
func main() {
	repository.InitDB()
	r := handler.SetupRouter()

	fmt.Println("Starting server on http://localhost:8080")
	fmt.Println("Swagger UI: http://localhost:8080/swagger/index.html")
	r.Run() // listen and serve on 0.0.0.0:8080
}
