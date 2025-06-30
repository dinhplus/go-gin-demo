package main

import (
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

func main() {
	repository.InitDB()
	r := handler.SetupRouter()

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080
}
